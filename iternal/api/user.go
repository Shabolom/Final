package api

import (
	"Graduation_Project/iternal/models"
	"Graduation_Project/iternal/service"
	"Graduation_Project/iternal/tools"
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type UserAPI struct {
}

func NewUserApi() *UserAPI {
	return &UserAPI{}
}

var userService = service.NewUserService()

// @Summary	регистрация пользователя с выдачей токена
// @Produce	json
// @Accept	json
// @Tags	Authorization
// @Param	ввод	body		models.Register	true	"ввести логин и пароль"
// @Success	200		{string}	string	"вы зарегестрировались"
// @Failure	400		{object}	models.Error
// @Failure	500		{object}	models.Error
// @Failure	404		{object}	models.Error
// @Failure	409		{object}	models.Error
// @Router		/api/user/register [post]
func (ua *UserAPI) Regis(c *gin.Context) {
	var user models.Register

	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	err = json.Unmarshal(data, &user)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	err, mess, code, result := userService.Register(user)

	if err != nil {
		tools.CreateError(code, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	jwtString, err, code := tools.JWTCreator(result)
	if err != nil {
		tools.CreateError(code, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	c.Writer.Header().Set("Authorization", jwtString)

	c.String(code, mess)
}

// @Summary	авторизация с выдачей токена в куках и хэдерсе
// @Produce	json
// @Accept	json
// @Tags	Authorization
// @Param	ввод	body		models.Register	true	"авторизация"
// @Success	200		{string}	string	"успешно авторизировались"
// @Failure	400		{object}	models.Error
// @Failure	500		{object}	models.Error
// @Failure	401		{object}	models.Error
// @Router		/api/user/login [post]
func (ua *UserAPI) Login(c *gin.Context) {
	var user models.Register

	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	err = json.Unmarshal(data, &user)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	result, err, code := userService.Login(user)
	if err != nil {
		tools.CreateError(code, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	jwtString, err, code := tools.JWTCreator(result)
	if err != nil {
		tools.CreateError(code, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	c.Writer.Header().Set("Authorization", jwtString)
	tools.CookieCreate(c, jwtString)

	c.String(code, "успешно авторизировались")
}

// @Summary	заносим заказ в историю покупок user,a
// @Security ApiKeyAuth
// @Produce	json
// @Accept	json
// @Tags	User
// @Param	ввод	body		models.Order	true	"ввести номер заказа"
// @Success	200		{string}	string	"номер заказа уже был загружен этим пользователем"
// @Success	202		{string}	string	"новый номер заказа принят в обработку"
// @Failure	400		{object}	models.Error
// @Failure	500		{object}	models.Error
// @Failure	401		{object}	models.Error
// @Failure	409		{object}	models.Error
// @Failure	422		{object}	models.Error
// @Router		/api/user/orders [post]
func (ua *UserAPI) PostOrder(c *gin.Context) {
	var order models.Order
	userID := tools.JwtParsUserID(c)

	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	err = json.Unmarshal(data, &order)
	if err != nil {
		tools.CreateError(http.StatusUnprocessableEntity, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	mess, code, err := userService.PostOrder(order, userID)

	if err != nil {
		tools.CreateError(code, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	c.String(code, mess)
}

// @Summary	получение всех заказов сделанных конкретным юзером
// @Security ApiKeyAuth
// @Produce	json
// @Tags	User
// @Success	204		{string}	string	"нет данных для ответа"
// @Success	200		{object}	[]models.ReqOrders
// @Failure	400		{object}	models.Error
// @Failure	500		{object}	models.Error
// @Failure	401		{object}	models.Error
// @Router		/api/user/orders [get]
func (ua *UserAPI) GetOrder(c *gin.Context) {

	userID := tools.JwtParsUserID(c)

	result, err, code := userService.GetOrder(userID)
	if err != nil {
		tools.CreateError(code, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	c.JSON(code, result)
}

// @Summary	получение баланса конкретного юзера
// @Security ApiKeyAuth
// @Produce	json
// @Tags	User
// @Success	200		{object}	domain.Balance
// @Failure	500		{object}	models.Error
// @Failure	401		{object}	models.Error
// @Router		/api/user/balance [get]
func (ua *UserAPI) GetBal(c *gin.Context) {

	userID := tools.JwtParsUserID(c)

	result, err, code := userService.GetBal(userID)
	if err != nil {
		tools.CreateError(code, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	c.JSON(code, result)
}

// @Summary	Запрос на списание средств
// @Security ApiKeyAuth
// @Accept	json
// @Produce	json
// @Tags	User
// @Param	ввод	body		models.Purchase	true	"введите номер и стоимость заказа"
// @Success	200		{string}	string	"успешно преобретено"
// @Failure	500		{object}	models.Error
// @Failure	401		{object}	models.Error
// @Failure	422		{object}	models.Error
// @Failure	402		{object}	models.Error
// @Router		/api/user/balance/withdraw [post]
func (ua *UserAPI) PostWithD(c *gin.Context) {
	var purch models.Purchase
	userID := tools.JwtParsUserID(c)

	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	err = json.Unmarshal(data, &purch)
	if err != nil {
		tools.CreateError(http.StatusBadRequest, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	mess, err, code := userService.PostWithD(purch, userID)

	c.String(code, mess)
}

// @Summary	Получение информации о выводе средств
// @Security ApiKeyAuth
// @Produce	json
// @Tags	User
// @Success	200		{object}	[]domain.History
// @Success	204		{string}	string "нет ни 1 покупки"
// @Failure	204		{object}	models.Error
// @Failure	401		{object}	models.Error
// @Failure	500		{object}	models.Error
// @Router		/api/user/withdrawals [get]
func (ua *UserAPI) GetWithD(c *gin.Context) {
	userID := tools.JwtParsUserID(c)

	result, err, code := userService.GetWithD(userID)
	if err != nil {
		tools.CreateError(code, err, c)
		log.WithField("component", "rest").Warn(err)
		return
	}

	c.JSON(code, result)
}
