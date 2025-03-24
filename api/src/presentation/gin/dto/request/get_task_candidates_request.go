package request

type GetTaskCandidates struct {
	Text string `json:"text" binding:"required,max=500"`
}
