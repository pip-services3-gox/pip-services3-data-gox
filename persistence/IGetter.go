package persistence

import "context"

// IGetter Interface for data processing components that can get data items.
//	Typed params:
//		- T IDataObject[T, K] any type that implemented
//			IDataObject interface of getting element
//		- K any type of id (key)
type IGetter[T IDataObject[T, K], K any] interface {

	// GetOneById a data items by its unique id.
	//	Parameters:
	//		- ctx context.Context
	//		- correlationId (optional) transaction id to trace execution through call chain.
	//		- id an id of item to be retrieved.
	//	Returns: T, error item or error
	GetOneById(ctx context.Context, correlationId string, id K) (item T, err error)
}
