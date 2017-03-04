package controllers

import (
	"net/http"
	"wukongServer/models"
	"wukongServer/wukong"
)

var (
	wk *wukong.Wukong
)

func init() {
	wk = wukong.NewWukong()

	//wk.AddIndexDocument(`此次百度收购将成中国互联网最大并购`)
	//wk.AddIndexDocument(`百度宣布拟全资收购91无线业务`)
	//wk.AddIndexDocument(`百度是中国最大的搜索引擎`)
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

		doc, err := wk.AddIndexDocument(content)
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
