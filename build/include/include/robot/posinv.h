#ifndef _POSINV_H_
#define _POSINV_H_
#ifdef __cplusplus
extern "C" {
#endif

#define iszero( x)	(fabs(x)<em?1:0)

struct _robot_posInv_t;

typedef struct {
	struct _robot_posInv_t *posinv;
	void *priv;
}robot_posInv_data;

typedef int (*posinv_init_func_t)(void);
typedef void (*joint2matrix_func_t)(double[][4],double[]);
typedef void (*matrix2joint_func_t)(double[],double[][4],int*);
typedef void (*joint2pose_func_t)(double[],double[]);
typedef void (*pose2joint_func_t)(double[],double[], int*);

typedef struct _robot_posInv_t{
	posinv_init_func_t init;
	joint2matrix_func_t joint2matrix;
	matrix2joint_func_t matrix2joint;
	joint2pose_func_t joint2pose;
	pose2joint_func_t pose2joint;
}robot_posInv_t;

extern robot_posInv_data posinv_data;

void baseMatrix_decTool(double matrix[4][4]);
void baseMatrix_addTool(double matrix[4][4]);

#define ConvertJointToMatrix(m,j)	do{posinv_data.posinv->joint2matrix(m,j);baseMatrix_addTool(m);}while(0)
#define ConvertMatrixToJoint(j,m,o)	do{double t_m[4][4];memcpy(t_m,m,sizeof(double)*16);baseMatrix_decTool(t_m);\
										posinv_data.posinv->matrix2joint(j,t_m,o);}while(0)
#define PositiveKinematic(p,j)	posinv_data.posinv->joint2pose(p,j)
#define InverseKinematic(j,p,o)	posinv_data.posinv->pose2joint(j,p,o)

void posinv_solver_init(void);
void common_joint2pose(double pos[6], double xp_joint[]);
void common_pose2joint(double xp_joint[],double pos[],int *outside);

void joint2matrx_tool(double matx[][4],double joint[]);
void joint2matrx_robot(double matx[][4],double joint[]);
void matrx2joint_tool(double joint[],double matx[][4],int *out);
void matrx2joint_robot(double joint[],double matx[][4],int *out);

double SelectInfiniteAxisAngle(double startAngle,double targetAngle,int num);
double SelectAngle(double d_v1,double d_v2,double joint);
int GetToolMatrix(double toolmatrix[4][4]);
void GetToolMatrix_XYZ(double toolmatrix[4][4]);

double CalculateRadSIN(double qjoint7, double null_midQ7, double null_ampQ7);
void GetNullSpaceLimit(double *null_midQ7, double *null_ampQ7, double qPoint[]);
void UpdateMovnJointPos(double xp_joint[], double delta_step, double matrix[][4], double null_midQ7, double null_ampQ7, int *outside);
void UpdateMovnSpeed(double *curSpeed, double maxSpeed, double RadSIN, double RadSIN_start, double RadSIN_target);

typedef double (*CalculateRadSIN_func_t)(double,double,double);
typedef void (*GetNullSpaceLimit_func_t)(double*,double*,double[]);
typedef void (*UpdateMovnJointPos_func_t)(double[],double,double[][4],double,double,int*);
typedef void (*UpdateMovnSpeed_func_t)(double*,double,double,double,double);

typedef struct _robot_movn70_t{
	CalculateRadSIN_func_t CalculateRadSIN;
	GetNullSpaceLimit_func_t GetNullSpaceLimit;
	UpdateMovnJointPos_func_t UpdateMovnJointPos;
	UpdateMovnSpeed_func_t UpdateMovnSpeed;
}robot_movn70_t;

extern robot_movn70_t movn70_solver;

#ifdef __cplusplus
}
#endif

#endif // _POSINV_H_

