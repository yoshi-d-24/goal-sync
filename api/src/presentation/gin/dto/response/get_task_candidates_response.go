package response

type GetTaskCandidatesResponse struct {
	TaskCandidates []TaskCandidate `json:"taskCandidates"`
}

type TaskCandidate struct {
	Name       string `json:"name"`
	MatcheRate string `json:"matcheRate"`
}
