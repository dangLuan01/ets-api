package v1dto

type GetAllCategoryParams struct {
	Page 	int32 `form:"page" binding:"required,min=1"`
	Limit 	int32 `form:"limit" binding:"required,min=1,max=50"`
}

type CategoryInputParams struct {
	ParentId 		*int 	`json:"parent_id" db:"parent_id"`
	Name			string 	`json:"name" db:"name"`
	Slug			*string	`json:"slug" db:"slug"`
	Type			string	`json:"type" db:"type"`
	Status 			int		`json:"status" db:"status"`
	IsFilterable 	int 	`json:"is_filterable" db:"is_filterable"`
	Priority 		int 	`json:"priority" db:"priority"`
}

type CategoryUpdateParams struct {
	Id 				int 	`json:"id" binding:"required"`
	ParentId 		*int 	`json:"parent_id" db:"parent_id"`
	Name			string 	`json:"name" db:"name"`
	Slug			*string	`json:"slug" db:"slug"`
	Type			string	`json:"type" db:"type"`
	Status 			int		`json:"status" db:"status"`
	IsFilterable 	int 	`json:"is_filterable" db:"is_filterable"`
	Priority 		int 	`json:"priority" db:"priority"`
}

type CategoryDTO struct {
	ParentId 		*int 	`json:"parent_id" db:"parent_id"`
	Name			string 	`json:"name" db:"name"`
	Slug			string	`json:"slug" db:"slug"`
	Type			string	`json:"type" db:"type"`
	Status 			int		`json:"status" db:"status"`
	IsFilterable 	int 	`json:"is_filterable" db:"is_filterable"`
	Priority 		int 	`json:"priority" db:"priority"`
}

type CategoryStructure struct {
	Id			int		`json:"id" db:"id"`
	Name		string	`json:"name" db:"name"`
	ParentId	*int	`json:"-" db:"parent_id"`
	Children	[]*CategoryStructure `json:"children,omitempty" db:"-"`
}

func (c *CategoryStructure) GetID() int {
    return c.Id
}

func (c *CategoryStructure) GetParentID() *int {
    return c.ParentId
}

func (c *CategoryStructure) GetChildren() *[]*CategoryStructure {
    return &c.Children
}