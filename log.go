package log

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/vovanec/log/logattrs"
)

func Initialize(opts ...Option) {

	var (
		conf    = newConfig(opts...)
		handler slog.Handler
	)

	logLevel, err := toSlogLevel(conf.level)
	if err != nil {
		panic(err)
	}

	if conf.textFormat {
		handler = slog.NewTextHandler(
			conf.output,
			&slog.HandlerOptions{Level: logLevel},
		)
	} else {
		handler = slog.NewJSONHandler(
			conf.output,
			&slog.HandlerOptions{Level: logLevel},
		)
	}

	slog.SetDefault(slog.New(handler))
}

func Debug(msg string, attr ...*logattrs.Attrs) {
	slog.Debug(msg, logattrs.Slice(attr)...)
}

func DebugContext(ctx context.Context, msg string, attr ...*logattrs.Attrs) {
	slog.DebugContext(ctx, msg, logattrs.SliceContext(ctx, attr)...)
}

func Info(msg string, attr ...*logattrs.Attrs) {
	slog.Info(msg, logattrs.Slice(attr)...)
}

func InfoContext(ctx context.Context, msg string, attr ...*logattrs.Attrs) {
	slog.InfoContext(ctx, msg, logattrs.SliceContext(ctx, attr)...)
}

func Warn(msg string, attr ...*logattrs.Attrs) {
	slog.Warn(msg, logattrs.Slice(attr)...)
}

func WarnContext(ctx context.Context, msg string, attr ...*logattrs.Attrs) {
	slog.WarnContext(ctx, msg, logattrs.SliceContext(ctx, attr)...)
}

func Error(msg string, attr ...*logattrs.Attrs) {
	slog.Error(msg, logattrs.Slice(attr)...)
}

func ErrorContext(ctx context.Context, msg string, attr ...*logattrs.Attrs) {
	slog.ErrorContext(ctx, msg, logattrs.SliceContext(ctx, attr)...)
}

func Panic(msg string, attr ...*logattrs.Attrs) {
	slog.Error(msg, logattrs.Slice(attr)...)
	doPanic(msg, attr...)
}

func PanicContext(ctx context.Context, msg string, attr ...*logattrs.Attrs) {
	slog.ErrorContext(ctx, msg, logattrs.SliceContext(ctx, attr)...)
	doPanic(msg, attr...)
}

func doPanic(msg string, attr ...*logattrs.Attrs) {
	if len(attr) > 0 {
		panic(fmt.Sprintf("%s, log attributes: %s", msg, attr[0]))
	}
	panic(msg)
}
