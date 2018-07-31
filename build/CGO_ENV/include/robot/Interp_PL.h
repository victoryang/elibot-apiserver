#ifndef _INTERP_PL_H_
#define _INTERP_PL_H_
//type: 0:加速区和减速区均普通 1:加速区普通减速区要并 2:加速区要并减速区普通 3: 加速区减速区都要并
typedef enum PL_TYPE{
	PL_FN_BN,
	PL_FN_BH,
	PL_FH_BN,
	PL_FH_BH	
}PLTYPE;
int MOV_interpolation_Dual_S(MoveData_t *md1, MoveData_t *md2);

char PreMovPL(MoveData_t *input_md,MoveData_t *output_md,int type);
void PreProcess_S(MoveData_t *md,int type);//initialize md->comm before interpolation
void ExtPreDataProcess_S(MoveData_t *md,int combineFlag);
double MOVJ_Stepinterpolation_S(double *poschip,MoveData_t *md,double *machinePos, int *isover);
MoveData_t * reUpdata_movjData_forStop(MoveData_t *md,double curPos[]);
void set_preview_pl(double v_pl);
double get_preview_pl(void);
#endif
