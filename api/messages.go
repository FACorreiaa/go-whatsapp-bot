package api

import (
	"context"
	"log"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types/events"
)

var wac *whatsmeow.Client

func HandleEvent(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		go HandleMessage(v)
	default:
		log.Printf("Unhandled event type: %T\n", v)
	}
}

func HandleMessage(messageEvent *events.Message) {
	messageContent := messageEvent.Message.GetConversation()
	if messageContent == "" {
		log.Println("Empty message content")
		return
	}
	response, err := OpenAI(messageContent)
	if err != nil {
		log.Printf("error generating response from OpenAI: %v\n", err)
		return
	}

	println(response)

	msg := &proto.Message{
		Conversation: &response,
	}

	_, err = wac.SendMessage(context.Background(), messageEvent.Info.Chat, msg)
	if err != nil {
		log.Printf("Failed to send message: %v\n", err)
	}

	//println(message)
}
