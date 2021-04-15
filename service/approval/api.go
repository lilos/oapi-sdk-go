package approval

import (
	"github.com/larksuite/oapi-sdk-go/api"
	"github.com/larksuite/oapi-sdk-go/api/core/request"
	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/core/config"
)

type Service struct {
	conf              *config.Config
	Instances         *InstanceService
}

func NewService(conf *config.Config) *Service {
	s := &Service{
		conf: conf,
	}
	s.Instances = newInstanceService(s)
	return s
}

type InstanceService struct {
	service *Service
}

func newInstanceService(service *Service) *InstanceService {
	return &InstanceService{
		service: service,
	}
}

type InstanceGetReqCall struct {
	ctx         *core.Context
	Instances   *InstanceService
	body        *InstanceGetReqBody
	queryParams map[string]interface{}
	optFns      []request.OptFn
}

func (rc *InstanceGetReqCall) Do() (*Instance, error) {
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &Instance{}
	req := request.NewRequest("approval/openapi/v2/instance/get", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.Instances.service.conf, req)
	return result, err
}

func (Instances *InstanceService) Get(ctx *core.Context, body *InstanceGetReqBody, optFns ...request.OptFn) *InstanceGetReqCall {
	return &InstanceGetReqCall{
		ctx:         ctx,
		Instances:   Instances,
		body:        body,
		queryParams: map[string]interface{}{},
		optFns:      optFns,
	}
}
