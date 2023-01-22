package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {

	panic("Fetch the tokens from local repo for app to run")

	botTokenName := "SLACK_BOT_TOKEN"
	botToken := "ge it fromtoken text file in local repo"

	appTokenName := "SLACK_APP_TOKEN"
	appToken := "ge it fromtoken text file in local repo"

	os.Setenv(botTokenName, botToken)
	os.Setenv(appTokenName, appToken)

	bot := slacker.NewClient(os.Getenv(botTokenName), os.Getenv(appTokenName))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob Calulator",
		Examples:    []string{"my yob is 1990"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)

			if err != nil {
				fmt.Println(err)
			}
			age := 2023 - yob
			r := fmt.Sprintf("Age is %d", age)
			fmt.Println(r)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)

	if err != nil {
		log.Fatal(err)
	}
}
