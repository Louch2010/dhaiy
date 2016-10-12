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
	length   int
	items    []*SortItem
	scoreMap map[interface{}]int
}

func NewSortSet() *SortSet {
	return &SortSet{
		length:   0,
		items:    make([]*SortItem, 10),
		scoreMap: make(map[interface{}]int),
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
	//如果已经存在，则修改score的值
	if _, ok := l.scoreMap[item.Data]; ok {
		l.scoreMap[item.Data] = item.Score
		for i := 0; i < l.length; i++ {
			if l.items[i].Data == item.Data {
				l.items[i].Score = item.Score
				break
			}
		}
		return
	}
	//scoreMap添加元素
	l.scoreMap[item.Data] = item.Score
	//items添加元素
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
	delete(l.scoreMap, item.Data)
	l.length--
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
	//不存在
	if item == nil {
		return nil
	}
	delete(l.scoreMap, item.Data)
	l.length--
	//将index后的数据进行复制
	for i := index; i < l.length; i++ {
		l.items[i] = l.items[i+1]
	}
	return item
}

//获取分值
func (l *SortSet) Score(data interface{}) int {
	if score, ok := l.scoreMap[data]; ok {
		return score
	}
	return -99
}

//获取排名
func (l *SortSet) Rank(data interface{}) int {
	for index, item := range l.GetItems() {
		if data == item.Data {
			return index + 1
		}
	}
	return 0
}

//获取分值范围内的数量
func (l *SortSet) Count(min, max int) int {
	count := 0
	for _, item := range l.GetItems() {
		if item.Score >= min && item.Score <= max {
			count++
		}
	}
	return count
}

//增加分数
func (l *SortSet) AddScore(data interface{}, add int) int {
	for _, item := range l.GetItems() {
		if data == item.Data {
			item.Score = item.Score + add
			l.scoreMap[item.Data] = item.Score
			//重新排序 TODO

			return item.Score
		}
	}
	return 0
}
