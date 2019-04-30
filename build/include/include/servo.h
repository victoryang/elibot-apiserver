
#ifndef __SERVO_H__
#define __SERVO_H__

typedef struct{
	char name[128];
	int (*servo_func_send)(int index);
	int (*servo_func_receive)(void);
	void (*servo_func_init)(void* priv);
	int (*servo_func_calc_encode)(int index);
}STRU_SERVO;

#define BUILDIN_SERVOCOMM(n)	static STRU_SERVO n##_buildin_servo \
	__attribute__((section("buildin_servo"),aligned(sizeof(void*)),__used__))= \
	{.name=#n,\
	.servo_func_init=n##_servo_comm_init,\
	.servo_func_send = n##_servo_comm_send,\
	.servo_func_receive=n##_servo_comm_receive,\
	.servo_func_calc_encode=n##_servo_calc_encode}
void servo_comm_continue(void);
int servo_comm_init(void *priv);
void servo_comm_exit(void);

int send_servo_comm_data(char *data,int num);
int get_servo_comm_data(char *data);
int servo_get_abs_encoder(int index, int *value);
void flush_uart_mem(void);

#endif

