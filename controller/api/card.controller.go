package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"platform-sample/controller/dto"
	"platform-sample/service"
	"strconv"
)

type CardController struct {
	service.CardService
}

func (CardController) NewCardController(service service.CardService) *CardController {
	return &CardController{service}
}

func (cardController *CardController) Init(e *echo.Group) {
	e.POST("", cardController.CreateCard)
	e.DELETE("", cardController.DeleteCard)
}

func (cardController *CardController) DeleteCard(c echo.Context) error {
	cardId, err := strconv.Atoi(c.QueryParam("cardId"))
	if err != nil {
		return ReturnApiFail(c, http.StatusBadRequest, ApiParameterError, err)
	}

	userId, err := strconv.Atoi(c.QueryParam("userId"))
	if err != nil {
		return ReturnApiFail(c, http.StatusBadRequest, ApiParameterError, err)
	}

	cardController.CardService.DeleteCard(cardId, userId)
	return ReturnApiSuccess(c, http.StatusNoContent, nil)
}

func (cardController *CardController) CreateCard(c echo.Context) error {
	cardDto := &dto.CardDto{}
	err := c.Bind(cardDto)
	if err != nil {
		return ReturnApiFail(c, http.StatusBadRequest, ApiParameterError, err)
	}
	card, err := cardController.CardService.CreateCard(cardDto)
	if err != nil {
		return ReturnApiFail(c, http.StatusInternalServerError, ApiQueryError, err)
	}
	return ReturnApiSuccess(c, http.StatusCreated, card)
}
