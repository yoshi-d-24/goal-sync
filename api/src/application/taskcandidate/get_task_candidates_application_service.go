package taskcandidate

import "context"

type IGenerativeApiCleint interface {
	GenerateText(ctx context.Context, prompt string) (string, error)
}

type GetTaskCandidatesApplicationService struct {
	generateApiClient IGenerativeApiCleint
}

type GetTaskCandidatesCommand struct {
	Text string
	Job  string
}

func NewGetTaskCandidatesApplicationService(generativeApiClient IGenerativeApiCleint) *GetTaskCandidatesApplicationService {
	return &GetTaskCandidatesApplicationService{generateApiClient: generativeApiClient}
}

// TODO: return TaskCandidate struct and validate the generated text
func (s *GetTaskCandidatesApplicationService) Execute(ctx context.Context, command GetTaskCandidatesCommand) (string, error) {
	prompt := "あなたは文章からタスクを抽出するプロフェッショナルです。\n" +
		"文章の筆者の職業は" + command.Job + "です。\n" +
		"文章内にこの職業のタスクらしきものがあれば、箇条書きで返却してください。\n" +
		"タスクらしきものがない場合、 'No Task Found' と返却してください。\n" +
		"以下がタスク抽出対象の文章です。\n" +
		command.Text

	text, err := s.generateApiClient.GenerateText(ctx, prompt)
	if err != nil {
		return "", err
	}
	return text, nil
}
