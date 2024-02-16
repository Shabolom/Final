package repository

import (
	"Graduation_Project/config"
	"Graduation_Project/iternal/domain"
	"errors"
	"github.com/gofrs/uuid"
	"net/http"
)

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) Register(user domain.Register) (error, string, int, domain.Register) {

	if _, err := ur.Find(user.Login); err == nil {
		return nil, "такой пользователь уже существует", http.StatusConflict, domain.Register{}
	}

	err := config.DB.
		Create(&user).
		Error

	if err != nil {
		return err, "", http.StatusBadRequest, domain.Register{}
	}

	return nil, "вы зарегестрировались", http.StatusOK, user
}

func (ur *UserRepo) Find(login string) (domain.Register, error) {
	var user domain.Register

	err := config.DB.
		Where("login = ?", login).
		First(&user).
		Error

	if err != nil {
		return domain.Register{}, err
	}

	return user, nil
}

func (ur *UserRepo) PostOrder(order domain.UserOrder) (string, int, error) {

	code := ur.FindOrder(order)
	if code == http.StatusOK {
		return "номер заказа уже был загружен вами", code, nil
	}

	if code == http.StatusConflict {
		return "номер заказа уже был загружен другим пользователем", code, nil
	}

	err := config.DB.Create(&order).
		Error
	if err != nil {
		return "", http.StatusBadRequest, err
	}

	return "новый номер заказа принят в обработку;", http.StatusAccepted, nil
}

func (ur *UserRepo) FindOrder(order domain.UserOrder) int {
	var postOrder domain.UserOrder

	err := config.DB.
		Where("order = ?", order.Order).
		First(&postOrder).
		Error

	if err != nil {
		return http.StatusAccepted
	}

	if postOrder.UserID != order.UserID {
		return http.StatusConflict
	}

	return http.StatusOK
}

func (ur *UserRepo) GetOrder(userID uuid.UUID) ([]domain.UserOrder, error, int) {
	var orders []domain.UserOrder

	err := config.DB.
		Where("user_id = ?", userID).
		Find(&orders).
		Error

	if len(orders) == 0 {
		return []domain.UserOrder{}, errors.New("у вас нет заказов"), http.StatusNoContent
	}

	if err != nil {
		return []domain.UserOrder{}, err, http.StatusInternalServerError
	}

	return orders, nil, http.StatusOK
}

func (ur *UserRepo) GetBal(userID uuid.UUID) (domain.Register, error, int) {
	var orders domain.Register

	err := config.DB.
		Where("id = ?", userID.String()).
		Find(&orders).
		Error

	if err != nil {
		return domain.Register{}, err, http.StatusNoContent
	}

	return orders, nil, http.StatusOK
}

func (ur *UserRepo) PostWithD(userID uuid.UUID, balance domain.Balance, history domain.History) (error, int) {
	var user domain.Register

	err := config.DB.
		Model(&user).
		Where("id = ?", userID.String()).
		Updates(domain.Register{
			LoyalTokens:     balance.LoyalTokens,
			UsedLoyalTokens: balance.UsedLoyalTokens,
		}).
		Error
	if err != nil {
		return err, http.StatusBadRequest
	}

	err = config.DB.
		Create(&history).
		Error
	if err != nil {
		return err, http.StatusInternalServerError
	}

	return nil, http.StatusOK
}

func (ur *UserRepo) GetWithD(userID uuid.UUID) ([]domain.History, error, int) {
	var userHistory []domain.History

	err := config.DB.
		Where("user_id = ?", userID).
		Find(&userHistory).
		Error

	if len(userHistory) == 0 {
		return []domain.History{}, nil, http.StatusNoContent
	}

	if err != nil {
		return []domain.History{}, err, http.StatusNoContent
	}

	return userHistory, nil, http.StatusOK
}
