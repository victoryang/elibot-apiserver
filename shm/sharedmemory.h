#ifndef SHAREDMEMORY_H
#define SHAREDMEMORY_H
#include "share/resource.h"
#include "crc.h"

int get_press_reset();

int32_t get_zero_encode(const int x);

void init_shm();

void GetAll();

#endif //SHAREDMEMORY_H