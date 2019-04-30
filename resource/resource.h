#ifndef RESOURCE_H
#define RESOURCE_H

#include "stdlib.h"

#include "share/resource.h"
#include "cJSON.h"

#define M_OUT_START 400

#define GetLocVarcRobLB 0
#define GetLocVariRobLI 1
#define GetLocVardRobLD 2
#define GetLocVardRobLP 3
#define GetLocVardRobLV 4

typedef void (*getLocVarFunc)(cJSON*, int, int, int);

char* elt_get_locvar(int datatype, int number, int start, int end);

#define GetSysVarcRobB 0
#define GetSysVariRobI 1
#define GetSysVardRobD 2
#define GetSysVardRobP 3
#define GetSysVardRobV 4

typedef void (*getSysVarFunc)(cJSON*, int, int);

char* elt_get_sysvar(int datatype, int start, int end);

int elt_get_io_in(int addr);

int elt_get_io_out(int addr);

int elt_get_io_vin(int addr);

int elt_get_io_vout(int addr);

char* elt_get_plc();

char* elt_get_main_program_name();

int elt_get_manual_speed_rate();

int elt_get_tool_num();

int elt_get_user_num();

char* elt_get_nv();

int elt_get_robot_state();

char* elt_get_cur_robot_pos();

char* elt_get_cur_robot_pose();

int elt_get_cur_program_line();

int elt_get_servo_enabled();

int elt_can_motor_run();

int elt_get_speed_modify_play();

int elt_get_robot_mode();

char* elt_get_resource();

int elt_get_remote_mode_status();

int elt_get_encryption_status();

int elt_get_encryption_remain_time();

char* elt_get_machine_code();

int elt_resource_init();

#endif //RESOURCE_H