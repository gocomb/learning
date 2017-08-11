package list

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

//单向链表single linked list
type sList struct {
	data        string
	listPtrNext *sList
}

//单向链表实例类型
type sListItem struct {
	sListHead *sList
	sListLast *sList
}

//双向链表
type dList struct {
	data        string
	listPtrPrev *dList
	listPtrNext *dList
}

//双向链表实例类型
type dListItem struct {
	dListHead *dList
	dListLast *dList
}

//单向链表实例工厂，生成单向链表实例
func sListFactory() *sListItem {
	return &sListItem{
		sListHead: new(sList),
		sListLast: new(sList),
	}
}

//双向链表实例工厂，生成单向链表实例
func dListFactory() *dListItem {
	return &dListItem{
		dListHead: new(dList),
		dListLast: new(dList),
	}
}

//生成并初始化单向链表
func newSList() *sListItem {
	newList := sListFactory()
	*newList.sListHead = sList{data: "first single linked list ", listPtrNext: newList.sListLast}
	*newList.sListLast = sList{data: "init last single linked list ", listPtrNext: nil}
	return newList
}

//生成并初始化双向链表
func newDList() *dListItem {
	newList := dListFactory()
	*newList.dListHead = dList{data: "init first single linked list ", listPtrNext: newList.dListLast, listPtrPrev: nil}
	*newList.dListLast = dList{data: "init last single linked list ", listPtrNext: nil, listPtrPrev: newList.dListHead}
	return newList
}

//单向链表尾部增加元素
func (p *sListItem) addSList(data string) {
	sListNew := new(sList)
	sListNew = &sList{
		data:        data,
		listPtrNext: nil,
	}
	temp := p.sListLast
	temp.listPtrNext = sListNew
	p.sListLast = sListNew
}

//双向链表尾部增加元素
func (p *dListItem) addDListEnd(data string) {
	dListNew := new(dList)
	dListNew = &dList{
		data:        data,
		listPtrNext: nil,
		listPtrPrev: new(dList),
	}
	temp := p.dListLast
	temp.listPtrNext = dListNew
	dListNew.listPtrPrev = temp
	p.dListLast = dListNew
}

//遍历单向链表
func (p *sListItem) traverseSList() {
	temp := p.sListHead
	fmt.Println(temp)
	condition := true
	for condition {
		temp = temp.listPtrNext
		fmt.Println(unsafe.Pointer(temp), temp)
		if temp == nil {
			condition = false
		}
	}

}

//正向遍历双向链表
func (p *dListItem) traverseSListForward() {
	temp := p.dListHead
	fmt.Println(temp)
	condition := true
	for condition {
		temp = temp.listPtrNext
		fmt.Println(unsafe.Pointer(temp), temp)
		if temp == nil {
			condition = false
		}
	}

}

//逆向遍历双向链表
func (p *dListItem) traverseSListReverse() {
	temp := p.dListLast
	fmt.Println(temp)
	condition := true
	for condition {
		temp = temp.listPtrPrev
		fmt.Println(unsafe.Pointer(temp), temp)
		if temp == nil {
			condition = false
		}
	}

}

//链表搜索
func getDateList(key string, list interface{}) error {
	fmt.Println(fmt.Sprintln(reflect.TypeOf(list)))
	if strings.Contains(fmt.Sprintln(reflect.TypeOf(list)), "sListItem") {
		return getDateSList(key, list)
	}
	if strings.Contains(fmt.Sprintln(reflect.TypeOf(list)), "dListItem") {
		return getDateDList(key, list)
	}
	return errors.New("can not match the list")
}

//链表搜索:单链表
func getDateSList(key string, list interface{}) error {
	sList, ok := list.(*sListItem)
	if !ok {
		return errors.New("list is wrong")
	}
	temp := sList.sListHead
	fmt.Println(temp)
	condition := true
	for condition {
		if temp.data == key {
			fmt.Println("get it !!!", unsafe.Pointer(temp), temp)
			return nil
		}
		temp = temp.listPtrNext
		fmt.Println(unsafe.Pointer(temp), temp)
		if temp == nil {
			condition = false
		}
	}
	return errors.New("can not get the key")
}

//链表搜索:双链表
func getDateDList(key string, list interface{}) error {
	sList, ok := list.(*dListItem)
	if !ok {
		return errors.New("list is wrong")
	}
	temp := sList.dListHead
	fmt.Println(temp)
	condition := true
	for condition {
		if temp.data == key {
			fmt.Println("get it !!!", unsafe.Pointer(temp), temp)
			return nil
		}
		temp = temp.listPtrNext
		fmt.Println(unsafe.Pointer(temp), temp)
		if temp == nil {
			condition = false
		}
	}
	return errors.New("can not get the key")
}
