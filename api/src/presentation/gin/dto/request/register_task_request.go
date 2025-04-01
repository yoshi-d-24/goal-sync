package request

type RegisterTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Dod         string `json:"dod" binding:"required"`
}
