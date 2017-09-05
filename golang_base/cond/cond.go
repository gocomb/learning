package cond

import (
	"sync"
	"fmt"
	"time"
)

/*
通常在处理高并发事务时会采用锁机制，但是由于各种原因，共享内容不满足需求，导致无法解锁，例如存在一个消费队列，生产者向其中添加产品，消费者从
其中获取产品消费，当生产者锁定此队列并准备向其中添加产品时，发现队列已经满了，这个时候消费者因为此队列也被锁，队列也不会被消费，形成死锁，当然
也可以将锁定放在循环之外，即无论是否列满，都先进行解锁，然后执行添加队列操作，判断是否添加队列操作是否完成，若没有完成则继续锁定，这样就可以避免
死锁，但是这样循环在消费者总是不消费的情况下仍会进行，浪费资源，这种情形，采用条件变量则会显得非常方便

它有三个函数: wait/signal/broadcast

1、条件变量总是与互斥量一起用
2、当条件不允许时，wait，会先解锁当前互斥量，然后堵塞当前线程，需要注意的是解锁互斥量和阻塞当前线程这两步操作是原子性的，也就是说，在其中不允许
任何其他线程锁定该互斥量
3、signal为单发通知
4、broadcast为广播



*/
type shareData struct {
	data int
	cond *sync.Cond
}



func cond_study(){
	s := shareData{
		data:0,
		cond:sync.NewCond(new(sync.RWMutex)),
	}
	NeverStop := make(chan struct{})
	go Add(&s)
	go Del(&s)
	<-NeverStop
	fmt.Printf("over")
}


func Add(s *shareData){
	t := time.NewTicker(time.Second)
	for range t.C {
		s.cond.L.Lock()
		s.data++
		if s.data > 10{
			fmt.Println("i am wrong,need wait")
			s.cond.Wait()
		}
		fmt.Printf("now add data,data is %d \n",s.data)
		s.cond.L.Unlock()
	}
}


func Del(s *shareData){
	t := time.NewTicker(3*time.Second)
	for range t.C {
		s.cond.L.Lock()
		if s.data > 10{
			fmt.Printf("now del data,data is %d \n",s.data)
			s.data = s.data - 10
			fmt.Println("i am right,send signal")
			s.cond.Signal()
		}
		s.cond.L.Unlock()
	}
}