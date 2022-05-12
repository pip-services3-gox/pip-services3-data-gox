package persistence

import (
	"context"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

// ISetter interface for data processing components that can set (create or update) data items.
//	Typed params:
//		- T cdata.ICloneable[T] any type that implemented
//			ICloneable interface of getting element
//		- K comparable type of id (key)
type ISetter[T cdata.ICloneable[T], K any] interface {

	// Set a data item. If the data item exists it updates it,
	// otherwise it create a new data item.
	//	Parameters:
	//		- ctx context.Context
	//		- correlationId string transaction id to trace execution through call chain.
	//		- item T is an item to be set.
	//	Returns: T, error updated item or error.
	Set(ctx context.Context, correlationId string, item T) (value T, err error)
}
