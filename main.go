package main

import (
	"context"
	"flag"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	gogpt "github.com/sashabaranov/go-gpt3"
	"log"
	"time"
)

var (
	telegramToken string
	openaiToken   string
	maxToken      int

	telegramClient *tgbotapi.BotAPI
	openAIClient   *gogpt.Client
)

func initFlags() {
	flag.StringVar(&telegramToken, "t", "", "telegram token")
	flag.StringVar(&openaiToken, "o", "", "openai token")
	flag.IntVar(&maxToken, "m", 2000, "max token count, default 2000")
	flag.Parse()
}

func initTelegram() {
	var err error
	telegramClient, err = tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Panic(err)
	}

	telegramClient.Debug = true

	log.Printf("Telegram Authorized on account %s", telegramClient.Self.UserName)
}

func initOpenAI() {
	openAIClient = gogpt.NewClient(openaiToken)
}

func replyMsg(chatId int64, msg string, msgId int) {
	message := tgbotapi.NewMessage(chatId, msg)
	message.ReplyToMessageID = msgId

	_, err := telegramClient.Send(message)

	if err != nil {
		log.Printf("send back failed. %v \n", err)
	}
}

func msgHandler(update *tgbotapi.Update) {
	log.Printf("[%s] %s\n", update.Message.From.UserName, update.Message.Text)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	req := gogpt.CompletionRequest{
		Model:     "text-davinci-003",
		MaxTokens: maxToken,
		Prompt:    update.Message.Text,
	}
	resp, err := openAIClient.CreateCompletion(ctx, req)
	if err != nil {
		log.Printf("wait openai reply fail. %v \n", err)
		replyMsg(update.Message.Chat.ID, "wait for openAI reply fail", update.Message.MessageID)

	}

	replyMsg(update.Message.Chat.ID, resp.Choices[0].Text, update.Message.MessageID)
}

func msgLoop() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := telegramClient.GetUpdatesChan(u)

	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message != nil { // If we got a message
			msgHandler(&update)
		}
	}
}

func main() {
	initFlags()
	initTelegram()
	initOpenAI()
	msgLoop()
}
