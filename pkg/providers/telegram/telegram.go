package telegram

import (
	"context"
	"fmt"
	"log"
	"os"
	fp "path/filepath"
	"strconv"

	"github.com/celestix/gotgproto"
	"github.com/celestix/gotgproto/ext"
	"github.com/celestix/gotgproto/sessionMaker"
	"github.com/gotd/td/telegram/uploader"
	"github.com/gotd/td/tg"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
)

type Telegram struct {
	client  *gotgproto.Client
	context *ext.Context
}

func New() (*Telegram, error) {
	godotenv.Load()

	apiID, err := strconv.Atoi(os.Getenv("TELEGRAM_API_ID"))
	if err != nil {
		return nil, err
	}

	var (
		apiHash  = os.Getenv("TELEGRAM_API_HASH")
		botToken = os.Getenv("TELEGRAM_BOT_TOKEN")
	)

	fmt.Println(botToken)

	client, err := gotgproto.NewClient(
		apiID,
		apiHash,
		gotgproto.ClientTypeBot(botToken),
		&gotgproto.ClientOpts{
			Session:          sessionMaker.SqlSession(sqlite.Open("kabarin_tg.sqlite")),
			DisableCopyright: true,
		},
	)
	if err != nil {
		log.Fatalln("failed to start client:", err)
	}

	context := client.CreateContext()

	return &Telegram{
		client:  client,
		context: context,
	}, nil
}

func (t *Telegram) SendText(text string) error {
	chatID, err := strconv.ParseInt(os.Getenv("TELEGRAM_CHAT_ID"), 10, 64)
	if err != nil {
		return err
	}

	if _, err := t.context.SendMessage(chatID, &tg.MessagesSendMessageRequest{
		Message: text,
	}); err != nil {
		return err
	}

	return nil
}

func (t *Telegram) SendFile(filepath string) error {
	chatID, err := strconv.ParseInt(os.Getenv("TELEGRAM_CHAT_ID"), 10, 64)
	if err != nil {
		return err
	}

	ctx := context.Background()
	f, err := uploader.NewUploader(t.context.Raw).FromPath(ctx, filepath)
	if err != nil {
		return err
	}

	filename := fp.Base(filepath)

	if _, err := t.context.SendMedia(chatID, &tg.MessagesSendMediaRequest{
		Media: &tg.InputMediaUploadedDocument{
			Attributes: []tg.DocumentAttributeClass{
				&tg.DocumentAttributeFilename{FileName: filename},
			},
			File:      f,
			ForceFile: true,
		},
	}); err != nil {
		return err
	}

	return nil
}

func (t *Telegram) Close() error {
	t.client.Stop()
	return nil
}
