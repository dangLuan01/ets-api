package app

import (
	"log"

	"github.com/dangLuan01/ets-api/internal/config"
	"github.com/dangLuan01/ets-api/internal/db"
	"github.com/dangLuan01/ets-api/internal/routes"
	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/dangLuan01/ets-api/internal/validation"
	"github.com/dangLuan01/ets-api/pkg/auth"
	"github.com/dangLuan01/ets-api/pkg/cache"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type Module interface {
	Routes() routes.Route
}

type Application struct {
	config *config.Config
	router *gin.Engine
	modules []Module
}

type ModuleContext struct {
	DB *goqu.Database
	Redis *redis.Client
}

func NewApplication(cfg *config.Config) (*Application, error) {
	mode := utils.GetEnv("MODE", "release")
	if mode == "release" {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)

	if err := validation.InitValidator(); err != nil {
		log.Fatalf("⛔ Validation init failed %v:", err)
	}
	
	r := gin.Default()
	
	if err := db.InitDB(); err != nil {
		log.Fatalf("⛔ Unable to connect to sql")
	}

	redisClient := config.NewRedisClient()
	cacheRedisService := cache.NewRedisCacheService(redisClient)

	tokenService := auth.NewJWTService(cacheRedisService)
	
	// s3Client := config.NewS3Client()
	// storeS3Service := s3.NewS3Service(s3Client)

	// factory, err := mail.NewProviderFactory(mail.ProviderResent)
	// if err != nil {
	// 	log.Fatalf("⛔ Unable to init mail:%s", err)
	// 	return nil, err	
	// }

	// mailService, err := mail.NewMailService(cfg, factory)
	// if err != nil {
	// 	log.Fatalf("⛔ Unable to init mail service:%s", err)
	// }

	// rabbitmqService, err := rabbitmq.NewRabitMQService(utils.GetEnv("RABBITMQ_URL",""))
	// if err != nil {
	// 	log.Fatalf("⛔ Unable to connect queue.")
	// }
	
	ctx := &ModuleContext{
		DB: db.DB,
		Redis: redisClient,
	}

	modules := []Module{
		NewUserModule(ctx),
		NewExamModule(ctx),
		// NewAuthModule(ctx, tokenService, cacheRedisService, mailService, rabbitmqService),
		// NewStreamingModule(ctx, cacheRedisService, storeS3Service),
		// NewVideoModule(ctx, storeS3Service),
		// NewAccountModule(ctx),
		// NewPaymentModule(ctx, cacheRedisService, rabbitmqService),
		// NewPartnerModule(ctx),
	}

	routes.RegisterRoute(r, tokenService, cacheRedisService ,getModuleRoutes(modules)...)

	return &Application{
		config: cfg,
		router: r,
		modules: modules,
	}, nil
}

func (a *Application) Run() error {
	
	return a.router.Run(a.config.ServerAddress)
}

func getModuleRoutes(modules []Module) []routes.Route {
	routeList := make([]routes.Route, len(modules))
	for i, module := range modules {
		routeList[i] = module.Routes()
	}

	return routeList
}

func LoadEnv() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}