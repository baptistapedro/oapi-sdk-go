// Package gray_test_open_sg code generated by oapi sdk gen
package larkgray_test_open_sg

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

// 构建业务域服务实例
func NewService(config *larkcore.Config) *GrayTestOpenSgService {
	g := &GrayTestOpenSgService{config: config}
	g.Moto = &moto{service: g}
	return g
}

// 业务域服务定义
type GrayTestOpenSgService struct {
	config *larkcore.Config
	Moto   *moto
}

// 资源服务定义
type moto struct {
	service *GrayTestOpenSgService
}

// 资源服务方法定义
func (m *moto) Create(ctx context.Context, req *CreateMotoReq, options ...larkcore.RequestOptionFunc) (*CreateMotoResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/gray_test_open_sg/v1/motos"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, m.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateMotoResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *moto) Get(ctx context.Context, req *GetMotoReq, options ...larkcore.RequestOptionFunc) (*GetMotoResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/gray_test_open_sg/v1/motos/:moto_id"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, m.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetMotoResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *moto) List(ctx context.Context, req *ListMotoReq, options ...larkcore.RequestOptionFunc) (*ListMotoResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/gray_test_open_sg/v1/motos"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, m.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListMotoResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (m *moto) ListByIterator(ctx context.Context, req *ListMotoReq, options ...larkcore.RequestOptionFunc) (*ListMotoIterator, error) {
	return &ListMotoIterator{
		ctx:      ctx,
		req:      req,
		listFunc: m.List,
		options:  options,
		limit:    req.Limit}, nil
}
