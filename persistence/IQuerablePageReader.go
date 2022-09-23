package persistence

import (
	"context"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

// IQuerablePageReader interface for data processing components that can query a page of data items.
//	Typed params:
//		- T any type
type IQuerablePageReader[T any] interface {

	// GetPageByQuery gets a page of data items using a query string.
	//	Parameters:
	//		- ctx context.Context	operation context
	//		- correlationId string transaction id to trace execution through call chain.
	//		- query string a query string
	//		- paging data.PagingParams paging parameters
	//		- sort  data.SortParams sort parameters
	//	Returns: data.DataPage[T], error receives list of items or error.
	GetPageByQuery(ctx context.Context, correlationId string,
		query string, paging cdata.PagingParams, sort cdata.SortParams) (page cdata.DataPage[T], err error)
}
