package main

import (
	"context"
	"fmt"
	"github.com/slack-io/slacker"
	"log"
	"os"

	"strconv"
)

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "")
	os.Setenv("SLACK_APP_TOKEN", "")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	definition := &slacker.CommandDefinition{
		Command:     "my yob is <year>",
		Description: "yob calculator",
		Handler: func(ctx *slacker.CommandContext) {
			year := ctx.Request().Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				panic(err)
			}
			age := 2021 - yob
			r := fmt.Sprintf("age is %d", age)
			ctx.Response().Reply(r)
		},
	}

	bot.AddCommand(definition)

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
