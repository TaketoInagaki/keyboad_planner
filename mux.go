package main

import (
	"context"
	"net/http"

	"github.com/TaketoInagaki/keyboard_planner/auth"
	"github.com/TaketoInagaki/keyboard_planner/clock"
	"github.com/TaketoInagaki/keyboard_planner/config"
	"github.com/TaketoInagaki/keyboard_planner/handler"
	"github.com/TaketoInagaki/keyboard_planner/service"
	"github.com/TaketoInagaki/keyboard_planner/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})
	v := validator.New()
	db, cleanup, err := store.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}
	clocker := clock.RealClocker{}
	r := store.Repository{Clocker: clocker}
	rcli, err := store.NewKVS(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}

	// 認証認可系
	jwter, err := auth.NewJWTer(rcli, clocker)
	if err != nil {
		return nil, cleanup, err
	}
	ru := &handler.RegisterUser{
		Service:   &service.RegisterUser{DB: db, Repo: &r},
		Validator: v,
	}
	mux.Post("/register", ru.ServeHTTP)
	l := &handler.Login{
		Service: &service.Login{
			DB:             db,
			Repo:           &r,
			TokenGenerator: jwter,
		},
		Validator: v,
	}
	mux.Post("/login", l.ServeHTTP)
	mux.Route("/admin", func(r chi.Router) {
		r.Use(handler.AuthMiddleware(jwter), handler.AdminMiddleware)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			_, _ = w.Write([]byte(`{"message": "admin only"}`))
		})
	})

	// タスク系
	at := &handler.AddTask{
		Service:   &service.CreateTask{DB: db, Repo: &r},
		Validator: v,
	}
	lt := &handler.ListTask{
		Service:   &service.ListTask{DB: db, Repo: &r},
		Validator: v,
	}
	mux.Route("/tasks", func(r chi.Router) {
		r.Use(handler.AuthMiddleware(jwter))
		r.Post("/", at.ServeHTTP)
		r.Get("/", lt.ServeHTTP)
	})

	// 振り返り系API
	cr := &handler.CreateOrEditReflection{
		Service:   &service.CreateOrEditReflection{DB: db, Repo: &r},
		Validator: v,
	}
	fr := &handler.FetchReflection{
		Service:   &service.FetchReflection{DB: db, Repo: &r},
		Validator: v,
	}
	mux.Route("/reflection", func(r chi.Router) {
		r.Use(handler.AuthMiddleware(jwter))
		r.Post("/", cr.ServeHTTP)
		r.Get("/", fr.ServeHTTP)
	})

	// 継続リスト系API
	cc := &handler.CreateOrEditContinuationList{
		Service:   &service.CreateOrEditContinuationList{DB: db, Repo: &r},
		Validator: v,
	}
	// fc := &handler.FetchContinuationList{
	// 	Service:   &service.FetchContinuationList{DB: db, Repo: &r},
	// 	Validator: v,
	// }
	mux.Route("/continuation", func(r chi.Router) {
		r.Use(handler.AuthMiddleware(jwter))
		r.Post("/", cc.ServeHTTP)
		// r.Get("/", fc.ServeHTTP)
	})

	return mux, cleanup, nil
}
