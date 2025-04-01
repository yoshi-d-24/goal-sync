package taskcandidate

import (
	"context"
	"strings"
)

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

type TaskCandidate struct {
	Name      string
	MatchRate string
}

func NewGetTaskCandidatesApplicationService(generativeApiClient IGenerativeApiCleint) *GetTaskCandidatesApplicationService {
	return &GetTaskCandidatesApplicationService{generateApiClient: generativeApiClient}
}

func (s *GetTaskCandidatesApplicationService) Execute(ctx context.Context, command GetTaskCandidatesCommand) ([]TaskCandidate, error) {
	prompt := "あなたは文章からタスクを抽出するプロフェッショナルです。\n" +
		"文章の筆者の職業は" + command.Job + "です。\n" +
		"文章内にこの職業のタスクらしきものがあれば、箇条書きで、連番,タスク名,職業とタスクの関連率(%)を返却してください。\n" +
		"関連率が50%を下回るものは除外してください。\n" +
		"以下に箇条書きの例を与えます。\n" +
		"1,ユーザー管理画面の実装,95%\n" +
		"2,バグ修正,95%\n" +
		"3,パソコンを立ち上げる,50%\n" +
		"ここまでが箇条書きの例です。\n" +
		"タスクらしきものがない場合、 'No Task Found' と返却してください。\n" +
		"以下がタスク抽出対象の文章です。\n" +
		command.Text

	text, err := s.generateApiClient.GenerateText(ctx, prompt)
	if err != nil {
		return []TaskCandidate{}, err
	}

	if strings.Contains(text, "No Task Found") {
		return []TaskCandidate{}, nil
	}

	lines := strings.Split(text, "\\n")
	println(text)
	var candidates []TaskCandidate
	for _, line := range lines {
		println(line)
		if strings.TrimSpace(line) == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}
		name := strings.TrimSpace(parts[1])
		matchRate := strings.TrimSpace(parts[2])
		candidates = append(candidates, TaskCandidate{
			Name:      name,
			MatchRate: matchRate,
		})
	}

	return candidates, nil
}
