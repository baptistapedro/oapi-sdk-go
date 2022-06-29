// Package calendar code generated by oapi sdk gen
package calendar

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/core"
)

// 构建业务域服务实例
func NewService(httpClient *http.Client, config *core.Config) *CalendarService {
	c := &CalendarService{httpClient: httpClient, config: config}
	c.Calendar = &calendar{service: c}
	c.CalendarAcl = &calendarAcl{service: c}
	c.CalendarEvent = &calendarEvent{service: c}
	c.CalendarEventAttendee = &calendarEventAttendee{service: c}
	c.CalendarEventAttendeeChatMember = &calendarEventAttendeeChatMember{service: c}
	c.ExchangeBinding = &exchangeBinding{service: c}
	c.Freebusy = &freebusy{service: c}
	c.Setting = &setting{service: c}
	c.TimeoffEvent = &timeoffEvent{service: c}
	return c
}

// 业务域服务定义
type CalendarService struct {
	httpClient                      *http.Client
	config                          *core.Config
	Calendar                        *calendar
	CalendarAcl                     *calendarAcl
	CalendarEvent                   *calendarEvent
	CalendarEventAttendee           *calendarEventAttendee
	CalendarEventAttendeeChatMember *calendarEventAttendeeChatMember
	ExchangeBinding                 *exchangeBinding
	Freebusy                        *freebusy
	Setting                         *setting
	TimeoffEvent                    *timeoffEvent
}

// 资源服务定义
type calendar struct {
	service *CalendarService
}
type calendarAcl struct {
	service *CalendarService
}
type calendarEvent struct {
	service *CalendarService
}
type calendarEventAttendee struct {
	service *CalendarService
}
type calendarEventAttendeeChatMember struct {
	service *CalendarService
}
type exchangeBinding struct {
	service *CalendarService
}
type freebusy struct {
	service *CalendarService
}
type setting struct {
	service *CalendarService
}
type timeoffEvent struct {
	service *CalendarService
}

// 资源服务方法定义
func (c *calendar) Create(ctx context.Context, req *CreateCalendarReq, options ...core.RequestOptionFunc) (*CreateCalendarResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, c.service.config, http.MethodPost,
		"/open-apis/calendar/v4/calendars", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendar) Delete(ctx context.Context, req *DeleteCalendarReq, options ...core.RequestOptionFunc) (*DeleteCalendarResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, c.service.config, http.MethodDelete,
		"/open-apis/calendar/v4/calendars/:calendar_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendar) Get(ctx context.Context, req *GetCalendarReq, options ...core.RequestOptionFunc) (*GetCalendarResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, c.service.config, http.MethodGet,
		"/open-apis/calendar/v4/calendars/:calendar_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendar) List(ctx context.Context, req *ListCalendarReq, options ...core.RequestOptionFunc) (*ListCalendarResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, c.service.config, http.MethodGet,
		"/open-apis/calendar/v4/calendars", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendar) Patch(ctx context.Context, req *PatchCalendarReq, options ...core.RequestOptionFunc) (*PatchCalendarResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, c.service.config, http.MethodPatch,
		"/open-apis/calendar/v4/calendars/:calendar_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendar) Primary(ctx context.Context, req *PrimaryCalendarReq, options ...core.RequestOptionFunc) (*PrimaryCalendarResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, c.service.config, http.MethodPost,
		"/open-apis/calendar/v4/calendars/primary", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PrimaryCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendar) Search(ctx context.Context, req *SearchCalendarReq, options ...core.RequestOptionFunc) (*SearchCalendarResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, c.service.config, http.MethodPost,
		"/open-apis/calendar/v4/calendars/search", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SearchCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendar) SearchCalendar(ctx context.Context, req *SearchCalendarReq, options ...core.RequestOptionFunc) (*SearchCalendarIterator, error) {
	return &SearchCalendarIterator{
		ctx:      ctx,
		req:      req,
		listFunc: c.Search,
		options:  options,
		limit:    req.Limit}, nil
}
func (c *calendar) Subscribe(ctx context.Context, req *SubscribeCalendarReq, options ...core.RequestOptionFunc) (*SubscribeCalendarResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, c.service.config, http.MethodPost,
		"/open-apis/calendar/v4/calendars/:calendar_id/subscribe", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SubscribeCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendar) Subscription(ctx context.Context, options ...core.RequestOptionFunc) (*SubscriptionCalendarResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, c.service.config, http.MethodPost,
		"/open-apis/calendar/v4/calendars/subscription", []core.AccessTokenType{core.AccessTokenTypeUser}, nil, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SubscriptionCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendar) Unsubscribe(ctx context.Context, req *UnsubscribeCalendarReq, options ...core.RequestOptionFunc) (*UnsubscribeCalendarResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, c.service.config, http.MethodPost,
		"/open-apis/calendar/v4/calendars/:calendar_id/unsubscribe", []core.AccessTokenType{core.AccessTokenTypeUser, core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UnsubscribeCalendarResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarAcl) Create(ctx context.Context, req *CreateCalendarAclReq, options ...core.RequestOptionFunc) (*CreateCalendarAclResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, c.service.config, http.MethodPost,
		"/open-apis/calendar/v4/calendars/:calendar_id/acls", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateCalendarAclResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarAcl) Delete(ctx context.Context, req *DeleteCalendarAclReq, options ...core.RequestOptionFunc) (*DeleteCalendarAclResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, c.service.config, http.MethodDelete,
		"/open-apis/calendar/v4/calendars/:calendar_id/acls/:acl_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteCalendarAclResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarAcl) List(ctx context.Context, req *ListCalendarAclReq, options ...core.RequestOptionFunc) (*ListCalendarAclResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, c.service.config, http.MethodGet,
		"/open-apis/calendar/v4/calendars/:calendar_id/acls", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListCalendarAclResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarAcl) ListCalendarAcl(ctx context.Context, req *ListCalendarAclReq, options ...core.RequestOptionFunc) (*ListCalendarAclIterator, error) {
	return &ListCalendarAclIterator{
		ctx:      ctx,
		req:      req,
		listFunc: c.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (c *calendarAcl) Subscription(ctx context.Context, req *SubscriptionCalendarAclReq, options ...core.RequestOptionFunc) (*SubscriptionCalendarAclResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, c.service.config, http.MethodPost,
		"/open-apis/calendar/v4/calendars/:calendar_id/acls/subscription", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SubscriptionCalendarAclResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEvent) Create(ctx context.Context, req *CreateCalendarEventReq, options ...core.RequestOptionFunc) (*CreateCalendarEventResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, c.service.config, http.MethodPost,
		"/open-apis/calendar/v4/calendars/:calendar_id/events", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateCalendarEventResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEvent) Delete(ctx context.Context, req *DeleteCalendarEventReq, options ...core.RequestOptionFunc) (*DeleteCalendarEventResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, c.service.config, http.MethodDelete,
		"/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteCalendarEventResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEvent) Get(ctx context.Context, req *GetCalendarEventReq, options ...core.RequestOptionFunc) (*GetCalendarEventResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, c.service.config, http.MethodGet,
		"/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetCalendarEventResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEvent) List(ctx context.Context, req *ListCalendarEventReq, options ...core.RequestOptionFunc) (*ListCalendarEventResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, c.service.config, http.MethodGet,
		"/open-apis/calendar/v4/calendars/:calendar_id/events", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListCalendarEventResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEvent) Patch(ctx context.Context, req *PatchCalendarEventReq, options ...core.RequestOptionFunc) (*PatchCalendarEventResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, c.service.config, http.MethodPatch,
		"/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &PatchCalendarEventResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEvent) Search(ctx context.Context, req *SearchCalendarEventReq, options ...core.RequestOptionFunc) (*SearchCalendarEventResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, c.service.config, http.MethodPost,
		"/open-apis/calendar/v4/calendars/:calendar_id/events/search", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SearchCalendarEventResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEvent) SearchCalendarEvent(ctx context.Context, req *SearchCalendarEventReq, options ...core.RequestOptionFunc) (*SearchCalendarEventIterator, error) {
	return &SearchCalendarEventIterator{
		ctx:      ctx,
		req:      req,
		listFunc: c.Search,
		options:  options,
		limit:    req.Limit}, nil
}
func (c *calendarEvent) Subscription(ctx context.Context, req *SubscriptionCalendarEventReq, options ...core.RequestOptionFunc) (*SubscriptionCalendarEventResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, c.service.config, http.MethodPost,
		"/open-apis/calendar/v4/calendars/:calendar_id/events/subscription", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SubscriptionCalendarEventResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEventAttendee) BatchDelete(ctx context.Context, req *BatchDeleteCalendarEventAttendeeReq, options ...core.RequestOptionFunc) (*BatchDeleteCalendarEventAttendeeResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, c.service.config, http.MethodPost,
		"/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees/batch_delete", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &BatchDeleteCalendarEventAttendeeResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEventAttendee) Create(ctx context.Context, req *CreateCalendarEventAttendeeReq, options ...core.RequestOptionFunc) (*CreateCalendarEventAttendeeResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, c.service.config, http.MethodPost,
		"/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateCalendarEventAttendeeResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEventAttendee) List(ctx context.Context, req *ListCalendarEventAttendeeReq, options ...core.RequestOptionFunc) (*ListCalendarEventAttendeeResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, c.service.config, http.MethodGet,
		"/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListCalendarEventAttendeeResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEventAttendee) ListCalendarEventAttendee(ctx context.Context, req *ListCalendarEventAttendeeReq, options ...core.RequestOptionFunc) (*ListCalendarEventAttendeeIterator, error) {
	return &ListCalendarEventAttendeeIterator{
		ctx:      ctx,
		req:      req,
		listFunc: c.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (c *calendarEventAttendeeChatMember) List(ctx context.Context, req *ListCalendarEventAttendeeChatMemberReq, options ...core.RequestOptionFunc) (*ListCalendarEventAttendeeChatMemberResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, c.service.config, http.MethodGet,
		"/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees/:attendee_id/chat_members", []core.AccessTokenType{core.AccessTokenTypeTenant, core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListCalendarEventAttendeeChatMemberResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (c *calendarEventAttendeeChatMember) ListCalendarEventAttendeeChatMember(ctx context.Context, req *ListCalendarEventAttendeeChatMemberReq, options ...core.RequestOptionFunc) (*ListCalendarEventAttendeeChatMemberIterator, error) {
	return &ListCalendarEventAttendeeChatMemberIterator{
		ctx:      ctx,
		req:      req,
		listFunc: c.List,
		options:  options,
		limit:    req.Limit}, nil
}
func (e *exchangeBinding) Create(ctx context.Context, req *CreateExchangeBindingReq, options ...core.RequestOptionFunc) (*CreateExchangeBindingResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, e.service.config, http.MethodPost,
		"/open-apis/calendar/v4/exchange_bindings", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateExchangeBindingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *exchangeBinding) Delete(ctx context.Context, req *DeleteExchangeBindingReq, options ...core.RequestOptionFunc) (*DeleteExchangeBindingResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, e.service.config, http.MethodDelete,
		"/open-apis/calendar/v4/exchange_bindings/:exchange_binding_id", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteExchangeBindingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (e *exchangeBinding) Get(ctx context.Context, req *GetExchangeBindingReq, options ...core.RequestOptionFunc) (*GetExchangeBindingResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, e.service.config, http.MethodGet,
		"/open-apis/calendar/v4/exchange_bindings/:exchange_binding_id", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetExchangeBindingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (f *freebusy) List(ctx context.Context, req *ListFreebusyReq, options ...core.RequestOptionFunc) (*ListFreebusyResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, f.service.config, http.MethodPost,
		"/open-apis/calendar/v4/freebusy/list", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListFreebusyResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *setting) GenerateCaldavConf(ctx context.Context, req *GenerateCaldavConfSettingReq, options ...core.RequestOptionFunc) (*GenerateCaldavConfSettingResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, s.service.config, http.MethodPost,
		"/open-apis/calendar/v4/settings/generate_caldav_conf", []core.AccessTokenType{core.AccessTokenTypeUser}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GenerateCaldavConfSettingResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *timeoffEvent) Create(ctx context.Context, req *CreateTimeoffEventReq, options ...core.RequestOptionFunc) (*CreateTimeoffEventResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, t.service.config, http.MethodPost,
		"/open-apis/calendar/v4/timeoff_events", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateTimeoffEventResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (t *timeoffEvent) Delete(ctx context.Context, req *DeleteTimeoffEventReq, options ...core.RequestOptionFunc) (*DeleteTimeoffEventResp, error) {
	// 发起请求
	rawResp, err := core.SendRequest(ctx, t.service.config, http.MethodDelete,
		"/open-apis/calendar/v4/timeoff_events/:timeoff_event_id", []core.AccessTokenType{core.AccessTokenTypeTenant}, req, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteTimeoffEventResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
