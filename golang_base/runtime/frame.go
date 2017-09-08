package runtime

import (
	"runtime"
	"fmt"
)


/*
主要介绍runtime包里面Caller、Callers、CallersFrames函数的用法
从字面上可以看出来，这些方法主要是为了获取调用者信息的，这些信息可以方便地用于记载日志、错误输出／记录等，非常实用
Caller



*/
func CallersStudy(){
	BeCalled()
}


func BeCalled(){
	pc,file,line,ok:=runtime.Caller(1)
	if ok{
		fmt.Printf("pc is %v\n",pc)
		fmt.Printf("file is %v\n",file)
		fmt.Printf("line is %v\n",line)
	}

	pcs := make([]uintptr,32)
	count := runtime.Callers(2,pcs)
	fmt.Printf("count is %v\n",count)
	fmt.Printf("pcs is %v\n",pcs)

	frames := runtime.CallersFrames(pcs[:count])
	for {
		frame,more:= frames.Next()
		fmt.Printf("call entry is %v,call function is %v,call file is %v,call line is %v\n",frame.Entry,frame.Function,frame.File,frame.Line)
		if !more{
			break
		}
	}
}