package main

import (
	"log"

	"github.com/joho/godotenv"
	GinHandler "github.com/yoshi-d-24/goal-sync/presentation/gin"
)

func main() {
	loadEnv()
	GinHandler.Start()
}

// .envを呼び出します。
func loadEnv() {
	err := godotenv.Load(".env")

	// もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
	if err != nil {
		log.Fatal("cannot read .env file", err)
	}
}
