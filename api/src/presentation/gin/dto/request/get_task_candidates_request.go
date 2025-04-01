package request

type GetTaskCandidatesRequest struct {
	Text string `json:"text" binding:"required,max=500"`
	Job  string `json:"job" binding:"required,max=50"`
}
