package v1dto

type QuestionParamsInput struct {
	ExamId 			int64 					`json:"exam_id" db:"-" binding:"required"`
	EntityType 		string 					`json:"entity_type" db:"-" binding:"required,oneof=SINGLE GROUP"`
	PartId 			int 					`json:"part_id" db:"-" binding:"required"`

	GroupId 		*int64 					`json:"group_id" db:"group_id" binding:"omitempty"`
	Part 			int 					`json:"part" db:"part" binding:"required"`
	QuestionText 	*string 				`json:"question_text" db:"question_text" binding:"omitempty"`
	ImageUrl 		*string 				`json:"image_url" db:"image_url" binding:"omitempty,max=255"`
	CorrectAnswer 	string 					`json:"correct_answer" db:"correct_answer" binding:"required"`
	OptionA 		*string 				`json:"option_a" db:"option_a" binding:"omitempty,max=255"`
	OptionB 		*string 				`json:"option_b" db:"option_b" binding:"omitempty,max=255"`
	OptionC 		*string 				`json:"option_c" db:"option_c" binding:"omitempty,max=255"`
	OptionD 		*string 				`json:"option_d" db:"option_d" binding:"omitempty,max=255"`
	AudioStartMs 	*int 					`json:"audio_start_ms" db:"audio_start_ms" binding:"omitempty"`
	AudioEndMs 		*int 					`json:"audio_end_ms" db:"audio_end_ms" binding:"omitempty"`
	SubOrder 		int 					`json:"sub_order" db:"sub_order" binding:"required"`
	Explanation 	*string 				`json:"explanation" db:"explanation" binding:"omitempty"`
	Transcript 		*string 				`json:"transcript" db:"transcript" binding:"omitempty"`
	Tags 			*string 				`json:"tags" db:"tags" binding:"omitempty"`
}

type QuestionParams struct {
	GroupId 		*int64 					`json:"group_id" db:"group_id" binding:"omitempty"`
	Part 			int 					`json:"part" db:"part" binding:"required"`
	QuestionText 	*string 				`json:"question_text" db:"question_text" binding:"omitempty"`
	ImageUrl 		*string 				`json:"image_url" db:"image_url" binding:"omitempty,max=255"`
	CorrectAnswer 	string 					`json:"correct_answer" db:"correct_answer" binding:"required"`
	OptionA 		*string 				`json:"option_a" db:"option_a" binding:"omitempty,max=255"`
	OptionB 		*string 				`json:"option_b" db:"option_b" binding:"omitempty,max=255"`
	OptionC 		*string 				`json:"option_c" db:"option_c" binding:"omitempty,max=255"`
	OptionD 		*string 				`json:"option_d" db:"option_d" binding:"omitempty,max=255"`
	AudioStartMs 	*int 					`json:"audio_start_ms" db:"audio_start_ms" binding:"omitempty"`
	AudioEndMs 		*int 					`json:"audio_end_ms" db:"audio_end_ms" binding:"omitempty"`
	SubOrder 		int 					`json:"sub_order" db:"sub_order" binding:"required"`
	Explanation 	*string 				`json:"explanation" db:"explanation" binding:"omitempty"`
	Transcript 		*string 				`json:"transcript" db:"transcript" binding:"omitempty"`
	Tags 			*string 				`json:"tags" db:"tags" binding:"omitempty"`
}


type ExamQuestionMappingInput struct {
	ExamId 			int64					`json:"exam_id" db:"exam_id" binding:"required"`
	EntityType 		string 					`json:"entity_type" db:"entity_type" binding:"required,oneof=SINGLE GROUP"`
	EntityId 		int64					`json:"entity_id" db:"entity_id" binding:"required"`
	OrderIndex 		int						`json:"order_index" db:"order_index" binding:"required"`
	PartId 			int						`json:"part_id" db:"part_id" binding:"required"`
}

type QuestionGroupParamsInput struct {
	ExamId 			int64 					`json:"exam_id" db:"-" binding:"required"`
	EntityType 		string 					`json:"entity_type" db:"-" binding:"required,oneof=SINGLE GROUP"`

	PartId 			int 					`json:"part_id" db:"part_id" binding:"required"`
	PassageText 	*string 				`json:"passage_text" db:"passage_text" binding:"omitempty"`
	ImageUrl 		*string 				`json:"image_url" db:"image_url" binding:"omitempty,max=255"`
	AudioStartMs 	*int 					`json:"audio_start_ms" db:"audio_start_ms" binding:"omitempty"`
	AudioEndMs 		*int 					`json:"audio_end_ms" db:"audio_end_ms" binding:"omitempty"`
	Explanation 	*string 				`json:"explanation" db:"explanation" binding:"omitempty"`
	Transcript 		*string 				`json:"transcript" db:"transcript" binding:"omitempty"`
	SubQuestions 	[]QuestionParams 		`json:"sub_questions" db:"-" binding:"required,dive"`
}