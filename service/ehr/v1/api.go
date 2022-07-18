// Package ehr code generated by oapi sdk gen
package larkehr

import (
	"bytes"
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

// 构建业务域服务实例
func NewService(config *larkcore.Config) *EhrService {
	e := &EhrService{config: config}
	e.Attachment = &attachment{service: e}
	e.Employee = &employee{service: e}
	return e
}

// 业务域服务定义
type EhrService struct {
	config     *larkcore.Config
	Attachment *attachment
	Employee   *employee
}

// 资源服务定义
type attachment struct {
	service *EhrService
}
type employee struct {
	service *EhrService
}

// 资源服务方法定义
func (a *attachment) Get(ctx context.Context, req *GetAttachmentReq, options ...larkcore.RequestOptionFunc) (*GetAttachmentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/ehr/v1/attachments/:token"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetAttachmentResp{ApiResp: apiResp}
	// 如果是下载，则设置响应结果
	if apiResp.StatusCode == http.StatusOK {
		resp.File = bytes.NewBuffer(apiResp.RawBody)
		resp.FileName = larkcore.FileNameByHeader(apiResp.Header)
		return resp, err
	}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *employee) List(ctx context.Context, req *ListEmployeeReq, options ...larkcore.RequestOptionFunc) (*ListEmployeeResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/ehr/v1/employees"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, e.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListEmployeeResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *employee) ListByIterator(ctx context.Context, req *ListEmployeeReq, options ...larkcore.RequestOptionFunc) (*ListEmployeeIterator, error) {
	return &ListEmployeeIterator{
		ctx:      ctx,
		req:      req,
		listFunc: e.List,
		options:  options,
		limit:    req.Limit}, nil
}
