#ifndef _ERRINFO_H_
#define _ERRINFO_H_

#ifdef __cplusplus
extern "C" {
#endif 

//ErrNo:´íÎóºÅ
//ErrLever:´íÎóµÈ¼¶
void _errInfor(int ErrNo,int ErrLever, int subNo, const char* fun, int line, const char *fmt, ...) __attribute__((format(printf, 6, 7)));
#define errInfor(ErrNo,ErrLever,subNo,fmt,...) _errInfor(ErrNo,ErrLever,subNo, __FUNCTION__, __LINE__, fmt, ##__VA_ARGS__)

int errInfor_init(const char *errlog);

#ifdef __cplusplus
}
#endif

#endif
