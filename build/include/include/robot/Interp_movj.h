#ifndef _INTERP_MOVJ_H_
#define _INTERP_MOVJ_H_

double SearchResult_4(double MinSpeed, double MaxSpeed, double Time, double Jcc, double distance);
int MOVJ_interpolation_S(MoveData_t *md);
char PreMovjPositionLevel(MoveData_t *inputmd, MoveData_t outputmd[3],int type);
int update_automovj_dataS(AutoMovjDataS amj[], STRU_MOV_COMMON *mov, int mov_done);
void MovjPreDataProcess_S(AutoMovjDataS automovjdata[],int mode);
void AutoMovjPreDataProcess_S(AutoMovjDataS automovjdata[],int mode);
void init_SpeedConstraintPara(void);
double GetJointMaxAcc(int axis);
double GetJointMaxJcc(int axis);
double GetSmallLength(int axis);
void SetMoveSmallLengthState(int sign);
int GetMoveSmallLengthState(void);
double GetDecFactor(void);
double GetMotionSmooth(int axis);
double Calculate_Pl_length(int axis, int IsSmallSegment, unsigned int pl_level);
double GetOverlapTime(double totaltime, unsigned int pl_level);
double Search_Maximum_Time(MoveData_t *md);
void Printf_Matrix_float(char *name, double *matrix, int Row, int Column);

#endif
