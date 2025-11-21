package rest

import "log/slog"

type usecase interface {
}

type Handler struct {
	log *slog.Logger
	//usecase
}
