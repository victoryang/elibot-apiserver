#ifndef _EXTERN_ENCODER_H_
#define _EXTERN_ENCODER_H_

int ext_encoder_init(const char *dev);
int get_ext_encoder(unsigned int ch, int *edata, int *z);
void ext_encoder_close(void);

#endif
