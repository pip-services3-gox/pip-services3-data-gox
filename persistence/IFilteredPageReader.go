package persistence

import (
	"context"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

// IFilteredPageReader is interface for data processing components
// that can retrieve a page of data items by a filter.
//	Typed params:
//		- T any type of getting element
type IFilteredPageReader[T data.ICloneable[T]] interface {

	// GetPageByFilter gets a page of data items using filter
	//	Parameters:
	//		- ctx context.Context
	//		- correlationId transaction id to trace execution through call chain.
	//		- filter  data.FilterParams filter parameters
	//		- paging data.PagingParams paging parameters
	//		- sort data.SortParams sort parameters
	//	Returns: *data.DataPage[T], error list of items or error.
	GetPageByFilter(ctx context.Context, correlationId string,
		filter data.FilterParams, paging data.PagingParams, sort data.SortParams) (page data.DataPage[T], err error)
}
