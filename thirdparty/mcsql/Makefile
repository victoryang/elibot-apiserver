CC = arm-linux-gnueabihf-gcc
AR = arm-linux-gnueabihf-ar
SHARED = -shared
RCS = rcs

CPPFLAGS += -I/root/mcserver/include/ -fPIC -Wall

.PHONY: all

all: libmcsql.so

libmcsql.so: mcsql.o mcquery.o mcmanager.o mcsqlmapper.o
	$(CC) $(SHARED) -o $@ $^

test: test_main.c mcsql.o mcquery.o mcmanager.o
	$(CC) $(CPPFLAGS) -o $@ $^ -lsqlitedb

mcsql.o: mcsql.c
	$(CC) $(CPPFLAGS) -o $@ -c $^ -lsqlitedb

mcquery.o: mcquery.c
	$(CC) $(CPPFLAGS) -o $@ -c $^ -lsqlitedb

mcmanager.o: mcmanager.c
	$(CC) $(CPPFLAGS) -o $@ -c $^ -lsqlitedb

mcsqlmapper.o: mcsqlmapper.c
	$(CC) $(CPPFLAGS) -o $@ -c $^ -lsqlitedb

clean:
	rm -f *.o libmcsql.so