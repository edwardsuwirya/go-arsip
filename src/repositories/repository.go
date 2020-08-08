package repositories

import "entitites"

type IRepository interface {
	Create(data interface{})
	Read() interface{}
}
type DocumentRepository interface {
	Create(d entitites.Document)
	CreateBulk(d []*entitites.Document) error
	SelectAll() ([]*entitites.Document, error)
	Find(no string) (*entitites.Document, error)
	FindByTitle(title string) (*entitites.Document, error)
}
