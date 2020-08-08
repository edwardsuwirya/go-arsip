package repositories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type JsonRepository struct {
	FileName string
	data     interface{}
}

func (j JsonRepository) Create(data interface{}) {
	b, err := json.Marshal(data)
	j.data = data
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(j.FileName, b, 0644)

	if err != nil {
		panic(err)
	}
}

func (j JsonRepository) Read() interface{} {
	content, err := ioutil.ReadFile(j.FileName)
	if err != nil {
		j.data = make([]interface{}, 0, 1)
		j.Create(j.data)
	} else {
		if err := json.Unmarshal(content, &j.data); err != nil {
			fmt.Println("error:", err)
			panic(err)
		}
	}
	return j.data
}
