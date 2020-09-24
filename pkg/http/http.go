package http

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/erikstmartin/erikbotdev/bot"
)

type Server struct {
	// this is the hub. we're calling it hug, deal with it
	hug *hub
}

func NewServer(addr, webPath string) *Server {
	return &Server{
		hug: newhub(),
	}
}

func (s *Server) Start(addr string, webPath string) error {
	go s.hug.run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	// http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(filepath.Join(bot.WebPath(), "public")))))
	// http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir(bot.MediaPath()))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(bot.WebPath(), "public", "index.html"))
	})

	return http.ListenAndServe(addr, nil)
}

func (s *Server) BroadcastMessage(msg Message) error {
	log.Printf("Broadcasting message %+v", msg)
	return s.hug.BroadcastMessage(msg)
}

// func (s *Server) BroadcastChatMessage(user *bot.User, msg string) error {
// 	m := &ChatMessage{
// 		User: user,
// 		Text: msg,
// 	}

// 	return s.hug.BroadcastMessage(m)
// }
