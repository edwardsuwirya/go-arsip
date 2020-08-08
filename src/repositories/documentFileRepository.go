package repositories

import (
	"entitites"
	"github.com/mitchellh/mapstructure"
	"strings"
)

type DocumentFileRepository struct {
	Repo IRepository
}

func (j *DocumentFileRepository) Create(d *entitites.Document) {
	docList, _ := j.SelectAll()
	docList = append(docList, d)
	j.CreateBulk(docList)
}

func (j *DocumentFileRepository) CreateBulk(d []*entitites.Document) error {
	j.Repo.Create(d)
	defer func() {
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	return nil
}

func (j *DocumentFileRepository) SelectAll() ([]*entitites.Document, error) {
	doc := j.Repo.Read()
	res := make([]*entitites.Document, 0, 1)
	for _, d := range doc.([]interface{}) {
		var result = entitites.Document{}
		mapstructure.Decode(d, &result)
		res = append(res, &result)
	}
	return res, nil
}

func (j *DocumentFileRepository) Find(no string) (*entitites.Document, error) {
	docList, _ := j.SelectAll()
	for _, a := range docList {
		if strings.HasPrefix(a.DocumentNo, no) {
			return a, nil
		}
	}
	return nil, nil
}
func (j *DocumentFileRepository) FindByTitle(title string) (*entitites.Document, error) {
	return nil, nil
}
