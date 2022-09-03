package context

import (
	"encoding/binary"

	"lab.draklowell.net/routine-runtime/internal"
	"lab.draklowell.net/routine-runtime/internal/word"
)

type Constant interface{}

type Context struct {
	bytecode  []byte
	offset    int
	variables []word.Word

	finished bool
	broken   bool
	ret      []word.Word

	machine   *internal.Machine
	entry     string
	constants []Constant
	lineMap   map[int]int

	order binary.ByteOrder
}

func NewContext(machine *internal.Machine, order binary.ByteOrder, bytecode []byte, constants []Constant, lineMap map[int]int, entry string) *Context {
	return &Context{
		machine:   machine,
		bytecode:  bytecode,
		constants: constants,
		variables: make([]word.Word, 256),
		lineMap:   lineMap,
		entry:     entry,
		order:     order,
	}
}

func (ctx *Context) GetLine() int {
	byteIndex := ctx.offset - 1
	for byteIndex >= 0 {
		if ctx.lineMap[byteIndex] != 0 {
			return ctx.lineMap[byteIndex]
		}
		byteIndex--
	}
	return 1
}

func (ctx *Context) GetCaller() string {
	return ctx.entry
}

func (ctx *Context) ExecuteCommand() error {
	if ctx.finished {
		return ErrContextFinished
	}

	command, finished, err := ctx.readCommand()
	if err != nil {
		return err
	}
	if finished {
		ctx.finished = true
		return nil
	}

	return ctx.executeCommand(command)
}

func (ctx *Context) Execute(breaker internal.BreakCallback) error {
	err := ctx.machine.Stack.PushLock()
	if err != nil {
		return err
	}
	defer ctx.machine.Stack.PopLock()

	for !ctx.IsFinished() && !breaker() {
		err := ctx.ExecuteCommand()
		if err != nil {
			return &ErrTraceBack{
				Base:   err,
				Caller: ctx.entry,
				Line:   ctx.GetLine(),
			}
		}
	}
	if !ctx.IsFinished() {
		ctx.broken = true
	}
	return nil
}

func (ctx *Context) GetReturn() ([]word.Word, error) {
	if ctx.finished {
		return ctx.ret, nil
	} else if ctx.broken {
		return nil, ErrContextBroken
	} else {
		return nil, ErrContextNotFinished
	}
}

func (ctx *Context) SetVariable(index uint8, value word.Word) {
	ctx.variables[index] = value
}

func (ctx *Context) GetVariable(index uint8) word.Word {
	value := ctx.variables[index]
	if value == nil {
		return word.None
	}

	return value
}

func (ctx *Context) IsFinished() bool {
	return ctx.finished
}
