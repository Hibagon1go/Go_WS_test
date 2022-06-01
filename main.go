package main

import (
	"go_ws_test/repository"
	"go_ws_test/sort"

	"github.com/gofiber/fiber/v2"
)

func main() {
	repository.SetupRedis()

	app := fiber.New()


    // ルーティング (別ファイルにするのも良さげ)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("users/:uuid", getUserList)
	app.Get("ranking", ranking)

	app.Listen(":8080")
}

func getUserList(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	userList, err := repository.GetUserList(uuid)
	if err != nil {
		panic(err)
	}

	userList = sort.RankingSort(userList)

	return c.JSON(userList)
}

func ranking(ctx *fiber.Ctx) error {
	result, err := repository.GetRankings()

	if err != nil {
		panic(err)
	}

	return ctx.JSON(result)
}