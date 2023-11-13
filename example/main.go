package main

import (
	"context"

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

type AppVersion struct {
	Major int
	Minor int
	Patch int
}

func (v AppVersion) LogAttrs() *logattrs.Attrs {
	return logattrs.Group("version",
		logattrs.FromMap(
			map[string]any{
				"major": v.Major,
				"minor": v.Minor,
				"patch": v.Patch,
			},
		),
	)
}

type Application struct {
	Name    string
	Version AppVersion
	Build   string
}

func (a Application) LogAttrs() *logattrs.Attrs {
	return logattrs.Group("application",
		logattrs.FromMap(
			map[string]any{
				"name":    a.Name,
				"version": a.Version,
				"build": logattrs.FromMap(
					map[string]any{
						"commit": a.Build,
					},
				),
			},
		))
}

func main() {
	log.Initialize(log.WithLevel(log.LevelDebug))

	app := Application{
		Name:  "myApp",
		Build: "af6392c",
		Version: AppVersion{
			Major: 1,
			Minor: 0,
			Patch: 0,
		},
	}

	ctx := logattrs.New().
		Add(app).
		Context(context.Background())

	doSomething(ctx)
}
