package approval

import (
	"github.com/larksuite/oapi-sdk-go/api"
	"github.com/larksuite/oapi-sdk-go/api/core/request"
	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/core/config"
)

type Service struct {
	conf                *config.Config
	Instances           *InstanceService
	SubscriptionService *SubscriptionService
}

func NewService(conf *config.Config) *Service {
	s := &Service{
		conf: conf,
	}
	s.Instances = newInstanceService(s)
	s.SubscriptionService = newSubscriptionService(s)
	return s
}

type InstanceService struct {
	service *Service
}

type SubscriptionService struct {
	service *Service
}

func newInstanceService(service *Service) *InstanceService {
	return &InstanceService{
		service: service,
	}
}

func newSubscriptionService(service *Service) *SubscriptionService {
	return &SubscriptionService{
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

type SubscriptionSubscribeReqCall struct {
	ctx           *core.Context
	subscriptions *SubscriptionService
	body          *SubscriptionSubscribeReqBody
	queryParams   map[string]interface{}
	optFns        []request.OptFn
}

func (rc *SubscriptionSubscribeReqCall) Do() (*SubscriptionSubscribeReqBody, error) {
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &SubscriptionSubscribeReqBody{}
	req := request.NewRequest("approval/openapi/v2/subscription/subscribe", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.subscriptions.service.conf, req)
	return result, err
}

func (Subscriptions *SubscriptionService) Subscribe(ctx *core.Context, body *SubscriptionSubscribeReqBody, optFns ...request.OptFn) *SubscriptionSubscribeReqCall {
	return &SubscriptionSubscribeReqCall{
		ctx:           ctx,
		subscriptions: Subscriptions,
		body:          body,
		queryParams:   map[string]interface{}{},
		optFns:        optFns,
	}
}
