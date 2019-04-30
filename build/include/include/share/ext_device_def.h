#ifndef _EXT_DEVICE_DATADEF_
#define _EXT_DEVICE_DATADEF_

#ifdef __cplusplus
extern "C" {
#endif
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
int yiputeng_init(void);

#ifdef __cplusplus
}
#endif
#endif 
