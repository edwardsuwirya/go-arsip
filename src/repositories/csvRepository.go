package repositories

import (
	"fmt"
	"github.com/jszwec/csvutil"
	"io/ioutil"
)

type CsvRepository struct {
	fileName string
	data     interface{}
}

func (c CsvRepository) Create() {
	b, err := csvutil.Marshal(c.data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(c.fileName, b, 0644)

	if err != nil {
		panic(err)
	}
}

func (c CsvRepository) Read() interface{} {
	content, err := ioutil.ReadFile(c.fileName)
	if err != nil {
		c.Create()
	}
	if err := csvutil.Unmarshal(content, &c.data); err != nil {
		fmt.Println("error:", err)
		panic(err)
	}
	return c.data
}
