package v1dto

type GetMenuParams struct {
	Limit 	int32 `form:"limit" binding:"omitempty,min=1,max=50"`
	Page 	int32 `form:"page" binding:"omitempty,min=1"`
	Type	string `form:"type" binding:"required"`
}

type MenuDTO struct {
	Id 			int			`json:"-" db:"id"`
	Name 		string 		`json:"name" db:"name"`
	Slug 		*string 	`json:"slug" db:"slug"`
	Type		string		`json:"type" db:"type"`
	ParentId	*int		`json:"-" db:"parent_id"`
	Children 	[]*MenuDTO 	`json:"children,omitempty" db:"-"`
}

type GetAllMenuParams struct {
	Page 	int32 `form:"page" binding:"required,min=1"`
	Limit 	int32 `form:"limit" binding:"required,min=1,max=50"`
}

type MenuInputParams struct {
	ParentId 		*int 	`json:"parent_id" db:"parent_id"`
	Name			string 	`json:"name" db:"name"`
	Slug			*string	`json:"slug" db:"slug"`
	Type			string	`json:"type" db:"type"`
	Status 			int		`json:"status" db:"status"`
	Priority 		int 	`json:"priority" db:"priority"`
}

type MenuUpdateParams struct {
	Id 				int 	`json:"id" binding:"required"`
	ParentId 		*int 	`json:"parent_id" db:"parent_id"`
	Name			string 	`json:"name" db:"name"`
	Slug			*string	`json:"slug" db:"slug"`
	Type			string	`json:"type" db:"type"`
	Status 			int		`json:"status" db:"status"`
	Priority 		int 	`json:"priority" db:"priority"`
}

func (m *MenuDTO) GetID() int {
    return m.Id
}

func (m *MenuDTO) GetParentID() *int {
    return m.ParentId
}

func (m *MenuDTO) GetChildren() *[]*MenuDTO {
    return &m.Children
}
