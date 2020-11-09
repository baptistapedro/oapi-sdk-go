package http

import (
	"github.com/larksuite/oapi-sdk-go/card"
	"github.com/larksuite/oapi-sdk-go/core/config"
	coremodel "github.com/larksuite/oapi-sdk-go/core/model"
	"net/http"
)

func Handle(conf *config.Config, request *http.Request, response http.ResponseWriter) {
	req, err := coremodel.ToOapiRequest(request)
	if err != nil {
		err = coremodel.NewOapiResponseOfErr(err).WriteTo(response)
		if err != nil {
			conf.GetLogger().Error(req.Ctx, err)
		}
		return
	}
	err = card.Handle(conf, req).WriteTo(response)
	if err != nil {
		conf.GetLogger().Error(req.Ctx, err)
	}
}
