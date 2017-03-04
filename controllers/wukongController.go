package controllers

import (
	"net/http"
	"wukongServer/models"
	"wukongServer/wukong"
)

var (
	wk *wukong.Wukong = wukong.DefaultEngine()
)

func init() {
}

type WukongController struct {
	BaseController
}

func (wc *WukongController) Documents() {
	wc.AjaxSuccess(models.Document{}.Get())
}

func (wc *WukongController) Document() {

	switch wc.Ctx.Request.Method {
	case http.MethodGet:
	case http.MethodPost:
		content := wc.GetString(`content`)
		if len(content) == 0 {
			wc.AjaxError(`content length is 0`)
			return
		}

		doc := models.Document{Content: content}
		err := wk.AddIndexDocument(&doc)
		if err == nil {
			wc.AjaxSuccess(doc)
		} else {
			wc.AjaxError(err)
		}

	default:
		wc.AjaxError(`unknown http method`)
	}
}

func (wc *WukongController) Search() {

	searchText := wc.GetString(`kw`)

	s := wk.SearchText(searchText)

	wc.AjaxSuccess(s)
}
