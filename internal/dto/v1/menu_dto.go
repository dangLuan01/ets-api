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
	ParentId	*int		`json:"-" db:"parent_id"`
	Children 	[]*MenuDTO 	`json:"children,omitempty" db:"-"`
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
