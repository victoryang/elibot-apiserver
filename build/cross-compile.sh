#!/bin/sh

# template of cross-compile
# make HAS_PKG_CONFIG=false CC=your-gcc CXX=your-g++ LD=your-gcc LDXX=your-g++ AR=your-ar \
# CPPFLAGS=-I/path/to/packages/include LDFLAGS=-L/path/to/packages/lib static

# e.g. make HAS_PKG_CONFIG=false CC=arm-xilinx-linux-gnueabi-gcc CXX=arm-xilinx-linux-gnueabi-g++ ...