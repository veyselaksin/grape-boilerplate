package genericresponse

import "github.com/gofiber/fiber/v2"

type GenericResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(ctx *fiber.Ctx, data interface{}, status int, msg ...string) error {
	if len(msg) == 0 {
		msg = append(msg, "Success")
	}

	return ctx.Status(status).JSON(GenericResponse{
		Success: true,
		Message: msg[0],
		Data:    data,
	})
}

func ErrorResponse(ctx *fiber.Ctx, status int, msg string, data ...interface{}) error {
	if len(data) == 0 {
		data = []interface{}{}
	}

	return ctx.Status(status).JSON(GenericResponse{
		Success: false,
		Message: msg,
		Data:    data,
	})
}

func RedirectResponse(ctx *fiber.Ctx, url string) error {
	return ctx.Redirect(url, fiber.StatusTemporaryRedirect)
}
