package errh

import (
	er "errors"
	"github.com/pkg/errors"
)

// ErrHandler 错误处理器
type ErrHandler struct {
	err                error
	defaultErrWrappers []func(err error) error
	matched            bool
}

func New(defaultErrWrappers ...func(err error) error) ErrHandler {
	return ErrHandler{
		defaultErrWrappers: defaultErrWrappers,
	}
}

func (e *ErrHandler) If(isTure bool) *ErrHandler {
	e.matched = isTure
	return e
}

// ErrIs 判断错误类型
func (e *ErrHandler) ErrIs(err error) *ErrHandler {
	return e.If(er.Is(e.err, err))
}

// ReplaceErr 替换错误
func (e *ErrHandler) ReplaceErr(err error) *ErrHandler {
	if e.matched {
		e.err = err
	}
	return e
}

func (e *ErrHandler) HasErr() bool {
	return e.err != nil
}

// TryToSetErr 如果结构体中err不为空 则设置err
func (e *ErrHandler) TryToSetErr(err error) {
	if e.HasErr() {
		return
	}
	e.err = err
}

// Err 返回错误
func (e *ErrHandler) Err(errWrappers ...func(err error) error) error {
	if e.err == nil {
		return nil
	}
	err := errors.New(e.err.Error())
	if len(errWrappers) == 0 {
		for _, errWrapper := range e.defaultErrWrappers {
			err = errWrapper(err)
		}
	} else {
		for _, errWrapper := range errWrappers {
			err = errWrapper(err)
		}
	}
	return err
}
