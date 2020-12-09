package demo
 
/*
#cgo LDFLAGS: -static -lpsapi -lstdc++ 
//注意这里引用的是mingw 的libpsapi.a，千万不要引用windows sdk下的Psapi.Lib，虽然最终调用的都是
//系统的psapi.dll，但函数导出符号不一样，编译不会通过！！！另外，使用mingw尽量使用-static静态链
//接C++库，不然应用运行时需要libstdc++-6.dll
#include "demo.h"
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)
 
/**
 * @description: 强杀一个服务
 * @param {pid：进程编号, serviceName 进程名称}
 * @return:
	r:操作结果e
	 err：异常
*/
func KillPID(pid int32, serviceName string) (ret uint32, err error) {
	c_serviceName := C.CString(serviceName)//go中开辟的内存传给c是安全的，c函数返回前地址不会变化
	defer C.free(unsafe.Pointer(c_serviceName))
 
	r, err := C.KillPID(C.uint(pid), c_serviceName)
	if err != nil {
		err = fmt.Errorf("KillPID failed errno:%s!", err)
	}
	ret = uint32(r)
	return ret, err
}
 
//导出到C中需要此注释
//export writeInfoLogln
func writeInfoLogln(log string) {
	fmt.Println(log)
}