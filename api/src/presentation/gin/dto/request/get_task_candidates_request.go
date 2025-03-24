package request

type GetTaskCandidates struct {
	Text string `json:"text" binding:"required,max=500"`
	Job  string `json:"job" binding:"required,max=50"`
}
