package persistence

import (
	"context"
)

// ISetter interface for data processing components that can set (create or update) data items.
//	Typed params:
//		- T any type
//		- K type of id (key)
type ISetter[T any, K any] interface {

	// Set a data item. If the data item exists it updates it,
	// otherwise it create a new data item.
	//	Parameters:
	//		- ctx context.Context
	//		- correlationId string transaction id to trace execution through call chain.
	//		- item T is an item to be set.
	//	Returns: T, error updated item or error.
	Set(ctx context.Context, correlationId string, item T) (value T, err error)
}
