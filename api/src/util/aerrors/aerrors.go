package aerrors

import (
	"errors"
	"fmt"

	"golang.org/x/xerrors"
)

func create(msg string) *ErrApp {
	return &ErrApp{
		frame: Caller(2),
		msg:   msg,
	}
}

func New(msg string) *ErrApp {
	return create(msg)
}

func Errorf(format string, a ...interface{}) *ErrApp {
	return create(fmt.Sprintf(format, a...))
}

// Wrap: エラーをラップしてErrAppを生成する。msgを指定しない場合はcodeをmsgとして保存する。
func Wrap(err error, msg ...string) *ErrApp {
	if err == nil {
		return nil
	}

	var m string
	if len(msg) != 0 {
		m = msg[0]
	} else {
		if e := AsErrApp(err); e != nil {
			e.msg = string(e.Code())
		}
	}

	e := create(m)
	e.next = err

	return e
}

func Wrapf(err error, format string, a ...interface{}) *ErrApp {
	e := create(fmt.Sprintf(format, a...))
	e.next = err
	return e
}

func AsErrApp(err error) *ErrApp {
	if err == nil {
		return nil
	}

	var e *ErrApp
	if errors.As(err, &e) {
		return e
	}
	return nil
}

func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

// ErrApp: アプリケーション用のエラー。外部APIのエラーはすべてErrAppに変換する
type ErrApp struct {
	next    error
	frame   Frame
	msg     string
	code    Code
	infoMsg string
}

func (e *ErrApp) Error() string {
	if e.next == nil {
		return e.msg
	}
	if e.msg != "" {
		return e.msg + ": " + e.next.Error()
	}
	return e.next.Error()
}

// Messagef: ユーザー向けのメッセージを保存する
func (e *ErrApp) Message(infoMsg string) *ErrApp {
	e.infoMsg = infoMsg
	return e
}

// Messagef: ユーザー向けのメッセージを保存する
func (e *ErrApp) Messagef(format string, a ...interface{}) *ErrApp {
	e.infoMsg = fmt.Sprintf(format, a...)
	return e
}

func (e *ErrApp) SetCode(c Code) *ErrApp {
	e.code = c
	return e
}

func (e *ErrApp) Code() Code {
	var errApp *ErrApp = e
	for errApp.code == "" {
		if err := AsErrApp(errApp.next); err != nil {
			errApp = err
			continue
		}
		return "not_defined_code"
	}

	return errApp.code
}

func (e *ErrApp) InfoMsg() string {
	var next *ErrApp = e
	for next.infoMsg == "" {
		if err := AsErrApp(next.next); err != nil {
			next = err
		} else {
			return ""
		}
	}

	return next.infoMsg
}

func (e *ErrApp) Format(f fmt.State, c rune) {
	xerrors.FormatError(e, f, c)
}

func (e *ErrApp) FormatError(p xerrors.Printer) error {
	p.Print(e.Error())
	if p.Detail() {
		e.frame.Format(p)
	}
	return e.next
}

func (e *ErrApp) Unwrap() error {
	return e.next
}

type Stack struct {
	Func string
	File string
	Line int
	Msg  string
}

func (e *ErrApp) StackTrace() []*Stack {
	n := e
	stacktrace := []*Stack{}
	for n != nil {
		f, file, line := n.frame.location()
		stacktrace = append(stacktrace, &Stack{
			Func: f,
			File: file,
			Line: line,
			Msg:  n.Error(),
		})

		nn := n.next
		if next, ok := nn.(*ErrApp); ok {
			n = next
		} else {
			n = nil
		}
	}

	return stacktrace
}
