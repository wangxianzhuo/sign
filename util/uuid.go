package util

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func UUID() string {
	return fmt.Sprintf("%v", uuid.NewV4().String())
}
