package persistence

import (
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

// IDataObject interface for data objects to operate inside persistence.
//	Example
//		type MyDataObject struct {
//			Id string
//			...
//		}
//		func (d MyDataObject) IsEqualId(id string) bool {
//			return d.Id == id
//		}
//		func (d MyDataObject) GetId() string {
//			return d.Id
//		}
//		func (d MyDataObject) IsZeroId() bool {
//			return d.Id == ""
//		}
//		func (d MyDataObject) WithId(id string) MyDataObject {
//			d.Id = id
//			return d
//		}
//		func (d MyDataObject) WithGeneratedId() MyDataObject {
//			d.Id = data.IdGenerator.NextLong()
//			return d
//		}
//		func (c MyDataObject) Clone() MyDataObject {
//			cloneObj := MyDataObject{
//				// Copy every attribute from this to cloneObj here.
//				...
//			}
//			return cloneObj
//		}
type IDataObject[T any, K any] interface {
	cdata.ICloneable[T]
	cdata.IIdentifiable[T, K]
}
