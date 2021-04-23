package approval

import (
	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/core/config"
	"github.com/larksuite/oapi-sdk-go/event"
)

type TaskEventHandler struct {
	Fn func(*core.Context, *TaskEvent) error
}

func (h *TaskEventHandler) GetEvent() interface{} {
	return &TaskEvent{}
}

func (h *TaskEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*TaskEvent))
}

func SetTaskEventHandler(conf *config.Config, fn func(ctx *core.Context, event *TaskEvent) error) {
	event.SetTypeHandler(conf, "approval_task", &TaskEventHandler{Fn: fn})
}
