package v1dto

type FilterStructure struct {
	Id			int		`json:"id" db:"id"`
	Name		string	`json:"name" db:"name"`
	ParentId	*int	`json:"-" db:"parent_id"`
	Type		string	`json:"type" db:"type"`
	Children	[]*FilterStructure `json:"children,omitempty" db:"-"`
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