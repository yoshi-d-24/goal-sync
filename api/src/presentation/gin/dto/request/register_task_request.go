package request

type RegisterTask struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Dod         string `json:"dod" binding:"required"`
}
