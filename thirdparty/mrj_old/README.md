 arm-linux-gnueabihf-gcc -I/root/mcserver/include/  -o mrj.o -c mrj.c

 arm-linux-gnueabihf-gcc -I/root/mcserver/include/ -g -Wall *.c -l m

# howto build dynamic and static library
 arm-linux-gnueabihf-gcc -o cJSON.o -c cJSON.c
 arm-linux-gnueabihf-ar rcs cJSON.a cJSON.o
 arm-linux-gnueabihf-gcc -shared -o cJSON.so cJSON.o



 arm-linux-gnueabihf-gcc -fPIC -I/root/mcserver/include/ -o mrj.o -Wall -c mrj.c -lm -lshare
 arm-linux-gnueabihf-gcc -shared -o mrj.so mrj.o

 # compiling and runtime library load(-WL,R)
 LDFLAGS = -L/var/xxx/lib -L/opt/mysql/lib -Wl,R/var/xxx/lib -Wl,R/opt/mysql/lib