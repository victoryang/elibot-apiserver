#ifndef INSTRUMENT_H
#define INSTRUMENT_H
void enable_instrument( void ) __attribute__ ((no_instrument_function));
void disable_instrument( void ) __attribute__ ((no_instrument_function));
void main_destructor( void ) __attribute__ ((no_instrument_function, destructor));
void __cyg_profile_func_enter( void *, void *) __attribute__ ((no_instrument_function));
void __cyg_profile_func_exit( void *, void *) __attribute__ ((no_instrument_function));
#endif
