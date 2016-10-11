package common

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	list := SortSet{}
	list.Add(3, "cccc")
	list.Add(4, "dddd")
	list.Add(2, "bbbb")
	list.Add(1, "aaaa")
	list.Add(5, "eeee")
	for index, item := range list.GetItems() {
		fmt.Println(index, " = ", item)
	}
}
