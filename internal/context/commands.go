package context

import "lab.draklowell.net/routine-runtime/common/word"

const (
	// -> None
	CommandPushNone = 0
	// -> -1
	CommandPushm1 = 1
	// -> 0
	CommandPush0 = 2
	// -> 1
	CommandPush1 = 3
	// -> 2
	CommandPush2 = 4
	// -> 3
	CommandPush3 = 5
	// -> -1.0
	CommandPushm1f = 6
	// -> 0.0
	CommandPush0f = 7
	// -> 1.0
	CommandPush1f = 8
	// -> 2.0
	CommandPush2f = 9
	// -> 3.0
	CommandPush3f = 10
	// -> constant
	CommandPushConstant = 11
	// -> variable
	CommandPushVariable = 12
	// -> variable0
	CommandPushVariable0 = 13
	// -> variable1
	CommandPushVariable1 = 14
	// -> variable2
	CommandPushVariable2 = 15
	// -> variable3
	CommandPushVariable3 = 16
	// element ->
	CommandPop = 17
	// variable ->
	CommandPopVariable = 18
	// variable0 ->
	CommandPopVariable0 = 19
	// variable1 ->
	CommandPopVariable1 = 20
	// variable2 ->
	CommandPopVariable2 = 21
	// variable3 ->
	CommandPopVariable3 = 22
	// ret0, ret1, ...retn ->
	CommandReturn = 23
	// ->
	CommandReturn0 = 24
	// ret0 ->
	CommandReturn1 = 25
	// ret0, ret1 ->
	CommandReturn2 = 26
	// entry, arg0, arg1, ...argn -> ret0, ret1, ...retn
	CommandInvokeDynamic = 27
	// arg0, arg1, ...argn -> ret0, ret1, ...retn
	CommandInvoke = 28
	// -> array
	CommandArrayNew = 29
	// array -> length
	CommandArrayLength = 30
	// array, index -> element
	CommandArrayGet = 31
	// array -> elementFirst
	CommandArrayGet0 = 32
	// array -> elementLast
	CommandArrayGetm1 = 33
	// array, index, element ->
	CommandArraySet = 34
	// array, elementFirst ->
	CommandArraySet0 = 35
	// array, elementLast ->
	CommandArraySetm1 = 36
	// element -> bool
	CommandCheckTypeNone = 37
	// element -> bool
	CommandCheckTypeInteger = 38
	// element -> bool
	CommandCheckTypeFloat = 39
	// element -> bool
	CommandCheckTypeBytes = 40
	// element -> bool
	CommandCheckTypeArray = 41
	// element -> bool
	CommandCheckTypeContainer = 42
	// float -> integer
	CommandFloatToInteger = 43
	// integer -> float
	CommandIntegerToFloat = 44
	// ->
	CommandGoto = 45
	// bool ->
	CommandGotoIf = 46
	// -> container
	CommandContainerNew = 47
	// container, key, element ->
	CommandContainerPutDynamic = 48
	// container, element ->
	CommandContainerPut = 49
	// container, key -> element
	CommandContainerGetDynamic = 50
	// container -> element
	CommandContainerGet = 51
	// container -> array
	CommandContainerKeys = 52
	// element -> bool
	CommandCheckTypeBoolean = 53
)

func (ctx *Context) executeCommand(command uint8) error {
	switch command {
	case CommandPushNone:
		return ctx.commandPushNone()
	case CommandPushm1:
		return ctx.commandPushIntegerStatic(-1)
	case CommandPush0:
		return ctx.commandPushIntegerStatic(0)
	case CommandPush1:
		return ctx.commandPushIntegerStatic(1)
	case CommandPush2:
		return ctx.commandPushIntegerStatic(2)
	case CommandPush3:
		return ctx.commandPushIntegerStatic(3)
	case CommandPushm1f:
		return ctx.commandPushFloatStatic(-1)
	case CommandPush0f:
		return ctx.commandPushFloatStatic(0)
	case CommandPush1f:
		return ctx.commandPushFloatStatic(1)
	case CommandPush2f:
		return ctx.commandPushFloatStatic(2)
	case CommandPush3f:
		return ctx.commandPushFloatStatic(3)
	case CommandPushConstant:
		return ctx.commandPushConstant()
	case CommandInvokeDynamic:
		return ctx.commandInvokeDynamic()
	case CommandInvoke:
		return ctx.commandInvoke()
	case CommandPushVariable:
		return ctx.commandPushVariable()
	case CommandPushVariable0:
		return ctx.commandPushVariableStatic(0)
	case CommandPushVariable1:
		return ctx.commandPushVariableStatic(1)
	case CommandPushVariable2:
		return ctx.commandPushVariableStatic(2)
	case CommandPushVariable3:
		return ctx.commandPushVariableStatic(3)
	case CommandPop:
		return ctx.commandPop()
	case CommandPopVariable:
		return ctx.commandPopVariable()
	case CommandPopVariable0:
		return ctx.commandPopVariableStatic(0)
	case CommandPopVariable1:
		return ctx.commandPopVariableStatic(1)
	case CommandPopVariable2:
		return ctx.commandPopVariableStatic(2)
	case CommandPopVariable3:
		return ctx.commandPopVariableStatic(3)
	case CommandReturn:
		return ctx.commandReturn()
	case CommandReturn0:
		return ctx.commandReturnStatic(0)
	case CommandReturn1:
		return ctx.commandReturnStatic(1)
	case CommandReturn2:
		return ctx.commandReturnStatic(2)
	case CommandArrayNew:
		return ctx.commandArrayNew()
	case CommandArrayLength:
		return ctx.commandArrayLength()
	case CommandArrayGet:
		return ctx.commandArrayGet()
	case CommandArrayGet0:
		return ctx.commandArrayGetStatic(0)
	case CommandArrayGetm1:
		return ctx.commandArrayGetStatic(-1)
	case CommandArraySet:
		return ctx.commandArraySet()
	case CommandArraySet0:
		return ctx.commandArraySetStatic(0)
	case CommandArraySetm1:
		return ctx.commandArraySetStatic(-1)
	case CommandCheckTypeNone:
		return ctx.commandCheckTypeStatic(word.TypeNone)
	case CommandCheckTypeInteger:
		return ctx.commandCheckTypeStatic(word.TypeInteger)
	case CommandCheckTypeFloat:
		return ctx.commandCheckTypeStatic(word.TypeFloat)
	case CommandCheckTypeBytes:
		return ctx.commandCheckTypeStatic(word.TypeBytes)
	case CommandCheckTypeArray:
		return ctx.commandCheckTypeStatic(word.TypeArray)
	case CommandCheckTypeContainer:
		return ctx.commandCheckTypeStatic(word.TypeContainer)
	case CommandCheckTypeBoolean:
		return ctx.commandCheckTypeStatic(word.TypeBoolean)
	case CommandFloatToInteger:
		return ctx.commandFloatToInteger()
	case CommandIntegerToFloat:
		return ctx.commandIntegerToFloat()
	case CommandGoto:
		return ctx.commandGoto()
	case CommandGotoIf:
		return ctx.commandGotoIf()
	case CommandContainerNew:
		return ctx.commandContainerNew()
	case CommandContainerPutDynamic:
		return ctx.commandContainerPutDynamic()
	case CommandContainerPut:
		return ctx.commandContainerPut()
	case CommandContainerGetDynamic:
		return ctx.commandContainerGetDynamic()
	case CommandContainerGet:
		return ctx.commandContainerGet()
	case CommandContainerKeys:
		return ctx.commandContainerKeys()
	}

	return nil
}
