#ifndef __RT_THREAD_H__
#define __RT_THREAD_H__

#include <pthread.h>

//PRREMPT_RT must under 50, use 50 as the priority of kernel tasklets and interrupt handler by default
void rt_thread_create(pthread_t *thread, int priority, void *(*start_routine)(void *), void *arg);

#endif
