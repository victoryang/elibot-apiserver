// Copyright (c) 2013-2018 elibot. All rights reserved.
// Contact: http://www.elibot.cn

#ifndef MCSERVER_DB_DEBUG_H
#define MCSERVER_DB_DEBUG_H
#define  DEBUG
#ifdef DEBUG
#include <stdio.h>
#define db_printf(fm, args...) printf("%s:%d "fm, __FUNCTION__, __LINE__, ##args)
#else
#define db_printf(...)
#endif

#endif //MCSERVER_DB_DEBUG_H
