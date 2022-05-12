package persistence

import (
	"context"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

// IQuerableReader interface for data processing components that can query a list of data items.
//	Typed params:
//		- T cdata.ICloneable[T] any type that implemented
//			ICloneable interface of getting element
type IQuerableReader[T cdata.ICloneable[T]] interface {

	// GetListByQuery gets a list of data items using a query string.
	//	Parameters:
	//		- ctx context.Context
	//		- correlationId  string transaction id to trace execution through call chain.
	//		- query string a query string
	//		- sort data.SortParams sort parameters
	// Returns []T, error list of items or error.
	GetListByQuery(ctx context.Context, correlationId string,
		query string, sort cdata.SortParams) (items []T, err error)
}
