package list

import (
	"testing"
	"strconv"
	"fmt"
)

//遍历单向链表
func TestTraverseSList(t *testing.T) {
	mySList := newSList()
	for i := 1; i <= 10; i++ {
		mySList.addSList("add" + strconv.Itoa(i))
	}
	mySList.traverseSList()
}

//正向遍历双向链表
func TestTraverseDList(t *testing.T) {
	myDList := newDList()
	for i := 1; i <= 10; i++ {
		myDList.addDListEnd("add" + strconv.Itoa(i))
	}
	myDList.traverseSListForward()
}

//反向遍历双向链表
func TestTraverRDList(t *testing.T) {
	myDList := newDList()
	for i := 1; i <= 10; i++ {
		myDList.addDListEnd("add" + strconv.Itoa(i))
	}
	myDList.traverseSListReverse()
}

//链表搜索
func TestGetDateList(t *testing.T) {
	mySList := newSList()
	myDList := newDList()
	for i := 1; i <= 10; i++ {
		mySList.addSList("add" + strconv.Itoa(i))
		myDList.addDListEnd("add" + strconv.Itoa(i))
	}
	fmt.Println(getDateList("add6", mySList))
	fmt.Println(getDateList("add6", myDList))
}
