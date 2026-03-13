package v1service

import (
	"encoding/json"
	"sort"
	"time"

	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
	repository "github.com/dangLuan01/ets-api/internal/repository/exam"
	"github.com/dangLuan01/ets-api/internal/utils"
	"github.com/doug-martin/goqu/v9"
)

type examService struct {
	repo repository.ExamRepository
}

func NewExamService(repo repository.ExamRepository) ExamService {
	return &examService{
		repo: repo,
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
	directions, err := rs.repo.FindDirectionByExamId(examId)
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