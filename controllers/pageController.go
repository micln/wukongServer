package controllers

import (
	"net/http"
	. "wukongServer/models"
)

type PageController struct {
	BaseController
}

func (c *PageController) Search() {
	kw := c.GetString(`kw`)

	s := wk.SearchText(kw)

	c.View(`page/search.html`, map[string]interface{}{
		`searchRequest`: map[string]interface{}{
			`kw`: kw,
		},
		`searchResult`: s,
	})

	return

}

func (c *PageController) Documents() {

	c.View(`page/documents.html`, map[string]interface{}{
		`documents`: Document{}.Get(),
	})
}

func (c *PageController) Document() {
	switch c.Ctx.Request.Method {
	case http.MethodPost:
		title := c.GetString(`title`)
		content := c.GetString(`content`)
		url := c.GetString(`url`)

		if len(content) < 5 {
			c.Error(`content length shouldn't less then 5`)
			return
		}

		if len(title) == 0 {
			title = content[:8]
		}

		doc := Document{
			Title:   title,
			Content: content,
			Url:     url,
		}
		err := wk.AddIndexDocument(&doc)
		if err == nil {
			c.Redirect(c.Ctx.Request.Header.Get(`referer`), 301)
		} else {
			c.Error(err.Error())
		}

	default:
		c.Error(`unknown http method`)
	}
}
