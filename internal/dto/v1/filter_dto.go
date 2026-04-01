package v1dto

type FilterStructure struct {
	Id			int		`json:"id" db:"id"`
	Name		string	`json:"name" db:"name"`
	ParentId	*int	`json:"-" db:"parent_id"`
	Type		string	`json:"type" db:"type"`
	Children	[]*FilterStructure `json:"children,omitempty" db:"-"`
}

type FilterExamParams struct {
	Search		*string	`form:"search" binding:"omitempty"`
	CategoryId	[]int	`form:"category_id" binding:"omitempty"`
	Page		int32	`form:"page" binding:"required,min=1"`
	Limit		int32	`form:"limit" binding:"required,max=50"`
}

func (m *FilterStructure) GetID() int {
    return m.Id
}

func (m *FilterStructure) GetParentID() *int {
    return m.ParentId
}

func (m *FilterStructure) GetChildren() *[]*FilterStructure {
    return &m.Children
}