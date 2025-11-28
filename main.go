package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIレスポンスの設計図
type beat struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// テストデータのスライス
var beats = []beat{
	{ID: "1", Title: "A"},
	{ID: "2", Title: "B"},
}

// APIハンドラーを設定する
// beatsをJSONで返す処理
func getBeats(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, beats)
}

// ルーターの登録
func main() {
	router := gin.Default()
	router.GET("/beats", getBeats)
	router.Run("localhost:8080")
}
