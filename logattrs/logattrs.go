package logattrs

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"
)

type (
	logAttrCtxKeyType struct{}
)

var logAttrCtxKey logAttrCtxKeyType

type LogAttrs interface {
	LogAttrs() *Attrs
}

type Attrs struct {
	attr map[string]slog.Attr
}

func (a *Attrs) Duration(key string, value time.Duration) *Attrs {
	return New().AddDuration(key, value)
}

func Context(ctx context.Context) *Attrs {
	if attr, ok := ctx.Value(logAttrCtxKey).(*Attrs); ok {
		return New().Add(attr)
	}
	return New()
}

func (a *Attrs) Has(key string) bool {
	_, found := a.attr[key]
	return found
}

func (a *Attrs) LogAttrs() *Attrs {
	return a
}

func (a *Attrs) Empty() bool {
	return a == nil || len(a.attr) == 0
}

func (a *Attrs) Add(attr LogAttrs) *Attrs {
	if attr == nil {
		return a
	}

	for _, attrs := range attr.LogAttrs().attr {
		a.addAttr(attrs)
	}

	return a
}

func (a *Attrs) AddError(err error) *Attrs {
	if errAttr, ok := err.(LogAttrs); ok {
		return a.Add(errAttr)
	}
	return a
}

func (a *Attrs) AddString(key, value string) *Attrs {
	return a.addAttr(slog.String(key, value))
}

func (a *Attrs) AddInt64(key string, value int64) *Attrs {
	return a.addAttr(slog.Int64(key, value))
}

func (a *Attrs) AddUint64(key string, value uint64) *Attrs {
	return a.addAttr(slog.Uint64(key, value))
}

func (a *Attrs) AddInt(key string, value int) *Attrs {
	return a.addAttr(slog.Int(key, value))
}

func (a *Attrs) AddFloat64(key string, value float64) *Attrs {
	return a.addAttr(slog.Float64(key, value))
}

func (a *Attrs) AddBool(key string, value bool) *Attrs {
	return a.addAttr(slog.Bool(key, value))
}

func (a *Attrs) AddTime(key string, value time.Time) *Attrs {
	return a.addAttr(slog.Time(key, value))
}

func (a *Attrs) AddDuration(key string, value time.Duration) *Attrs {
	return a.addAttr(slog.Duration(key, value))
}

func (a *Attrs) AddAny(key string, value any) *Attrs {
	return a.addAttr(slog.Any(key, value))
}

// AddContext adds log attributes (if there are any) from passed context.
func (a *Attrs) AddContext(ctx context.Context) *Attrs {
	if attr, ok := ctx.Value(logAttrCtxKey).(*Attrs); ok {
		return a.Add(attr)
	}
	return a
}

func (a *Attrs) AddGroup(key string, g *Attrs) *Attrs {
	return a.addAttr(slog.Group(key, g.asSlice()...))
}

// Context returns new context with log attributes and parent context.
// New context inherits log attributes (if any) from parent, but doesn't
// override any existing attributes.
func (a *Attrs) Context(ctx context.Context) context.Context {
	return context.WithValue(ctx, logAttrCtxKey, Context(ctx).Add(a))
}

func (a *Attrs) String() string {
	var parts []string
	for _, a := range a.attr {
		parts = append(parts, fmt.Sprintf("[%s: %s]", a.Key, a.Value))
	}
	return strings.Join(parts, " ")
}

func (a *Attrs) asSlice() []any {
	var ret []any
	for _, a := range a.attr {
		ret = append(ret, a)
	}
	return ret
}

func (a *Attrs) addAttr(attr slog.Attr) *Attrs {
	a.attr[attr.Key] = attr
	return a
}

func Slice(attr []*Attrs) []any {
	if len(attr) > 0 {
		return attr[0].asSlice()
	}
	return nil
}

func SliceContext(ctx context.Context, attr []*Attrs) []any {
	// log attrs from context have lower precedence than ones explicitly provided.
	attrsFromContext := Context(ctx)
	if len(attr) > 0 {
		for _, a := range attr[0].attr {
			attrsFromContext.addAttr(a)
		}
	}
	return attrsFromContext.asSlice()
}
