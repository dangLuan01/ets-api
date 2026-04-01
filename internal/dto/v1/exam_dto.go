package v1dto

import (
	"mime/multipart"

	"github.com/dangLuan01/ets-api/internal/models"
)

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
	CategoryIds 		[]int 		`json:"category_ids" db:"-" binding:"required"`
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
	CategoryIds 	[]int 			`json:"category_ids" db:"-"`
	AudioFullUrl 	*string			`json:"audio_full_url" db:"audio_full_url"`
	Status 			*int 			`json:"status" db:"status" binding:"required,oneof=0 1"`
}

type ExamStructure struct {
	ExamId 			int					`json:"exam_id"`
	ExamName 		string 				`json:"exam_name"`
	CertCode   		string 				`json:"cert_code"`
	Blueprint 		[]SkillDTO 			`json:"blueprint"`
}

type SkillDTO struct {
	SkillId 		int 				`json:"skill_id"`
	SkillCode 		string 				`json:"skill_code"`
	SkillName 		string 				`json:"skill_name"`
	Parts 			[]PartDTO 			`json:"parts"`
}

type PartDTO struct {
	PartId 			int 				`json:"part_id"`
	PartName 		string 				`json:"part_name"`
	PartNumber 		int 				`json:"part_number"`
}

type ExamPart struct {
	ExamId 		int 						`json:"exam_id"`
	PartId 		int							`json:"part_id"`
	Direction   *models.Direction			`json:"direction,omitempty"`
	Items		[]ExamQuestionMappingDTO 	`json:"items"`
}

type ExamQuestionMappingDTO struct {
	EntityType 		string 					`json:"entity_type" db:"entity_type"`
	EntityId 		int 					`json:"entity_id" db:"entity_id"`
	OrderIndex 		int						`json:"order_index" db:"order_index"`
	QuestionData 	*models.Question 		`json:"question_data,omitempty"`
	GroupData 		*models.QuestionGroup	`json:"group_data,omitempty"`
}

type ExamImportInputParams struct {
	ExamId 		int 						`form:"exam_id" binding:"required"`
	File 		multipart.FileHeader  		`form:"file" binding:"required,file_ext=xlsx,maxfile=100"`
}

type UpdateQuestionSingleInputParams struct {
	// ExamId 			int 					`json:"exam_id" binding:"required"`
	QuestionId 		int 					`json:"question_id" binding:"required"`
	QuestionText 	*string 				`json:"question_text,omitempty"`
	ImageUrl 		*string 				`json:"image_url,omitempty"`
	CorrectAnswer 	string 					`json:"correct_answer" binding:"required"`
	OptionA 		*string 				`json:"option_a,omitempty"`
	OptionB 		*string 				`json:"option_b,omitempty"`
	OptionC 		*string 				`json:"option_c,omitempty"`
	OptionD 		*string 				`json:"option_d,omitempty"`
	SubOrder 		int 					`json:"sub_order" binding:"required"`
	AudioStartMs 	*int 					`json:"audio_start_ms,omitempty"`
	AudioEndMs 		*int 					`json:"audio_end_ms,omitempty"`
	Explanation 	*string 				`json:"explanation,omitempty"`
	Transcript 		*string 				`json:"transcript,omitempty"`
	Tags 			*string 				`json:"tags,omitempty"`
}

type UpdateQuestionGroupInputParams struct {
	// ExamId 		int 						`json:"exam_id" binding:"required"`
	// PartId 		int 						`json:"part_id" binding:"required"`
	GroupId 		int 					`json:"group_id" binding:"required"`
	PassageText 	*string 				`json:"passage_text,omitempty"`
	ImageUrl 		*string 				`json:"image_url,omitempty"`
	AudioStartMs 	*int 					`json:"audio_start_ms,omitempty"`
	AudioEndMs 		*int 					`json:"audio_end_ms,omitempty"`
	Explanation 	*string 				`json:"explanation,omitempty"`
	Transcript 		*string 				`json:"transcript,omitempty"`
	SubQuestions 	[]struct {
		QuestionId 		int 				`json:"question_id" binding:"required"`
		QuestionText 	*string 			`json:"question_text,omitempty"`
		CorrectAnswer 	string 				`json:"correct_answer" binding:"required"`
		OptionA 		*string 			`json:"option_a,omitempty"`
		OptionB 		*string 			`json:"option_b,omitempty"`
		OptionC 		*string 			`json:"option_c,omitempty"`
		OptionD 		*string 			`json:"option_d,omitempty"`
		SubOrder 		int 				`json:"sub_order" binding:"required"`
		Explanation 	*string 			`json:"explanation,omitempty"`
	} `json:"sub_questions" binding:"required"`
}

type ExamDTO struct {
	Id 				int 					`json:"id" db:"id"`
	Title 			string 					`json:"title" db:"title"`
	Year 			int 					`json:"year" db:"year"`
	TotalTime 		int 					`json:"total_time" db:"total_time"`
	TotalQuestion	int						`json:"total_question" db:"total_question"`
	Thumbnail 		*string 				`json:"thumbnail" db:"thumbnail"`
	UpdatedAt 		string 					`json:"-" db:"updated_at"`
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