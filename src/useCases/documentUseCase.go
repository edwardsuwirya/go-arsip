package useCases

import (
	"entitites"
	"repositories"
)

type IDocumentUseCase interface {
	CreateDocument(doc entitites.Document)
	GetTotalDocument() int
	FindAll() []*entitites.Document
	FindDocumentByNumber(docNo string) []*entitites.Document
}

type DocumentUseCase struct {
	Repo repositories.DocumentFileRepository
}

func (d *DocumentUseCase) FindAll() []*entitites.Document {
	list, err := d.Repo.SelectAll()
	if err != nil {
		return nil
	}
	return list
}

func (d *DocumentUseCase) GetTotalDocument() int {
	list, err := d.Repo.SelectAll()
	if err != nil {
		return -1
	}
	return len(list)
}
func (d *DocumentUseCase) CreateDocument(doc entitites.Document) {
	d.Repo.Create(&doc)
}

func (d *DocumentUseCase) FindDocumentByNumber(docNo string) []*entitites.Document {
	if len(docNo) == 0 {
		return d.FindAll()
	} else {
		res, err := d.Repo.Find(docNo)
		if err != nil {
			return nil
		}
		var docResult []*entitites.Document
		if res != nil {
			docResult = []*entitites.Document{res}
		} else {
			docResult = []*entitites.Document{}
		}
		return docResult
	}
}
