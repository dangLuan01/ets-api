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
	repo repositoryExam.ExamRepository
	repoPartDirection repositoryPartDirection.PartDirectionRepository
	repoQuestion repositoryQuestion.QuestionRepository
}

func NewExamService(repo repositoryExam.ExamRepository, repoPartDirection repositoryPartDirection.PartDirectionRepository, repoQuestion repositoryQuestion.QuestionRepository) ExamService {
	return &examService{
		repo: repo,
		repoPartDirection: repoPartDirection,
		repoQuestion: repoQuestion,
	}
}

func (rs *examService) FindExamById(examId int) (models.Exam, error) {

	exam, err := rs.repo.FindExamById(examId)
	if err != nil {
		return models.Exam{}, err
	}

	sections, err := rs.repo.FindExamQuestionMappingById(examId)
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
	directions, err := rs.repoPartDirection.FindDirectionByExamId(examId)
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

		questions, err := rs.repo.FindQuesionByIds(singleIDs)
		
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

		groups, err := rs.repo.FindGroupQuestionByIds(groupIDs)
		
		if err != nil {
			return models.Exam{}, err
		}

		for _, g := range groups {
			gCopy := g
			groupMap[g.Id] = &gCopy
		}

		subQuestions, err := rs.repo.FindSubQuesionByGroupIds(groupIDs)

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

	skillsMater, err := rs.repo.FindSkillsByCertId(exam.CertificateId)
	if err != nil {
		return models.Exam{}, err
	}

	partsMaster, err := rs.repo.FindPartsByCertId(exam.CertificateId)
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

func (rs *examService) CalculateScoreExam(params v1dto.QuestionAnswerInputParams) (v1dto.DetailExamScore, error) {

	questionIds 	:= make([]int, 0, len(params.Answers))
	userAnswerMap 	:= make(map[int]string)

	for _, ans := range params.Answers {
		questionIds = append(questionIds, ans.QuestionId)
		userAnswerMap[ans.QuestionId] = ans.SelectedAnswer
	}

	correctAnswer, err := rs.repo.GetCorrectAnswersWithSkillContext(params.ExamId, questionIds)
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
	
	exam, _ := rs.repo.FindExamById(params.ExamId)
	conversionTable, _ := rs.repo.GetScoreConversionTable(exam.CertificateId)

	finalSkillScores := make(map[int]int)
	totalScore := 0

	for skillId, correctCount := range rawScores {
		scaled := utils.LookupScaledScore(conversionTable, skillId, correctCount)
		finalSkillScores[skillId] = scaled
		totalScore += scaled
	}

	err = rs.repo.SaveAttemptWithAnswers(models.UserAttempt{
		UserId: 1,
		ExamId: params.ExamId,
		StartTime: time.Now().Format(time.RFC3339),
		EndTime: time.Now().Format(time.RFC3339),
		TotalScore: totalScore,
		ListeningScore: finalSkillScores[1],
		ReadingScore: finalSkillScores[2],
	}, detailsAnswers)
	if err != nil {
		return v1dto.DetailExamScore{}, err
	}

	return v1dto.DetailExamScore{
		TotalScore: totalScore,
		RawScore: rawScores,
		ScaledScore: finalSkillScores,
	}, nil
}

func (rs *examService) GetAllExams(params v1dto.GetAllExamParams) ([]models.ExamModel, int64, error) {
	return rs.repo.FindAllExams(params)
}

func (rs *examService) CreateExam(params v1dto.CreateExamInputParams) error {
	return rs.repo.CreateExam(params)
}

func (rs *examService) EditExamById(id int) (models.ExamModel, error) {
	return rs.repo.GetExamById(id)
}

func (rs *examService) UpdateExam(params v1dto.UpdateExamInputParams) error {
	updateData := goqu.Record{}
	
	if params.Description != nil {
		updateData["description"] = params.Description
	}
	if params.Thumbnail != nil {
		updateData["thumbnail"] = params.Thumbnail
	}
	if params.Category != nil {
		updateData["category"] = params.Category
	}
	if params.AudioFullUrl != nil {
		updateData["audio_full_url"] = params.AudioFullUrl
	}
	if params.Status != nil {
		updateData["status"] = params.Status
	}
	
	updateData["title"] = params.Title
	updateData["year"] = params.Year
	updateData["cert_id"] = params.CertificateId
	updateData["total_question"] = params.TotalQuestion
	updateData["total_time"] = params.TotalTime

	return rs.repo.UpdateExam(params.Id, updateData)
}

func (rs *examService) CreatePartDirection(params v1dto.CreatePartDirectionInputParams) error {
	return rs.repoPartDirection.CreatePartDirection(params)
}

func (rs *examService) UpdatePartDirection(params v1dto.UpdatePartDirectionInputParams) error {
	return rs.repoPartDirection.UpdatePartDirection(params)
}

func (rs *examService) GetExamStructure(examId int) (v1dto.ExamStructure, error) {
	exam, err := rs.repo.FindExamById(examId)
	if err != nil {
		return v1dto.ExamStructure{}, err
	}

	sections, err := rs.repo.FindExamQuestionMappingById(examId)
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
	directions, err := rs.repoPartDirection.FindDirectionByExamId(examId)
	if err == nil {
		for i := range directions {
			d := &directions[i]
			directionMap[d.PartId] = *d
		}
	}

	skillsMater, err := rs.repo.FindSkillsByCertId(exam.CertificateId)
	if err != nil {
		return v1dto.ExamStructure{}, err
	}

	partsMaster, err := rs.repo.FindPartsByCertId(exam.CertificateId)
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

func (rs *examService) GetExamPart(examId int, partId int) (v1dto.ExamPart, error) {
	var directionMap models.Direction
	
	sections, err := rs.repo.FindExamQuestionMappingByPartId(examId, partId)
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
	
	direction, err := rs.repoPartDirection.FindDirectionByExamIdAndPartId(examId, partId)
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

		questions, err := rs.repo.FindQuesionByIds(singleIDs)
		
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

		groups, err := rs.repo.FindGroupQuestionByIds(groupIDs)
		
		if err != nil {
			return v1dto.ExamPart{}, err
		}

		for _, g := range groups {
			gCopy := g
			groupMap[g.Id] = &gCopy
		}

		subQuestions, err := rs.repo.FindSubQuesionByGroupIds(groupIDs)

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

func (rs *examService) UpdateQuestionSingle(params v1dto.UpdateQuestionSingleInputParams) error {
	return rs.repo.UpdateQuestionSingle(params)
}

func (rs *examService) UpdateQuestionGroup(params v1dto.UpdateQuestionGroupInputParams) error {
	return rs.repo.UpdateQuestionGroup(params)
}

func (rs *examService) ImportExamQuestionFromExcel(ctx *gin.Context, params v1dto.ExamImportInputParams) error {
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
			
			questionId, err := rs.repoQuestion.CreateQuestion(questionSingle)
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

			err = rs.repoQuestion.CreateQuestionMapping(questionMapping)
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
			groupId, err := rs.repoQuestion.CreateQuestionGroup(*group)
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

			err = rs.repoQuestion.CreateQuestions(paramsQuestion)
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

			err = rs.repoQuestion.CreateQuestionGroupMapping(paramsQuestionGroupMapping)
			if err != nil {
				return err
			}
		}
	}

	os.Remove(dst)
	
	return  nil
}