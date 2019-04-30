#ifndef _EXECUTE_H_
#define _EXECUTE_H_

#include <robot/instruction.h>

void instruction_exec_init(void);
void instruction_exec_exit(void);
void wait_for_running(void);
void exec_info_thread_cleanup(void *arg);
int check_thread_init(void);
void servo_on_of_exit(void );
void send_servo_comm_signal(void);
STRU_INSTRUCT* get_preInstruct(void);
void set_servoEnable(int isEnable);
//////////////////////////////////////////////////////////////////////
int plc_exec_action(void);
int synchronous_abz_pulse(int opt);
void arcct_exec(double len,double residue_len);
void set_stop_action(int num);

int put_movedata_queue(MoveData_t m_md);

unsigned char get_stop_action(void);
#endif
