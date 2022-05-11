package persistence

import (
	"context"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

// IFilteredReader interface for data processing components that can
// retrieve a list of data items by filter.
//	Typed params:
//		- T any type of getting element
type IFilteredReader[T data.ICloneable[T]] interface {

	// GetListByFilter gets a list of data items using filter
	//	Parameters:
	//		- correlationId string transaction id to trace execution through call chain.
	//		- filter data.FilterParams filter parameters
	//		- sort  data.SortParams sort parameters
	//	Returns: []T, error receives list of items or error.
	GetListByFilter(ctx context.Context, correlationId string,
		filter data.FilterParams, sort data.SortParams) (items []T, err error)
}
