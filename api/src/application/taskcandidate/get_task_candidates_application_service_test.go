package taskcandidate_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yoshi-d-24/goal-sync/application/taskcandidate"
)

// IGenerativeApiCleint のモック
type mockGenerativeApiClient struct {
	GenerateTextFunc func(ctx context.Context, prompt string) (string, error)
}

func (m *mockGenerativeApiClient) GenerateText(ctx context.Context, prompt string) (string, error) {
	if m.GenerateTextFunc != nil {
		return m.GenerateTextFunc(ctx, prompt)
	}
	return "", errors.New("GenerateTextFunc not implemented")
}

func TestGetTaskCandidatesApplicationService_Execute(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name               string
		command            taskcandidate.GetTaskCandidatesCommand
		mockGenerateText   func(ctx context.Context, prompt string) (string, error)
		expectedCandidates []taskcandidate.TaskCandidate
		expectedError      error
	}{
		{
			name: "正常系: タスク候補が抽出される場合",
			command: taskcandidate.GetTaskCandidatesCommand{
				Text: "ユーザー管理画面とバグ修正をお願いします。",
				Job:  "エンジニア",
			},
			mockGenerateText: func(ctx context.Context, prompt string) (string, error) {
				// 実際のAPI応答は改行コードを含む可能性があるため、\n で表現
				return "1,ユーザー管理画面の実装,95%\\n2,バグ修正,95%", nil
			},
			expectedCandidates: []taskcandidate.TaskCandidate{
				{Name: "ユーザー管理画面の実装", MatcheRate: "95%"},
				{Name: "バグ修正", MatcheRate: "95%"},
			},
			expectedError: nil,
		},
		{
			name: "正常系: No Task Found が返される場合",
			command: taskcandidate.GetTaskCandidatesCommand{
				Text: "特に何もありません。",
				Job:  "エンジニア",
			},
			mockGenerateText: func(ctx context.Context, prompt string) (string, error) {
				return "No Task Found", nil
			},
			expectedCandidates: []taskcandidate.TaskCandidate{},
			expectedError:      nil,
		},
		{
			name: "異常系: GenerateText がエラーを返す場合",
			command: taskcandidate.GetTaskCandidatesCommand{
				Text: "エラーが発生するテキスト",
				Job:  "エンジニア",
			},
			mockGenerateText: func(ctx context.Context, prompt string) (string, error) {
				return "", errors.New("API error")
			},
			expectedCandidates: []taskcandidate.TaskCandidate{},
			expectedError:      errors.New("API error"),
		},
		{
			name: "正常系: 不正な形式の行が含まれる場合（無視されること）",
			command: taskcandidate.GetTaskCandidatesCommand{
				Text: "不正な形式を含むテキスト",
				Job:  "エンジニア",
			},
			mockGenerateText: func(ctx context.Context, prompt string) (string, error) {
				return "1,正しいタスク,90%\\n不正な行\\n3,別の正しいタスク,80%", nil
			},
			expectedCandidates: []taskcandidate.TaskCandidate{
				{Name: "正しいタスク", MatcheRate: "90%"},
				{Name: "別の正しいタスク", MatcheRate: "80%"},
			},
			expectedError: nil,
		},
		{
			name: "正常系: 空行が含まれる場合（無視されること）",
			command: taskcandidate.GetTaskCandidatesCommand{
				Text: "空行を含むテキスト",
				Job:  "エンジニア",
			},
			mockGenerateText: func(ctx context.Context, prompt string) (string, error) {
				return "1,タスク1,90%\\n\\n2,タスク2,85%", nil
			},
			expectedCandidates: []taskcandidate.TaskCandidate{
				{Name: "タスク1", MatcheRate: "90%"},
				{Name: "タスク2", MatcheRate: "85%"},
			},
			expectedError: nil,
		},
		{
			name: "正常系: API応答の末尾に改行がある場合",
			command: taskcandidate.GetTaskCandidatesCommand{
				Text: "末尾改行テキスト",
				Job:  "エンジニア",
			},
			mockGenerateText: func(ctx context.Context, prompt string) (string, error) {
				return "1,タスクA,92%\\n2,タスクB,88%\\n", nil // 末尾に \n
			},
			expectedCandidates: []taskcandidate.TaskCandidate{
				{Name: "タスクA", MatcheRate: "92%"},
				{Name: "タスクB", MatcheRate: "88%"},
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := &mockGenerativeApiClient{
				GenerateTextFunc: tt.mockGenerateText,
			}
			// モジュールパスが正しいか確認してください
			service := taskcandidate.NewGetTaskCandidatesApplicationService(mockClient)

			candidates, err := service.Execute(ctx, tt.command)

			// アサーションの順序をエラーチェック -> 成功時の値チェック に変更
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
				// エラーが期待される場合、candidates の内容はチェックしないか、空であることを期待する
				assert.Empty(t, candidates)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedCandidates, candidates)
			}
		})
	}
}
