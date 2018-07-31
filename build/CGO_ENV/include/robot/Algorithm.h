#ifndef _ALGORITHM_H_
#define _ALGORITHM_H_

#include <robot/interpolation.h>
#include <robot/posinv.h>
#include <robot/math_func.h>

#define P_TYPE_PULSE	0
#define P_TYPE_POSE	1
#define ForeReadNt 0   //前瞻的最大段数,because of mode teach ,then continue running ,it must be > 0 zxc2017.1.8

void ChangePVariableTOMatrix(int pType[3],double P1[],double P2[], double P3[],double M1[4][4],double M2[4][4],double M3[4][4]);

int ShengjinDistinguishingMeans(double a,double b,double c, double d, double retx[]);

//int IsMovThrough(int TargetPType,double startJoint[],double targetPoint[]);
int IsMacrocephalicPoint(int TargetPType,double startJoint[],double targetPoint[]);
int IsCanMOVL(double curjiont[],double targetjoint[]);

void SetMovCommInterpData_S(AutoMovCommonS *movData, MoveData_t *md);

void PoseCalculate(AutoMovCommonS *interdata);

void CalLineData_S(int pType,MoveData_t  *md,AutoMovlInterpDataS *xyzData);

int ThreePointSituation(int pType[],double P1[],double P2[], double P3[], int coop);
void hcmatrix(double ff[3],double si,double matrix[4][4],double starMax[4][4]);

int CalCircularArc_S(int pType[3],double startJoint[],double midJoint[],double targetJoint[],AutoMovcDataS *movcData); //外部译码用圆弧参数计算函数
void CalCircularArc_S_interp(MoveData_t *md,AutoMovcInterpDataS *movcData);//插补用圆弧参数计算函数

double CalNextX(double currentX, double aa, double stepLength, double ratio);
void CalParabolaArc_S(int pType[3],double startJoint[],double midJoint[],double targetJoint[],AutoMovsDataS *movsData); 
void CalParabolaArc_S_interp(MoveData_t *md,AutoMovsInterpDataS *movsData);
void CalSinData_S(int pType[3],MoveData_t *md,AutoMovxInterpDataS *movxData); //外部译码用圆弧参数计算函数
double calc_matrix_dist(double start[4][4],double target[4][4]);

double SAccelerateCalculateNewFeedrate(int base);

void FirstRoundPreProcess_S(int base);
int SAccelerateCalculateNewAcc(int base);
AutoMovCommonS *get_acc_move_data_update(unsigned int offset);

void bubble_sort(double *nArray,int len);
void GearspeedReCalculate(MoveData_t *CurCoordinate);

#endif // ALGORITHM_H

