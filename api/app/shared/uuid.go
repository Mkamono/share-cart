package shared

import (
	"strings"

	"github.com/google/uuid"
)

// uuid generates a new UUID without hyphens.
func Uuid() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
