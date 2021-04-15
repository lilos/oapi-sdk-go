package approval

import (
	"github.com/larksuite/oapi-sdk-go/api/core/tools"
	"github.com/larksuite/oapi-sdk-go/event/core/model"
)

type TaskEvent struct {
	*model.BaseEvent
	Event *TaskEventData `json:"event"`
}

type UserService struct {
	service *Service
}

type Instance struct {
	ApprovalCode    string         `json:"approval_code,omitempty"`
	ApprovalName    string         `json:"approval_name,omitempty"`
	StartTime       int            `json:"start_time,omitempty"`
	EndTime         int            `json:"end_time,omitempty"`
	UserId          string         `json:"user_id,omitempty"`
	OpenId          string         `json:"open_id,omitempty"`
	DepartmentId    string         `json:"department_id,omitempty"`
	Status          string         `json:"status,omitempty"`
	//Form            []*Form        `json:"form,omitempty"`
	//TaskList        []*TaskList    `json:"task_list,omitempty"`
	//CommentList     []*CommentList `json:"comment_list,omitempty"`
	//Timeline        []*Timeline    `json:"timeline,omitempty"`
	ForceSendFields []string       `json:"-"`
}

func (s *Instance) MarshalJSON() ([]byte, error) {
	type cp Instance
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type InstanceGetReqBody struct {
	InstanceCode    string   `json:"instance_code,omitempty"`
	Locale          string   `json:"locale,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *InstanceGetReqBody) MarshalJSON() ([]byte, error) {
	type cp InstanceGetReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type TaskEventData struct {
	AppId           string   `json:"app_id"`
	ApprovalCode    string   `json:"approval_code"`
	InstanceCode    string   `json:"instance_code"`
	OpenId          string   `json:"open_id"`
	OperateTime     string   `json:"operate_time"`
	Status          string   `json:"status"`
	TaskId          string   `json:"task_id"`
	TenantKey       string   `json:"tenant_key"`
	Type            string   `json:"type"`
	UserId          string   `json:"user_id"`
	ForceSendFields []string `json:"-"`
}

func (s *TaskEventData) MarshalJSON() ([]byte, error) {
	type cp TaskEventData
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}
