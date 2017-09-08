package rate

import (
	"golang.org/x/time/rate"
	"context"
	"fmt"
	"time"
)


/*
令牌桶，限速的利器，与传统的漏桶限速不同，它可以应对突然间的流量增加

r := rate.NewLimiter(1,10)
的意思是每秒一次，桶的大小为10
r.Wait(ctx)每次循环时都会重新计算离下一次请求的间隙，可以通俗的理解，如果桶中还有剩余的令牌，则不堵塞，直接运行，
如果桶中没有令牌了，则需要遵循一秒一次的设定，也就是NewLimiter第一个参数的值

可以通过r.SetLimit(rate.Limit(5))调速度，为每秒钟五次，数值越大，速度越快
*/

func RateLimitor(){
	r := rate.NewLimiter(1,10)
	ctx := context.Background()
	times := 0
	for {
		r.Wait(ctx)
		times++
		if times > 10{
			r.SetLimit(rate.Limit(2))
		}
		if times > 20{
			r.SetLimit(rate.Limit(5))
		}
		if times > 30{
			time.Sleep(1*time.Second) //两秒钟之后桶就满了
			times = times - 30
		}
		fmt.Printf("rate limitor,time is %d\n",times)
	}
}


//TODO 自己写一个令牌桶算法实现
