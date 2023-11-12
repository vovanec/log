package logattrs

import (
	"log/slog"
	"time"
)

func New() *Attrs {
	return &Attrs{attr: make(map[string]slog.Attr)}
}

func Error(err error) *Attrs {
	return New().AddError(err)
}

func String(key, value string) *Attrs {
	return New().AddString(key, value)
}

func Int64(key string, value int64) *Attrs {
	return New().AddInt64(key, value)
}

func Uint64(key string, value uint64) *Attrs {
	return New().AddUint64(key, value)
}

func Int(key string, value int) *Attrs {
	return New().AddInt(key, value)
}

func Float64(key string, value float64) *Attrs {
	return New().AddFloat64(key, value)
}

func Bool(key string, value bool) *Attrs {
	return New().AddBool(key, value)
}

func Time(key string, value time.Time) *Attrs {
	return New().AddTime(key, value)
}
