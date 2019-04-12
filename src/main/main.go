package main

import (
	"chat"
	"fmt"
	"logger"
	"net"
	"os"
	"os/signal"
	"syscall"
	"tao"
)

// ChatServer is the chatting server.
type ChatServer struct {
	*tao.Server
}

// NewChatServer returns a ChatServer.
func NewChatServer() *ChatServer {
	onConnectOption := tao.OnConnectOption(func(conn tao.WriteCloser) bool {
		logger.Infoln("on connect")
		return true
	})
	onErrorOption := tao.OnErrorOption(func(conn tao.WriteCloser) {
		logger.Infoln("on error")
	})
	onCloseOption := tao.OnCloseOption(func(conn tao.WriteCloser) {
		logger.Infoln("close chat client")
	})
	return &ChatServer{
		tao.NewServer(onConnectOption, onErrorOption, onCloseOption),
	}
}

func main() {
	defer logger.Start().Stop()

	tao.Register(chat.ChatMessage, chat.DeserializeMessage, chat.ProcessMessage)

	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", "0.0.0.0", 12345))
	if err != nil {
		logger.Fatalln("listen error", err)
	}
	chatServer := NewChatServer()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c
		chatServer.Stop()
	}()

	logger.Infoln(chatServer.Start(l))
}
