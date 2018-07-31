#ifndef __PLC__H__
#define __PLC__H__

#include <share/resource.h>

#ifdef __cplusplus
extern "C"{
#endif /* __cplusplus */ 
//P指针
#define   CSP_Pn_MAX	127
//Counter  
#define		Highspeed_Counter_index	244
#define		Highspeed_Counter_index_Dcounting	250
#define		Highspeed_Counter_index_Dphase		253
#define		Highspeed_Counter_Num	12
#define		Counter_index_16Bit		0
#define		Counter_index_32Bit		200

/*****************************************************************************\
	PLC define
\*****************************************************************************/
#define   _X_num      128    // 输入端口，    编号：X0-X127
#define   _Y_num      128    // 输出端口，    编号：Y0-Y127
#define   _M_num      1536   // 中间M继电器，编号：M0-M1535
#define   _M8xxx_num  256    // 中间M8xxx继电器， 编号：M8000-M8256
#define   _S_num      1000   // 中间S继电器，编号：S0-S999
#define   _T_num      256    // 定时器，      编号：T0-T199:100ms, T200-T245:10ms, T246-249:1ms, T250-255:100ms
#define   _C_num      256    // 计数器，      编号：C0-C255
#define   _D_num      8256    // 位元件 D，    编号：D0-D8256
#define	  _AD_num		8	//模拟输入端口	

#define TIMER_index_1MS			246
#define TIMER_1MS_NUM			4
#define TIMER_index_10MS		200
#define TIMER_10MS_NUM			46
#define TIMER_index_100MS		0
#define TIMER_100MS_NUM			200
#define TIMER_index_100MS_C		250
#define TIMER_100MS_NUM_C		6
//PLC位元件地址偏移量
#define OFFSET_S 0
#define OFFSET_X 0x400
#define OFFSET_Y 0x500
#define OFFSET_T 0x600
#define OFFSET_M 0x800
#define OFFSET_C 0xE00
#define OFFSET_M8xxx 0xF00

//=============================================================================
#define   _X_BYTE     (_X_num + 7) / 8    // 48个输入端口，所占内存字节数
#define   _Y_BYTE     (_Y_num + 7) / 8    // 48个输出端口，所占内存字节数
#define   _M_BYTE     (_M_num + 7) / 8    // 256个中间M继电器，所占内存字节数
#define   _S_BYTE     (_S_num + 7) / 8    // 256个中间S继电器，所占内存字节数
#define   _T_BYTE     (_T_num + 7) / 8    // 32个定时器，所占内存字节数
#define   _C_BYTE     (_C_num + 7) / 8    // 32个计数器，所占内存字节数
#define   _M8xxx_BYTE (_M8xxx_num+7)/8    // 24个M8xxx继电器，所占内存字节数

typedef enum{
	plc_input=0,
	plc_output,
	plc_M,
	plc_S,
	plc_T,
	plc_C,
	plc_M8,
	plc_D,
}plc_mem_enum;

typedef struct{
	void *input;
	void *output;
	void *M;
	void *D;
}plc_mem_t;

typedef enum {
	PLC_RUNST_RUN=0x09,
	PLC_RUNST_STOP=0x0a,
}plc_run_state;


typedef struct{
	unsigned char *rom_addr;
	unsigned int rom_size;
	plc_mem_t plcmem;

	void (*save_code)(unsigned char*, unsigned int);
	void (*refresh_input)(unsigned char*);
	void (*refresh_output)(const unsigned char*);
}plc_device_t;


typedef struct{
	plc_device_t dev;
	void *access_mutex;
	unsigned long last1ms, last10ms, last100ms;
} plc_device_priv_t;


void plc_init(plc_device_t *plcdev);
void plc_cycle(void);
void* plc_getMem(plc_mem_enum type);
//return tx_buf size
int plc_download(const unsigned char *rx_buf, unsigned int nrx, unsigned char *tx_buf);
void plc_reset_ram(void);
void create_plc_client(unsigned int  port);
int plc_exec_bookprog(void);
int plc_exec_comm_user(int comm_sock);
#ifdef __cplusplus
}
#endif /* __cplusplus */ 


#endif

