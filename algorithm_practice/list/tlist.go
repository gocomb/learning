package list

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"unsafe"
	"bytes"
)

//单向链表single linked list
type sList struct {
	data string
	listPtrNext *sList
}


func NewsList(item string) *sList{
	return &sList{
		data:item,
		listPtrNext:nil,
	}
}

//单向链表实例类型
type sListItem struct {
	sListHead *sList
	sListLast *sList
}

//双向链表
type dList struct {
	data string
	listPtrPrev *dList
	listPtrNext *dList
}

func NewdList(item string) *dList{
	return &dList{
		data:item,
		listPtrNext:nil,
		listPtrPrev:nil,
	}
}

//双向链表实例类型
type dListItem struct {
	dListHead *dList
	dListLast *dList
}


//生成并初始化单向链表
func newSList() *sListItem {
	return &sListItem{
		sListHead:nil,
		sListLast:nil,
	}
}

//生成并初始化双向链表
func newDList() *dListItem {
	return &dListItem{
		dListHead:nil,
		dListLast:nil,
	}
}

func(s *sList)Equal(m *sList)bool{
	if m == nil{
		return false
	}
	return s.data == m.data
}

func(s *sList)DeepEqual(m *sList)bool{
	temp := map[*sList]bool{
		s:true,
	}
	return temp[m]
}

//单向链表尾部增加元素
func (p *sListItem) addSList(n *sList) {
	if p.sListHead == nil{
		p.sListHead = n
		p.sListLast = n
		return
	}
	p.sListLast.listPtrNext = n
	p.sListLast = n
}

//双向链表尾部增加元素
func (p *dListItem) addDListEnd(n *dList) {
	if p.dListHead == nil{
		p.dListHead = n
		p.dListLast = n
		return
	}
	p.dListLast.listPtrNext = n
	n.listPtrPrev = p.dListLast
	p.dListLast = n
}

//遍历单向链表
func (p *sListItem) traverseSList() {
	temp := p.sListHead
	for temp.listPtrNext!=nil {
		fmt.Println(unsafe.Pointer(temp),temp.data)
		temp = temp.listPtrNext
	}
	fmt.Println(unsafe.Pointer(temp),temp.data)
}


//链表删除
func (p *sListItem) Delete(n *sList){
	if p.sListHead == nil{
		return
	}
	if p.sListHead.Equal(n){
		if p.sListHead.listPtrNext == nil{ //only one Node
			p.sListHead = nil
			p.sListLast = nil
			return
		}
		p.sListHead = p.sListHead.listPtrNext
		return
	}
	item := p.sListHead
	for item.listPtrNext != nil{
		if item.listPtrNext.Equal(n) {
			if item.listPtrNext.listPtrNext == nil{
				p.sListLast = item
			}
			item.listPtrNext = item.listPtrNext.listPtrNext
			return
		}
		item = item.listPtrNext
	}
}

//递归链表反转单向链表
func (p *sListItem) Reverse(){
	if p.sListHead ==nil || p.sListHead.listPtrNext == nil{
		return
	}
	temp:= p.sListHead
	p.sListHead = p.sListHead.listPtrNext
	p.Reverse()
	temp.listPtrNext = nil
	p.sListLast.listPtrNext = temp
	p.sListLast = temp
	return
}

//单向链表找环 快慢指针
func (p *sListItem)hasLoopByFS() bool{
	if p.sListHead == nil{
		return false
	}
	fast := p.sListHead
	slow := p.sListHead
	for fast!=nil && fast.listPtrNext!= nil{
		fmt.Println(fast.data,slow.data)
		fmt.Println(fast,slow)
		fast = fast.listPtrNext.listPtrNext
		slow = slow.listPtrNext
		if fast == slow {
			return true
		}
	}
	return false
}



//单向链表找环入口 快慢指针
func (p *sListItem)hasLoopEntryByFS() *sList{
	if p.sListHead == nil{
		return nil
	}
	fast := p.sListHead
	slow := p.sListHead
	for fast!=nil && fast.listPtrNext!= nil{
		fast = fast.listPtrNext.listPtrNext
		slow = slow.listPtrNext
		if fast == slow {
			break
		}
	}

	if fast ==nil || fast.listPtrNext == nil {
		return nil
	}
	fast = p.sListHead
	for fast != slow{
		fast = fast.listPtrNext
		slow = slow.listPtrNext
	}
	return fast
}

//正向遍历双向链表
func (p *dListItem) traverseDListForward() {
	temp := p.dListHead
	for temp.listPtrNext!=nil {
		fmt.Println(unsafe.Pointer(temp), temp.data)
		temp = temp.listPtrNext
	}
	fmt.Println(unsafe.Pointer(temp), temp.data)
}

//逆向遍历双向链表
func (p *dListItem) traverseDListReverse() {
	temp := p.dListLast
	for temp.listPtrPrev!=nil {
		fmt.Println(unsafe.Pointer(temp), temp.data)
		temp = temp.listPtrPrev
	}
	fmt.Println(unsafe.Pointer(temp), temp.data)
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

//XXX:单链表搜索和双链表搜索有区别吗？
//链表搜索:单链表
func getDateSList(key string, list interface{}) error {
	sList, ok := list.(*sListItem)
	if !ok {
		return errors.New("list is wrong")
	}
	temp := sList.sListHead
	for temp.listPtrNext!=nil {
		fmt.Println(unsafe.Pointer(temp), temp.data)
		if temp.data == key {
			fmt.Println("get it !!!", unsafe.Pointer(temp), temp.data)
			return nil
		}
		temp = temp.listPtrNext
	}
	fmt.Println(unsafe.Pointer(temp), temp)
	return errors.New("can not get the key")
}

//链表搜索:双链表
func getDateDList(key string, list interface{}) error {
	sList, ok := list.(*dListItem)
	if !ok {
		return errors.New("list is wrong")
	}
	temp := sList.dListHead
	for temp.listPtrNext!=nil {
		fmt.Println(unsafe.Pointer(temp), temp.data)
		if temp.data == key {
			fmt.Println("get it !!!", unsafe.Pointer(temp), temp.data)
			return nil
		}
		temp = temp.listPtrNext
	}
	fmt.Println(unsafe.Pointer(temp), temp.data)
	return errors.New("can not get the key")
}

func (s *sListItem) String()string{
	buf := bytes.Buffer{}
	item := s.sListHead
	if item == nil{
		return "Queue{}"
	}
	buf.WriteString(fmt.Sprintf("sList{%v",item.data))
	for item.listPtrNext != nil{
		item = item.listPtrNext
		buf.WriteString(fmt.Sprintf(" %v",item.data))
	}
	return buf.String()+"}"
}


func (s *dListItem) String()string{
	buf := bytes.Buffer{}
	item := s.dListHead
	if item == nil{
		return "sList{}"
	}
	buf.WriteString(fmt.Sprintf("sList{%v",item.data))
	for item.listPtrNext != nil{
		item = item.listPtrNext
		buf.WriteString(fmt.Sprintf(" %v",item.data))
	}
	return buf.String()+"}"
}




