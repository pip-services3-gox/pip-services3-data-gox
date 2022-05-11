package persistence

import "github.com/pip-services3-gox/pip-services3-commons-gox/data"

// IDataObject interface for data objects to operate inside persistence.
//	Example
//		type MyDataObject struct {
//			...
//			id string
//		}
//
//		func (c *MyDataObject) GetId() string {
//			return c.id
//		}
//		func (c *MyDataObject) SetId(id string) string {
//			c.id = id
//		}
//		func (c *MyDataObject) Clone() *MyDataObject {
//			cloneObj := new(MyDataObject)
//			// Copy every attribute from this to cloneObj here.
//			...
//			return cloneObj
//		}
type IDataObject[T any, K any] interface {
	data.ICloneable[T]
	data.IIdentifiable[T, K]
}
