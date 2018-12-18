CC = arm-linux-gnueabihf-gcc
AR = arm-linux-gnueabihf-ar
SHARED = -shared
RCS = rcs

CPPFLAGS += -I/root/mcserver/include/ -fPIC -Wall

.PHONY: all

all: libmrj.so

libmrj.so: mrj.o mcnv.o mcresource.o mcplc.o mcvars.o cJSON.o
	$(CC) $(SHARED) -o $@ $^

test: main.c mrj.o mcnv.o mcresource.o mcvars.o cJSON.o
	$(CC) $(CPPFLAGS) -o $@ $^ -lm -lrobresource

mrj.o: mrj.c
	$(CC) $(CPPFLAGS) -o $@ -c $^ -lm -lrobresource

mcnv.o: mcnv.c
	$(CC) $(CPPFLAGS) -o $@ -c $^ -lm -lrobresource

mcresource.o: mcresource.c
	$(CC) $(CPPFLAGS) -o $@ -c $^ -lm -lrobresource

mcplc.o: mcplc.c
	$(CC) $(CPPFLAGS) -o $@ -c $^ -lm -lrobresource

mcvars.o: mcvars.c
	$(CC) $(CPPFLAGS) -o $@ -c $^ -lm -lrobresource

cJSON.o: cJSON.c
	$(CC) -fPIC -Wall -o $@ -c $^

clean:
	rm -f *.o libmrj.so
