package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/vlamai/papir_discord/config"
	"github.com/vlamai/papir_discord/message"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	jira, conn := config.NewJira()
	defer conn.Close()

	parser := config.NewParser(jira)

	dg := config.NewDiscordBot()

	dg.AddHandler(ParseMessages(parser))
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err := dg.Open()
	if err != nil {
		log.Fatalf("error opening connection | %v", err)
	}
	log.Println("Bot is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	select {
	case <-sc:
		err = dg.Close()
		if err != nil {
			log.Fatalf("close session : %v", err)
		}
	}
}

func ParseMessages(p message.Parser) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "ping" {
			_, err := s.ChannelMessageSend(m.ChannelID, "Pong!")
			if err != nil {
				log.Printf("send message | %v", err)
			}
		}

		result := p.Parse(m.Content)

		for _, text := range result {
			_, err := s.ChannelMessageSend(m.ChannelID, text)
			if err != nil {
				log.Printf("send message | %v", err)
			}
		}
	}
}
