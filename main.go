package main

import (
	"fmt"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

func main() {
	fmt.Println("tf.Version = ", tf.Version())
}
