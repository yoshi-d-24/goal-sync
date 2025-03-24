package ginhandler

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	TaskApplicationService "github.com/yoshi-d-24/goal-sync/application/task"
	TaskCandidateApplicationService "github.com/yoshi-d-24/goal-sync/application/taskcandidate"
	Gemini "github.com/yoshi-d-24/goal-sync/infrastructure/gemini"
	GormCore "github.com/yoshi-d-24/goal-sync/infrastructure/gorm/core"
	GormTaskRepository "github.com/yoshi-d-24/goal-sync/infrastructure/gorm/task"
	Request "github.com/yoshi-d-24/goal-sync/presentation/gin/dto/request"
)

func Start() {
	r := gin.Default()

	db := GormCore.CreateDB()
	taskRepository := GormTaskRepository.NewGormTaskRepository(GormCore.CreateDB())
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDb.Close()

	r.POST("/task-candidates", func(c *gin.Context) {
		var json Request.GetTaskCandidates

		if err := c.Copy().ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx := context.Background()

		geminiApiClient, err := Gemini.NewGeminiApiClient(ctx)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		getTaskCandidateService := TaskCandidateApplicationService.NewGetTaskCandidatesApplicationService(geminiApiClient)

		command := TaskCandidateApplicationService.GetTaskCandidatesCommand{
			Text: json.Text,
			Job:  json.Job,
		}

		candidates, err := getTaskCandidateService.Execute(ctx, command)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]string{"candidates": candidates})
	})

	r.POST("/task", func(c *gin.Context) {
		var json Request.RegisterTask

		if err := c.Copy().ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		regsterTaskApplicatoinService := TaskApplicationService.NewRegisterTaskApplicationService(taskRepository)

		command := TaskApplicationService.RegisterTaskCommand{
			Title:       json.Title,
			Description: json.Description,
			Dod:         json.Dod,
		}
		if err := regsterTaskApplicatoinService.Execute(command); err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	})

	r.Run(":8080") // デフォルトで :8080
}
