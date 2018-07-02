#ifndef NV_H
#define NV_H

int get_axis_count();

char* get_main_file();

int32_t get_sys_technics();

int32_t get_cur_coordinate();

int32_t get_mannual_speed_rate();

int32_t get_cycle_mode();

int32_t get_tool_number();

int32_t get_user_number();

int32_t get_zero_encode(const int x);

double get_system_period();

int32_t get_sys_ctrl_mode();

double get_origin(const int x);

int init_nv_ram();

#endif //NV_H