package wukong

import (
	"wukongServer/models"

	"github.com/huichen/wukong/engine"
	"github.com/huichen/wukong/types"
	. "github.com/micln/go-utils"
)

type Wukong struct {
	engine.Engine

	numDocument uint64
	documents   map[uint64]string
}

var (
	wukong *Wukong
)

func NewWukong() *Wukong {
	wk := &Wukong{
		Engine:    engine.Engine{},
		documents: make(map[uint64]string),
	}

	wk.Engine.Init(types.EngineInitOptions{
		SegmenterDictionaries: "dictionary.simple.txt",
	})

	//	load from db
	go func() {
		docs := models.Document{}.Get()
		for i := range docs {
			wk.IndexDocument(docs[i].Id, types.DocumentIndexData{
				Content: docs[i].Content,
			}, false)
		}
		wk.FlushIndex()
	}()

	return wk
}

func init() {
	//wukong = NewWukong()
}

func (wk *Wukong) AddIndexDocument(document string) (doc *models.Document, err error) {

	doc = &models.Document{Content: document}

	err = doc.Save()
	if err != nil {
		return
	}

	docId := doc.Id
	wk.documents[docId] = document

	wk.IndexDocument(docId, types.DocumentIndexData{
		Content: document,
	}, false)

	go wk.FlushIndex()

	return
}

func (wk *Wukong) nextId() uint64 {
	wk.numDocument += 1
	return wk.numDocument
}

func (wk *Wukong) SearchText(text string) SearchResults {

	results := SearchResults{
		SearchResponse: wk.Search(types.SearchRequest{Text: text}),
		Documents:      make(map[uint64]string),
	}

	for idx := range results.Docs {
		docId := results.Docs[idx].DocId
		results.Documents[docId] = wk.documents[docId]
	}

	return results
}

func (wk *Wukong) ToJson() string {
	return JsonEncode(wk.Engine)
}

func OnClose() {
	wukong.Engine.Close()
}

func intsMin(args ...int) (min int) {
	min = 1<<31 - 1
	for i := range args {
		if args[i] < min {
			min = args[i]
		}
	}
	return
}
