package web

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

func Web() {
	app := fiber.New()

	app.Get("/async", func(ctx *fiber.Ctx) error {
		go takeTime("Maria")
		return ctx.Send([]byte(time.Now().String()))
	})

	app.Listen(":8080")
}

func takeTime(s string) {
	fmt.Println(time.Now())
	time.Sleep(5 * time.Second)
	fmt.Printf("this is a notification by: %s\n", s)

}
