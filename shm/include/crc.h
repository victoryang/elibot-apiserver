#ifndef CRC_H
#define CRC_H
#include<stdint.h>

static uint32_t crc32_table[256];
void crc32_init(void)
{
	uint32_t c;
	unsigned int i, j;	

	for (i = 0; i < 256; i++){
		c = (uint32_t)i;
		for (j = 0; j < 8; j++){
			if (c & 1)
				c = 0xedb88320L ^ (c >> 1);
			else
				c = c >> 1;
		}
		crc32_table[i] = c;
	}
}

uint32_t crc32(uint32_t crc,unsigned char *buffer, unsigned int size)
{
	unsigned int i;
	for (i = 0; i < size; i++){
		crc = crc32_table[(crc ^ buffer[i]) & 0xff]^(crc >> 8);
	}
	return crc;
}

#endif //CRC_H