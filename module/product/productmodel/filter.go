package productmodel

type Filter struct {
	Status      int    `json:"status,omitempty" form:"status"`
	Name        string `json:"name,omitempty" form:"name"`
	Description string `json:"description,omitempty" form:"description"`
}
