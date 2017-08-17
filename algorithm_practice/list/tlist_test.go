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
		mySList.addSList(NewsList("add" + strconv.Itoa(i)))
	}
	mySList.traverseSList()
}

//正向遍历双向链表
func TestTraverseDList(t *testing.T) {
	myDList := newDList()
	for i := 1; i <= 10; i++ {
		myDList.addDListEnd(NewdList("add" + strconv.Itoa(i)))
	}
	myDList.traverseDListForward()
}

//反向遍历双向链表
func TestTraverRDList(t *testing.T) {
	myDList := newDList()
	for i := 1; i <= 10; i++ {
		myDList.addDListEnd(NewdList("add" + strconv.Itoa(i)))
	}
	myDList.traverseDListReverse()
}

//链表搜索
func TestGetDateList(t *testing.T) {
	mySList := newSList()
	myDList := newDList()
	for i := 1; i <= 10; i++ {
		mySList.addSList(NewsList("add" + strconv.Itoa(i)))
		myDList.addDListEnd(NewdList("add" + strconv.Itoa(i)))
	}
	fmt.Println(getDateList("add6", mySList))
	fmt.Println(getDateList("add6", myDList))
}

//单向链表反转
func TestSListReverse(t *testing.T){
	mySList := newSList()
	for i := 1; i <= 10; i++ {
		mySList.addSList(NewsList("add" + strconv.Itoa(i)))
	}
	mySList.Reverse()
	t.Log(mySList)
}



//单向链表删除
func TestSListDelete(t *testing.T){
	mySList := newSList()
	for i := 1; i <= 10; i++ {
		mySList.addSList(NewsList("add" + strconv.Itoa(i)))
	}
	t.Log(mySList)
	mySList.Delete(NewsList("add10"))
	t.Log(mySList)
}



//找环
func TestSListFindLoop(t *testing.T){
	mySList := newSList()
	for i := 1; i <= 10; i++ {
		mySList.addSList(NewsList("add" + strconv.Itoa(i)))
	}
	mySList.sListLast.listPtrNext = mySList.sListHead.listPtrNext.listPtrNext
	t.Log(mySList.hasLoopByFS())
	t.Log(mySList.hasLoopEntryByFS())
}
