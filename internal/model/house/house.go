package house

import (
	"booking_system/internal/idgen"
	"errors"
	"fmt"
	"strings"
	"time"
)

type House struct {
	ID           uint64
	Title        string
	Description  string
	Address      string
	Photos       []string
	PricePerDay  float64
	OwnerID      uint64
	Amenities    Amenities
	rating       float64
	createdAt    time.Time
	isActive     bool
	CheckInHour  int
	CheckOutHour int
}

type Amenities struct {
	Wifi            bool // Доступ в интернет
	Kitchen         bool // Кухня, где гости могут готовить
	FreeParking     bool // Бесплатная парковка
	AirConditioning bool // Кондиционер
	Washer          bool // Стиральная машина
	Dryer           bool // Сушильная машина
	Dishwasher      bool // Посудомоечная машина
	TV              bool // Телевизор
	Pool            bool // Бассейн
	HotTub          bool // Горячая ванна
	IndoorFireplace bool // Камин
	PetFriendly     bool // Разрешено с питомцами
	SmokeAlarm      bool // Датчик дыма
	FirstAidKit     bool // Аптечка первой помощи
	HairDryer       bool // Фен
	Iron            bool // Утюг
}

const alpha = 0.3

// NewHouse возвращает структуру с новым домом
func NewHouse(title, description, address string, photos []string, price float64, ownerID uint64, amenities Amenities, checkIn, checkOut int) (*House, error) {

	if err := CheckHouseData(title, description, address, price); err != nil {
		fmt.Println(err)
		return nil, err
	}

	id := idgen.GenerateID()
	if id == 0 {
		err := errors.New("не удалось сгенерировать ID дома")
		fmt.Println(err)
		return nil, err
	}

	return &House{
		ID:           id,
		Title:        title,
		Description:  description,
		Address:      address,
		Photos:       photos,
		PricePerDay:  price,
		OwnerID:      ownerID,
		rating:       0,
		isActive:     true,
		createdAt:    time.Now(),
		Amenities:    amenities,
		CheckInHour:  checkIn,
		CheckOutHour: checkOut,
	}, nil
}

// CheckHouseData проверяет корректность переданных данных
func CheckHouseData(title, description, address string, price float64) error {

	var errs []error

	if title == "" {
		errs = append(errs, errors.New("не заполнено название дома"))
	}
	if description == "" {
		errs = append(errs, errors.New("не заполнено описание"))
	}
	if address == "" {
		errs = append(errs, errors.New("не указан адрес"))
	}
	if price == 0 {
		errs = append(errs, errors.New("не указана цена за ночь"))
	}

	if len(errs) != 0 {
		return errors.Join(errs...)
	}

	return nil

}

// Rating возвращает рейтинг дома
func (h *House) Rating() float64 {
	return h.rating
}

// CreatedAt возвращает дату создания объявления
func (h *House) CreatedAt() time.Time {
	return h.createdAt
}

// Deactivate делает дом неактивным для аренды
func (h *House) Deactivate() {
	h.isActive = false
}

// Activate активирует дом для аренды
func (h *House) Activate() {
	h.isActive = true
}

// IsActive возвращает статус активности объявления об аренде
func (h *House) IsActive() bool {
	return h.isActive
}

// UpdateRating обновляет рейтинг дома
func (h *House) UpdateRating(newRating float64) {

	if h.rating == 0 {
		h.rating = newRating
		return
	}

	h.rating = alpha*newRating + (1-alpha)*h.rating
}

// ListAvailable возвращает список удобств, которые доступны в данном доме
func (a *Amenities) ListAvailable() []string {

	var list []string

	if a.Wifi {
		list = append(list, "Wi-Fi")
	}
	if a.Kitchen {
		list = append(list, "Кухня")
	}
	if a.FreeParking {
		list = append(list, "Бесплатная парковка")
	}
	if a.AirConditioning {
		list = append(list, "Кондиционер")
	}
	if a.Washer {
		list = append(list, "Стиральная машина")
	}
	if a.Dryer {
		list = append(list, "Сушильная машина")
	}
	if a.Dishwasher {
		list = append(list, "Посудомоечная машина")
	}
	if a.TV {
		list = append(list, "Телевизор")
	}
	if a.Pool {
		list = append(list, "Бассейн")
	}
	if a.HotTub {
		list = append(list, "Горячая ванна")
	}
	if a.IndoorFireplace {
		list = append(list, "Камин")
	}
	if a.PetFriendly {
		list = append(list, "Можно с питомцами")
	}
	if a.SmokeAlarm {
		list = append(list, "Датчик дыма")
	}
	if a.FirstAidKit {
		list = append(list, "Аптечка")
	}
	if a.HairDryer {
		list = append(list, "Фен")
	}
	if a.Iron {
		list = append(list, "Утюг")
	}
	return list
}

// String возвращает строковое представление информации о доме.
func (h *House) String() string {
	return fmt.Sprintf(
		"%s\nАдрес: %s\nОписание: %s\nЦена за ночь: %.2f₽\nРейтинг: %.1f\nВладелец: %d\nФото: %d шт.\nУдобства: %s\nЗаезд с %02d:00, выезд до %02d:00",
		h.Title,
		h.Address,
		h.Description,
		h.PricePerDay,
		h.Rating(),
		h.OwnerID,
		len(h.Photos),
		strings.Join(h.Amenities.ListAvailable(), ", "),
		h.CheckInHour,
		h.CheckOutHour,
	)
}
