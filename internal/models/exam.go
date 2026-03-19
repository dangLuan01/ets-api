package models

// --- TẦNG DANH MỤC (MASTER DATA) ---

type Certificate struct {
	Id 				int 					`json:"id" db:"id"`
	Code 			string 					`json:"code" db:"code"`
	Name 			string 					`json:"name" db:"name"`
	Description 	*string 				`json:"description" db:"description"`
	Status 			int 					`json:"status" db:"status"`
}

type SkillMaster struct {
	Id 				int 					`json:"id" db:"id"`
	CertId 			int 					`json:"cert_id" db:"cert_id"`
	Code 			string 					`json:"code" db:"code"`
	Name 			string 					`json:"name" db:"name"`
	OrderIndex      int     				`json:"order_index" db:"order_index"`
	Status 			int 					`json:"status" db:"status"`
}

type PartMaster struct {
	Id 				int 					`json:"id" db:"id"`
	SkillId 		int 					`json:"skill_id" db:"skill_id"`
	PartNumber 		int 					`json:"part_number" db:"part_number"`
	Name 			string 					`json:"name" db:"name"`
	Status 			int 					`json:"status" db:"status"`
}

// --- TẦNG RESPONSE (API JSON) ---

type ExamModel struct {
	Id 				int 					`json:"id" db:"id"`
	CertificateId	int						`json:"cert_id" db:"cert_id"`
	Title 			string 					`json:"title" db:"title"`
	Year 			int 					`json:"year" db:"year"`
	Category 		*string 				`json:"category" db:"category"`
	TotalTime 		int 					`json:"total_time" db:"total_time"`
	TotalQuestion	int						`json:"total_question" db:"total_question"`
	Description 	*string 				`json:"description" db:"description"`
	Thumbnail 		*string 				`json:"thumbnail" db:"thumbnail"`
	AudioFullUrl 	*string					`json:"audio_full_url" db:"audio_full_url"`
	Status 			int 					`json:"status" db:"status"`
	CreatedAt 		string 					`json:"created_at" db:"created_at"`
}

type Exam struct {
	Id 				int 					`json:"exam_id" db:"id"`
	CertificateId	int						`json:"-" db:"cert_id"`
	Title 			string 					`json:"title" db:"title"`
	Year 			int 					`json:"year" db:"year"`
	Category 		*string 				`json:"category" db:"category"`
	TotalTime 		int 					`json:"total_time" db:"total_time"`
	TotalQuestion	int						`json:"total_question" db:"total_question"`
	Description 	*string 				`json:"description" db:"description"`
	Thumbnail 		*string 				`json:"thumbnail" db:"thumbnail"`
	AudioFullUrl 	*string					`json:"audio_full_url" db:"audio_full_url"`
	Status 			int 					`json:"status" db:"status"`
	CreatedAt 		string 					`json:"created_at" db:"created_at"`
	Skills			[]ExamSkill				`json:"skills"`
}

type ExamSkill struct {
    SkillId         int         			`json:"-"`
    SkillCode       string      			`json:"skill_code"`
    SkillName       string      			`json:"skill_name"`
    Parts           []ExamPart  			`json:"parts"`
}

type ExamPart struct {
	PartId 			int						`json:"-"`
	PartNumber      int                     `json:"part_number"`
    PartName        string                  `json:"part_name"`
	Direction   	*Direction				`json:"direction,omitempty"`
	Items			[]ExamQuestionMapping 	`json:"items"`
}

type Direction struct {
	Id              int             		`json:"-" db:"id"`
	Text 			string 					`json:"text" db:"direction_text"`
	ExamId	        int						`json:"-" db:"exam_id"`
	PartId	        int						`json:"-" db:"part_id"`
	AudioStart 		int 					`json:"audio_start_ms" db:"audio_start_ms"`
	AudioEnd 		int 					`json:"audio_end_ms" db:"audio_end_ms"`
	ExampleRaw		[]byte					`json:"-" db:"example_data"`
	Example			*ExampleData 			`json:"example,omitempty"`
}

type ExampleData struct {
	Explanation 	string					`json:"explanation"`
	ImageUrl 		string 					`json:"image_url"`
	CorrectOption 	string 					`json:"correct_option"`
	AudioStartMs 	int						`json:"audio_start_ms"`
	AudioEndMs 		int						`json:"audio_end_ms"`
}

type ExamQuestionMapping struct {
	Id 				int						`json:"-" db:"id"`
	ExamId 			int 					`json:"exam_id" db:"exam_id"`
	EntityType 		string 					`json:"entity_type" db:"entity_type"`
	EntityId 		int 					`json:"entity_id" db:"entity_id"`
	OrderIndex 		int						`json:"order_index" db:"order_index"`
	PartId 			int 					`json:"part_id" db:"part_id"`
	QuestionData 	*Question 				`json:"question_data,omitempty"`
	GroupData 		*QuestionGroup			`json:"group_data,omitempty"`
}

type QuestionGroup struct {
	Id 				int 					`json:"-" db:"id"`
	PartId 			int 					`json:"-" db:"part_id"`
	PassageText 	*string 				`json:"passage_text" db:"passage_text"`
	ImageUrl 		*string 				`json:"image_url" db:"image_url"`
	AudioStart 		*int 					`json:"audio_start_ms" db:"audio_start_ms"`
	AudioEnd 		*int 					`json:"audio_end_ms" db:"audio_end_ms"`
	Transcript		*string					`json:"transcript" db:"transcript"`
	Explanation		*string					`json:"explanation" db:"explanation"`
	SubQuestions	[]Question 				`json:"sub_questions"`
}

type Question struct {
	Id 				int 					`json:"question_id" db:"id"`
	GroupId 		*int 					`json:"-" db:"group_id"`
	QuestionText 	*string 				`json:"question_text" db:"question_text"`
	ImageUrl 		*string 				`json:"image_url" db:"image_url"`
	AudioStart 		*int 					`json:"audio_start_ms" db:"audio_start_ms"`
	AudioEnd 		*int 					`json:"audio_end_ms" db:"audio_end_ms"`
	OptionA 		*string 				`json:"-" db:"option_a"`
	OptionB 		*string 				`json:"-" db:"option_b"`
	OptionC 		*string 				`json:"-" db:"option_c"`
	OptionD 		*string 				`json:"-" db:"option_d"`
	CorrectAnswer 	*string 				`json:"correct_answer" db:"correct_answer"`
	DisplayNumber	int						`json:"display_number"`
	SubOrder		int						`json:"sub_order" db:"sub_order"`
	Explanation 	*string 				`json:"explanation" db:"explanation"`
	Transcript 		*string 				`json:"transcript" db:"transcript"`
	Tags			*string 				`json:"-" db:"tags"`
	Options			map[string]*string 		`json:"options"`
}

type UserAnswer struct {
	AttemptId 		int 					`db:"attempt_id"`
	QuestionId 		int 					`db:"question_id"`
	SelectedAnswer 	*string 					`db:"selected_answer"`
	IsCorrect 		bool 					`db:"is_correct"`
}

type UserAttempt struct {
	UserId 			int						`db:"user_id"`
	ExamId 			int						`db:"exam_id"`
	StartTime 		string					`db:"start_time"`
	EndTime			string					`db:"end_time"`
	TotalScore 		int						`db:"total_score"`
	ListeningScore 	int						`db:"listening_score"`
	ReadingScore 	int						`db:"reading_score"`
}

type ScoreConversion struct {
	SkillId 		int 					`db:"skill_id"`
	RawScore 		int 					`db:"raw_score"`
	ScaledScore 	int 					`db:"scaled_score"`
}