#ifndef LOGMESSAGE_H
#define LOGMESSAGE_H

#include <syslog.h>

#define LOG(...)	syslog(LOG_USER|LOG_INFO, __VA_ARGS__)
#define LOG_debug(...)	syslog(LOG_USER|LOG_DEBUG, __VA_ARGS__)
#define LOG_err(...)	syslog(LOG_USER|LOG_ERR, __VA_ARGS__)
#define LOG_warn(...)	syslog(LOG_USER|LOG_WARN, __VA_ARGS__)

#define SHOW_FAILED(x)	LOG_err("%s@%d failed %s\n", __FUNCTION__, __LINE__, #x)
#define ASSERT_AND_RETV(x, v) if(!(x)){SHOW_FAILED(x);return v;}
#define ASSERT_AND_RET(x) if(!(x)){SHOW_FAILED(x);}


#ifdef DEBUG
#define DPRINTF(fm, args...) printf("%s:%d "fm, __FUNCTION__, __LINE__, ##args)
#else
#define DPRINTF(...)
#endif

#define DUMP_VECTOR6(p)	DPRINTF("[%lf,%lf,%lf,%lf,%lf,%lf]", p[0],p[1],p[2],p[3],p[4],p[5]);

#endif //logmessage.h
