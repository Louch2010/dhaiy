package common

import (
	"fmt"
)

//有序项
type SortItem struct {
	Score int         //分值
	Data  interface{} //数据
}

//有序集合
type SortSet struct {
	length int
	items  []*SortItem
}

//长度
func (l SortSet) Len() int {
	return l.length
}

//增加
func (l SortSet) Add(score int, data interface{}) {
	fmt.Println("-------------")
	fmt.Println("长度为：", l.Len())
	fmt.Println("容量为：", len(l.items))
	tmp := &SortItem{
		Score: score,
		Data:  data,
	}
	l.AddItem(tmp)
}

//增加元素
func (l SortSet) AddItem(item *SortItem) {
	if l.Len() == 0 {
		fmt.Println("0000000")
		l.length++
		l.items = make([]*SortItem, 10)
		l.items[0] = item
		fmt.Println("长度为：", l.Len())
		fmt.Println("容量为：", len(l.items))
		return
	}
	//扩容
	if l.length == l.Len() {
		tmp := l.items
		//扩为2倍
		//l.items = [len(tmp) * 2]*SortItem{}
		l.items = make([]*SortItem, len(tmp)*2)
		//复制
		for i := 0; i < len(tmp); i++ {
			l.items[i] = tmp[i]
		}
	}
	//排序
	for i := 0; i < l.length; i++ {
		if l.items[i].Score > item.Score {
			//复制
			for j := l.length; j > i; j-- {
				l.items[j] = l.items[j-1]
			}
			l.items[i] = item
		}
	}
	l.length++
}

//获取元素
func (l SortSet) GetItem(index int) *SortItem {
	if index >= l.Len() {
		return nil
	}
	return l.items[index]
}

//获取内容
func (l SortSet) GetItems() []*SortItem {
	return l.items[0:l.length]
}

//根据index删除
func (l SortSet) RemoveItem(index int) *SortItem {
	if index >= l.Len() {
		return nil
	}
	item := l.items[index]
	for i := index; i < l.Len(); i++ {
		l.items[i] = l.items[i+1]
	}
	return item
}

//根据内容删除
func (l SortSet) RemoveData(data interface{}) *SortItem {
	index := 0
	var item *SortItem
	for i := 0; i < l.Len(); i++ {
		if l.items[i].Data == data {
			index = i
			item = l.items[i]
			break
		}
	}
	//将index后的数据进行复制
	for i := index; i < l.Len(); i++ {
		l.items[i] = l.items[i+1]
	}
	return item
}
