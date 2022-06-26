#!/bin/bash

compile:
	cd src;                      \
	go build main.go;            \
	cd ..;                       \
	mv src/main mandelbrot;      \
	chmod u+x mandelbrot;        \

clean:
	rm mandelbrot
