# Build

## Install cross compile tools on Ubuntu
	sudo apt-get install -y libc6-armel-cross libc6-dev-armel-cross
	sudo apt-get install -y binutils-arm-linux-gnueabi binutils-arm-linux-gnueabihf
	sudo apt-get install -y libncurses5-dev

## For different arch
### SYSV interpreted by /lib/ld-linux-armhf.so.x
	sudo apt-get install -y gcc-arm-linux-gnueabihf
	sudo apt-get install -y g++-arm-linux-gnueabihf

### Or GNU/LINUX interpreted by /lib/ld-linux.so.x
	sudo apt-get install -y gcc-arm-linux-gnueabi
	sudo apt-get install -y g++-arm-linux-gnueabi

## Install build dependencies
	cp lib/libsqlitedb /usr/lib/gcc-cross/arm-linux-gnueabihf/5/
	cp lib/libz.so.1.2.8 /usr/arm-linux-gnueabihf/lib/ 
	ln -s /usr/arm-linux-gnueabihf/lib/libz.so.1.2.8 /usr/arm-linux-gnueabihf/lib/libz.so.1
	cp lib/libshare.a /usr/arm-linux-gnueabihf/lib/
	mkdir /root/mcserver
	tar -xvf include/include.tar.gz -C /root/mcserver/
	cp include/config.h /root/mcserver

## Run build
	CC=arm-linux-gnueabihf-gcc GOOS=linux GOARCH=arm GOARM=6 CGO_ENABLED=1 go build -v -o /root/elibot-server