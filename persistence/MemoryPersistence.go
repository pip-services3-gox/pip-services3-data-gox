package persistence

import (
	"context"
	"math/rand"
	"sort"
	"sync"
	"time"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	"github.com/pip-services3-gox/pip-services3-components-gox/log"
)

// TODO:: fix comments after complete

// MemoryPersistence abstract persistence component that stores data in memory.
//
//	This is the most basic persistence component that is only
//	able to store data items of any type. Specific CRUD operations
//	over the data items must be implemented in child struct by
//	accessing Items property and calling Save method.
//
//	The component supports loading and saving items from another data source.
//	That allows to use it as a base struct for file and other types
//	of persistence components that cache all data in memory.
//
//	References:
//		*:logger:*:*:1.0    ILogger components to pass log messages
//	Example:
//		type MyMemoryPersistence struct {
//			MemoryPersistence
//		}
//		func (c * MyMemoryPersistence) GetByName(correlationId string, name string)(item interface{}, err error) {
//			for _, v := range c.Items {
//				if v.name == name {
//					item = v
//					break
//				}
//			}
//			return item, nil
//		});
//
//		func (c * MyMemoryPersistence) Set(correlatonId: string, item: MyData, callback: (err) => void): void {
//			c.Items = append(c.Items, item);
//			c.Save(correlationId);
//		}
//
//		persistence := NewMyMemoryPersistence();
//		err := persistence.Set("123", MyData{ name: "ABC" })
//		item, err := persistence.GetByName("123", "ABC")
//		fmt.Println(item)   // Result: { name: "ABC" }
//	implements IReferenceable, IOpenable, ICleanable
type MemoryPersistence[T cdata.ICloneable[T]] struct {
	Logger      *log.CompositeLogger
	Items       []T
	Loader      ILoader[T]
	Saver       ISaver[T]
	Lock        sync.RWMutex
	opened      bool
	MaxPageSize int
}

// NewMemoryPersistence creates a new instance of the MemoryPersistence
//	Parameters: prototype reflect.Type type of contained data
//	Return *MemoryPersistence a MemoryPersistence
func NewMemoryPersistence[T cdata.ICloneable[T]]() *MemoryPersistence[T] {
	c := &MemoryPersistence[T]{}
	c.Logger = log.NewCompositeLogger()
	c.Items = make([]T, 0, 10)
	return c
}

// SetReferences references to dependent components.
//	Parameters:
//		- ctx context.Context
//		- references refer.IReferences references to locate the component dependencies.
func (c *MemoryPersistence[T]) SetReferences(ctx context.Context, references refer.IReferences) {
	c.Logger.SetReferences(ctx, references)
}

// IsOpen checks if the component is opened.
// 	Returns true if the component has been opened and false otherwise.
func (c *MemoryPersistence[T]) IsOpen() bool {
	c.Lock.RLock()
	defer c.Lock.RUnlock()
	return c.opened
}

// Open the component.
//	Parameters:
//		- ctx context.Context
//		- correlationId  string (optional) transaction id to trace execution through call chain.
//	Returns: error or null no errors occured.
func (c *MemoryPersistence[T]) Open(ctx context.Context, correlationId string) error {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	if err := c.load(ctx, correlationId); err != nil {
		return err
	}
	c.opened = true
	return nil
}

func (c *MemoryPersistence[T]) load(ctx context.Context, correlationId string) error {
	if c.Loader == nil {
		return nil
	}

	items, err := c.Loader.Load(ctx, correlationId)
	if err == nil && items != nil {
		//c.Items = items
		c.Items = make([]T, len(items))
		for i, v := range items {
			c.Items[i] = v.Clone()
		}
		length := len(c.Items)
		c.Logger.Trace(ctx, correlationId, "Loaded %d items", length)
	}
	return err
}

// Closes component and frees used resources.
// Parameters:
//  - correlationId string
//  (optional) transaction id to trace execution through call chain.
// Retruns: error or nil if no errors occured.
func (c *MemoryPersistence[T]) Close(ctx context.Context, correlationId string) error {
	err := c.Save(ctx, correlationId)
	c.Lock.Lock()
	defer c.Lock.Unlock()
	c.opened = false
	return err
}

// Saves items to external data source using configured saver component.
// Parameters:
//   - correlationId string
//   (optional) transaction id to trace execution through call chain.
// Return error or null for success.
func (c *MemoryPersistence[T]) Save(ctx context.Context, correlationId string) error {
	c.Lock.RLock()
	defer c.Lock.RUnlock()

	if c.Saver == nil {
		return nil
	}

	err := c.Saver.Save(ctx, correlationId, c.Items)
	if err == nil {
		length := len(c.Items)
		c.Logger.Trace(ctx, correlationId, "Saved %d items", length)
	}
	return err
}

// Clears component state.
// Parameters:
//  - correlationId string
//  (optional) transaction id to trace execution through call chain.
//  Returns error or null no errors occured.
func (c *MemoryPersistence[T]) Clear(ctx context.Context, correlationId string) error {
	if err := c.Save(ctx, correlationId); err != nil {
		return err
	}
	c.Lock.Lock()
	defer c.Lock.Unlock()

	c.Items = make([]T, 0, 5)
	c.Logger.Trace(ctx, correlationId, "Cleared items")

	return nil
}

// Gets a page of data items retrieved by a given filter and sorted according to sort parameters.
// cmethod shall be called by a func (imp* IdentifiableMemoryPersistence) getPageByFilter method from child struct that
// receives FilterParams and converts them into a filter function.
// Parameters:
//   - correlationId string
//   transaction id to trace execution through call chain.
//   - filter func(interface{}) bool
//   (optional) a filter function to filter items
//   - paging *cdata.PagingParams
//   (optional) paging parameters
//   - sortFunc func(a, b interface{}) bool
//   (optional) sorting compare function func Less (a, b interface{}) bool  see sort.Interface Less function
//   - selectFunc func(in interface{}) (out interface{})
// (optional) projection parameters
// Return cdata.DataPage, error
// data page or error.
func (c *MemoryPersistence[T]) GetPageByFilter(ctx context.Context, correlationId string,
	filterFunc func(T) bool,
	paging cdata.PagingParams,
	sortFunc func(T, T) bool,
	selectFunc func(T) T) (page cdata.DataPage[T], err error) {

	c.Lock.RLock()
	defer c.Lock.RUnlock()

	items := make([]T, 0, len(c.Items))

	// Apply filtering
	if filterFunc != nil {
		for _, v := range c.Items {
			if filterFunc(v) {
				items = append(items, v.Clone())
			}
		}
	} else {
		for _, v := range c.Items {
			items = append(items, v.Clone())
		}
	}

	// Apply sorting
	if sortFunc != nil {
		localSort := sorter[T]{items: items, compFunc: sortFunc}
		sort.Sort(localSort)
	}

	// Extract a page
	skip := paging.GetSkip(-1)
	take := paging.GetTake((int64)(c.MaxPageSize))
	var total int64
	if paging.Total {
		total = (int64)(len(items))
	}
	if skip > 0 {
		_len := (int64)(len(items))
		if skip >= _len {
			skip = _len
		}
		items = items[skip:]
	}
	if (int64)(len(items)) >= take {
		items = items[:take]
	}

	// Get projection
	if selectFunc != nil {
		for i, v := range items {
			items[i] = selectFunc(v)
		}
	}

	c.Logger.Trace(ctx, correlationId, "Retrieved %d items", len(items))

	return *cdata.NewDataPage[T](items, int(total)), nil
}

// Gets a list of data items retrieved by a given filter and sorted according to sort parameters.
// This method shall be called by a func (c * IdentifiableMemoryPersistence) GetListByFilter method from child struct that
// receives FilterParams and converts them into a filter function.
// Parameters:
//   - correlationId string
//   (optional) transaction id to trace execution through call chain.
//   - filter func(interface{}) bool
//   (optional) a filter function to filter items
//   - sortFunc func(a, b interface{}) bool
//   (optional) sorting compare function func Less (a, b interface{}) bool  see sort.Interface Less function
//   - selectFunc func(in interface{}) (out interface{})
//   (optional) projection parameters
// Returns  []interface{},  error
// array of items and error
func (c *MemoryPersistence[T]) GetListByFilter(ctx context.Context, correlationId string,
	filterFunc func(T) bool,
	sortFunc func(T, T) bool,
	selectFunc func(T) T) ([]T, error) {

	c.Lock.RLock()
	defer c.Lock.RUnlock()

	// Apply filter
	items := make([]T, 0, len(c.Items))

	// Apply filtering
	if filterFunc != nil {
		for _, v := range c.Items {
			if filterFunc(v) {
				items = append(items, v.Clone())
			}
		}
	} else {
		for _, v := range c.Items {
			items = append(items, v.Clone())
		}
	}

	if len(items) == 0 {
		return nil, nil
	}

	// Apply sorting
	if sortFunc != nil {
		localSort := sorter[T]{items: items, compFunc: sortFunc}
		sort.Sort(localSort)
	}

	// Get projection
	if selectFunc != nil {
		for i, v := range items {
			items[i] = selectFunc(v)
		}
	}

	c.Logger.Trace(ctx, correlationId, "Retrieved %d items", len(items))

	return items, nil
}

// Gets a random item from items that match to a given filter.
// This method shall be called by a func (c* IdentifiableMemoryPersistence) GetOneRandom method from child type that
// receives FilterParams and converts them into a filter function.
// Parameters:
//   - correlationId string
//   (optional) transaction id to trace execution through call chain.
//   - filter   func(interface{}) bool
//   (optional) a filter function to filter items.
// Returns: interface{}, error
// random item or error.
func (c *MemoryPersistence[T]) GetOneRandom(ctx context.Context, correlationId string,
	filterFunc func(T) bool) (result T, err error) {

	c.Lock.RLock()
	defer c.Lock.RUnlock()

	// Apply filter
	items := make([]T, 0, len(c.Items))

	// Apply filtering
	if filterFunc != nil {
		for _, v := range c.Items {
			if filterFunc(v) {
				items = append(items, v.Clone())
			}
		}
	} else {
		for _, v := range c.Items {
			items = append(items, v.Clone())
		}
	}
	rand.Seed(time.Now().UnixNano())

	var item *T = nil
	if len(items) > 0 {
		item = &items[rand.Intn(len(items))]
	}

	if item != nil {
		c.Logger.Trace(ctx, correlationId, "Retrieved a random item")
	} else {
		c.Logger.Trace(ctx, correlationId, "Nothing to return as random item")
	}

	return *item, nil
}

// Creates a data item.
// Returns:
//   - correlationId string
//   (optional) transaction id to trace execution through call chain.
//   - item  string
//   an item to be created.
// Returns:  interface{}, error
// created item or error.
func (c *MemoryPersistence[T]) Create(ctx context.Context, correlationId string,
	item T) (T, error) {

	c.Lock.Lock()

	c.Items = append(c.Items, item.Clone())

	c.Logger.Trace(ctx, correlationId, "Created item")

	c.Lock.Unlock()

	if err := c.Save(ctx, correlationId); err != nil {
		return item.Clone(), err
	}

	return item.Clone(), nil
}

// Deletes data items that match to a given filter.
// this method shall be called by a func (c* IdentifiableMemoryPersistence) DeleteByFilter method from child struct that
// receives FilterParams and converts them into a filter function.
// Parameters:
//   - correlationId  string
//   (optional) transaction id to trace execution through call chain.
//   - filter  filter func(interface{}) bool
//   (optional) a filter function to filter items.
// Retruns: error
// error or nil for success.
func (c *MemoryPersistence[T]) DeleteByFilter(ctx context.Context, correlationId string,
	filterFunc func(T) bool) (err error) {

	c.Lock.Lock()

	deleted := 0
	for i := 0; i < len(c.Items); {
		if filterFunc(c.Items[i]) {
			if i == len(c.Items)-1 {
				c.Items = c.Items[:i]
			} else {
				c.Items = append(c.Items[:i], c.Items[i+1:]...)
			}
			deleted++
		} else {
			i++
		}
	}
	c.Lock.Unlock()

	if deleted == 0 {
		return nil
	}

	c.Logger.Trace(ctx, correlationId, "Deleted %s items", deleted)

	return c.Save(ctx, correlationId)
}

// Gets a count of data items retrieved by a given filter.
// this method shall be called by a func (imp* IdentifiableMemoryPersistence) getCountByFilter method from child struct that
// receives FilterParams and converts them into a filter function.
// Parameters:
//  - correlationId string
//  transaction id to trace execution through call chain.
//  - filter func(interface{}) bool
//  (optional) a filter function to filter items
// Return int, error
// data count or error.
func (c *MemoryPersistence[T]) GetCountByFilter(ctx context.Context, correlationId string,
	filterFunc func(T) bool) (count int64, err error) {

	c.Lock.RLock()
	defer c.Lock.RUnlock()

	// Apply filtering
	if filterFunc != nil {
		for _, v := range c.Items {
			if filterFunc(v) {
				count++
			}
		}
	} else {
		count = 0
	}
	c.Logger.Trace(ctx, correlationId, "Find %d items", count)
	return count, nil
}
