#ifndef _COMPILE_H_
#define _COMPILE_H_
#include <robot/instruction.h>

STRU_INSTRUCT* DOUT_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* MOUT_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* DIN_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* WAIT_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* WAIT_forever_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* PULSE_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* AOUT_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* AIN_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* JUMP_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* CALL_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* TIMER_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* RET_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* END_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* RUNIOCUTIN_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* PAUSE_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* CLEAR_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* INC_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* DEC_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* AND_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* OR_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* NOT_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* XOR_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* GETPOS_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* SET_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* SETP_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* ADD_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* ADDP_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* SUB_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* SUBP_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* MUL_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* MULP_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* DIV_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* DIVP_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* MOD_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* MODP_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* SETE_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* GETE_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* DIST_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* MOVJ_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* MOVL_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* MOVC_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* MOVCA_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* MOVFW_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* MOVS_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* MOVML_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* PREMMOV_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* MMOVE_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* MOVN_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* TRACK_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* VISION_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* AXISDISABLE_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);

STRU_INSTRUCT* ARCON_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* ARCOFF_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* ARCSET_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* ARCCTS_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* ARCCTE_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* SFTOF_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* SFTON_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* WVOF_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* WVON_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* MOVR_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* MOVML_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* MOVX_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* CCOOD_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* MFRAME_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* HAND_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* HSEN_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* LASERON_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* LASEROFF_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* LASERTRACKON_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* LASERTRACKOFF_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* LASERSEARCH_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);

STRU_INSTRUCT* TOOL_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* STITCHON_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* STITCHOFF_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* SETJOINT_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);

STRU_INSTRUCT* SPECIAL_exec_func(STRU_INSTRUCT *instruct, FILE *pfout);
STRU_INSTRUCT* WAIT_exec_onetime_func(STRU_INSTRUCT *instruct, FILE *pfout);

int is_cal_Pvar(STRU_INSTRUCT* instruct);

void enter_move_instruction(void);
void exit_move_instruction(void);
void begin_compile_instruction	(void);
void ClearMovCQueue(void);
void clear_movca_queue(void);
void reset_arcct_data(void);
void set_arcct_start_data(double cur,double vol);
void set_arcct_end_data(double cur,double vol,unsigned char exe,double dist);
void set_analog_cur_double(double cur);
void set_analog_vol_double(double vol);
void set_arcct_type(unsigned char c);
int arc_sys_start(int num);
int arc_sys_stop(int num, int force);
int arc_machine_alarm(unsigned int time);
void reset_stop_action(void);
STRU_INSTRUCT* get_pre_two_instr_data(unsigned int num);

//arc_tool_switch(unsigned int num ,unsigned char onoff);

STRU_INSTRUCT* get_stop_instruct(void);
int load_file_byname(const char *file, FILE *pfout);

void set_current_instr(STRU_INSTRUCT *instr);
void set_preInstruct(STRU_INSTRUCT *instr);

#endif
