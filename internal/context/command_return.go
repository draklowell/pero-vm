package context

import "lab.draklowell.net/routine-runtime/common/word"

func (ctx *Context) commandReturnStatic(retCount int) error {
	result := make([]word.Word, retCount)
	for i := 0; i < retCount; i++ {
		word, err := ctx.machine.Stack.Pop()
		if err != nil {
			return err
		}
		result[i] = word
		i++
	}

	ctx.ret = result
	ctx.finished = true
	return nil
}

func (ctx *Context) commandReturn() error {
	retCount, err := ctx.readU1()
	if err != nil {
		return err
	}
	return ctx.commandReturnStatic(int(retCount))
}
