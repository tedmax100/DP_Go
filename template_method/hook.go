package main

type Hook interface {
	Execute(func(...interface{}))
}

type Hooked struct {
	Hook
}

func (t *Hooked) Execute(hookFuncList ...func(args ...interface{})) {
	for _, hookFunc := range hookFuncList {
		hookFunc()
	}
	// fmt.Println("Hooked.Executed")
}
