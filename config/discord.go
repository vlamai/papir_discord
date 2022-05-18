package config

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func NewDiscordBot() *discordgo.Session {
	dg, err := discordgo.New("Bot " + getEnvVar("PAPIR_DISCORD_TOKEN"))
	if err != nil {
		log.Fatalf("error creating Discord session | %v", err)
	}
	return dg
}
