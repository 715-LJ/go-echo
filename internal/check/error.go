package check

import (
	"fmt"
	"log/slog"
)

func IfError(err error) bool {
	if err == nil {
		return false
	}
	slog.Error(fmt.Sprintf("%v", err))
	return true
}
