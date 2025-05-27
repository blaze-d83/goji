package server

import (
	"net/http"

	"github.com/blaze-d83/goji/internal/config"
)

func New(cfg *config.Config) *http.Server {

	// Declare a router
	// For example:
	// r := chi.NewRouter()

	return &http.Server{
		// Addr:                         "",
		// Handler:                      nil,
		// DisableGeneralOptionsHandler: false,
		// TLSConfig:                    &tls.Config{},
		// ReadTimeout:                  0,
		// ReadHeaderTimeout:            0,
		// WriteTimeout:                 0,
		// IdleTimeout:                  0,
		// MaxHeaderBytes:               0,
		// TLSNextProto:                 map[string]func(*http.Server, *tls.Conn, http.Handler){},
		// ConnState: func(net.Conn, http.ConnState) {
		// 	panic("TODO")
		// },
		// ErrorLog: &log.Logger{},
		// BaseContext: func(net.Listener) context.Context {
		// 	panic("TODO")
		// },
		// ConnContext: func(ctx context.Context, c net.Conn) context.Context {
		// 	panic("TODO")
		// },
		// HTTP2:     &http.HTTP2Config{},
		// Protocols: &http.Protocols{},
	}
}
