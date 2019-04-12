#ifndef __TECHNICS_H__
#define __TECHNICS_H__
#include <robot/instruction.h>
#include <robot/interpolation.h>

enum arc_io_num {alloff,hanjie,song_qi,song_si,allon};

enum stop_type{normal_stop,pause_stop,modechange_stop,error_stop};

struct struTechnics{
	int name;
	int (*func_init)(void);
	int (*func_start)(int num);
	int (*func_stop)(int num, int force);
	int (*func_machine_alarm)(unsigned int flg);
};

#define BUILDIN_TECHNIC(n)	static struct struTechnics n##_buildin_technic \
	__attribute__((section("buildin_technic"),aligned(sizeof(void*)),__used__))= \
	{.name=n,\
	.func_init=n##_init,\
	.func_start = n##_start,\
	.func_stop=n##_stop,\
	.func_machine_alarm=n##_machine_alarm}

struct last_stop_infor{
	int s_num; //technics number
	int line;//行号
	int bline;//NOP行号
	unsigned char is_action; //停止时是否工艺在执行中
	enum stop_type s_type; //停止类型
	double stop_pos[AXIS_COUNT];//停止时的位置
	STRU_INSTRUCT *stop_instru;//停止时的指令令位置
	struct struTechnics *technic_infor;
};

int technic_init(void);
void technic_start(unsigned int force);
void technic_stop(enum stop_type type, STRU_INSTRUCT* instru, int force);

void reset_stop_action(void);
void set_stop_action(int num);
unsigned char get_stop_action(void);
STRU_INSTRUCT* get_stop_instruct(void);
int continue_technic(unsigned int prog_line,int* errline);
void get_technic_stop_pos(double *pos);
STRU_INSTRUCT* set_stop_instruct(STRU_INSTRUCT* instru);
int technic_on(MoveData_t *md);

#endif

