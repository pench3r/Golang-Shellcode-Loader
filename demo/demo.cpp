#include "demo.h"
#include <Windows.h>
#include <Psapi.h>
#include <errno.h>
#include <stdio.h>
#include "_cgo_export.h"//go中的导出函数的声明，此文件自动构建
 
void WriteLog(const char* log){
    _GoString_  gs;//go中的string类型导出到_cgo_export.h文件中，类似的还有切片
    gs.p = log;
    gs.n = strlen(log);
    writeInfoLogln(gs);//go中导出打印日志的方法
}
int KillPID(unsigned int pid, char* srvName)
{
    char log[300];
    int bRet = 0;
    HANDLE proc = OpenProcess(PROCESS_TERMINATE, FALSE, pid);
    if (proc)
    {
        bRet = TerminateProcess(proc, 2);
        CloseHandle(proc);
    }else{
         errno = GetLastError();
         return bRet;
    }
    sprintf(log, "Killing srv:%s", srvName);
    WriteLog(log);
    bRet = 1;
    return bRet;
}