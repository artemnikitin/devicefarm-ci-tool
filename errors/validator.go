package errors

import "log"

// Validate finish an app with exit code 1 in case of error
func Validate(err error, text string) {
	if err != nil {
		log.Fatal(text, err)
	}
}
