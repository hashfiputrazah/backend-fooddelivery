package helpers

import (
	"github.com/google/uuid"
)

func UUIDGen() string {
	var id = uuid.New()

	return id.String()
}
