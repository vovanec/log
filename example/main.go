package main

import (
	"context"
	"log/slog"

	"github.com/vovanec/log"
	"github.com/vovanec/log/logattrs"
)

func doSomethingElse(ctx context.Context) {
	log.InfoContext(ctx, "doing something else",
		logattrs.String("function.name", "doSomethingElse"))
}

func doSomething(ctx context.Context) {
	log.InfoContext(ctx, "doing something",
		logattrs.String("function.name", "doSomething"))
	doSomethingElse(ctx)
}

func main() {
	log.Initialize(slog.LevelInfo)

	ctx := logattrs.String("app.name", "main").
		Context(context.Background())
	doSomething(ctx)
}
