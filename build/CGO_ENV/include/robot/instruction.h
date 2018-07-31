//*******************************************************************  
// Copyright	 
// Description:   
// Update:	 
// Date 			 Name			Description  
// ==========	 ==============    ===================================	
//	
// Known issues: 
//*******************************************************************


#ifndef _INSTRUCTION_H_
#define _INSTRUCTION_H_

#include <stdint.h>
#include <robot/size.h>
#include <robot/robotresource.h>

#define JOBFILE_EXT_NAME	".jbi"

//#define ARCWELDING
//#undef ARCWELDING

#define INSTRCT_MAX_LENGTH 228			   //指令最长的字节数(即MOVJ的字节数)

typedef int instruct_handle;
#define INVALID_HANDLE	((instruct_handle)(-1))
#define IS_VALID_HANDLE(hd)	((hd)>=0)

/*
结构名:STRU_CONST_VAR
说明:常量,变量结构类型
*/
typedef struct _STRU_CONST_VAR{
	int (*getValueFunc)(struct _STRU_CONST_VAR *StruConstVar, double *dbVal);
	int (*setValueFunc)(struct _STRU_CONST_VAR *StruConstVar, double *dbVal);
	union{
		double value;	//for const
		struct{
			RobotPos *c_val;	//for Cxxx address
			unsigned int index;	//for B, I, D, R, P
			unsigned int subindex; //for Pxxx(subindex)
		};
	};
}STRU_CONST_VAR;

typedef struct _STRU_IO_VAR{	//OT, OG, OGH, IT, IG, IGH
	int (*getIOFunc)(struct _STRU_IO_VAR *ioVar);
	int (*setIOFunc)(struct _STRU_IO_VAR *ioVar, int value);
	double (*getIOFunc_var)(struct _STRU_IO_VAR *ioVar);
	int (*setIOFunc_var)(struct _STRU_IO_VAR *ioVar, double value);
	STRU_CONST_VAR var;
}STRU_IO_VAR;

#define INSTRUCT_MAGIC	0x54534E49
#define VALID_INSTRUCT(v)	((v) && (v->head.magic==INSTRUCT_MAGIC))

enum instr_type{
	instru_unknown=0,
	instru_io,
	instru_ctrl,
	instru_cal,
	instru_mov,
	instru_mov_step,
	instru_arc,
	instru_technic,
	instru_other,
};

enum instr_exec_type{
	exec_unknown=0,
	exec_realtime,
	exec_waittime,
};

enum mov_type{
	mov_unkn=0,
	mov_joint,
	mov_line,
	mov_arc,
	mov_Parabola,
};
//指令头部信息
typedef struct {
	unsigned int magic;
	const char *jobname;			//所在的JOB程序号 
	unsigned int nopLine;
	unsigned int nline;		//所在的JOB行数(程序点号)
	enum instr_type type;
	enum instr_exec_type e_type;
}STRU_INSTRUCT_HEAD;

struct _STRU_INSTRUCT;
typedef struct _STRU_INSTRUCT* (*instruct_func_t)(struct _STRU_INSTRUCT *instruct, FILE *pfout);
typedef int (*dump_func_t)(struct _STRU_INSTRUCT *instruct, FILE *pfout);

//机器人指令结构
typedef struct _STRU_INSTRUCT{
	STRU_INSTRUCT_HEAD head; 	//指令头部信息
	instruct_func_t func;  //译码函数(compile)
#ifdef INSTRUCTION_DUMP
	dump_func_t dump_func;
#endif
	unsigned int size;			//instruct interbuf size
}STRU_INSTRUCT;

//IF,UNTIL,ENWAIT
typedef struct {
	STRU_CONST_VAR StrucObj;		  //目标
	uint8_t ucObjType;				  //0:other, 1: IN, IG,B,LB,I,LI,D,LD,R,LR
	uint8_t ucCMP_OP; 				 //比较操作符,0:other, 1:=,<>,>,<,>=,<=,		   
	uint8_t ucExpected;				 //0:other,1:B,LB,byte type constnant ,On/OFF,I,LI,D,LD,R,LR,Real //type constant
	STRU_CONST_VAR StrucExpected;	  //期望值
}STRU_IF;
typedef STRU_IF STRU_UNTIL;

//I/O 命令
typedef struct {
	STRU_IO_VAR io; 		  //输出目标IO		
	STRU_CONST_VAR value;		  //输出值	
}STRU_IO;	//include dout, din

//I/O 命令
typedef struct {
	STRU_CONST_VAR aio; 		  //输出目标AIO		
	STRU_CONST_VAR value;		  //输出值	
}STRU_AIO;	//include AOUT AIN

typedef struct {
	STRU_IO_VAR asf_num;
	STRU_CONST_VAR ac;		//
	STRU_CONST_VAR av;		//
}STRU_ARC;

typedef struct{
	STRU_CONST_VAR dist;
	STRU_ARC arc;
}STRU_ARCCT;

typedef struct {
	STRU_IO_VAR wv_num;
}STRU_WVON;

typedef struct {
	unsigned char type; //sport 1 , mov weld 2
	STRU_CONST_VAR time;
	STRU_CONST_VAR s1;
	STRU_CONST_VAR s2;
}STRU_STITCHON;

typedef struct {
	STRU_CONST_VAR is_on;
}STRU_LASER;
typedef STRU_LASER STRU_LASERTRACK;

typedef struct {
	unsigned char num;
	STRU_CONST_VAR value;
}STRU_HAND;

typedef struct {
	STRU_HAND date;
	STRU_CONST_VAR time;
}STRU_HSEN;

typedef struct {
	STRU_CONST_VAR pose;
}STRU_SFTON;

typedef struct {
	unsigned char disable;
	STRU_CONST_VAR axis;
}STRU_AXISDISABLE;

struct _STRU_cond;
typedef int (*cond_func)(struct _STRU_cond *instruct);
typedef int (*logic_func)(int src1, int src2);

#define IS_COND_EMPTY(c)	((c)->cond==NULL)

typedef struct _STRU_cond{
	STRU_IO_VAR src1;		//操作数1, IN#或者B
	STRU_IO_VAR src2;	//操作数2, B
	cond_func cond;
	logic_func logic;		//| & or null
}STRU_cond;

typedef struct {
	STRU_CONST_VAR timeout;		   //T	(Units: 0.01 seconds)	
	STRU_cond cond[1];	//maybe more than 1. depends on .cond.logc==NULL
}STRU_WAIT;

typedef struct {
	STRU_CONST_VAR time;		   //T	(Units: 0.01 seconds)	
	STRU_IO_VAR dest;		//操作数1, IN#或者B
	STRU_CONST_VAR value;			   //输入数据存储目标 B LB
}STRU_PULSE;


//Control 命令结构定义
typedef struct {
	instruct_handle instr_hd;
	STRU_cond cond[1];	//IF
}STRU_JUMP_CALL;

typedef struct {	
	STRU_CONST_VAR time;		   //T	(Units: 0.01 seconds)	
}STRU_TIMER;

typedef struct {	
	STRU_INSTRUCT_HEAD struHead;	   //命令头部信息
	uint32_t uiDes; 					   //Lable指示的位置
}STRU_LABLE;

//Operating Insturction
typedef struct{
	STRU_CONST_VAR target;	//must be B,LB,I,LI,D,LD
	uint32_t count;	//-1 for all
}STRU_CLEAR;

//single operating, INC, DEC
typedef struct {
	STRU_CONST_VAR data;
}STRU_SOP;

//double operating, SET, ADD, SUB, MUL, DIV,MOD, AND, OR, NOT, XOR
typedef struct {
	STRU_CONST_VAR data1;
	STRU_CONST_VAR data2;
}STRU_DOP;

typedef struct {
	STRU_CONST_VAR dis;
	STRU_DOP cpose;
}STRU_DIST;

typedef struct {
	STRU_CONST_VAR tf_num;
}STRU_TOOL;

typedef struct {	
	STRU_CONST_VAR data1; 		 //设置的目标
	STRU_CONST_VAR data2; 		 //要设置的值
	uint32_t index;
}STRU_DOPE;

typedef struct {	
	STRU_CONST_VAR pnum; 		 //设置的目标
	double pos[AXIS_COUNT]; 		 //要设置的值
}STRU_SETJOINT;


//Move instruction
//ENWAIT instructions
typedef struct {	
	STRU_CONST_VAR struWaitTime;		 
}STRU_ENWAIT;

#define MOV_WAIT_NULL	0
#define MOV_NWAIT		1
#define MOV_ENWAIT		2

typedef enum {
	speed_v,
	speed_vr,
	speed_av,
	speed_avr
}SPEED_TYPE;

typedef enum{
	none,ext_first,ext_second,ext_all
}COOP_TYPE;

typedef struct{
	STRU_CONST_VAR pos;
	int speed_type; // 0:v; 1:vr; 
	union{
		STRU_CONST_VAR v;
		STRU_CONST_VAR avr;
		STRU_CONST_VAR av;
		STRU_CONST_VAR vr;
		STRU_CONST_VAR vj;
	};
	STRU_CONST_VAR tool;
	STRU_CONST_VAR acc;
	STRU_CONST_VAR dec;
	STRU_CONST_VAR wait_t;
	STRU_CONST_VAR coop;
	double pl;	//length
	uint32_t wait;
	STRU_cond cond[1];
}STRU_MOV_COMMON;	//for movj

typedef struct {
	STRU_MOV_COMMON common;
	unsigned char hasFPT;
}STRU_MOV;	//for movl, movc, movs

typedef struct {
	STRU_CONST_VAR  initPhase; //相位
	STRU_CONST_VAR waveLength; //波长
	STRU_CONST_VAR amplitude; //振幅
	STRU_MOV common;
}STRU_MOVX;

typedef struct {
	unsigned int n_refp;
	STRU_MOV common;
}STRU_MOVR;

typedef struct {
	STRU_CONST_VAR sub;
	STRU_CONST_VAR tool_dist;
	STRU_CONST_VAR measure_dist;
	STRU_MOV common;
}STRU_MMOVE;

/*typedef struct {	
	STRU_CONST_VAR struIncPosVar;		 //P 000 to 127
	uint8_t ucIncPosVarType;				 //0:other,1:P,LP,EX,LEX

	STRU_CONST_VAR struIncBPVar;			
	uint8_t ucIncBPVarType;				 //0:other,1:BP,LBP

	STRU_CONST_VAR struIncEXVar;		 //EX 000 to 127
	uint8_t ucIncEXVarType;				 //0:other,1:EX,LEX

	STRU_CONST_VAR struV_VR_VE_VJ;		 //V,VR,VE,VJ
	uint8_t ucVType;						 //0:other,	1:V,2:VR,3:VE,
	
	STRU_CONST_VAR struSingleIncEXVar;	 //仅仅有外部轴的增量 EX 000 to 127
	uint8_t ucSingleIncEXVarType;			 //仅仅有外部轴的增量 0:other,1:EX,LEX
	
	STRU_CONST_VAR struPL;					 
	uint8_t bPL_IF;								//0:other,1:PLEVEL,SPDL,CR
	
	uint8_t bNWAIT_IF;					  
	uint8_t ucUNTIL_NSRCH_Type; ///1:UNTIL, 2:SRCH, 5:NSRCH
	STRU_UNTIL struUNTIL;
	STRU_ENWAIT struENWAIT;
	STRU_CONST_VAR struCoord;			 //坐标 BF,RF,TF,UF,MTF 0:other
	uint8_t ucCoordType;					 //  0:other, BF,RF,TF,UF,MTF

	STRU_CONST_VAR struACC;			   //ACC	
	uint8_t bACC_IF;							   //出现ACC		1:ACC,0:other

	STRU_CONST_VAR stru_DEC;				 //DEC	
	uint8_t DEC_IF;							   //出现DEC		1:DEC,0:other

	uint8_t bPlusIMOV_IF; 			  //+IMOV

}STRU_IMOV;*/


// V: Speed: 0.1 mm to 1500.0 mm/s
// VJ:Speed: 0.01% to 100.00%
// VR:0.1 degrees to 180.0 degrees/s
// VE:Speed: 0.01% to 100.00%
typedef struct {
	STRU_CONST_VAR vj;
	STRU_CONST_VAR v;
	STRU_CONST_VAR vr;
	STRU_CONST_VAR ve;
}STRU_SPEED;

typedef struct {
	enum Robot_Coordinate rc;
	STRU_CONST_VAR nCood;
}STRU_CCOOD;

typedef struct {
	STRU_CONST_VAR nCood;
	STRU_CONST_VAR orig;
	STRU_CONST_VAR xx;
	STRU_CONST_VAR xy;
}STRU_MFRAME;

typedef struct {
	char jobname[MAX_LABEL_STR_LENTH];
	//STRU_CONST_VAR eps;
	STRU_CCOOD coord;
	STRU_MOV common;
}STRU_MOVML;

typedef struct {
	unsigned char state; //0:start 1:getdata 2:end
	STRU_CONST_VAR nTrack;
}STRU_TRACK;

typedef struct{
	unsigned char state; //0:run 1:trigger 2:getdata 3:clean
	STRU_CONST_VAR nVision;
}STRU_VISION;


instruct_handle get_label(const char *jobname, const char *label);
char* add_label(const char *jobname, const char *label, instruct_handle instrhd);

STRU_INSTRUCT * get_next_instruct(STRU_INSTRUCT *instr);
STRU_INSTRUCT *get_next_mov_instruct(STRU_INSTRUCT *instr);
STRU_INSTRUCT *get_instruct_from_line(unsigned int line);
STRU_INSTRUCT *get_preinstruct_from_line(unsigned int line);
STRU_INSTRUCT * get_current_instruct(void);
instruct_handle get_current_instruct_handle(void);
STRU_INSTRUCT* get_instruct(instruct_handle handle);
instruct_handle instruct_request(void);
void instruct_release(void);
void clean_instruct(STRU_INSTRUCT *instr);
void clean_instructhd(instruct_handle instrhd);

void instruct_reset(void);
void* instruct_buffer_request(unsigned int size);
void instruct_buffer_release(unsigned int size);
void* get_instruct_end(void);

#endif
