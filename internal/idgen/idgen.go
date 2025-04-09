package idgen

import (
	"github.com/sony/sonyflake"
)

// GenerateID генерирует уникальный ID
func GenerateID() uint64 {

	sf := sonyflake.NewSonyflake(sonyflake.Settings{})

	if sf == nil {
		return 0
	}

	id, _ := sf.NextID()
	return id
}
