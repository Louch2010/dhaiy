package common

import (
	"fmt"
	//"math/rand"
	//"strconv"
	"testing"
)

func TestSort(t *testing.T) {
	list := NewSortSet()
	list.Add(3, "cccc")
	list.Add(2, "bbbb")
	list.Add(4, "dddd")
	//	list.Add(4, "dddd")
	//	list.Add(2, "bbbb")
	//	list.Add(1, "aaaa")
	//	list.Add(5, "eeee")
	list.RemoveData("bbbb")
	//	for i := 0; i < 400; i++ {
	//		r := rand.Intn(10)
	//		list.Add(r, "v_"+strconv.Itoa(r))
	//	}
	for index, item := range list.GetItems() {
		fmt.Println(index, " = ", item)
	}
}
