#include "share/resource.h"
#include "crc.h"

struct nvram_data_s *nvram_data;
uint32_t crc = 0xFFFFFFFF;

char* get_main_file() {
	return GetMainfile();
}

int32_t get_zero_encode(const int x) {
	return Get_ZeroEncode(x);
}

int32_t get_cur_coordinate() {
	return GetCurCoordinate();
}

int watch_nv() { 
	int c = crc32(crc, (unsigned char *)(void *)nvram_data, sizeof(struct nvram_data_s));
	return c;
}

void init_nv_ram() {
	init_nvram(NULL);
	crc32_init();
}