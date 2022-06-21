package persistence

import (
	"context"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

// IdentifiableFilePersistence is an abstract persistence component that stores data in flat files
// and implements a number of CRUD operations over data items with unique ids.
// The data items must implement IDataObject interface
//
// In basic scenarios child classes shall only override GetPageByFilter,
// GetListByFilter or DeleteByFilter operations with specific filter function.
// All other operations can be used out of the box.
//
// In complex scenarios child classes can implement additional operations by
// accessing cached items via IdentifiableFilePersistence._items property and calling Save method
// on updates.
//	Important:
//		- this component is a thread save!
//		- the data items must implement IDataObject interface
//
//	see JsonFilePersister
//	see MemoryPersistence
//
//	Configuration parameters:
//		- path: path to the file where data is stored
//		- options:
//		- max_page_size: Maximum number of items returned in a single page (default: 100)
//	References:
//		- *:logger:*:*:1.0 (optional)  ILogger components to pass log messages
//	Typed params:
//		- T cdata.IDataObject[T, K] any type that implemented
//			IDataObject interface of getting element
//		- K any type if id (key)
//	Examples:
//		type MyFilePersistence struct {
//			IdentifiableFilePersistence[MyData, MyId]
//		}
//      func NewMyFilePersistence(path string)(mfp *MyFilePersistence) {
//  		mfp = MyFilePersistence{}
//  		mfp.IdentifiableFilePersistence = *NewJsonPersister[MyData](path)
//  		return mfp
//      }
//		func (c *MyFilePersistence) composeFilter(filter cdata.FilterParams)(func (item interface{})bool) {
//			if &filter == nil {
//				filter = NewFilterParams()
//			}
//			name := filter.GetAsNullableString("name");
//			return func (item MyData) bool {
//				if name != "" && dummy.Name != name {
//  				return false
//				}
//				return true
//			}
//		}
//		func (c *MyFilePersistence) GetPageByFilter(ctx context.Context, correlationId string,
//			filter FilterParams, paging PagingParams) (page cdata.DataPage[MyData], err error){
//  			return c.GetPageByFilter(correlationId, c.composeFilter(filter), paging, nil, nil)
//		}
//
//		persistence := NewMyFilePersistence("./data/data.json")
//
//		_, err := persistence.Create(context.Background(), "123", { Id: "1", Name: "ABC" })
//		if (err != nil) {
//  		panic()
//		}
//      page, err := persistence.GetPageByFilter(context.Background(), "123", *NewFilterParamsFromTuples("Name", "ABC"), nil)
//      if err != nil {
//  		panic("Error")
//  	}
//		_data, _ := page.Data()
//      fmt.Println(_data)   // Result: { Id: "1", Name: "ABC" )
type IdentifiableFilePersistence[T any, K any] struct {
	IdentifiableMemoryPersistence[T, K]
	Persister *JsonFilePersister[T]
}

// NewIdentifiableFilePersistence creates a new instance of the persistence.
//	Typed params:
//		- T cdata.IDataObject[T, K] any type that implemented
//			IDataObject interface of getting element
//		- K any type if id (key)
//	Parameters:
//		- persister (optional) a persister component that loads and saves data from/to flat file.
//	Returns: *IdentifiableFilePersistence pointer on new IdentifiableFilePersistence
func NewIdentifiableFilePersistence[T any, K any](persister *JsonFilePersister[T]) *IdentifiableFilePersistence[T, K] {
	c := &IdentifiableFilePersistence[T, K]{}
	if persister == nil {
		persister = NewJsonFilePersister[T]("")
	}
	c.IdentifiableMemoryPersistence = *NewIdentifiableMemoryPersistence[T, K]()
	c.Loader = persister
	c.Saver = persister
	c.Persister = persister
	return c
}

// Configure component by passing configuration parameters.
//	Parameters:
//		- ctx context.Context
//		- config *config.ConfigParams configuration parameters to be set.
func (c *IdentifiableFilePersistence[T, K]) Configure(ctx context.Context, config *config.ConfigParams) {
	c.Configure(ctx, config)
	c.Persister.Configure(ctx, config)
}
