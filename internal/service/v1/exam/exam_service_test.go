package v1service

import (
	"testing"

	v1dto "github.com/dangLuan01/ets-api/internal/dto/v1"
	"github.com/dangLuan01/ets-api/internal/models"
)
type mockExamRepo struct {
    saveCalled bool
}

func (m *mockExamRepo) FindExamById(examId int) (models.Exam, error) {
    return models.Exam{
        Id:            examId,
        CertificateId: 1,
    }, nil
}

func (m *mockExamRepo) GetCorrectAnswersWithSkillContext(
    examId int,
    ids []int,
) ([]v1dto.QuestionWithSkill, error) {
    return []v1dto.QuestionWithSkill{
        {
            QuestionId:    1,
            SkillId:       1,
            CorrectAnswer: "A",
        },
        {
            QuestionId:    2,
            SkillId:       1,
            CorrectAnswer: "B",
        },
        {
            QuestionId:    3,
            SkillId:       2,
            CorrectAnswer: "C",
        },
    }, nil
}

func (m *mockExamRepo) GetScoreConversionTable(certId int) ([]models.ScoreConversion, error) {
    return []models.ScoreConversion{
        {SkillId: 1, RawScore: 2, ScaledScore: 200},
        {SkillId: 2, RawScore: 0, ScaledScore: 100},
    }, nil
}

func (m *mockExamRepo) SaveAttemptWithAnswers(
    attempt models.UserAttempt,
    answers []models.UserAnswer,
) error {
    m.saveCalled = true
    return nil
}

/* Các method không dùng → stub rỗng */

func (m *mockExamRepo) FindExamQuestionMappingById(int) ([]models.ExamQuestionMapping, error) {
    return nil, nil
}
func (m *mockExamRepo) FindQuesionByIds([]int) ([]models.Question, error) {
    return nil, nil
}
func (m *mockExamRepo) FindGroupQuestionByIds([]int) ([]models.QuestionGroup, error) {
    return nil, nil
}
func (m *mockExamRepo) FindSubQuesionByGroupIds([]int) ([]models.Question, error) {
    return nil, nil
}
func (m *mockExamRepo) FindDirectionByExamId(int) ([]models.Direction, error) {
    return nil, nil
}
func (m *mockExamRepo) FindSkillsByCertId(int) ([]models.SkillMaster, error) {
    return nil, nil
}
func (m *mockExamRepo) FindPartsByCertId(int) ([]models.PartMaster, error) {
    return nil, nil
}

func (m *mockExamRepo) FindAllExams(params v1dto.GetAllExamParams) ([]models.ExamModel, int64,  error) {
    return []models.ExamModel{}, 0, nil
}

func (m *mockExamRepo) CreateExam(params v1dto.CreateExamInputParams) error {
    return nil
}

func (m *mockExamRepo) GetExamById(examId int) (models.ExamModel, error) {
    return models.ExamModel{}, nil
}

func (m *mockExamRepo) UpdateExam(examId int, params v1dto.UpdateExamInputParams) error {
    return nil
}

func (m *mockExamRepo) CreatePartDirection(params v1dto.CreatePartDirectionInputParams) error {
    return nil
}

func (m *mockExamRepo) FindExamQuestionMappingByPartId(examId int, partId int) ([]v1dto.ExamQuestionMappingDTO, error) {
    return nil, nil
}

func (m *mockExamRepo) UpdateQuestionSingle(params v1dto.UpdateQuestionSingleInputParams) error {
    return nil
}

func (m *mockExamRepo) UpdateQuestionGroup(params v1dto.UpdateQuestionGroupInputParams) error {
   return nil
}

func (m *mockExamRepo) FindFilterStructure() ([]*v1dto.FilterStructure, error) {
    return nil, nil
}

func (m *mockExamRepo) FindExamsByFilter(params v1dto.FilterExamParams) ([]v1dto.ExamDTO, int64, error) {
    return nil, 0, nil
}

func TestCalculateScoreExam_Success(t *testing.T) {
    repo := &mockExamRepo{}
    service := &examService{repo: repo}

    params := v1dto.QuestionAnswerInputParams{
        ExamId: 1,
        Answers: []v1dto.UserAnswerInput{
            {QuestionId: 1, SelectedAnswer: "A"}, // đúng
            {QuestionId: 2, SelectedAnswer: "B"}, // đúng
            {QuestionId: 3, SelectedAnswer: "D"}, // sai
        },
    }

    result, err := service.CalculateScoreExam(params)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    // Raw score
    if result.RawScore[1] != 2 {
        t.Errorf("expected raw score skill 1 = 2, got %d", result.RawScore[1])
    }
    if result.RawScore[2] != 0 {
        t.Errorf("expected raw score skill 2 = 0, got %d", result.RawScore[2])
    }

    // Scaled score
    if result.ScaledScore[1] != 200 {
        t.Errorf("expected scaled score skill 1 = 200, got %d", result.ScaledScore[1])
    }
    if result.ScaledScore[2] != 100 {
        t.Errorf("expected scaled score skill 2 = 100, got %d", result.ScaledScore[2])
    }

    // Total score
    if result.TotalScore != 300 {
        t.Errorf("expected total score = 300, got %d", result.TotalScore)
    }

    // Save attempt
    if !repo.saveCalled {
        t.Error("expected SaveAttemptWithAnswers to be called")
    }
}
