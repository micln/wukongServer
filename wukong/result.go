package wukong

import (
	"wukongServer/models"

	"github.com/huichen/wukong/types"
)

type SearchResults struct {
	types.SearchResponse
	Documents map[uint64]*models.Document
}
