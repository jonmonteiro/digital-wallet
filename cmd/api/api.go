package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmonteiro/picpay-like/core/domain/user"
	"github.com/jmonteiro/picpay-like/core/domain/wallet"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Route("/api/v1", func(api chi.Router) {

		// ===== WALLET =====
		walletStore := wallet.NewStore(s.db)
		walletService := wallet.NewWalletService(walletStore)

		// ===== USER =====
		userStore := user.NewStore(s.db)
		userSvc := user.NewUserService(userStore, walletStore)
		userHdlr := user.NewHandler(userSvc)
		userHdlr.RegisterRoutes(api)

		// ===== WALLET ROUTES =====
		walletHdlr := wallet.NewHandler(userStore, walletService)
		walletHdlr.RegisterRoutes(api)

		// ===== TRANSACTION =====
		// TODO: Implementar transaction handler
	})

	// Servir arquivos estáticos (opcional)
	r.Handle("/*", http.FileServer(http.Dir("static")))

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, r)
}
