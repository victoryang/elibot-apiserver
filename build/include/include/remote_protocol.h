#ifndef __REMOTE_PROTOCOL_H__
#define __REMOTE_PROTOCOL_H__

#include <stdint.h>
#include "hash_define.h"

#ifdef __cplusplus
extern "C" {
#endif 


struct robot_remote_head_s{
	uint32_t magic;
	char payload[0];
};

struct robot_remote_start_head_s{
	uint32_t magic;
	int packnmb;
	int len;
	char payload[0];
};

struct robot_remote_part_head_s{
	uint32_t magic;
	int len;
	char payload[0];
};

struct robot_pos_s{
	uint32_t nAxis;
	uint32_t pad;
	double machinePos[0];	//\u673a\u5668\u4eba\u4f4d\u7f6e ,\u5173\u8282\u89d2\u5ea6
};

struct robot_status_s{
	uint16_t robotMode:2; //\u673a\u5668\u4eba\u63a7\u5236\u6a21\u5f0f
	uint16_t cycleMode:2; //\u673a\u5668\u4eba\u5faa\u73af\u72b6\u6001, CYCLE_MODE_...
	uint16_t curCoordinate:3; //\u5f53\u524d\u5750\u6807\u7cfb	
	uint16_t robotState:3;	//\u5f53\u524d\u8fd0\u884c\u72b6\u6001
	uint16_t servoReady:1;
	uint16_t robSpeedNo:2; //speed grade, 0-3
	//int16_t :3
	uint16_t currentLine; //current run line No.
	char project_name[0];
};

#define RB_POS_MAGIC	HASH4_S('P', 'O', 'S', '_', sizeof(struct robot_pos_s))
#define RB_STATUS_MAGIC	HASH4_S('S', 'T', 'A', 'T', sizeof(struct robot_status_s))
#define RB_SHARE_RSC_MAGIC(n) HASH4_S('R', 'S', 'R', n, PACK_SIZE)
#define RB_SHARE_RSCX_MAGIC(n) HASH4_S('X', 'O', 'R', n, PACK_SIZE)
#define RB_NVRAM_MAGIC HASH4_S('N', 'V', 'R', 'M', NVRAM_SIZE)

#define ROBOT_REMOTE_SEND_PERIOD		100		//ms
#define ROBOT_REMOTE_FORCE_SEND_PERIOD	2000	//ms

#define DEFAULT_REMOTE_PORT	8056

#ifdef __cplusplus
}
#endif

#endif