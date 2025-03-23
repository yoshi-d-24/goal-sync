package ginhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	TaskApplicationService "github.com/yoshi-d-24/goal-sync/application/task"
	GormCore "github.com/yoshi-d-24/goal-sync/infrastructure/gorm/core"
	GormTaskRepository "github.com/yoshi-d-24/goal-sync/infrastructure/gorm/task"
	Request "github.com/yoshi-d-24/goal-sync/presentation/gin/dto/request"
)

func Start() {
	r := gin.Default()

	taskRepository := GormTaskRepository.NewGormTaskRepository(GormCore.CreateDB())

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
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
