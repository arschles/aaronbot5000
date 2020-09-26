package http

import (
	"log"
	"net/http"
	"os/user"
	"path/filepath"

	"github.com/erikstmartin/erikbotdev/bot"
)

type Server struct {
	// this is the hub. we're calling it hug, deal with it
	addr    string
	webPath string
	hug     *hub
}

func NewServer(addr, webPath string) *Server {
	return &Server{
		addr:    addr,
		webPath: webPath,
		hug:     newhub(),
	}
}

func (s *Server) Start() error {
	go s.hug.run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(s.hug, w, r)
	})

	// http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(filepath.Join(bot.WebPath(), "public")))))
	// http.Handle("/media/", http.StripPrefix("/media/", http.FileServer(http.Dir(bot.MediaPath()))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(bot.WebPath(), "public", "index.html"))
	})

	return http.ListenAndServe(s.addr, nil)
}

func (s *Server) BroadcastMessage(msg Message) error {
	log.Printf("Broadcasting message %+v", msg)
	return s.hug.BroadcastMessage(msg)
}

func (s *Server) BroadcastChatMessage(user *user.User, msg string) error {
	m := &ChatMessage{
		User: user,
		Text: msg,
	}

	return s.hug.BroadcastMessage(m)
}
