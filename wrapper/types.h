#ifndef __pero_types

#define peroVersion "22s0906"

#define peroNullPointer -1
#define peroGCFrequent 0
#define peroGCRare 1
#define peroTypeNone 0
#define peroTypeInteger 1
#define peroTypeFloat 2
#define peroTypeArray 3
#define peroTypeContainer 4
#define peroTypeBytes 5
#define peroTypeBoolean 6

typedef unsigned char (*peroBreaker) ();

typedef struct peroDynamicRoutine { void* data; int length; } peroDynamicRoutine;
typedef peroDynamicRoutine* (*peroDynamicLoader) (char* entry);

typedef int* (*peroNativeRoutine) (int vmPtr, int* arguments, int argumentSize, int* retSize);

static unsigned char peroBreakerBridge(peroBreaker f) { return f(); }
static peroDynamicRoutine* peroDynamicLoaderBridge(peroDynamicLoader f, char* entry) { return f(entry); }
static int* peroNativeRoutineBridge(peroNativeRoutine f, int vmPtr, int* arguments, int argumentSize, int* retSize) { return f(vmPtr, arguments, argumentSize, retSize); }

#define __pero_types
#endif
