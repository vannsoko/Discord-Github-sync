package discord

import (
	"log"
	"github.com/bwmarrin/discordgo"
)

var (
    bot *discordgo.Session
    BotToken string
    err error
)


func InitBot(token string) {
    log.Println("Bot started")
    bot, err = discordgo.New("Bot " + token)
    if err != nil {
        log.Fatalln("func InitBot: error creating Discord bot:", err)
    }
    //bot.AddHandler()
    // TODO implementing the discord bot handling

    err = bot.Open()
    if err != nil {
        log.Fatalln("func InitBot: error opening connection:", err)
    }

    BotToken = token
}

func CloseBot() {
    log.Println("Bot is closing")
    if bot != nil {
        bot.Close()
    }
}
