package listener

import "syscall/js"

type Listener struct {
	element js.Value
}

func New(id string) *Listener {
	return &Listener{
		element: js.Global().Get("document").Call("getElementById", id),
	}
}

func (l *Listener) OnClick(cb func()) {
	l.element.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		cb()

		return nil
	}))
}
