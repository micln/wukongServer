package wukong

import "github.com/huichen/wukong/types"

type SearchResults struct {
	types.SearchResponse
	Documents map[uint64]string
}
