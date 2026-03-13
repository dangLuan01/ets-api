package v1dto

type GetAllExamParams struct {
	Page int32 `form:"page" binding:"omitempty,min=1"`
	Limit int32 `form:"limit" binding:"omitempty,min=1,max=50"`
}

type QuestionAnswerInputParams struct {
	ExamId 		int 				`json:"exam_id"`
	Answers 	[]UserAnswerInput 	`json:"answers"`
}

type UserAnswerInput struct {
	QuestionId 		int 			`json:"question_id"`
	SelectedAnswer 	string 			`json:"selected_answer"`
}

type QuestionWithSkill struct {
	QuestionId 		int 			`json:"question_id" db:"question_id"`
	SkillId 		int				`json:"skill_id" db:"skill_id"`
	CorrectAnswer 	string 			`json:"correct_answer" db:"correct_answer"`
}

type DetailExamScore struct {
	TotalScore 		int 			`json:"total_score"`
	RawScore 		map[int]int 	`json:"raw_score"`
	ScaledScore 	map[int]int 	`json:"scaled_score"`
}

type DetailExamScoreDTO struct {
	TotalScore 		int 			`json:"total_score"`
	RawScore 		map[string]int 	`json:"raw_score"`
	ScaledScore 	map[string]int 	`json:"scaled_score"`
}

type CreateExamInputParams struct {
	CertificateId	int				`json:"cert_id" db:"cert_id" binding:"required"`
	Title 			string 			`json:"title" db:"title" binding:"required"`
	Year 			int 			`json:"year" db:"year" binding:"required"`
	TotalQuestion 	int 			`json:"total_question" db:"total_question" binding:"required"`
	TotalTime 		int 			`json:"total_time" db:"total_time" binding:"required"`
	Description 	*string 		`json:"description" db:"description" binding:"omitempty"`
	Thumbnail 		*string 		`json:"thumbnail" db:"thumbnail" binding:"omitempty"`
	Category 		*string 		`json:"category" db:"category" binding:"omitempty"`
	AudioFullUrl 	*string			`json:"audio_full_url" db:"audio_full_url" binding:"omitempty"`
}

type UpdateExamInputParams struct {
	Id				int				`json:"id" db:"id" binding:"required"`
	CertificateId	int				`json:"cert_id" db:"cert_id" binding:"required"`
	Title 			string 			`json:"title" db:"title" binding:"required"`
	Year 			int 			`json:"year" db:"year" binding:"required"`
	TotalQuestion 	int 			`json:"total_question" db:"total_question" binding:"required"`
	TotalTime 		int 			`json:"total_time" db:"total_time" binding:"required"`
	Description 	*string 		`json:"description" db:"description"`
	Thumbnail 		*string 		`json:"thumbnail" db:"thumbnail"`
	Category 		*string 		`json:"category" db:"category"`
	AudioFullUrl 	*string			`json:"audio_full_url" db:"audio_full_url"`
	Status 			*int 			`json:"status" db:"status" binding:"required,oneof=0 1"`
}

func MapDetailExamScoreDTO(params DetailExamScore) *DetailExamScoreDTO {
	rawScoreMap := make(map[string]int)
    scaledScoreMap := make(map[string]int)

	for skillId, countScore := range params.RawScore {
		skillName := formatSkill(skillId)
		rawScoreMap[skillName] = countScore
	}

	for skillId, score := range params.ScaledScore {
		skillName := formatSkill(skillId)
		scaledScoreMap[skillName] = score
	}

	return &DetailExamScoreDTO{
		TotalScore: params.TotalScore,
		RawScore: rawScoreMap,
		ScaledScore: scaledScoreMap,
	}
}

func formatSkill(skillId int) string {
	switch skillId {
	case 1:
		return "listening"
	default :
		return "reading"
	}
}