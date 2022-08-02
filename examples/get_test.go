package examples

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	endpoints, err := GetEndpoints()
	if err != nil {
		panic(err)
	}

	fmt.Println(err)
	fmt.Println(endpoints)
}
