package booking

import (
	"booking_system/internal/idgen"
	"booking_system/internal/model/house"
	"errors"
	"fmt"
	"time"
)

type Status string

const (
	StatusPending   Status = "pending"
	StatusConfirmed Status = "confirmed"
	StatusCancelled Status = "cancelled"
)

type Booking struct {
	ID         uint64
	UserID     uint64
	HouseID    uint64
	CheckIn    time.Time
	CheckOut   time.Time
	TotalPrice float64
	status     Status
	createdAt  time.Time
}

// NewBooking создает новое бронирование
func NewBooking(userID uint64, h *house.House, checkInDate, checkOutDate time.Time) (*Booking, error) {

	if err := validateBookingData(userID, h, checkInDate, checkOutDate); err != nil {
		return nil, err
	}

	id := idgen.GenerateID()
	if id == 0 {
		err := errors.New("не удалось сгенерировать ID бронирования")
		fmt.Println(err)
		return nil, err
	}

	checkIn := addHourToDate(checkInDate, h.CheckInHour)
	checkOut := addHourToDate(checkOutDate, h.CheckOutHour)

	b := &Booking{
		ID:        id,
		UserID:    userID,
		HouseID:   h.ID,
		CheckIn:   checkIn,
		CheckOut:  checkOut,
		status:    StatusPending,
		createdAt: time.Now(),
	}

	b.RecalculateTotal(h.PricePerDay)
	return b, nil
}

// validateBookingData проверяет корректность переданных данных
func validateBookingData(userID uint64, h *house.House, checkIn, checkOut time.Time) error {

	var errs []error

	if userID == 0 {
		errs = append(errs, errors.New("не указан ID пользователя"))
	}
	if h == nil {
		errs = append(errs, errors.New("не выбран дом для бронирования"))
	}
	if checkOut.Before(checkIn) || checkIn.Before(time.Now()) || checkIn.Equal(checkOut) {
		errs = append(errs, errors.New("некорректные даты заезда/выезда"))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

// addHourToDate добавляет время к дате
func addHourToDate(date time.Time, hour int) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), hour, 0, 0, 0, date.Location())
}

// Status возвращает текущий статус бронирования
func (b *Booking) Status() Status {
	return b.status
}

// CreatedAt возвращает дату создания бронирования
func (b *Booking) CreatedAt() time.Time {
	return b.createdAt
}

// Confirm переводит бронирование в статус подтверждено
func (b *Booking) Confirm() {
	if b.status == StatusPending {
		b.status = StatusConfirmed
	}
}

// Cancel отменяет бронирование
func (b *Booking) Cancel() {
	if b.status != StatusCancelled {
		b.status = StatusCancelled
	}
}

// IsConfirmed возвращает true, если бронирование подтверждено
func (b *Booking) IsConfirmed() bool {
	return b.status == StatusConfirmed
}

// IsCancelled возвращает true, если бронирование отменено
func (b *Booking) IsCancelled() bool {
	return b.status == StatusCancelled
}

// Nights возвращает количество ночей между датами заезда и выезда
func (b *Booking) Nights() int {
	checkIn := time.Date(b.CheckIn.Year(), b.CheckIn.Month(), b.CheckIn.Day(), 0, 0, 0, 0, b.CheckIn.Location())
	checkOut := time.Date(b.CheckOut.Year(), b.CheckOut.Month(), b.CheckOut.Day(), 0, 0, 0, 0, b.CheckOut.Location())

	days := int(checkOut.Sub(checkIn).Hours()) / 24
	if days <= 0 {
		return 1
	}
	return days
}

// RecalculateTotal пересчитывает сумму бронирования на основе количества ночей и цены за ночь
func (b *Booking) RecalculateTotal(pricePerNight float64) {

	b.TotalPrice = float64(b.Nights()) * pricePerNight

}

// String возвращает строковое представление бронирования.
func (b *Booking) String() string {
	return fmt.Sprintf("Бронирование %d\nПользователь: %d\nДом: %d\nС %s по %s\nНочей: %d\nСтатус: %s\nСумма: %.2f",
		b.ID,
		b.UserID,
		b.HouseID,
		b.CheckIn.Format("02.01.2006 15:04"),
		b.CheckOut.Format("02.01.2006 15:04"),
		b.Nights(),
		b.status.StatusText(),
		b.TotalPrice,
	)
}

// StatusText возвращает представление статуса на русском языке.
func (s Status) StatusText() string {

	if s == StatusPending {
		return "Ожидает подтверждения"
	}

	if s == StatusConfirmed {
		return "Подтверждено"
	}

	if s == StatusCancelled {
		return "Отменено"
	}

	return "Неизвестный статус"

}
