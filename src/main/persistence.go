package main

import (
	"fmt"
	"github.com/jszwec/csvutil"
	"io/ioutil"
)

func CreateCsv(filename string, a []Arsip) {
	b, err := csvutil.Marshal(a)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fileName, b, 0644)

	if err != nil {
		panic(err)
	}
}

func ReadCsv(filename string) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		CreateCsv(fileName, daftarArsip)
	}
	if err := csvutil.Unmarshal(content, &daftarArsip); err != nil {
		fmt.Println("error:", err)
		panic(err)
	}
}
