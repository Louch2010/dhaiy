package common

import (
	"fmt"
	"strconv"
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

func NewSortSet() SortSet {
	return SortSet{
		length: 0,
		items:  make([]*SortItem, 10),
	}
}

//长度
func (l *SortSet) Len() int {
	return l.length
}

//增加
func (l *SortSet) Add(score int, data interface{}) {
	tmp := &SortItem{
		Score: score,
		Data:  data,
	}
	l.AddItem(tmp)
}

//增加元素
func (l *SortSet) AddItem(item *SortItem) {
	if l.length == 0 {
		l.length++
		l.items[0] = item
		return
	}
	//扩容
	if len(l.items)-1 <= l.Len() {
		tmp := l.items
		//扩为2倍
		l.items = make([]*SortItem, len(tmp)*2)
		//复制
		for i := 0; i < len(tmp); i++ {
			l.items[i] = tmp[i]
		}
		fmt.Println("扩容后容量：", len(l.items))
	}
	//排序
	for i := 0; i < l.length; i++ {
		//如果原数组最后一个值比要插入的值小，则要插入的值放在最后即可，且跳出循环
		if i == l.length-1 && l.items[i].Score < item.Score {
			l.items[i+1] = item
			break
		}
		//如果当前值比要插入的值大，则从当前值开始，将数组剩余的部分全部后移一位，且将要插入的值放在当前位置
		if l.items[i].Score >= item.Score {
			//复制
			for j := l.length; j > i; j-- {
				l.items[j] = l.items[j-1]
			}
			l.items[i] = item
			break
		}
	}
	l.length++
}

//获取元素
func (l *SortSet) GetItem(index int) *SortItem {
	if index >= l.length {
		return nil
	}
	return l.items[index]
}

//获取内容
func (l *SortSet) GetItems() []*SortItem {
	return l.items[0:l.length]
}

//根据index删除
func (l *SortSet) RemoveItem(index int) *SortItem {
	if index >= l.length {
		return nil
	}
	item := l.items[index]
	for i := index; i < l.Len(); i++ {
		l.items[i] = l.items[i+1]
	}
	return item
}

//根据内容删除
func (l *SortSet) RemoveData(data interface{}) *SortItem {
	index := 0
	var item *SortItem
	for i := 0; i < l.length; i++ {
		if l.items[i].Data == data {
			index = i
			item = l.items[i]
			break
		}
	}
	//将index后的数据进行复制
	for i := index; i < l.length; i++ {
		l.items[i] = l.items[i+1]
	}
	return item
}

func (l *SortSet) Test() {
	//fmt.Println("容量为：", len(l.items), "，长度为：", l.length)
	sb := ""
	for i := 0; i < l.Len(); i++ {
		sb = sb + strconv.Itoa(l.items[i].Score) + ","
	}
	fmt.Println("items:", sb)
}
