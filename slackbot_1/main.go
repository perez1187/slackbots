package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shomali11/slacker"
)

func printCommandEvent(analiticsChannel <-chan *slacker.CommandEvent) {

	for event := range analiticsChannel {
		fmt.Println("Command Event")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {

	// because it is a test, tokens are paste here
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5092455568849-5073272973894-sPg9KSAwwduA6TsWfktWQKoJ")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A052576C69L-5073245929270-72cdb01f106e09eba48b75a9af1ec4d6a0c0e0281c7a59878e2aebd2f7e414b8")

	// create slack client
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvent(bot.CommandEvents())

	// the most important part is Command
	bot.Command("ping", &slacker.CommandDefinition{
		//this func create automaticalyy, bc -> botCtx, request and response
		Handler: func(bc slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
			w.Reply("pong")
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
