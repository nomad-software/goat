package log

import (
	"context"
	"fmt"
	"log/slog"
)

const (
	dateFormat = "15:04:05"
)

type handler struct {
	level slog.Level
}

func (h *handler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.level
}

func (h *handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *handler) WithGroup(name string) slog.Handler {
	return h
}

func (h *handler) Handle(ctx context.Context, r slog.Record) error {
	fmt.Printf(
		"%s %-5s %s",
		r.Time.Format(dateFormat),
		r.Level.String(),
		r.Message,
	)

	// WTF????
	// TODO use https://github.com/sirupsen/logrus
	r.Attrs(func(a slog.Attr) bool {
		fmt.Print(" ")
		fmt.Print(a)
		return true
	})

	fmt.Println("")

	return nil
}

func init() {
	handler := &handler{
		level: slog.LevelInfo,
	}

	log := slog.New(handler)

	slog.SetDefault(log)
}

func SetLevel(lvl slog.Level) {
	handler := &handler{
		level: lvl,
	}

	log := slog.New(handler)

	slog.SetDefault(log)
}
