package log

import (
	"io"
	"os"
)

type config struct {
	textFormat bool
	output     io.Writer
	level      Level
}

type Option func(c *config)

func WithTextFormatter() Option {
	return func(c *config) {
		c.textFormat = true
	}
}

func WithOutput(o io.Writer) Option {
	return func(c *config) {
		c.output = o
	}
}

func WithLevel(l Level) Option {
	return func(c *config) {
		c.level = l
	}
}

func newConfig(opts ...Option) *config {
	conf := &config{
		output: os.Stderr,
		level:  LevelInfo,
	}

	for _, opt := range opts {
		opt(conf)
	}

	return conf
}
