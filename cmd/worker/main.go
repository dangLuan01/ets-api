package main

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/dangLuan01/ets-api/internal/app"
	"github.com/dangLuan01/ets-api/internal/config"
	"github.com/dangLuan01/ets-api/internal/db"
	"github.com/dangLuan01/ets-api/pkg/cache"
	"github.com/doug-martin/goqu/v9"
)

type Worker struct {
	db		*goqu.Database
	cache 	cache.RedisCacheService
}

func NewWorker(db *goqu.Database, cache cache.RedisCacheService) *Worker {
	return &Worker{
		db: 	db,
		cache: 	cache,
	}
}

func (w *Worker) Start() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		func ()  {
			defer func()  {
				if r := recover(); r != nil {
					//log.Printf("🔥 [Worker Error] Phát hiện Panic: %v. Tự phục hồi để chạy tiếp chu kỳ sau...", r)
				}
			}()

			ctx, cancel := context.WithTimeout(context.Background(), 8 *time.Second)
			w.executeSync(ctx)
			
			defer cancel()	
		}()
	}
}

func (w *Worker) executeSync(ctx context.Context) {
	exists, _ := w.cache.Exits(ctx, "buffered:exam:clicks:sync")
	if !exists {
		exists, _ := w.cache.Exits(ctx, "buffered:exam:clicks")
		if !exists {return}

		err := w.cache.Rename(ctx, "buffered:exam:clicks", "buffered:exam:clicks:sync")
		if err != nil {return}
	}

	data, err := w.cache.HGetAll(ctx, "buffered:exam:clicks:sync")
	if err != nil || len(data) == 0 {
		return 
	}
	
	tx, err := w.db.Begin()
	if err != nil {return}

	for slug, countStr := range data {
		count, _ := strconv.Atoi(countStr)

		_, err := tx.Update("exams").Set(goqu.Record{
			"count": goqu.L("count + ?", count),
		}).Where(goqu.C("slug").Eq(slug)).Executor().ExecContext(ctx)

		if err != nil {
			tx.Rollback()
			return
		}
	}

	if err := tx.Commit(); err != nil {return}

	w.cache.Clear(ctx, "buffered:exam:clicks:sync")
}

func main() {
	app.LoadEnv()
	
	cfg := config.NewConfig()

	redisClient := config.NewRedisClient()
	cacheRedisService := cache.NewRedisCacheService(redisClient)

	if err := db.InitDB(cfg); err != nil {
		log.Fatalf("⛔ Unable to connect to sql")
	}

	worker := NewWorker(db.DB, cacheRedisService)
	worker.Start()
}