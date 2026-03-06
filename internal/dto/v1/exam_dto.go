package v1dto

type QuestionAnswerInputParams struct {
	ExamId 		int 				`json:"exam_id"`
	Answers 	[]UserAnswerInput 	`json:"answers"`
}

type UserAnswerInput struct {
	QuestionId 		int 	`json:"question_id"`
	SelectedAnswer 	string 	`json:"selected_answer"`
}

type QuestionWithSkill struct {
	QuestionId 		int 	`json:"question_id" db:"question_id"`
	SkillId 		int		`json:"skill_id" db:"skill_id"`
	CorrectAnswer 	string 	`json:"correct_answer" db:"correct_answer"`
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
		return "Listening"
	default :
		return "Reading"
	}
}