#ifndef __AIO_H__
#define __AIO_H__

#include <stdint.h>

int analog_io_init(const char *aio_dev);
void analog_io_deinit(void);
void set_analog_io(unsigned int ch, uint32_t v);
uint32_t get_analog_io(unsigned int ch);
void set_analog_io_double(unsigned int ch, double v);
double get_analog_io_double(unsigned int ch);

#endif
