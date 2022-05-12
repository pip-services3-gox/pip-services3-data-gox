package persistence

import (
	"context"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

// IFilteredPageReader is interface for data processing components
// that can retrieve a page of data items by a filter.
//	Typed params:
//		- T cdata.ICloneable[T] any type that implemented
//			ICloneable interface of getting element
type IFilteredPageReader[T cdata.ICloneable[T]] interface {

	// GetPageByFilter gets a page of data items using filter
	//	Parameters:
	//		- ctx context.Context
	//		- correlationId transaction id to trace execution through call chain.
	//		- filter  data.FilterParams filter parameters
	//		- paging data.PagingParams paging parameters
	//		- sort data.SortParams sort parameters
	//	Returns: data.DataPage[T], error list of items or error.
	GetPageByFilter(ctx context.Context, correlationId string,
		filter cdata.FilterParams, paging cdata.PagingParams, sort cdata.SortParams) (page cdata.DataPage[T], err error)
}
