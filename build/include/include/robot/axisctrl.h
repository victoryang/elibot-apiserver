#ifndef __AXIS_CTRL_H__
#define __AXIS_CTRL_H__

#include <list.h>
#include <share/sendpos.h>

typedef struct axis_fun_struct
{
	struct list_head list;
	char *mode;
	//return <0.error; 0.don't support dynamic; >1.support dynamic(axis number), period unit us
	int (*init)(void** priv, const char *device, unsigned int period);
	void (*later_init)(void* priv);
	void (*exit)(void* priv);
	//return 0 for nonblock, 1 for block
	int (*sendpos)(void* priv, sendpos_t *data);
	void (*stop)(void* priv);
	int (*get_abspos)(void* priv, int *pos,unsigned int bchannel,int force);
	int (*get_servomode)(void* priv, int *mod,unsigned int bchannel);
	//return bit 1 for alarm,  number of channel
	unsigned int (*get_alarm)(void* priv, unsigned int nchannel);
	void (*clear_alarm)(void* priv, unsigned int bchannel);
	void (*servo_enable)(void* priv, unsigned int bchannel);
	int  (*get_break)(void* priv, unsigned int bchannel);
	void (*set_output)(void* priv, unsigned int bchannel, unsigned int io_num,unsigned int val); //for test
	int (*get_input)(void* priv, unsigned int bchannel, unsigned int io_num);//for test
	//for inverse dynacic
	void (*change_operation_mode)(void* priv, enum OperationMode mode);
	//for xie zuo
	void (*set_emergency_stop)(void* priv, unsigned char em_stop);
	unsigned int (*get_io_input)(void *priv, int slot_no);//low 16bit enable
	void (*set_io_output)(void *priv, int slot_no, unsigned int iostat);//low 16bit enable
	int (*get_alarm_no)(void* priv, unsigned int nchannel, void *pdat);
	int (*clr_alarm_no)(void* priv, unsigned int nchannel);
	int (*update_servo_firmware)(void* priv, unsigned int servo_no, char *firmware_path);
	int (*get_update_stat)(void* priv, unsigned int servo_no);
	void *priv;
}axis_fun_t;

struct invdyn_func_struct{
	void (*inverse_dynamic_init)(void);
	void (*inverse_dynamic_deinit)(void);
	void (*inverse_dynamic_stop)(sendpos_t *data);
	void (*inverse_dynamic_ctrl)(sendpos_t *data);
	void (*process_feedback_data)(sendpos_t *data);
	int (*get_dynamic_mode)(void);
	void (*set_dynamic_mode)(int isIdentification);
	void (*set_dynamic_paras)(enum SetDynParas option, unsigned int axis_num, double *paras);
	void (*set_dynamic_ctrlmode)(enum OperationMode mode);
	void (*inverse_dynamic_sync_motiondata)(sendpos_t *data);
};

#define BUILDIN_POSFUN(ft)	static axis_fun_t* ft##_buildin_axisfun \
	__attribute__((section("buildin_axisfun"),__used__,aligned(sizeof(void*)))) = {&ft}

int axis_ctrl_init(const char *pos_output);
void axis_output_exit(void);
int axis_output_init(const char *pos_output,char *m_mode);

void axis_output_later_init(void);
void axis_ctrl_exit(void);
void SendPosition(sendpos_t *data);
void StopPosition(void);//call when interpolation stop
unsigned int axis_get_alarm(unsigned int nchannel);
void axis_clear_alarm(unsigned int bchannel);
void axis_servo_enable(unsigned int bchannel);
int axis_get_break(unsigned int nchannel);
void set_emergency_stop(unsigned char em_stop);
void post_readEnCode_cond(void);
void update_MachinePos_by_encode(int abspos[]);
void set_abz2abspulse(void);
//opt: 0-synch by abz(servo), 1-synch by abs(software),2-only display  pulse
int synchronous_abz_pulse(int opt);
void set_servoEnable(int isEnable);

int read_encode_init(void );
void read_encode_exit(void );
int GetAbsPosition(int *pos,unsigned int bchannel,int force);
int GetEncoder(int *ecdcnt,unsigned int axis);
void reset_teach_mod(void);
void pos2pulse(double *pos, int abspos[]);
void pulse2pos(int abspos[], double *pos);
void set_posdata2array(void);
void change_teach_mode(enum OperationMode mode);

#endif
