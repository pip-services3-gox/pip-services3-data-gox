package persistence

import (
	"context"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

// IFilteredReader interface for data processing components that can
// retrieve a list of data items by filter.
//	Typed params:
//		- T any type
type IFilteredReader[T any] interface {

	// GetListByFilter gets a list of data items using filter
	//	Parameters:
	//		- ctx context.Context	operation context
	//		- correlationId string transaction id to trace execution through call chain.
	//		- filter data.FilterParams filter parameters
	//		- sort  data.SortParams sort parameters
	//	Returns: []T, error receives list of items or error.
	GetListByFilter(ctx context.Context, correlationId string,
		filter cdata.FilterParams, sort cdata.SortParams) (items []T, err error)
}
