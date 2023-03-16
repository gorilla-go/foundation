package try_catch

import (
	"errors"
	"reflect"
)

type CatchAble struct {
	tryAction    func()
	catchArr     []*func(error)
	errorNameMap map[*func(error)]error
	finalFunc    func()
}

func Try(f func()) *CatchAble {
	return &CatchAble{
		tryAction:    f,
		catchArr:     []*func(error){},
		errorNameMap: make(map[*func(error)]error),
		finalFunc:    nil,
	}
}

func (a *CatchAble) Catch(err error, handler func(error)) *CatchAble {
	a.catchArr = append(a.catchArr, &handler)
	a.errorNameMap[&handler] = err
	return a
}

func (a *CatchAble) Finally(finalFunc func()) *CatchAble {
	a.finalFunc = finalFunc
	return a
}

func (a *CatchAble) Exec() {
	if a.finalFunc != nil {
		defer a.finalFunc()
	}
	defer func() {
		err := recover()
		if err != nil {
			errType := reflect.TypeOf(err)
			var baseHandler func(error) = nil

			for _, handler := range a.catchArr {
				if _, ok := a.errorNameMap[handler]; !ok {
					continue
				}
				targetErrType := reflect.TypeOf(a.errorNameMap[handler])
				if targetErrType.String() == "*exception.exception" {
					baseHandler = *handler
					continue
				}
				if errType == targetErrType {
					(*handler)(err.(error))
					return
				}
			}
			if baseHandler != nil {
				if errType.String() == "string" {
					baseHandler(errors.New(err.(string)))
					return
				}
				baseHandler(err.(error))
				return
			}
			panic("uncaught errors.")
		}
	}()
	a.tryAction()
}

func ThrowNew(e error) {
	panic(e)
}
