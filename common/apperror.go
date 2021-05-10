package common

import (
	"fmt"
	"log"
)

//AppError -- Application Error Structure
type AppError struct {
	ErrorCode   string
	ErrorMsg    string
	ErrorDetail string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s - %s", e.ErrorCode, e.ErrorMsg)
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)
}
