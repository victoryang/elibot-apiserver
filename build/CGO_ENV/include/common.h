#ifndef _COMMON_H_
#define _COMMON_H_

#include <stdint.h>

#ifndef FALSE
#define FALSE	0
#endif

#ifndef TRUE
#define TRUE	1
#endif

#define SET_BIT(a, b, v)			(a)=((a)&~(1<<(b)))|(((v)&1)<<(b))
#define SET_BIT4(a, b, v)			(a)=((a)&~(0xf<<(b)))|(((v)&0xf)<<(b))

#define GET_ARRAY_BIT(a, b)		(((a)[(b)>>3]>>((b)&7))&1)
#define GET_ARRAY_BIT4(a, b)		(((a)[(b)>>1]>>(((b)&1)*4))&0xf)
#define GET_ARRAY_BYTE(a, b)		((a)[b])

#define SET_ARRAY_BIT(a, b, v)		SET_BIT((a)[(b)>>3], b&7, v)
#define SET_ARRAY_BIT4(a, b, v)		SET_BIT4((a)[(b)>>1], (b&1)*4, v)
#define SET_ARRAY_BYTE(a, b, v)		((a)[b]=v)

#define ARRAY_SIZE(x)	(sizeof(x)/sizeof(*(x)))
#define min(a,b)	((a)<(b)?(a):(b))
#define max(a,b)	((a)>(b)?(a):(b))

#define BOUND(i, v, x)	((v)>(x)?(x):(v)<(i)?(i):(v))

#define QUEUE_SIZE(h, t, m)	(((h)-(t))&((m)-1))
#define ROUND_ADD(a, b, m)	(((a)+(b))&((m)-1))
#define ROUND_INC(x, m)	ROUND_ADD(x,1,m)

#define ROUND_RED(a, b, m)	(((a)-(b))&((m)-1))
#define ROUND_DEC(x, m)	ROUND_RED(x,1,m)

typedef struct{
	unsigned char *name;
	unsigned int blocksize;
	unsigned int inc_step;		//realloc size
	void *data;
	unsigned int cnt;
	unsigned int max;
}dynamic_array;

#define DEFINE_DYARRAY(na, bs, step)	dynamic_array na={.name=#na, .blocksize=(bs), .inc_step=(step),}
#define DYARRAY_HEAD(da)	(da).data
#define DYARRAY_COUNT(da)	(da).cnt

#define DEFINE_DYMEMORY(name, step)		DEFINE_DYARRAY(name, 1, step)

void* dynamic_array_inc(dynamic_array *da);
void* dynamic_array_request(dynamic_array *da, unsigned int size);
void dynamic_array_release(dynamic_array *da, unsigned int size);
void dynamic_array_free(dynamic_array *da);

void strtrim_space(char *dst);
char * skip_blank(const char *s, int *n);
int wordcmp(char *str1, char *str2, char **p);
void crc32_init(void);
uint32_t crc32(uint32_t crc,unsigned char *buffer, unsigned int size);

#include <pthread.h>

struct binary_semaphore{
    pthread_mutex_t mutex;
    pthread_cond_t cvar;
    int v;
};

void binSem_init(struct binary_semaphore *p, int val);
void binSem_post(struct binary_semaphore *p);
void binSem_wait(struct binary_semaphore *p);
void binSem_wait_noclear(struct binary_semaphore *p);
void binSem_clear(struct binary_semaphore *p);

void msleep(unsigned int ms);
unsigned long getMsTimeInterval(struct timeval *from, struct timeval *to);

#define USTIMER_NUM		128
void ustimer_start(unsigned int n);
int ustimer_stop(unsigned int n);
void ustimer_print(unsigned int n);

#endif
