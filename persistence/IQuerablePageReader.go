package persistence

import (
	"context"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

// IQuerablePageReader interface for data processing components that can query a page of data items.
//	Typed params:
//		- T any type of getting element
type IQuerablePageReader[T data.ICloneable[T]] interface {

	// GetPageByQuery gets a page of data items using a query string.
	//	Parameters:
	//		- ctx context.Context
	//		- correlationId string transaction id to trace execution through call chain.
	//		- query string a query string
	//		- paging data.PagingParams paging parameters
	//		- sort  data.SortParams sort parameters
	//	Returns: *data.DataPage[T], error receives list of items or error.
	GetPageByQuery(ctx context.Context, correlationId string,
		query string, paging data.PagingParams, sort data.SortParams) (page data.DataPage[T], err error)
}
