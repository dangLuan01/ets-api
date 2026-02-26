package models

type Exam struct {
	Id 				int 					`json:"exam_id" db:"id"`
	Title 			string 					`json:"title" db:"title"`
	Year 			int 					`json:"year" db:"year"`
	Category 		*string 				`json:"category" db:"category"`
	TotalTime 		int 					`json:"total_time" db:"total_time"`
	Description 	*string 				`json:"description" db:"description"`
	Thumbnail 		*string 				`json:"thumbnail" db:"thumbnail"`
	AudioFullUrl 	*string					`json:"audio_full_url" db:"audio_full_url"`
	Status 			int 					`json:"status" db:"status"`
	CreatedAt 		string 					`json:"created_at" db:"created_at"`
	Sections		[]ExamPart				`json:"sections"`
}

type ExamPart struct {
	Part 			int						`json:"part"`
	Direction   	Direction				`json:"direction"`
	Items			[]ExamQuestionMapping 	`json:"items,omitempty"`
}

type Direction struct {
	Text 			string 					`json:"text" db:"direction_text"`
	ExamId	        int						`json:"-" db:"exam_id"`
	Part	        int						`json:"-" db:"part"`
	AudioStart 		int 					`json:"audio_start_ms" db:"audio_start_ms"`
	AudioEnd 		int 					`json:"audio_end_ms" db:"audio_end_ms"`
	ExmapleRaw		[]byte					`json:"-" db:"example_data"`
	Exmaple			*ExampleData 			`json:"exmaple,omitempty"`
}

type ExampleData struct {
	Explanation 	string					`json:"explanation"`
	ImageUrl 		string 					`json:"image_url"`
	CorrectOption 	string 					`json:"correct_option"`
	AudioStartMs 	int						`json:"audio_start_ms"`
	AudioEndMs 		int						`json:"audio_end_ms"`
}

type ExamQuestionMapping struct {
	Id 				int						`json:"id" db:"id"`
	ExamId 			int 					`json:"exam_id" db:"exam_id"`
	EntityType 		string 					`json:"entity_type" db:"entity_type"`
	EntityId 		int 					`json:"entity_id" db:"entity_id"`
	OrderIndex 		int						`json:"order_index" db:"order_index"`
	Part 			int 					`json:"part" db:"part"`
	QuestionData 	*Question 				`json:"question_data,omitempty"`
	GroupData 		*QuestionGroup			`json:"group_data,omitempty"`
}

type QuestionGroup struct {
	Id 				int 					`json:"group_id" db:"id"`
	Part 			int 					`json:"part" db:"part"`
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
	Difficulty 		*int 					`json:"-" db:"difficulty"`
	Tags			*string 				`json:"-" db:"tags"`
	Options			map[string]*string 		`json:"options"`
}