package handlers

import (
	"errors"
	"log"
)

// 包装错误
func PanicIfUserError(err error) {
	if err != nil {
		err = errors.New("userService--" + err.Error())
		log.Println("PanicIfUserError: ", err)
		panic(err)
	}
}

func PanicIfVideoError(err error) {
	if err != nil {
		err = errors.New("videoService--" + err.Error())
		log.Println("PanicIfVideoError: ", err)
		panic(err)
	}
}