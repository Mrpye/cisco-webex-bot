// @title cisco-webex-bot
// @version 1.0
// @description cisco-webex-bot is a CLI application written in Golang that gives the ability to send notifications to a Webex room that the bot is subscribed to via a rest api

// @contact.url https://github.com/Mrpye/cisco-webex-bot

// @license.name Apache 2.0 licensed
// @license.url https://github.com/Mrpye/cisco-webex-bot/blob/main/LICENSE

// @host localhost:8080
// @BasePath /
// @query.collection.format multi
package main

import "github.com/Mrpye/cisco-webex-bot/cmd"

func main() {
	cmd.Execute()
}
