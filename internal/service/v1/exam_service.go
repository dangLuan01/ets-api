package v1service

import (
	"encoding/json"
	"sort"

	"github.com/dangLuan01/ets-api/internal/models"
	"github.com/dangLuan01/ets-api/internal/repository"
)

type examService struct {
	repo repository.ExamRepository
}

func NewExamService(repo repository.ExamRepository) ExamService {
	return &examService{
		repo: repo,
	}
}

func (rs *examService) FindExamById(examId string) (models.Exam, error) {

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

	directionMap 	:= make(map[int]models.Direction)
	questionMap 	:= make(map[int]models.Question)
	groupMap 		:= make(map[int]*models.QuestionGroup)

	directions, err := rs.repo.FindDirectionByExamId(examId)
	if err == nil {
		for i := range directions {
			d := &directions[i]
			if len(d.ExmapleRaw) > 0 {
				var ex models.ExampleData
				if err := json.Unmarshal(d.ExmapleRaw, &ex); err == nil {
					d.Exmaple = &ex
				}
			}
			directionMap[d.Part] = *d
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
				for i := range g.SubQuestions {
					g.SubQuestions[i].DisplayNumber = s.OrderIndex + i
				}

				s.GroupData = g
			}
		}
	}

	sectionsByPart := make(map[int][]models.ExamQuestionMapping)

	for _, s := range sections {
		sectionsByPart[s.Part] = append(sectionsByPart[s.Part], s)
	}

	parts := make([]int, 0, len(sectionsByPart))
	for part := range sectionsByPart {
		parts = append(parts, part)
	}

	if _, ok := directionMap[0]; ok {
		parts = append(parts, 0)
	}
	sort.Ints(parts)
	
	var examParts []models.ExamPart
	for _, part := range parts {
		examParts = append(examParts, models.ExamPart{
			Part: part,
			Direction: directionMap[part],
			Items: sectionsByPart[part],
		})
	}
	
	exam.Sections = examParts

	return exam, nil
}