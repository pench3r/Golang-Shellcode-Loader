
#if !defined _DEMO_H_
#define _DEMO_H_
 
#ifdef __cplusplus
extern "C" { //导出C接口
#endif
 
int KillPID(unsigned int pid, char* srvName);//注意函数签名中不要带有C++的元素
 
#ifdef __cplusplus
}
#endif
 
#endif