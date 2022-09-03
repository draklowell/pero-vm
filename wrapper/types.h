#ifndef __rrt_types
#define __rrt_types

#define rrtNullPointer -1
#define rrtGCFrequent 0
#define rrtGCRare 1

typedef unsigned char (*rrtBreaker) ();

typedef struct rrtDynamicRoutine { void* data; int length; } rrtDynamicRoutine;
typedef rrtDynamicRoutine* (*rrtDynamicLoader) (char* entry);

typedef int* (*rrtNativeRoutine) (int vmPtr, int* arguments, int argumentSize, int* retSize);

static unsigned char rrtBreakerBridge(rrtBreaker f) { return f(); }
static rrtDynamicRoutine* rrtDynamicLoaderBridge(rrtDynamicLoader f, char* entry) { return f(entry); }
static int* rrtNativeRoutineBridge(rrtNativeRoutine f, int vmPtr, int* arguments, int argumentSize, int* retSize) { return f(vmPtr, arguments, argumentSize, retSize); }

#endif
