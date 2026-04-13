package v1service

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"time"

	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	repositoryExam "github.com/dangLuan01/ets-api/internal/repository/exam"
	repositoryPartDirection "github.com/dangLuan01/ets-api/internal/repository/part_direction"
	repositoryQuestion "github.com/dangLuan01/ets-api/internal/repository/question"
	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type examService struct {
	db *goqu.Database
	repo repositoryExam.ExamRepository
	repoPartDirection repositoryPartDirection.PartDirectionRepository
	repoQuestion repositoryQuestion.QuestionRepository
}

func NewExamService(repo repositoryExam.ExamRepository, DB *goqu.Database, repoPartDirection repositoryPartDirection.PartDirectionRepository, repoQuestion repositoryQuestion.QuestionRepository) ExamService {
	return &examService{
		db: DB,
		repo: repo,
		repoPartDirection: repoPartDirection,
		repoQuestion: repoQuestion,
	}
}

func (es *examService) FindExamById(examId int) (models.Exam, error) {

	exam, err := es.repo.FindExamById(examId)
	if err != nil {
		return models.Exam{}, err
	}

	sections, err := es.repo.FindExamQuestionMappingById(examId)
	if err != nil {
		return models.Exam{}, err
	}

	singleIDs 	:= make([]int, 0)
	groupIDs	:= make([]int, 0)

	for _, s := range sections {
		switch s.EntityType {
		case "SINGLE":
			singleIDs = append(singleIDs, s.EntityId)
		case "GROUP":
			groupIDs = append(groupIDs, s.EntityId)
		}
	}

	
	questionMap 	:= make(map[int]models.Question)
	groupMap 		:= make(map[int]*models.QuestionGroup)

	directionMap 	:= make(map[int]models.Direction)
	directions, err := es.repoPartDirection.FindDirectionByExamId(examId)
	if err == nil {
		for i := range directions {
			d := &directions[i]
			if len(d.ExampleRaw) > 0 {
				var ex models.ExampleData
				if err := json.Unmarshal(d.ExampleRaw, &ex); err == nil {
					d.Example = &ex
				}
			}
			directionMap[d.PartId] = *d
		}
	}

	if len(singleIDs) > 0 {

		questions, err := es.repo.FindQuesionByIds(singleIDs)
		
		if err != nil {
			return models.Exam{}, err
		}

		for i := range questions {
			
			q := &questions[i]

			opts := map[string]*string{
				"A": q.OptionA,
				"B": q.OptionB,
				"C": q.OptionC,
				"D": q.OptionD,
			}
			q.Options = opts
			questionMap[q.Id] = *q
		}
	}

	if len(groupIDs) > 0 {

		groups, err := es.repo.FindGroupQuestionByIds(groupIDs)
		
		if err != nil {
			return models.Exam{}, err
		}

		for _, g := range groups {
			gCopy := g
			groupMap[g.Id] = &gCopy
		}

		subQuestions, err := es.repo.FindSubQuesionByGroupIds(groupIDs)

		if err != nil { 
			return models.Exam{}, err 
		}

		for i := range subQuestions {
			q := &subQuestions[i]

			q.Options = map[string]*string{
				"A": q.OptionA,
				"B": q.OptionB,
				"C": q.OptionC,
				"D": q.OptionD,
			}

			g := groupMap[*q.GroupId]
			g.SubQuestions = append(g.SubQuestions, *q)
			groupMap[*q.GroupId] = g
		}
	}

	for i := range sections {
		s := &sections[i]		

		switch s.EntityType {
		case "SINGLE":
			if q, ok := questionMap[s.EntityId]; ok {
				q.DisplayNumber = s.OrderIndex
				qCopy := q
				s.QuestionData = &qCopy
			}

		case "GROUP":
			if g, ok := groupMap[s.EntityId]; ok {
				for j := range g.SubQuestions {
					g.SubQuestions[j].DisplayNumber = s.OrderIndex + j
				}

				s.GroupData = g
			}
		}
	}

	skillsMater, err := es.repo.FindSkillsByCertId(exam.CertificateId)
	if err != nil {
		return models.Exam{}, err
	}

	partsMaster, err := es.repo.FindPartsByCertId(exam.CertificateId)
	if err != nil {
		return models.Exam{}, err
	}

	skillMasterMap := make(map[int]models.SkillMaster)
	for _, sm := range skillsMater {
		skillMasterMap[sm.Id] = sm
	}

	partMasterMap := make(map[int]models.PartMaster)
    for _, pm := range partsMaster {
        partMasterMap[pm.Id] = pm
    }

	sectionsByPart := make(map[int][]models.ExamQuestionMapping)

	for _, s := range sections {
		sectionsByPart[s.PartId] = append(sectionsByPart[s.PartId], s)
	}

	partIDSet := make(map[int]bool)
	for partID := range sectionsByPart {
		partIDSet[partID] = true
	}
	for partID := range directionMap {
		partIDSet[partID] = true
	}

	partsBySkill := make(map[int][]models.ExamPart)

	for partID := range partIDSet {
		pm, ok := partMasterMap[partID]
		if !ok {
			continue
		}

		var dir *models.Direction
		if d, exist := directionMap[partID]; exist {
			dCopy := d
			dir = &dCopy
		}

		items := sectionsByPart[partID]
        if items == nil {
            items = []models.ExamQuestionMapping{}
        }

		examPart := models.ExamPart{
			PartId: partID,
			PartNumber: pm.PartNumber,
			PartName: pm.Name,
			Direction: dir,
			Items: items,
		}

		partsBySkill[pm.SkillId] = append(partsBySkill[pm.SkillId], examPart)
	}

	var examSkills []models.ExamSkill

	for skillID, examParts := range partsBySkill {
        sm, ok := skillMasterMap[skillID]
        if !ok { continue }

        // Sắp xếp các Part bên trong Skill theo PartNumber (0, 1, 2, 3...)
        sort.Slice(examParts, func(i, j int) bool {
            return examParts[i].PartNumber < examParts[j].PartNumber
        })

        examSkills = append(examSkills, models.ExamSkill{
            SkillId:   skillID,
            SkillCode: sm.Code,
            SkillName: sm.Name,
            Parts:     examParts,
        })
    }

	sort.Slice(examSkills, func(i, j int) bool {
        return skillMasterMap[examSkills[i].SkillId].OrderIndex < skillMasterMap[examSkills[j].SkillId].OrderIndex
    })

	exam.Skills = examSkills

	return exam, nil
}

func (es *examService) CalculateScoreExam(ctx *gin.Context, params v1dto.QuestionAnswerInputParams) (v1dto.DetailExamScore, error) {

	questionIds 	:= make([]int, 0, len(params.Answers))
	userAnswerMap 	:= make(map[int]string)

	for _, ans := range params.Answers {
		questionIds = append(questionIds, ans.QuestionId)
		userAnswerMap[ans.QuestionId] = ans.SelectedAnswer
	}

	correctAnswer, err := es.repo.GetCorrectAnswersWithSkillContext(params.ExamId, questionIds)
	if err != nil {
		return v1dto.DetailExamScore{}, err
	}
	
	rawScores := make(map[int]int)
	var detailsAnswers []models.UserAnswer
	
	for _, ca := range correctAnswer {
		isCorrect := false
		if _, ok := rawScores[ca.SkillId]; !ok {
        	rawScores[ca.SkillId] = 0
    	}

		if userAnswerMap[ca.QuestionId] == ca.CorrectAnswer {
			isCorrect = true
			rawScores[ca.SkillId]++
		}

		var selectedAnswer *string
		if ans, ok := userAnswerMap[ca.QuestionId]; ok && ans != "" {
			selectedAnswer = &ans
		} else {
			selectedAnswer = nil
		}

		detailsAnswers = append(detailsAnswers, models.UserAnswer{
			QuestionId: ca.QuestionId,
			SelectedAnswer: selectedAnswer,
			IsCorrect: isCorrect,
		})
	}
	
	exam, _ := es.repo.FindExamById(params.ExamId)
	conversionTable, _ := es.repo.GetScoreConversionTable(exam.CertificateId)

	finalSkillScores := make(map[int]int)
	totalScore := 0

	for skillId, correctCount := range rawScores {
		scaled := utils.LookupScaledScore(conversionTable, skillId, correctCount)
		finalSkillScores[skillId] = scaled
		totalScore += scaled
	}

	if userLoged, exists := utils.GetUserLogged(ctx); exists {
	 
		if err = es.repo.SaveAttemptWithAnswers(models.UserAttempt{
			UserId: userLoged.UserUUID.String(),
			ExamId: params.ExamId,
			StartTime: time.Now().Format(time.DateTime),
			EndTime: time.Now().Format(time.DateTime),
			TotalScore: totalScore,
			ListeningScore: finalSkillScores[1],
			ReadingScore: finalSkillScores[2],
		}, detailsAnswers); err != nil {
			return v1dto.DetailExamScore{}, err
		}
	}

	return v1dto.DetailExamScore{
		TotalScore: totalScore,
		RawScore: rawScores,
		ScaledScore: finalSkillScores,
	}, nil
}

func (es *examService) GetAllExams(params v1dto.GetAllExamParams) ([]models.ExamModel, int64, error) {
	return es.repo.FindAllExams(params)
}

func (es *examService) CreateExam(params v1dto.CreateExamInputParams) error {
	return es.repo.CreateExam(params)
}

func (es *examService) EditExamById(id int) (models.ExamModel, error) {
	return es.repo.GetExamById(id)
}

func (es *examService) UpdateExam(params v1dto.UpdateExamInputParams) error {
	tx, err := es.db.Begin()
	if err != nil {
		return err
	}

	defer func ()  {
		if err != nil {
			_ = tx.Rollback()
		}
	}()
	
	if params.Target != nil {
		targetExam, err := es.repo.GetTargetExamQuestions(params.Target.TargetExamId, params.Target.TargetPartId)
		if err != nil {
			return err
		}

		err = es.repo.DeleteExamQuestions(tx, params.Id)
		if err != nil {
			return err
		}

		var rows []models.TargetExamMapping

		for _, target := range targetExam {
			rows = append(rows, models.TargetExamMapping {
				ExamId: params.Id,
				EntityType: target.EntityType,
				EntityId: target.EntityId,
				PartId: target.PartId,
				OrderIndex: target.OrderIndex,
			})
		}

		if err = es.repo.InsertExamQuestions(tx, rows); err != nil {
			return err
		}
	}
	
	updateData := goqu.Record{
        "title":          params.Title,
        "year":           params.Year,
        "cert_id":        params.CertificateId,
        "total_question": params.TotalQuestion,
        "total_time":     params.TotalTime,
    }

    if params.Description != nil {
        updateData["description"] = params.Description
    }
    if params.Thumbnail != nil {
        updateData["thumbnail"] = params.Thumbnail
    }
    if params.AudioFullUrl != nil {
        updateData["audio_full_url"] = params.AudioFullUrl
    }
    if params.Status != nil {
        updateData["status"] = params.Status
    }

	if err = es.repo.UpdateExam(tx, params.Id, updateData); err != nil {
		return err
	}

	if err := es.repo.DeleteExamCategories(tx, params.Id); err != nil {
		return err
	}

	// INSERT categories (batch insert)
    if len(params.CategoryIds) > 0 {
        rows := make([]goqu.Record, 0, len(params.CategoryIds))
        for _, cateId := range params.CategoryIds {
            rows = append(rows, goqu.Record{
                "exam_id":     params.Id,
                "category_id": cateId,
            })
        }

        if err := es.repo.InsertExamCategories(tx, rows); err != nil {
            return err
        }
    }
	
	return tx.Commit()
}

func (es *examService) CreatePartDirection(params v1dto.CreatePartDirectionInputParams) error {
	return es.repoPartDirection.CreatePartDirection(params)
}

func (es *examService) UpdatePartDirection(params v1dto.UpdatePartDirectionInputParams) error {
	return es.repoPartDirection.UpdatePartDirection(params)
}

func (es *examService) GetExamStructure(examId int) (v1dto.ExamStructure, error) {
	exam, err := es.repo.FindExamById(examId)
	if err != nil {
		return v1dto.ExamStructure{}, err
	}

	sections, err := es.repo.FindExamQuestionMappingById(examId)
	if err != nil {
		return v1dto.ExamStructure{}, err
	}

	singleIDs 	:= make([]int, 0)
	groupIDs	:= make([]int, 0)

	for _, s := range sections {
		switch s.EntityType {
		case "SINGLE":
			singleIDs = append(singleIDs, s.EntityId)
		case "GROUP":
			groupIDs = append(groupIDs, s.EntityId)
		}
	}

	directionMap 	:= make(map[int]models.Direction)
	directions, err := es.repoPartDirection.FindDirectionByExamId(examId)
	if err == nil {
		for i := range directions {
			d := &directions[i]
			directionMap[d.PartId] = *d
		}
	}

	skillsMater, err := es.repo.FindSkillsByCertId(exam.CertificateId)
	if err != nil {
		return v1dto.ExamStructure{}, err
	}

	partsMaster, err := es.repo.FindPartsByCertId(exam.CertificateId)
	if err != nil {
		return v1dto.ExamStructure{}, err
	}

	skillMasterMap := make(map[int]models.SkillMaster)
	for _, sm := range skillsMater {
		skillMasterMap[sm.Id] = sm
	}

	partMasterMap := make(map[int]models.PartMaster)
    for _, pm := range partsMaster {
        partMasterMap[pm.Id] = pm
    }

	sectionsByPart := make(map[int][]models.ExamQuestionMapping)

	for _, s := range sections {
		sectionsByPart[s.PartId] = append(sectionsByPart[s.PartId], s)
	}

	partIDSet := make(map[int]bool)
	for partID := range sectionsByPart {
		partIDSet[partID] = true
	}
	for partID := range directionMap {
		partIDSet[partID] = true
	}

	partsBySkill := make(map[int][]v1dto.PartDTO)

	for partID := range partIDSet {
		pm, ok := partMasterMap[partID]
		if !ok {
			continue
		}

		examPart := v1dto.PartDTO{
			PartId: partID,
			PartNumber: pm.PartNumber,
			PartName: pm.Name,
		}

		partsBySkill[pm.SkillId] = append(partsBySkill[pm.SkillId], examPart)
	}

	var examSkills []v1dto.SkillDTO

	for skillID, examParts := range partsBySkill {
        sm, ok := skillMasterMap[skillID]
        if !ok { continue }

        // Sắp xếp các Part bên trong Skill theo PartNumber (0, 1, 2, 3...)
        sort.Slice(examParts, func(i, j int) bool {
            return examParts[i].PartNumber < examParts[j].PartNumber
        })

        examSkills = append(examSkills, v1dto.SkillDTO{
            SkillId:   skillID,
            SkillCode: sm.Code,
            SkillName: sm.Name,
            Parts:     examParts,
        })
    }

	sort.Slice(examSkills, func(i, j int) bool {
        return skillMasterMap[examSkills[i].SkillId].OrderIndex < skillMasterMap[examSkills[j].SkillId].OrderIndex
    })

	return v1dto.ExamStructure{
		ExamId: exam.Id,
		ExamName: exam.Title,
		CertCode: exam.CertCode,
		Blueprint: examSkills,
	}, nil
}

func (es *examService) GetExamPart(examId int, partId int) (v1dto.ExamPart, error) {
	var directionMap models.Direction
	
	sections, err := es.repo.FindExamQuestionMappingByPartId(examId, partId)
	if err != nil {
		return v1dto.ExamPart{}, err
	}

	singleIDs 	:= make([]int, 0)
	groupIDs	:= make([]int, 0)

	for _, s := range sections {
		switch s.EntityType {
		case "SINGLE":
			singleIDs = append(singleIDs, s.EntityId)
		case "GROUP":
			groupIDs = append(groupIDs, s.EntityId)
		}
	}

	questionMap 	:= make(map[int]models.Question)
	groupMap 		:= make(map[int]*models.QuestionGroup)
	
	direction, err := es.repoPartDirection.FindDirectionByExamIdAndPartId(examId, partId)
	if err == nil {
		d := &direction
		if len(d.ExampleRaw) > 0 {
			var ex models.ExampleData
			if err := json.Unmarshal(d.ExampleRaw, &ex); err == nil {
				d.Example = &ex
			}
		}
		directionMap = *d
	}

	if len(singleIDs) > 0 {

		questions, err := es.repo.FindQuesionByIds(singleIDs)
		
		if err != nil {
			return v1dto.ExamPart{}, err
		}

		for i := range questions {
			
			q := &questions[i]

			opts := map[string]*string{
				"A": q.OptionA,
				"B": q.OptionB,
				"C": q.OptionC,
				"D": q.OptionD,
			}
			q.Options = opts
			questionMap[q.Id] = *q
		}
	}

	if len(groupIDs) > 0 {

		groups, err := es.repo.FindGroupQuestionByIds(groupIDs)
		
		if err != nil {
			return v1dto.ExamPart{}, err
		}

		for _, g := range groups {
			gCopy := g
			groupMap[g.Id] = &gCopy
		}

		subQuestions, err := es.repo.FindSubQuesionByGroupIds(groupIDs)

		if err != nil { 
			return v1dto.ExamPart{}, err 
		}

		for i := range subQuestions {
			q := &subQuestions[i]

			q.Options = map[string]*string{
				"A": q.OptionA,
				"B": q.OptionB,
				"C": q.OptionC,
				"D": q.OptionD,
			}

			g := groupMap[*q.GroupId]
			g.SubQuestions = append(g.SubQuestions, *q)
			groupMap[*q.GroupId] = g
		}
	}

	for i := range sections {
		s := &sections[i]		

		switch s.EntityType {
		case "SINGLE":
			if q, ok := questionMap[s.EntityId]; ok {
				q.DisplayNumber = s.OrderIndex
				qCopy := q
				s.QuestionData = &qCopy
			}

		case "GROUP":
			if g, ok := groupMap[s.EntityId]; ok {
				for j := range g.SubQuestions {
					g.SubQuestions[j].DisplayNumber = s.OrderIndex + j
				}

				s.GroupData = g
			}
		}
	}

	return v1dto.ExamPart{
		ExamId: examId,
		PartId: partId,
		Direction: &directionMap,
		Items: sections,
	}, nil
}

func (es *examService) UpdateQuestionSingle(params v1dto.UpdateQuestionSingleInputParams) error {
	return es.repo.UpdateQuestionSingle(params)
}

func (es *examService) UpdateQuestionGroup(params v1dto.UpdateQuestionGroupInputParams) error {
	return es.repo.UpdateQuestionGroup(params)
}

func (es *examService) ImportExamQuestionFromExcel(ctx *gin.Context, params v1dto.ExamImportInputParams) error {
	exePath, _ := os.Getwd()
	assetsPath := filepath.Join(exePath, "assets")

	os.MkdirAll(assetsPath, 0755)

	fileName := fmt.Sprintf("%s_%s.xlsx", params.File.Filename, time.Now().Format("20060102150405"))
	dst := filepath.Join(assetsPath, fileName)

	if err := ctx.SaveUploadedFile(&params.File, dst); err != nil {
		return err
	}

	f, err := excelize.OpenFile(dst)
	if err != nil {
		return err
	}
	defer f.Close()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return err
	}

	getSafeCell := func(r []string, index int) string {
		if index < len(r) { 
			return r[index] 
		}

		return ""
    }

	groupTracker := make(map[string]*v1dto.QuestionGroupParamsInput)
	
	for i, row := range rows {
		if i == 0 {
			continue
		}
		
		orderNoStr := getSafeCell(row, 0)
		partIdStr := getSafeCell(row, 1)
		PartStr := getSafeCell(row, 2)
		tempGroupId := getSafeCell(row, 3)
		imageUrl := getSafeCell(row, 4)
		passageText := getSafeCell(row, 5)
		questionText := getSafeCell(row, 6)
		optionA := getSafeCell(row, 7)
		optionB := getSafeCell(row, 8)
		optionC := getSafeCell(row, 9)
		optionD := getSafeCell(row, 10)
		correctAnswer := getSafeCell(row, 11)
		audioStartMsStr := getSafeCell(row, 12)
		audioEndMsStr := getSafeCell(row, 13)
		explanation := getSafeCell(row, 14)
		transcript := getSafeCell(row, 15)

		partId, _ := strconv.Atoi(partIdStr)
		part, _ := strconv.Atoi(PartStr)
		orderNo, _ := strconv.Atoi(orderNoStr)
		audioStartMs, _ := strconv.Atoi(audioStartMsStr)
		audioEndMs, _ := strconv.Atoi(audioEndMsStr)

		if tempGroupId == "" {
			questionSingle := v1dto.QuestionParamsInput{
				ExamId: int64(params.ExamId),
				EntityType: "SINGLE",
				PartId:	partId,
				Part: part,
				QuestionText: &questionText,
				OptionA: &optionA,
				OptionB: &optionB,
				OptionC: &optionC,
				OptionD: &optionD,
				CorrectAnswer: correctAnswer,
				Explanation: &explanation,
				Transcript: &transcript,
				AudioStartMs: &audioStartMs,
				AudioEndMs: &audioEndMs,
				ImageUrl: &imageUrl,
				SubOrder: orderNo,
			}
			
			questionId, err := es.repoQuestion.CreateQuestion(questionSingle)
			if err != nil {
				return err
			}
			questionMapping := v1dto.ExamQuestionMappingInput{
				ExamId: int64(params.ExamId),
				EntityType: "SINGLE",
				EntityId: questionId,
				OrderIndex: orderNo,
				PartId: partId,
			}

			err = es.repoQuestion.CreateQuestionMapping(questionMapping)
			if err != nil {
				return err
			}
		} else {
			if _, exists := groupTracker[tempGroupId]; !exists {
				newGroup := &v1dto.QuestionGroupParamsInput{
					ExamId: int64(params.ExamId),
					PartId: partId,
					EntityType: "GROUP",
					PassageText: &passageText,
					ImageUrl: &imageUrl,
					AudioStartMs: &audioStartMs,
					AudioEndMs: &audioEndMs,
					Explanation: &explanation,
					Transcript: &transcript,
					SubQuestions: []v1dto.QuestionParams{},
				}
				groupTracker[tempGroupId] = newGroup
			}

			subQuestionRaw := v1dto.QuestionParams{
				Part: part,
				QuestionText: &questionText,
				OptionA: &optionA,
				OptionB: &optionB,
				OptionC: &optionC,
				OptionD: &optionD,
				CorrectAnswer: correctAnswer,
				SubOrder: orderNo,
			}

			groupTracker[tempGroupId].SubQuestions = append(groupTracker[tempGroupId].SubQuestions, subQuestionRaw)
		}
	}

	if len(groupTracker) > 0 {
		for _, group := range groupTracker {
			groupId, err := es.repoQuestion.CreateQuestionGroup(*group)
			if err != nil {
				return err
			}

			var paramsQuestion []v1dto.QuestionParamsInput
			for _, subQuestion := range group.SubQuestions {
				paramsQuestion = append(paramsQuestion, v1dto.QuestionParamsInput{
					GroupId: &groupId,
					Part: subQuestion.Part,
					QuestionText: subQuestion.QuestionText,
					ImageUrl: subQuestion.ImageUrl,
					CorrectAnswer: subQuestion.CorrectAnswer,
					OptionA: subQuestion.OptionA,
					OptionB: subQuestion.OptionB,
					OptionC: subQuestion.OptionC,
					OptionD: subQuestion.OptionD,
					AudioStartMs: subQuestion.AudioStartMs,
					AudioEndMs: subQuestion.AudioEndMs,
					SubOrder: subQuestion.SubOrder,
					Explanation: subQuestion.Explanation,
					Transcript: subQuestion.Transcript,
					Tags: subQuestion.Tags,
				})
			}

			err = es.repoQuestion.CreateQuestions(paramsQuestion)
			if err != nil {
				return err
			}

			paramsQuestionGroupMapping := v1dto.ExamQuestionMappingInput{
				ExamId: int64(params.ExamId),
				EntityType: "GROUP",
				EntityId: groupId,
				OrderIndex: group.SubQuestions[0].SubOrder,
				PartId: group.PartId,
			}

			err = es.repoQuestion.CreateQuestionGroupMapping(paramsQuestionGroupMapping)
			if err != nil {
				return err
			}
		}
	}

	os.Remove(dst)
	
	return  nil
}

func (es *examService) GetFilterStructure() ([]*v1dto.FilterStructure, error) {

	filterStructure, err := es.repo.FindFilterStructure()
	if err != nil {
		return nil, err
	}
	
	return utils.BuildTree(filterStructure), nil
}

func (es *examService) FilterExam(params v1dto.FilterExamParams) ([]v1dto.ExamFilterDTO, int64, error) {
	return es.repo.FindExamsByFilter(params)
}

func (es *examService) GetFeaturedExams(params v1dto.ExamFeaturedParams) (v1dto.FeaturedDTO, int64, error) {
	raw, total, err := es.repo.FindFeaturedExams(params)
	if err != nil {
		return v1dto.FeaturedDTO{}, 0, err
	}
	
	resultMap := make(map[string]*v1dto.FeaturedDTO)

	for _, r := range raw {
		if _, ok := resultMap[r.Type]; !ok {
			resultMap[r.Type]  = &v1dto.FeaturedDTO{
				Type: r.Type,
				Name: r.Name,
				CateDescription: r.CateDescription,
				Exams: []v1dto.ExamFeaturedDTO{},
			}
		}

		resultMap[r.Type].Exams = append(resultMap[r.Type].Exams, v1dto.ExamFeaturedDTO{
			Id: r.Id,
			Title: r.Title,
			Year: r.Year,
			TotalTime: r.TotalTime,
			TotalQuestion: r.TotalQuestion,
			Thumbnail: r.Thumbnail,
		})
	}

	var result v1dto.FeaturedDTO

	for _, r := range resultMap {
		result = *r
		break
	}

	return result, total, nil
}