#ifndef _EXT_VISION_H_
#define _EXT_VISION_H_

#define INumber 50
#define BNumber 50
#define VNumber 50

typedef struct _VsDt
{
	unsigned int typenum;		//目标物体标号
	float X;               //目标物体在固定坐标系X轴的坐标值（mm）
	float Y;               //目标物体在固定坐标系Y轴的坐标值（mm）
	float Z;				//目标物体在固定坐标系Z轴的坐标值（mm）
	float AngleX;			//目标物体相对于固定坐标系x轴旋转的角度（rad）
	float AngleY;			//目标物体相对于固定坐标系y轴旋转的角度（rad）
	float AngleZ;          //目标物体相对于固定坐标系z轴旋转的角度（rad）
	int EncoderVal;       //目标物体在经过拍照后，记录初始时刻的编码器值
}VisualData;

void vision_clean_data(int num);
int vision_get_data(int num);
int vision_run(int flnb);

int vision_init(void);
int vision_trigger(int num);
void vision_exit();
#endif
