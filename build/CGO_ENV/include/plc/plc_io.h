#ifndef __PLC_IO_H__
#define __PLC_IO_H__

#include <list.h>
#include <stdint.h>

typedef struct plcio_fun_struct
{
	struct list_head list;
	char *mode;
	int (*init)(struct plcio_fun_struct* pft, const char *device);
	void (*exit)(struct plcio_fun_struct* pft);
	//return 0 or -1
	int (*readio)(struct plcio_fun_struct* pft, uint8_t io[], unsigned int byteofio);
	void (*writeio)(struct plcio_fun_struct* pft, const uint8_t io[], unsigned int byteofio);
	void *priv;
}plcio_fun_t;

#define BUILDIN_PLCIOFUN(ft)	static plcio_fun_t* ft##_buildin_plciofun \
	__attribute__((section("buildin_plciofun"),__used__,aligned(sizeof(void*)))) = {&ft}

int plcio_output_init(const char *pos_output);
void plcio_output_exit(void);
void plc_readio(uint8_t io[], unsigned int byteofio);
void plc_writeio(const uint8_t io[], unsigned int byteofio);

#define DEFAULT_NETIO_PORT	8051

#endif
