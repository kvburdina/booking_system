package user

import (
	"booking_system/internal/idgen"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          uint64
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	password    string
	rating      float64
	Bio         string
	created_at  time.Time
	isOwner     bool
	isTourist   bool
}

const alfa = 0.3

// NewUser создает нового пользователя
func NewUser(name, lastname, email, phone, pass, bio string, owner, tourist bool) (*User, error) {

	if err := CheckUserData(name, lastname, email, phone, pass, owner, tourist); err != nil {
		fmt.Println(err)
		return nil, err
	}

	id := idgen.GenerateID()
	if id == 0 {
		err := errors.New("не удалось сгенерировать ID пользователя")
		fmt.Println(err)
		return nil, err
	}

	u := &User{
		ID:          id,
		FirstName:   name,
		LastName:    lastname,
		Email:       email,
		PhoneNumber: phone,
		rating:      0,
		Bio:         bio,
		created_at:  time.Now(),
	}

	if u.HashPassword(pass) != nil {
		err := errors.New("возникла ошибка при хэшировании пароля")
		fmt.Println(err)
		return nil, err
	}

	return u, nil
}

// CheckUserData проверяет корректность переданных данных
func CheckUserData(name, lastname, email, phone, pass string, owner, tourist bool) error {

	var errs []error

	if name == "" {
		errs = append(errs, errors.New("не заполнено имя пользователя"))
	}
	if lastname == "" {
		errs = append(errs, errors.New("не заполнена фамилия пользователя"))
	}
	if email == "" {
		errs = append(errs, errors.New("не заполнена почта пользователя"))
	}
	if phone == "" {
		errs = append(errs, errors.New("не заполнен телефон пользователя"))
	}
	if pass == "" {
		errs = append(errs, errors.New("не заполнен пароль пользователя"))
	}
	if !owner && !tourist {
		errs = append(errs, errors.New("не выбраны роли пользователя"))
	}

	if len(errs) != 0 {
		return errors.Join(errs...)
	}

	return nil

}

// HashPassword хэширует переданный пароль.
func (u *User) HashPassword(pass string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.password = string(hashed)
	return nil
}

// CheckPassword сравнивает хэшированный пароль с переданным.
func (u *User) CheckPassword(pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(pass))
	return err == nil
}

// UpdateRating обновляет рейтинг пользователя по методу экспоненциального сглаживания
func (u *User) UpdateRating(newRat float64) {

	if u.rating == 0 {
		u.rating = newRat
		return
	}
	u.rating = alfa*newRat + (1-alfa)*u.rating

}

// CreatedAt возвращает дату создания пользователя
func (u *User) CreatedAt() time.Time {
	return u.created_at
}

// Rating возвращает рейтинг пользователя
func (u *User) Rating() float64 {
	return u.rating
}
