#ifndef _EXT_DEVICE_DATADEF_
#define _EXT_DEVICE_DATADEF_

#ifdef __cplusplus
extern "C" {
#endif 
//for chuangxiang zhikong
typedef struct {
	int manufacID;
	char func_type;//'f':all func or m:micro func
	char IPaddres[16];
	int serverport;
	int isused;
	int mode;//trace mode
	double accuracy;
	double distance;//from tracer to torch,fixed
	double freq;//30ms~60ms
	int schcnt;//search count
	int sta;//sta data
	double ex; //X error
	double ez; //Z error
}weld_offset;

//for keyence
typedef struct {
	unsigned int isused;
	unsigned int tool_measure;
	unsigned int tool_work;
	unsigned int analog_port;
	double valid_dist;
	double min_analog;
	double max_analog;
	double min_dist;
	double max_dist;
}DIST_SENSER_INFOR;

/**for keyence********************************************************/
void init_keyence(void);
double get_senser_distance(void);
//<0 is error
int get_tool_senser(double tool[]);
//<0 is error
int get_tool_work(double tool[]);
/**for keyence********************************************************/

void init_robotiq(void);
void catch_object(unsigned char pos);

#ifdef __cplusplus
}
#endif
#endif 
