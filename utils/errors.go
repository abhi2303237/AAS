package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type Error struct {
	Code    int32
	Message string
}

func (he *Error) Error() string {
	return he.Message + " : StatusCode:" + fmt.Sprintf("%v: ", he.Code)
}

func SendConceptDescriptionError(c *fiber.Ctx, code int, message string) error {

	petErr := Error{
		Code:    int32(code),
		Message: message,
	}

	return c.Status(code).JSON(petErr)
}

func GetError(code int, message string) error {
	return &Error{
		Code:    int32(code),
		Message: message,
	}
}

func HandleFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func HandleError(err error) {
	if err != nil {
		log.Error(err)
	}
}
