#if !defined _DEMO_H_
#define _DEMO_H_
 
#ifdef __cplusplus
extern "C" { //导出C接口
#endif
 
int faker_window();
char* http_download(char *payload_uri);
 
#ifdef __cplusplus
}
#endif
 
#endif