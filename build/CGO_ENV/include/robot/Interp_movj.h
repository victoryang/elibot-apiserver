#ifndef _INTERP_MOVJ_H_
#define _INTERP_MOVJ_H_

double SearchResult_4(double MinSpeed, double MaxSpeed, double Time, double Jcc, double distance);
int MOVJ_interpolation_S(MoveData_t *md);
int MOVJ_interpolation_Dual_S(MoveData_t *md, MoveData_t *md1);
//int PreMovjPositionLevel(AutoMovjDataS inputmovj[], AutoMovjDataS outputmovj[3][AXIS_COUNT],int type);
char PreMovjPositionLevel(MoveData_t *inputmd, MoveData_t outputmd[3],int type);
int update_automovj_dataS(AutoMovjDataS amj[], STRU_MOV_COMMON *mov, int mov_done);
void MovjPreDataProcess_S(AutoMovjDataS automovjdata[],int mode);
void set_movj_dec_data(AutoMovjDataS automovjdata[]);
void set_movj_dec_data_c3(AutoMovjDataS automovjdata[]);
void set_movj_dec_data_c0(AutoMovjDataS automovjdata[]);
void set_movj_dec_data_c1(AutoMovjDataS automovjdata[]);


#endif
