package context

import "lab.draklowell.net/routine-runtime/word"

const (
	CommandPushNone           = 0
	CommandPushm1             = 1
	CommandPush0              = 2
	CommandPush1              = 3
	CommandPush2              = 4
	CommandPush3              = 5
	CommandPushm1f            = 6
	CommandPush0f             = 7
	CommandPush1f             = 8
	CommandPush2f             = 9
	CommandPush3f             = 10
	CommandPushConstant       = 11
	CommandPushVariable       = 12
	CommandPushVariable0      = 13
	CommandPushVariable1      = 14
	CommandPushVariable2      = 15
	CommandPushVariable3      = 16
	CommandPop                = 17
	CommandPopVariable        = 18
	CommandPopVariable0       = 19
	CommandPopVariable1       = 20
	CommandPopVariable2       = 21
	CommandPopVariable3       = 22
	CommandReturn             = 23
	CommandReturn0            = 24
	CommandReturn1            = 25
	CommandReturn2            = 26
	CommandInvokeDynamic      = 27
	CommandInvoke             = 28
	CommandArrayNew           = 29
	CommandArrayLength        = 30
	CommandArrayGet           = 31
	CommandArrayGet0          = 32
	CommandArrayGetm1         = 33
	CommandArraySet           = 34
	CommandArraySet0          = 35
	CommandArraySetm1         = 36
	CommandCheckTypeNone      = 37
	CommandCheckTypeInteger   = 38
	CommandCheckTypeFloat     = 39
	CommandCheckTypeBytes     = 40
	CommandCheckTypeArray     = 41
	CommandCheckTypeContainer = 42
	CommandFloatToInteger     = 43
	CommandIntegerToFloat     = 44
	CommandGoto               = 45
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
	case CommandFloatToInteger:
		return ctx.commandFloatToInteger()
	case CommandIntegerToFloat:
		return ctx.commandIntegerToFloat()
	case CommandGoto:
		return ctx.commandGoto()
	}

	return nil
}
