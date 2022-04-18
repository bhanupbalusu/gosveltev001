package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	//template render engine
	engine := html.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views: engine, //set as render engine
	})

	app.Static("/public", "./public")
	app.Get("/", mainPage)
	app.Post("/input", inputPage)
	app.Get("/formui", formuiPage)
	app.Listen(":3000")
}

func mainPage(c *fiber.Ctx) error {
	c.Accepts("html")
	name := c.FormValue("fname")
	selected := c.FormValue("cars")
	fmt.Println(name)
	fmt.Println(selected)

	//This function will be see different soon
	return c.Render("mainpage", nil)
}

func inputPage(c *fiber.Ctx) error {
	c.Accepts("html")
	name := c.FormValue("fname")
	selected := c.FormValue("cars")
	fmt.Println("From inputPage -> ", name)
	fmt.Println("From inputPage -> ", selected)

	//This function will be see different soon
	return c.Render("inputpage", nil)
}

func formuiPage(c *fiber.Ctx) error {
	//This function will be see different soon
	return c.Render("formuipage", nil)
}
