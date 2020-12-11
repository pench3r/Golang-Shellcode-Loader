## Golang Shellcode Loader

在widnows下使用golang编译，准备好windows下的GCC编译器，编译选项：

`CGO_ENABLED=1 GOARCH=386 GOOS=windows`

使用golang实现的shellcode加载器，通过http拉取远端shellcode进行执行

已经编译好的`winloader_x86.exe`会拉取`https://raw.githubusercontent.com/pench3r/pench3r.github.io/master/example/windows_x86_shellcode.bin`中的十六进制shellcode，该shellcode执行后会弹一个窗口

![poc](https://github.com/pench3r/Golang-Shellcode-Loader/blob/main/poc.png)

ps: `raw.githubusercontent.com`域名被污染需要绑定host
