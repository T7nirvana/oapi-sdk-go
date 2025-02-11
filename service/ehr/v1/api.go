// Code generated by lark suite oapi sdk gen
package v1

import (
	"github.com/larksuite/oapi-sdk-go/api"
	"github.com/larksuite/oapi-sdk-go/api/core/request"
	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/core/config"
	"io"
)

type Service struct {
	conf        *config.Config
	Attachments *AttachmentService
	Employees   *EmployeeService
}

func NewService(conf *config.Config) *Service {
	s := &Service{
		conf: conf,
	}
	s.Attachments = newAttachmentService(s)
	s.Employees = newEmployeeService(s)
	return s
}

type AttachmentService struct {
	service *Service
}

func newAttachmentService(service *Service) *AttachmentService {
	return &AttachmentService{
		service: service,
	}
}

type EmployeeService struct {
	service *Service
}

func newEmployeeService(service *Service) *EmployeeService {
	return &EmployeeService{
		service: service,
	}
}

type AttachmentGetReqCall struct {
	ctx         *core.Context
	attachments *AttachmentService
	pathParams  map[string]interface{}
	optFns      []request.OptFn
	result      io.Writer
}

func (rc *AttachmentGetReqCall) SetToken(token string) {
	rc.pathParams["token"] = token
}
func (rc *AttachmentGetReqCall) SetResponseStream(result io.Writer) {
	rc.result = result
}

func (rc *AttachmentGetReqCall) Do() (io.Writer, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetResponseStream())
	req := request.NewRequest("/open-apis/ehr/v1/attachments/:token", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, nil, rc.result, rc.optFns...)
	err := api.Send(rc.ctx, rc.attachments.service.conf, req)
	return rc.result, err
}

func (attachments *AttachmentService) Get(ctx *core.Context, optFns ...request.OptFn) *AttachmentGetReqCall {
	return &AttachmentGetReqCall{
		ctx:         ctx,
		attachments: attachments,
		pathParams:  map[string]interface{}{},
		optFns:      optFns,
	}
}

type EmployeeListReqCall struct {
	ctx         *core.Context
	employees   *EmployeeService
	queryParams map[string]interface{}
	optFns      []request.OptFn
}

func (rc *EmployeeListReqCall) SetView(view string) {
	rc.queryParams["view"] = view
}
func (rc *EmployeeListReqCall) SetStatus(status ...int) {
	rc.queryParams["status"] = status
}
func (rc *EmployeeListReqCall) SetType(type_ ...int) {
	rc.queryParams["type"] = type_
}
func (rc *EmployeeListReqCall) SetStartTime(startTime int64) {
	rc.queryParams["start_time"] = startTime
}
func (rc *EmployeeListReqCall) SetEndTime(endTime int64) {
	rc.queryParams["end_time"] = endTime
}
func (rc *EmployeeListReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}
func (rc *EmployeeListReqCall) SetUserIds(userIds string) {
	rc.queryParams["user_ids"] = userIds
}
func (rc *EmployeeListReqCall) SetPageToken(pageToken string) {
	rc.queryParams["page_token"] = pageToken
}
func (rc *EmployeeListReqCall) SetPageSize(pageSize int) {
	rc.queryParams["page_size"] = pageSize
}

func (rc *EmployeeListReqCall) Do() (*EmployeeListResult, error) {
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &EmployeeListResult{}
	req := request.NewRequest("/open-apis/ehr/v1/employees", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.employees.service.conf, req)
	return result, err
}

func (employees *EmployeeService) List(ctx *core.Context, optFns ...request.OptFn) *EmployeeListReqCall {
	return &EmployeeListReqCall{
		ctx:         ctx,
		employees:   employees,
		queryParams: map[string]interface{}{},
		optFns:      optFns,
	}
}
