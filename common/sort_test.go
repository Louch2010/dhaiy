package common

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func TestSort(t *testing.T) {
	list := NewSortSet()
	//	list.Add(3, "cccc")
	//	list.Add(2, "bbbb2")
	//	list.Add(4, "dddd")
	//	list.Add(4, "dddd")
	//	list.Add(2, "bbbb")
	//	list.Add(1, "aaaa")
	//	list.Add(5, "eeee")
	for i := 0; i < 40; i++ {
		list.Add(rand.Intn(20), "v_"+strconv.Itoa(i))
	}
	for index, item := range list.GetItems() {
		fmt.Println(index, " = ", item)
	}
}
