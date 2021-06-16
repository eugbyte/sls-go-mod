package util

import (
	"encoding/json"
	"fmt"

	colors "github.com/TwinProduction/go-color"
)

func Trace(obj interface{}) {
	bytes, err := (json.MarshalIndent(obj, "", "\t"))
	if err != nil {
		fmt.Println((err))
	}
	fmt.Println(colors.Blue, string(bytes), colors.Reset)
}

func LogError(obj interface{}) {
	bytes, err := (json.MarshalIndent(obj, "", "\t"))
	if err != nil {
		fmt.Println((err))
	}
	fmt.Println(colors.Red, string(bytes), colors.Reset)
}
