package log

import (
	"context"
	"fmt"
	"log/slog"
)

const (
	dateFormat = "15:04:05" // Log date format.
)

// Initialise the logger with the custom handler.
func init() {
	SetLevel(slog.LevelInfo)
}

// handler is a custom log handler.
type handler struct {
	level slog.Level
}

// Enabled implements the slog.Handler interface.
func (h *handler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.level
}

// WithAttrs implements the slog.Handler interface.
func (h *handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

// WithGroup implements the slog.Handler interface.
func (h *handler) WithGroup(name string) slog.Handler {
	return h
}

// Handle implements the slog.Handler interface.
func (h *handler) Handle(ctx context.Context, r slog.Record) error {
	fmt.Printf(
		"%-5s %s",
		// r.Time.Format(dateFormat),
		r.Level.String(),
		r.Message,
	)

	if r.NumAttrs() > 0 {
		fmt.Print(" ")
		r.Attrs(func(a slog.Attr) bool {
			if len(a.Key) > 0 {
				fmt.Printf("[ %s %s ] ", a.Key, a.Value)
			} else {
				fmt.Printf("[ %s ] ", a.Value)
			}
			return true
		})
	}

	fmt.Print("\n")

	return nil
}

// SetLevel sets the level of the custom logger.
func SetLevel(lvl slog.Level) {
	handler := &handler{
		level: lvl,
	}

	log := slog.New(handler)

	slog.SetDefault(log)
}
