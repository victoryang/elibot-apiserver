#ifndef _HASH_DEFINE_H_
#define _HASH_DEFINE_H_

#define HASH1(v, h)	(((h)*33)+(unsigned long)(v))
#define HASH4(a,b,c,d)	HASH1(d,HASH1(a,HASH1(b,HASH1(a,0))))
#define HASH4_S(a,b,c,d,s)	HASH4((a)+((s)>>24), (b)+((s)>>16), (c)+((s)>>8), (d)+(s))

#endif
