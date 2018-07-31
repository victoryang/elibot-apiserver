#ifndef _WEAVING_H_
#define _WEAVING_H_
#include "../../config.h"

#ifdef __cplusplus
extern "C" {
#endif 
typedef enum _weavType
{
	SINGLE=0,
	TRIANGLE,
	L_TYPE
}weaveType;
typedef enum _ComplexityMode
{
	Simple=0,
	complex
}ComplexityMode;

typedef struct _WeavingInfor{
	ComplexityMode mode;//Simple/complex
	weaveType type;//single type, triangle type, L-type ,and so on
	double amplitude;//摆动幅度
	double vertical;//壁方向长度
	double horizontal;//水平方向长度
	double frequency;//频率
	double angle;//两壁夹角
	int direct;// swing direct 1:froward 0:backward
	double stoptime[4];//every section stop time
	char note[32];//describe note
	unsigned char useable;//当前工艺号参数是否可用
}WeavingInfor;

#define WEAVE_TECHNICS_COUNT		8

typedef struct _WeavingPara{
	unsigned char enable;
	int tech_num;//0-(WEAVE_TECHNICS_COUNT]
	double refp[2][AXIS_COUNT];
	WeavingInfor weave_infor[WEAVE_TECHNICS_COUNT];
}WeavingPara;

#ifdef __cplusplus
}
#endif

#endif
