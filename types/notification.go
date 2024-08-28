package types

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Notification struct {
	Level   string // info | warn | error | success
	Message string
}

func NewNotice(ctx context.Context, level, message string) {
	runtime.EventsEmit(ctx, "notification", Notification{Level: level, Message: message})
}
