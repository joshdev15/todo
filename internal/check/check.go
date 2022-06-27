package check

import (
	"fmt"
	"log"
)

var (
	errorList = []string{"bucket already exists"}
)

func Err(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrWithMsg(err error, msg string) {
	if err != nil {
		log.Fatalf("%v: %v", msg, err)
		panic(err)
	}
}

func CompareErr(err error) {
	if err != nil {
		errMsg := err.Error()
		for _, v := range errorList {
			if errMsg != v {
				fmt.Println(err)
				panic(err)
			}
		}
	}
}
