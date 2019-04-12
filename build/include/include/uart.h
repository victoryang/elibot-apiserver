#ifndef __RS485_H__
#define __RS485_H__

#define uart_num	5
//void rs485_closed(void);
int uart_open(const char *uart_name);
int uart_Send(int fd, char *send_buf,int data_len);
int uart_recv(int fd, char *rcv_buf,int data_len);
int uart_setopt(int fd,int nSpeed, int nBits, char nEvent, int nStop);
void uart_Close(int fd);

#endif

