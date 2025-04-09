package review

import (
	"booking_system/internal/idgen"
	"errors"
	"fmt"
	"time"
)

type Review struct {
	ID         uint64
	BookingID  uint64
	AuthorID   uint64
	SubjectID  uint64 // ID дома или пользователя
	IsForHouse bool   // true — отзыв о доме, false — отзыв о пользователе
	Rating     int
	Comment    string
	createdAt  time.Time
}

// NewReview создает новый отзыв.
func NewReview(bookingID, authorID, subjectID uint64, isForHouse bool, rating int, comment string) (*Review, error) {
	if err := validateReviewData(rating, comment); err != nil {
		return nil, err
	}

	id := idgen.GenerateID()
	if id == 0 {
		err := errors.New("не удалось сгенерировать ID отзыва")
		fmt.Println(err)
		return nil, err
	}

	return &Review{
		ID:         id,
		BookingID:  bookingID,
		AuthorID:   authorID,
		SubjectID:  subjectID,
		IsForHouse: isForHouse,
		Rating:     rating,
		Comment:    comment,
		createdAt:  time.Now(),
	}, nil
}

// validateReviewData проверяет корректность данных для отзыва.
func validateReviewData(rating int, comment string) error {

	var errs []error

	if rating < 1 || rating > 5 {
		errs = append(errs, errors.New("рейтинг должен быть от 1 до 5"))
	}
	if comment == "" {
		errs = append(errs, errors.New("не заполнен текст отзыва"))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

// CreatedAt возвращает дату создания отзыва.
func (r *Review) CreatedAt() time.Time {
	return r.createdAt
}

// String возвращает текстовое представление отзыва
func (r *Review) String() string {

	subj := "дом"
	if !r.IsForHouse {
		subj = "пользователя"
	}

	return fmt.Sprintf("Отзыв %d\nНа: %s (ID %d)\nОценка: %d\nКомментарий: %s",
		r.ID, subj, r.SubjectID, r.Rating, r.Comment)

}
