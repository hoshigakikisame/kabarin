package telegram

import (
	"bufio"
	"context"
	"os"
	"path/filepath"
	"strconv"

	"github.com/hoshigakikisame/kabarin/pkg/utils"
	"github.com/joho/godotenv"

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

func (t *Telegram) SendText(text string) error {
	receiverID := os.Getenv("TELEGRAM_RECEIVER_ID")

	if _, err := t.sender.Resolve(receiverID).Text(context.Background(), text); err != nil {
		return err
	}

	return nil
}

func (t *Telegram) SendFile(filePath string) error {
	receiverID := os.Getenv("TELEGRAM_RECEIVER_ID")

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	// Get file size
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	f, err := t.up.Upload(context.Background(), uploader.NewUpload(fileInfo.Name(), bufio.NewReader(file), fileInfo.Size()))
	if err != nil {
		return err
	}

	media := message.UploadedDocument(f).ForceFile(true).Filename(filepath.Base(filePath))

	if _, err := t.sender.Resolve(receiverID).Media(context.Background(), media); err != nil {
		return err
	}

	return nil
}

func (t *Telegram) Close() error {
	if err := t.stop(); err != nil {
		panic(err)
	}
	return nil
}
