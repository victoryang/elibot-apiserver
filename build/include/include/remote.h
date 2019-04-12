#ifndef __REMOTE_H__
#define __REMOTE_H__

#include "remote_protocol.h"

int create_remote_server(void);
void exit_remote_server(void);
void open_udpcomm(void);
void open_iobus(void);

#endif
