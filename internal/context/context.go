package context

import (
	"encoding/binary"

	"lab.draklowell.net/routine-runtime/internal"
	"lab.draklowell.net/routine-runtime/internal/utils"
	"lab.draklowell.net/routine-runtime/word"
)

type Constant interface{}

type Context struct {
	stack *utils.Stack[word.Word]

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
		stack:     utils.NewStack[word.Word](machine.StackSize),
		bytecode:  bytecode,
		constants: constants,
		variables: make([]word.Word, 256),
		lineMap:   lineMap,
		entry:     entry,
		order:     order,
	}
}

func (ctx *Context) GetLine() int {
	byteIndex := ctx.offset
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
		return ErrorContextFinished
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
	for !ctx.IsFinished() && !breaker() {
		err := ctx.ExecuteCommand()
		if err != nil {
			return &ErrorTraceBack{
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
		return nil, ErrorContextBroken
	} else {
		return nil, ErrorContextNotFinished
	}
}

func (ctx *Context) SetVariable(index uint8, value word.Word) {
	ctx.variables[index] = value
}

func (ctx *Context) GetVariable(index uint8) word.Word {
	return ctx.variables[index]
}

func (ctx *Context) IsFinished() bool {
	return ctx.finished
}
