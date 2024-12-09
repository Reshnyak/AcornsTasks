package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Reshnyak/AcornsTask/internal/gateway/gen"
	"github.com/gin-gonic/gin"
)

func AlicePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req gen.AliceRequest
		if err := c.BindJSON(&req); err != nil {
			log.Printf("Ошибка разбора JSON: %v\n", err)
			c.JSON(http.StatusBadRequest, gen.ErrorResponse{
				Error: "Неверный формат запроса",
			})
			return
		}

		// Логика обработки запроса от Алисы
		fmt.Println(req.Request.Command)
		fmt.Println(req.Session.UserId)

		// Пример формирования ответа
		response := gen.AliceResponse{
			Response: gen.AliceResponsePayload{
				Text:       "Привет, я бот!",
				Tts:        "Здравствуйте! Это мы, хоров+одо в+еды.",
				EndSession: true,
			},
			Session: gen.AliceResponseSession{
				SessionId: "2eac4854-fce721f3-b845abba-20d60",
				MessageId: 4,
				UserId:    "AC9WC3DF6FCE052E45A4566A48E6B7193774B84814CE49A922E163B8B29881DC",
			},
			// TodoLists: []string{"Сделать уборку", "Купить продукты"},
		}

		c.JSON(http.StatusOK, response)
	}
}
