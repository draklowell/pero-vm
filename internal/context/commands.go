package context

import "lab.draklowell.net/pero-core/common/word"

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
	// value -> integer
	CommandToInteger = 43
	// value -> float
	CommandToFloat = 44
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
	// value1, value2 -> integer
	CommandIAdd = 54
	// value1, value2 -> integer
	CommandISub = 55
	// value1, value2 -> integer
	CommandIMul = 56
	// value1, value2 -> integer
	CommandIDiv = 57
	// value1, value2 -> integer
	CommandIRem = 58
	// value1, value2 -> integer
	CommandFAdd = 59
	// value1, value2 -> integer
	CommandFSub = 60
	// value1, value2 -> integer
	CommandFMul = 61
	// value1, value2 -> integer
	CommandFDiv = 62
	// value1, value2 -> integer
	CommandFRem = 63
	// value -> newValue
	CommandNeg = 64
	// value1, value2 -> boolean
	CommandAnd = 65
	// value1, value2 -> boolean
	CommandOr = 66
	// value -> boolean
	CommandNot = 67
	// value1, value2 -> integer
	CommandBAnd = 68
	// value1, value2 -> integer
	CommandBOr = 69
	// value1, value2 -> integer
	CommandBXor = 70
	// value -> integer
	CommandBNot = 71
	// value1, value2 -> integer
	CommandShiftLeft = 72
	// value1, value2 -> integer
	CommandShiftRight = 73
	// value1, value2 -> boolean
	CommandIGT = 74
	// value1, value2 -> boolean
	CommandIGE = 75
	// value1, value2 -> boolean
	CommandILT = 76
	// value1, value2 -> boolean
	CommandILE = 77
	// value1, value2 -> boolean
	CommandIEQ = 78
	// value1, value2 -> boolean
	CommandINE = 79
	// value1, value2 -> boolean
	CommandFGT = 80
	// value1, value2 -> boolean
	CommandFGE = 81
	// value1, value2 -> boolean
	CommandFLT = 82
	// value1, value2 -> boolean
	CommandFLE = 83
	// value1, value2 -> boolean
	CommandFEQ = 84
	// value1, value2 -> boolean
	CommandFNE = 85
	// element -> bool
	CommandCheckTypeNotNone = 86
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
	case CommandToInteger:
		return ctx.commandToInteger()
	case CommandToFloat:
		return ctx.commandToFloat()
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
	case CommandIAdd:
		return ctx.commandMathIAdd()
	case CommandISub:
		return ctx.commandMathISub()
	case CommandIMul:
		return ctx.commandMathIMul()
	case CommandIDiv:
		return ctx.commandMathIDiv()
	case CommandIRem:
		return ctx.commandMathIRem()
	case CommandFAdd:
		return ctx.commandMathFAdd()
	case CommandFSub:
		return ctx.commandMathFSub()
	case CommandFMul:
		return ctx.commandMathFMul()
	case CommandFDiv:
		return ctx.commandMathFDiv()
	case CommandFRem:
		return ctx.commandMathFRem()
	case CommandNeg:
		return ctx.commandMathNeg()
	case CommandAnd:
		return ctx.commandLogicalAnd()
	case CommandOr:
		return ctx.commandLogicalOr()
	case CommandNot:
		return ctx.commandLogicalNot()
	case CommandBAnd:
		return ctx.commandBitwiseAnd()
	case CommandBOr:
		return ctx.commandBitwiseOr()
	case CommandBXor:
		return ctx.commandBitwiseXor()
	case CommandBNot:
		return ctx.commandBitwiseNot()
	case CommandShiftLeft:
		return ctx.commandBitwiseShiftLeft()
	case CommandShiftRight:
		return ctx.commandBitwiseShiftRight()
	case CommandIGT:
		return ctx.commandCompareIGreater()
	case CommandIGE:
		return ctx.commandCompareIGreaterEquals()
	case CommandILT:
		return ctx.commandCompareILower()
	case CommandILE:
		return ctx.commandCompareILowerEquals()
	case CommandIEQ:
		return ctx.commandCompareIEquals()
	case CommandINE:
		return ctx.commandCompareINotEquals()
	case CommandFGT:
		return ctx.commandCompareFGreater()
	case CommandFGE:
		return ctx.commandCompareFGreaterEquals()
	case CommandFLT:
		return ctx.commandCompareFLower()
	case CommandFLE:
		return ctx.commandCompareFLowerEquals()
	case CommandFEQ:
		return ctx.commandCompareFEquals()
	case CommandFNE:
		return ctx.commandCompareFNotEquals()
	case CommandCheckTypeNotNone:
		return ctx.commandCheckNotTypeStatic(word.TypeNone)
	}

	return nil
}
