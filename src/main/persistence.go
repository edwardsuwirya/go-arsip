package main

import (
	"encoding/json"
	"fmt"
	"github.com/jszwec/csvutil"
	"io/ioutil"
)

type CsvPersistence struct {
}

func (c CsvPersistence) Create(a []Arsip) {
	b, err := csvutil.Marshal(a)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fileName, b, 0644)

	if err != nil {
		panic(err)
	}
}

func (c CsvPersistence) Read() {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		c.Create(daftarArsip)
	}
	if err := csvutil.Unmarshal(content, &daftarArsip); err != nil {
		fmt.Println("error:", err)
		panic(err)
	}
}

type JsonPersistence struct {
}

func (j JsonPersistence) Create(a []Arsip) {
	b, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(fileName, b, 0644)

	if err != nil {
		panic(err)
	}
}

func (j JsonPersistence) Read() {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		j.Create(daftarArsip)
	} else {
		if err := json.Unmarshal(content, &daftarArsip); err != nil {
			fmt.Println("error:", err)
			panic(err)
		}
	}

}
