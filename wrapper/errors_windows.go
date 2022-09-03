package main

/*
#include "processthreadsapi.h"

static long getThreadId() {
	return GetCurrentThreadId();
}
*/
import "C"

func getThreadId() int64 {
	return int64(C.getThreadId())
}
