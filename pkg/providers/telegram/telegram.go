package telegram

import (
	"bytes"
	"context"
	"os"
	"strconv"

	"github.com/hoshigakikisame/kabarin/pkg/utils"
	"github.com/joho/godotenv"
	"github.com/projectdiscovery/gologger"

	"github.com/gotd/contrib/bg"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/message"
	"github.com/gotd/td/telegram/uploader"
	"github.com/gotd/td/tg"
)

type Telegram struct {
	api    *tg.Client
	sender *message.Sender
	up     *uploader.Uploader
	stop   bg.StopFunc
}

func New() (*Telegram, error) {
	godotenv.Load()
	utils.ValidateEnvVars("TELEGRAM_API_ID", "TELEGRAM_API_HASH", "TELEGRAM_BOT_TOKEN", "TELEGRAM_RECEIVER_ID")

	var (
		apiID, _ = strconv.Atoi(os.Getenv("TELEGRAM_API_ID"))
		apiHash  = os.Getenv("TELEGRAM_API_HASH")
		botToken = os.Getenv("TELEGRAM_BOT_TOKEN")
	)

	ctx := context.Background()

	client := telegram.NewClient(apiID, apiHash, telegram.Options{
		SessionStorage: &telegram.FileSessionStorage{
			Path: "./sessions/telegram.json",
		},
	})

	stop, err := bg.Connect(client)
	if err != nil {
		return nil, err
	}

	status, err := client.Auth().Status(ctx)
	if err != nil {
		return nil, err
	}

	if !status.Authorized {
		if _, err := client.Auth().Bot(ctx, botToken); err != nil {
			return nil, err
		}
	}

	api := tg.NewClient(client)
	sender := message.NewSender(api)
	up := uploader.NewUploader(api)

	return &Telegram{
		api:    api,
		sender: sender,
		up:     up,
		stop:   stop,
	}, nil
}

func (t *Telegram) SendText(text *string) error {
	receiverID := os.Getenv("TELEGRAM_RECEIVER_ID")

	if _, err := t.sender.Resolve(receiverID).Text(context.Background(), *text); err != nil {
		return err
	}

	return nil
}

func (t *Telegram) SendFile(fileName *string, data *[]byte) error {
	receiverID := os.Getenv("TELEGRAM_RECEIVER_ID")

	f, err := t.up.Upload(context.Background(), uploader.NewUpload(*fileName, bytes.NewReader(*data), int64(len(*data))))
	if err != nil {
		return err
	}

	media := message.UploadedDocument(f).ForceFile(true).Filename(*fileName)

	if _, err := t.sender.Resolve(receiverID).Media(context.Background(), media); err != nil {
		return err
	}

	return nil
}

func (t *Telegram) Close() error {
	if err := t.stop(); err != nil {
		gologger.Error().Msgf("Error stopping telegram client: %s", err)
	}
	return nil
}
