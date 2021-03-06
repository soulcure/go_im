package chat

import (
	"context"
	"logger"
	"tao"
)

const (
	// ChatMessage is the message number of chat message.
	ChatMessage int16 = 106
)

// Message defines the chat message.
type Message struct {
	Content string
}

// MessageNumber returns the message number.
func (cm Message) MessageNumber() int16 {
	return ChatMessage
}

// Serialize Serializes Message into bytes.
func (cm Message) Serialize() ([]byte, error) {
	return []byte(cm.Content), nil
}

// DeserializeMessage deserializes bytes into Message.
func DeserializeMessage(data []byte) (message tao.Message, err error) {
	if data == nil {
		return nil, tao.ErrNilData
	}
	content := string(data)
	msg := Message{
		Content: content,
	}
	return msg, nil
}

// ProcessMessage handles the Message logic.
func ProcessMessage(ctx context.Context, conn tao.WriteCloser) {
	logger.Infof("ProcessMessage")
	s, ok := tao.ServerFromContext(ctx)
	if ok {
		msg := tao.MessageFromContext(ctx)
		s.Broadcast(msg)
	}
}
