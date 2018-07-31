#!/bin/sh

# install gcc-arm tool chain
sudo apt-get update >/dev/null
sudo apt-get install -y cd
sudo apt-get install -y libc6-armel-cross libc6-dev-armel-cross >/dev/null
sudo apt-get install -y binutils-arm-linux-gnueabi binutils-arm-linux-gnueabihf >/dev/null
sudo apt-get install -y libncurses5-dev >/dev/null
sudo apt-get install -y gcc-arm-linux-gnueabihf >/dev/null
sudo apt-get install -y g++-arm-linux-gnueabihf >/dev/null