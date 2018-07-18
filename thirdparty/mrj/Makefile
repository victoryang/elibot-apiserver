CC = arm-linux-gnueabihf-gcc
AR = arm-linux-gnueabihf-ar
SHARED = -shared
RCS = rcs

CPPFLAGS += -I/root/mcserver/include/ -fPIC -Wall

.PHONY: all

all: libmrj.so

libmrj.so: mrj.o items.o cJSON.o
	$(CC) $(SHARED) -o $@ $^

mrj.o: mrj.c
	$(CC) $(CPPFLAGS) -o $@ -c $^ -lm -lshare

items.o: items.c
	$(CC) $(CPPFLAGS) -o $@ -c $^ -lm -lshare

cJSON.o: cJSON.c
	$(CC) -fPIC -Wall -o $@ -c $^

clean:
	rm -f *.o libmrj.so