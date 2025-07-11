package discord

import (
	"bytes"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/hoshigakikisame/kabarin/pkg/utils"
	"github.com/joho/godotenv"
	"github.com/projectdiscovery/gologger"
)

type Discord struct {
	client    *discordgo.Session
	channelID string
}

func New() (*Discord, error) {
	godotenv.Load()

	if err := utils.ValidateEnvVars("DISCORD_TOKEN", "CHANNEL_ID"); err != nil {
		return nil, err
	}

	var (
		botToken  = os.Getenv("DISCORD_TOKEN")
		channelID = os.Getenv("CHANNEL_ID")
	)

	parsedBotToken := "Bot " + botToken

	discord, err := discordgo.New(parsedBotToken)

	discord.Identify.Intents = discordgo.IntentsGuildMessageTyping

	if err != nil {
		return nil, err
	}

	return &Discord{
		client:    discord,
		channelID: channelID,
	}, nil

}

func (dc *Discord) SendText(text *string) error {

	if _, err := dc.client.ChannelMessageSend(dc.channelID, *text); err != nil {
		gologger.Error().Msgf("Failed to send message: %v", err)
		return err
	}

	return nil
}

func (dc *Discord) SendFile(fileName *string, data *[]byte) error {

	fileData := bytes.NewReader(*data)

	if _, err := dc.client.ChannelFileSend(dc.channelID, *fileName, fileData); err != nil {
		gologger.Error().Msgf("Failed to send file %s: %v", *fileName, err)
		return err
	}

	return nil
}

func (dc *Discord) Close() error {
	if err := dc.client.Close(); err != nil {
		gologger.Fatal().Msg(err.Error())
		return err
	}
	return nil
}
