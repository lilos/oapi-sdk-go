// Code generated by lark suite oapi sdk gen
package v1

import (
	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/core/config"
	"github.com/larksuite/oapi-sdk-go/event"
)

type TaskUpdatedEventHandler struct {
	Fn func(*core.Context, *TaskUpdatedEvent) error
}

func (h *TaskUpdatedEventHandler) GetEvent() interface{} {
	return &TaskUpdatedEvent{}
}

func (h *TaskUpdatedEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*TaskUpdatedEvent))
}

func SetTaskUpdatedEventHandler(conf *config.Config, fn func(ctx *core.Context, event *TaskUpdatedEvent) error) {
	event.SetTypeHandler(conf, "task.task.updated_v1", &TaskUpdatedEventHandler{Fn: fn})
}

type TaskCommentUpdatedEventHandler struct {
	Fn func(*core.Context, *TaskCommentUpdatedEvent) error
}

func (h *TaskCommentUpdatedEventHandler) GetEvent() interface{} {
	return &TaskCommentUpdatedEvent{}
}

func (h *TaskCommentUpdatedEventHandler) Handle(ctx *core.Context, event interface{}) error {
	return h.Fn(ctx, event.(*TaskCommentUpdatedEvent))
}

func SetTaskCommentUpdatedEventHandler(conf *config.Config, fn func(ctx *core.Context, event *TaskCommentUpdatedEvent) error) {
	event.SetTypeHandler(conf, "task.task.comment.updated_v1", &TaskCommentUpdatedEventHandler{Fn: fn})
}
