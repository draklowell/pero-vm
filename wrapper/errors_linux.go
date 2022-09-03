package main

/*
#include "pthread.h"

static long getThreadId() {
	return pthread_self();
}
*/
import "C"

func getThreadId() int64 {
	return int64(C.getThreadId())
}
