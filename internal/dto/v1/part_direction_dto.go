package v1dto

import "encoding/json"

type CreatePartDirectionInputParams struct {
	ExamId 			int					`json:"exam_id" db:"exam_id" binding:"required"`
	PartId 			int					`json:"part_id" db:"part_id" binding:"required"`
	Direction 		string 				`json:"direction_text" db:"direction_text" binding:"required"`
	AudioStartMs 	int 				`json:"audio_start_ms" db:"audio_start_ms"`
	AudioEndMs 		int 				`json:"audio_end_ms" db:"audio_end_ms"`
	ExampleData 	json.RawMessage 	`json:"example_data" db:"example_data"`
}

type UpdatePartDirectionInputParams struct {
	ExamId 			int					`json:"exam_id" db:"exam_id" binding:"required"`
	PartId 			int					`json:"part_id" db:"part_id" binding:"required"`
	Direction 		string 				`json:"direction_text" db:"direction_text" binding:"required"`
	AudioStartMs 	int 				`json:"audio_start_ms" db:"audio_start_ms"`
	AudioEndMs 		int 				`json:"audio_end_ms" db:"audio_end_ms"`
	ExampleData 	json.RawMessage 	`json:"example_data" db:"example_data"`
}