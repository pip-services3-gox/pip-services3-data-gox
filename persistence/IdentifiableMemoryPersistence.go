package persistence

import (
	"context"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	refl "github.com/pip-services3-gox/pip-services3-commons-gox/reflect"
	"github.com/pip-services3-gox/pip-services3-components-gox/log"
	"reflect"
)

/*
Abstract persistence component that stores data in memory
and implements a number of CRUD operations over data items with unique ids.
The data items must have Id field.

In basic scenarios child structs shall only override GetPageByFilter,
GetListByFilter or DeleteByFilter operations with specific filter function.
All other operations can be used out of the box.

In complex scenarios child structes can implement additional operations by
accessing cached items via c.Items property and calling Save method
on updates.

See MemoryPersistence

Configuration parameters

- options:
    - max_page_size:       Maximum number of items returned in a single page (default: 100)

 References

- *:logger:*:*:1.0     (optional) ILogger components to pass log messages

 Examples

  type MyMemoryPersistence struct{
  	IdentifiableMemoryPersistence
  }
      func composeFilter(filter: FilterParams) (func (item interface{}) bool ) {
          if &filter == nil {
  			filter = NewFilterParams()
  		}
          name := filter.getAsNullableString("Name");
          return func(item interface{}) bool {
  			dummy, ok := item.(MyData)
              if (*name != "" && ok && item.Name != *name)
                  return false;
              return true;
          };
      }

      func (mmp * MyMemoryPersistence) GetPageByFilter(correlationId string, filter FilterParams, paging PagingParams) (page DataPage, err error) {
          tempPage, err := c.GetPageByFilter(correlationId, composeFilter(filter), paging, nil, nil)
  		dataLen := int64(len(tempPage.Data))
  		data := make([]MyData, dataLen)
  		for i, v := range tempPage.Data {
  			data[i] = v.(MyData)
  		}
  		page = *NewMyDataPage(&dataLen, data)
  		return page, err}

      persistence := NewMyMemoryPersistence();

  	item, err := persistence.Create("123", { Id: "1", Name: "ABC" })
  	...
  	page, err := persistence.GetPageByFilter("123", NewFilterParamsFromTuples("Name", "ABC"), nil)
  	if err != nil {
  		panic("Error can't get data")
  	}
      fmt.Prnitln(page.data)         // Result: { Id: "1", Name: "ABC" }
  	item, err := persistence.DeleteById("123", "1")
  	...

*/
// extends MemoryPersistence  implements IConfigurable, IWriter, IGetter, ISetter
type IdentifiableMemoryPersistence[T IDataObject[T, K], K any] struct {
	MemoryPersistence[T]
}

const IdentifiableMemoryPersistenceConfigParamOptionsMaxPageSize = "options.max_page_size"

// Creates a new empty instance of the persistence.
// Parameters:
//  - prototype reflect.Type
//  data type of contains items
// Return * IdentifiableMemoryPersistence
// created empty IdentifiableMemoryPersistence
func NewIdentifiableMemoryPersistence[T IDataObject[T, K], K any]() (c *IdentifiableMemoryPersistence[T, K]) {
	c = &IdentifiableMemoryPersistence[T, K]{
		MemoryPersistence: *NewMemoryPersistence[T](),
	}
	c.Logger = log.NewCompositeLogger()
	c.MaxPageSize = 100
	return c
}

// Configures component by passing configuration parameters.
// Parameters:
//  - config  *config.ConfigParams
//  configuration parameters to be set.
func (c *IdentifiableMemoryPersistence[T, K]) Configure(ctx context.Context, config *config.ConfigParams) {
	c.MaxPageSize = config.GetAsIntegerWithDefault(IdentifiableMemoryPersistenceConfigParamOptionsMaxPageSize, c.MaxPageSize)
}

// Gets a list of data items retrieved by given unique ids.
// Parameters:
//   - correlationId string
//   (optional) transaction id to trace execution through call chain.
//   - ids  []interface{}
//   ids of data items to be retrieved
// Returns  []interface{}, error
// data list or error.
func (c *IdentifiableMemoryPersistence[T, K]) GetListByIds(ctx context.Context, correlationId string,
	ids []K) ([]T, error) {

	filter := func(item T) bool {
		exist := false
		for _, id := range ids {
			if item.IsEqualId(id) {
				exist = true
				break
			}
		}
		return exist
	}
	return c.GetListByFilter(ctx, correlationId, filter, nil, nil)
}

// Gets a data item by its unique id.
// Parameters:
//   - correlationId  string
//   (optional) transaction id to trace execution through call chain.
//   - id interface{}
//   an id of data item to be retrieved.
// Returns:  interface{}, error
// data item or error.
func (c *IdentifiableMemoryPersistence[T, K]) GetOneById(ctx context.Context, correlationId string, id K) (T, error) {

	c.Lock.RLock()
	defer c.Lock.RUnlock()

	for _, item := range c.Items {
		if item.IsEqualId(id) {
			c.Logger.Trace(ctx, correlationId, "Retrieved item %s", id)
			return item.Clone(), nil
		}
	}

	c.Logger.Trace(ctx, correlationId, "Cannot find item by %s", id)

	var defaultObject T
	return defaultObject, nil
}

// Get index by "Id" field
// return index number
func (c *IdentifiableMemoryPersistence[T, K]) GetIndexById(id K) int {
	c.Lock.RLock()
	defer c.Lock.RUnlock()

	for i, item := range c.Items {
		if item.IsEqualId(id) {
			return i
		}
	}
	return -1
}

// Creates a data item.
// Returns:
//   - correlationId string
//   (optional) transaction id to trace execution through call chain.
//   - item  string
//   an item to be created.
// Returns:  interface{}, error
// created item or error.
func (c *IdentifiableMemoryPersistence[T, K]) Create(ctx context.Context, correlationId string, item T) (T, error) {
	c.Lock.Lock()

	newItem := item.Clone()
	if newItem.IsZeroId() {
		newItem = newItem.WithGeneratedId()
	}
	c.Items = append(c.Items, newItem)

	c.Lock.Unlock()
	c.Logger.Trace(ctx, correlationId, "Created item %s", newItem.GetId())

	if err := c.Save(ctx, correlationId); err != nil {
		return newItem.Clone(), err
	}

	return newItem.Clone(), nil
}

// Sets a data item. If the data item exists it updates it,
// otherwise it create a new data item.
// Parameters:
//   - correlationId string
//   (optional) transaction id to trace execution through call chain.
//   - item  interface{}
//   a item to be set.
// Returns:  interface{}, error
// updated item or error.
func (c *IdentifiableMemoryPersistence[T, K]) Set(ctx context.Context, correlationId string, item T) (T, error) {
	newItem := item.Clone()
	if newItem.IsZeroId() {
		newItem = newItem.WithGeneratedId()
	}

	index := c.GetIndexById(item.GetId())

	c.Lock.Lock()
	if index < 0 {
		c.Items = append(c.Items, newItem)
	} else {
		c.Items[index] = newItem
	}

	c.Lock.Unlock()
	c.Logger.Trace(ctx, correlationId, "Set item %s", newItem.GetId())

	if err := c.Save(ctx, correlationId); err != nil {
		return newItem.Clone(), err
	}

	return newItem.Clone(), nil
}

// Updates a data item.
// Parameters:
//   - correlationId string
//   (optional) transaction id to trace execution through call chain.
//   - item  interface{}
//   an item to be updated.
// Returns:   interface{}, error
// updated item or error.
func (c *IdentifiableMemoryPersistence[T, K]) Update(ctx context.Context, correlationId string, item T) (T, error) {
	var defaultObject T

	index := c.GetIndexById(item.GetId())
	if index < 0 {
		c.Logger.Trace(ctx, correlationId, "Item %s was not found", item.GetId())
		return defaultObject, nil
	}
	newItem := item.Clone()

	c.Lock.Lock()
	c.Items[index] = newItem
	c.Lock.Unlock()

	c.Logger.Trace(ctx, correlationId, "Updated item %s", item.GetId())

	if err := c.Save(ctx, correlationId); err != nil {
		return newItem.Clone(), err
	}

	return newItem.Clone(), nil
}

// Updates only few selectFuncected fields in a data item.
// Parameters:
//   - correlationId string
//   (optional) transaction id to trace execution through call chain.
//   - id interface{}
//   an id of data item to be updated.
//   - data  cdata.AnyValueMap
//   a map with fields to be updated.
// Returns: interface{}, error
// updated item or error.
func (c *IdentifiableMemoryPersistence[T, K]) UpdatePartially(ctx context.Context, correlationId string,
	id K, data cdata.AnyValueMap) (T, error) {

	var defaultObject T

	index := c.GetIndexById(id)
	if index < 0 {
		c.Logger.Trace(ctx, correlationId, "Item %s was not found", id)
		return defaultObject, nil
	}

	c.Lock.Lock()

	newItem := c.Items[index].Clone()

	if reflect.ValueOf(newItem).Kind() == reflect.Map {
		refl.ObjectWriter.SetProperties(newItem, data.Value())
	} else {
		var intPointer any = newItem
		if reflect.TypeOf(newItem).Kind() != reflect.Pointer {
			objPointer := reflect.New(reflect.TypeOf(newItem))
			objPointer.Elem().Set(reflect.ValueOf(newItem))
			intPointer = objPointer.Interface()
		}
		refl.ObjectWriter.SetProperties(intPointer, data.Value())
		if _newItem, ok := reflect.ValueOf(intPointer).Elem().Interface().(T); ok {
			newItem = _newItem
		}
	}

	c.Items[index] = newItem

	c.Lock.Unlock()
	c.Logger.Trace(ctx, correlationId, "Partially updated item %s", id)

	if err := c.Save(ctx, correlationId); err != nil {
		return newItem.Clone(), err
	}

	return newItem.Clone(), nil
}

// Deleted a data item by it's unique id.
// Parameters:
//   - correlationId string
//   (optional) transaction id to trace execution through call chain.
//   - id interface{}
//   an id of the item to be deleted
// Retruns:  interface{}, error
// deleted item or error.
func (c *IdentifiableMemoryPersistence[T, K]) DeleteById(ctx context.Context, correlationId string, id K) (T, error) {

	var defaultObject T

	index := c.GetIndexById(id)
	if index < 0 {
		c.Logger.Trace(ctx, correlationId, "Item %s was not found", id)
		return defaultObject, nil
	}

	c.Lock.Lock()

	oldItem := c.Items[index]
	if index == len(c.Items) {
		c.Items = c.Items[:index-1]
	} else {
		c.Items = append(c.Items[:index], c.Items[index+1:]...)
	}

	c.Lock.Unlock()

	c.Logger.Trace(ctx, correlationId, "Deleted item by %s", id)

	if err := c.Save(ctx, correlationId); err != nil {
		return oldItem, err
	}
	return oldItem, nil
}

// Deletes multiple data items by their unique ids.
// Parameters:
//   - correlationId  string
//   (optional) transaction id to trace execution through call chain.
//   - ids []interface{}
//   ids of data items to be deleted.
// Returns: error
// error or null for success.
func (c *IdentifiableMemoryPersistence[T, K]) DeleteByIds(ctx context.Context, correlationId string, ids []K) error {
	filterFunc := func(item T) bool {
		for _, id := range ids {
			if item.IsEqualId(id) {
				return true
			}
		}
		return false
	}

	return c.DeleteByFilter(ctx, correlationId, filterFunc)
}
