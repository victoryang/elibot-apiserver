#ifndef CXZK_H
#define CXZK_H

#define LASER_ON_OFF_FAILED -1
#define SEARCH_FAILED -2
#define SEAM_NOT_FOUND -3
#define TRACE_FAILED -4
#define TCP_SEND_ERR -5
#define TCP_RECV_ERR -6
#define SOCKET_FAILED -7
#define CONNECT_ERR -8

#define LASER_IS_ON 1
#define LASER_IS_OFF 7
#define SEARCHING_SEAM 2
#define SEAM_FOCUSED 3
#define TRACE_KEEP_ON 4
#define TRACE1_IS_ON 5

enum sm_func {LASER_ON=0,SEARCH,TRACE_KEEP,TRACE1_ON,TRACE_OFF,LASER_OFF} ;

int smtrk_cxzk_init(void);
int cmd_cxzk_laser_switch(unsigned char is_on);
int cmd_cxzk_trace_switch(unsigned char is_on);
int cmd_cxzk_search(void);
void smtrk_cxzk_exit(void);

#endif // SEAM_TRACKING_FUNC_H
