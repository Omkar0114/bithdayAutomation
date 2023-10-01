package main

import (
	"fmt"
	"reflect"
	"time"
)

func main(){

	t := time.Now()

	date := fmt.Sprintf("%02d/%02d/%02d", int(t.Month()), int(t.Day()), int(t.Year()))

	fmt.Println(date)
}