// Package approval code generated by oapi sdk gen
package larkapproval

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

// 构建业务域服务实例
func NewService(config *larkcore.Config) *ApprovalService {
	a := &ApprovalService{config: config}
	a.Approval = &approval{service: a}
	a.ExternalApproval = &externalApproval{service: a}
	a.ExternalInstance = &externalInstance{service: a}
	a.ExternalTask = &externalTask{service: a}
	a.Instance = &instance{service: a}
	a.InstanceComment = &instanceComment{service: a}
	a.Task = &task{service: a}
	return a
}

// 业务域服务定义
type ApprovalService struct {
	config           *larkcore.Config
	Approval         *approval
	ExternalApproval *externalApproval
	ExternalInstance *externalInstance
	ExternalTask     *externalTask
	Instance         *instance
	InstanceComment  *instanceComment
	Task             *task
}

// 资源服务定义
type approval struct {
	service *ApprovalService
}
type externalApproval struct {
	service *ApprovalService
}
type externalInstance struct {
	service *ApprovalService
}
type externalTask struct {
	service *ApprovalService
}
type instance struct {
	service *ApprovalService
}
type instanceComment struct {
	service *ApprovalService
}
type task struct {
	service *ApprovalService
}

// 资源服务方法定义
func (a *approval) Create(ctx context.Context, req *CreateApprovalReq, options ...larkcore.RequestOptionFunc) (*CreateApprovalResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/approvals"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateApprovalResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (a *approval) Get(ctx context.Context, req *GetApprovalReq, options ...larkcore.RequestOptionFunc) (*GetApprovalResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/approvals/:approval_code"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetApprovalResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (a *approval) Subscribe(ctx context.Context, req *SubscribeApprovalReq, options ...larkcore.RequestOptionFunc) (*SubscribeApprovalResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/approvals/:approval_code/subscribe"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SubscribeApprovalResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (a *approval) Unsubscribe(ctx context.Context, req *UnsubscribeApprovalReq, options ...larkcore.RequestOptionFunc) (*UnsubscribeApprovalResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/approvals/:approval_code/unsubscribe"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UnsubscribeApprovalResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *externalApproval) Create(ctx context.Context, req *CreateExternalApprovalReq, options ...larkcore.RequestOptionFunc) (*CreateExternalApprovalResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/external_approvals"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, e.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateExternalApprovalResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *externalInstance) Check(ctx context.Context, req *CheckExternalInstanceReq, options ...larkcore.RequestOptionFunc) (*CheckExternalInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/external_instances/check"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, e.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CheckExternalInstanceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *externalInstance) Create(ctx context.Context, req *CreateExternalInstanceReq, options ...larkcore.RequestOptionFunc) (*CreateExternalInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/external_instances"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, e.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateExternalInstanceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *externalTask) List(ctx context.Context, req *ListExternalTaskReq, options ...larkcore.RequestOptionFunc) (*ListExternalTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/external_tasks"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, e.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListExternalTaskResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *externalTask) ListByIterator(ctx context.Context, req *ListExternalTaskReq, options ...larkcore.RequestOptionFunc) (*ListExternalTaskIterator, error) {
	return &ListExternalTaskIterator{
		ctx:      ctx,
		req:      req,
		listFunc: e.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (i *instance) AddSign(ctx context.Context, req *AddSignInstanceReq, options ...larkcore.RequestOptionFunc) (*AddSignInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/instances/add_sign"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &AddSignInstanceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (i *instance) Cancel(ctx context.Context, req *CancelInstanceReq, options ...larkcore.RequestOptionFunc) (*CancelInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/instances/cancel"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CancelInstanceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (i *instance) Cc(ctx context.Context, req *CcInstanceReq, options ...larkcore.RequestOptionFunc) (*CcInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/instances/cc"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CcInstanceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (i *instance) Create(ctx context.Context, req *CreateInstanceReq, options ...larkcore.RequestOptionFunc) (*CreateInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/instances"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateInstanceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (i *instance) Get(ctx context.Context, req *GetInstanceReq, options ...larkcore.RequestOptionFunc) (*GetInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/instances/:instance_id"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetInstanceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (i *instance) List(ctx context.Context, req *ListInstanceReq, options ...larkcore.RequestOptionFunc) (*ListInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/instances"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListInstanceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (i *instance) ListByIterator(ctx context.Context, req *ListInstanceReq, options ...larkcore.RequestOptionFunc) (*ListInstanceIterator, error) {
	return &ListInstanceIterator{
		ctx:      ctx,
		req:      req,
		listFunc: i.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (i *instance) Preview(ctx context.Context, req *PreviewInstanceReq, options ...larkcore.RequestOptionFunc) (*PreviewInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/instances/preview"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PreviewInstanceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (i *instance) Query(ctx context.Context, req *QueryInstanceReq, options ...larkcore.RequestOptionFunc) (*QueryInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/instances/query"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryInstanceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (i *instance) SearchCc(ctx context.Context, req *SearchCcInstanceReq, options ...larkcore.RequestOptionFunc) (*SearchCcInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/instances/search_cc"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SearchCcInstanceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (i *instance) SpecifiedRollback(ctx context.Context, req *SpecifiedRollbackInstanceReq, options ...larkcore.RequestOptionFunc) (*SpecifiedRollbackInstanceResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/instances/specified_rollback"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SpecifiedRollbackInstanceResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (i *instanceComment) Create(ctx context.Context, req *CreateInstanceCommentReq, options ...larkcore.RequestOptionFunc) (*CreateInstanceCommentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/instances/:instance_id/comments"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateInstanceCommentResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (i *instanceComment) Delete(ctx context.Context, req *DeleteInstanceCommentReq, options ...larkcore.RequestOptionFunc) (*DeleteInstanceCommentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/instances/:instance_id/comments/:comment_id"
	apiReq.HttpMethod = http.MethodDelete
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteInstanceCommentResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (i *instanceComment) List(ctx context.Context, req *ListInstanceCommentReq, options ...larkcore.RequestOptionFunc) (*ListInstanceCommentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/instances/:instance_id/comments"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListInstanceCommentResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (i *instanceComment) Remove(ctx context.Context, req *RemoveInstanceCommentReq, options ...larkcore.RequestOptionFunc) (*RemoveInstanceCommentResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/instances/:instance_id/comments/remove"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, i.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &RemoveInstanceCommentResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *task) Approve(ctx context.Context, req *ApproveTaskReq, options ...larkcore.RequestOptionFunc) (*ApproveTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/tasks/approve"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ApproveTaskResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *task) Query(ctx context.Context, req *QueryTaskReq, options ...larkcore.RequestOptionFunc) (*QueryTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/tasks/query"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant, larkcore.AccessTokenTypeUser}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryTaskResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *task) QueryByIterator(ctx context.Context, req *QueryTaskReq, options ...larkcore.RequestOptionFunc) (*QueryTaskIterator, error) {
	return &QueryTaskIterator{
		ctx:      ctx,
		req:      req,
		listFunc: t.Query,
		options:  options,
		limit:    req.Limit}, nil
}
func (t *task) Reject(ctx context.Context, req *RejectTaskReq, options ...larkcore.RequestOptionFunc) (*RejectTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/tasks/reject"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &RejectTaskResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *task) Search(ctx context.Context, req *SearchTaskReq, options ...larkcore.RequestOptionFunc) (*SearchTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/tasks/search"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SearchTaskResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *task) Transfer(ctx context.Context, req *TransferTaskReq, options ...larkcore.RequestOptionFunc) (*TransferTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/approval/v4/tasks/transfer"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &TransferTaskResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
