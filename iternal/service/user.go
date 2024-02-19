package service

import (
	"Graduation_Project/iternal/domain"
	"Graduation_Project/iternal/models"
	"Graduation_Project/iternal/repository"
	"Graduation_Project/iternal/tools"
	"errors"
	"github.com/gofrs/uuid"
	"net/http"
	"time"
)

type UserService struct {
}

func NewUserService() *UserService {
	return &UserService{}
}

var userRepo = repository.NewUserRepo()

func (us *UserService) Register(user models.Register) (error, string, int, domain.Register) {
	id, err := uuid.NewV4()
	if err != nil {
		return err, "", http.StatusInternalServerError, domain.Register{}
	}

	password, err := tools.HashPassword(user.Password)
	if err != nil {
		return err, "", http.StatusInternalServerError, domain.Register{}
	}

	userEntity := domain.Register{
		Login:    user.Login,
		Password: password,
	}
	userEntity.LoyalTokens = 1000
	userEntity.ID = id

	err, mess, code, result := userRepo.Register(userEntity)

	if err != nil {
		return err, mess, http.StatusBadRequest, domain.Register{}
	}

	return nil, mess, code, result
}

func (us *UserService) Login(user models.Register) (domain.Register, error, int) {

	result, err := userRepo.Find(user.Login)

	if err != nil {
		return domain.Register{}, errors.New("не верный логин или пароль"), http.StatusUnauthorized
	}

	if !tools.CheckPasswordHash(user.Password, result.Password) {
		return domain.Register{}, errors.New("не верный логин или пароль"), http.StatusUnauthorized
	}

	return result, nil, http.StatusOK
}

func (us *UserService) PostOrder(order models.Order, userID uuid.UUID) (string, int, error) {
	orderID, err := uuid.NewV4()
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	orderEntity := domain.UserOrder{
		UserID:    userID,
		UserOrder: order.Number,
	}
	orderEntity.ID = orderID

	mess, code, err := userRepo.PostOrder(orderEntity)
	if err != nil {
		return mess, code, err
	}

	return mess, code, nil
}

func (us *UserService) GetOrder(userID uuid.UUID) ([]models.ReqOrders, error, int) {
	var reqOrders []models.ReqOrders
	result, err, code := userRepo.GetOrder(userID)

	if err != nil {
		return []models.ReqOrders{}, err, code
	}

	for _, order := range result {
		reqOrderEntity := models.ReqOrders{
			Order: order.UserOrder,
			When:  order.CreatedAt,
		}
		reqOrders = append(reqOrders, reqOrderEntity)
	}

	return reqOrders, nil, code
}

func (us *UserService) GetBal(userID uuid.UUID) (domain.Balance, error, int) {

	result, err, code := userRepo.GetBal(userID)

	userBalance := domain.Balance{
		LoyalTokens:     result.LoyalTokens,
		UsedLoyalTokens: result.UsedLoyalTokens,
	}

	if err != nil {
		return domain.Balance{}, err, code
	}

	return userBalance, nil, code
}

func (us *UserService) PostWithD(purch models.Purchase, userID uuid.UUID) (string, error, int) {

	result, err, code := us.GetBal(userID)
	if err != nil {
		return "", err, code
	}

	if purch.Product == 0 {
		return "", errors.New("не верный ключ в поле json"), http.StatusBadRequest
	}

	if result.LoyalTokens < purch.Sum {
		return "недостаточно средств", nil, http.StatusPaymentRequired
	}

	result = domain.Balance{
		LoyalTokens:     result.LoyalTokens - purch.Sum,
		UsedLoyalTokens: result.UsedLoyalTokens + purch.Sum,
	}

	histEntity := domain.History{
		UserID:      userID,
		CreatedAt:   time.Now(),
		Sum:         purch.Sum,
		OrderNumber: purch.Product,
	}

	err, code = userRepo.PostWithD(userID, result, histEntity)
	if err != nil {
		return "", err, code
	}

	return "запрос успешно обработан", nil, code
}

func (us *UserService) GetWithD(userID uuid.UUID) ([]domain.History, error, int) {

	result, err, code := userRepo.GetWithD(userID)
	if err != nil {
		return []domain.History{}, err, code
	}

	return result, nil, code
}
