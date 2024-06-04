package shared

import "log"

func ErrorHandling(err error) {
	if err != nil {
		log.Default().Printf("error: %v", err)
		panic(err)
	}
}
