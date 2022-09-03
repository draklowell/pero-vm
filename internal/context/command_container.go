package context

import "lab.draklowell.net/routine-runtime/internal/word"

func (ctx *Context) commandContainerNew() error {
	return ctx.machine.Stack.Push(word.NewContainer())
}

func (ctx *Context) commandContainerPut() error {
	keyIndex, err := ctx.readU2()
	if err != nil {
		return err
	}

	key, err := ctx.getConstantString(keyIndex)
	if err != nil {
		return err
	}

	value, err := ctx.machine.Stack.Pop()
	if err != nil {
		return err
	}

	container, err := ctx.popContainer()
	if err != nil {
		return err
	}

	return container.Set(key, value)
}

func (ctx *Context) commandContainerPutDynamic() error {
	value, err := ctx.machine.Stack.Pop()
	if err != nil {
		return err
	}

	key, err := ctx.popString()
	if err != nil {
		return err
	}

	container, err := ctx.popContainer()
	if err != nil {
		return err
	}

	container.Set(key, value)
	return nil
}

func (ctx *Context) commandContainerGet() error {
	keyIndex, err := ctx.readU2()
	if err != nil {
		return err
	}

	key, err := ctx.getConstantString(keyIndex)
	if err != nil {
		return err
	}

	container, err := ctx.popContainer()
	if err != nil {
		return err
	}

	value, err := container.Get(key)
	if err != nil {
		return err
	}

	ctx.machine.Stack.Push(value)
	return nil
}

func (ctx *Context) commandContainerGetDynamic() error {
	key, err := ctx.popString()
	if err != nil {
		return err
	}

	container, err := ctx.popContainer()
	if err != nil {
		return err
	}

	value, err := container.Get(key)
	if err != nil {
		return err
	}

	ctx.machine.Stack.Push(value)
	return nil
}

func (ctx *Context) commandContainerKeys() error {
	container, err := ctx.popContainer()
	if err != nil {
		return err
	}

	keys := container.GetKeys()
	result := word.NewArray(len(keys))
	for i, key := range keys {
		value, err := word.NewBytes([]byte(key))
		if err != nil {
			return err
		}
		result.Set(i, value)
	}

	ctx.machine.Stack.Push(result)
	return nil
}
