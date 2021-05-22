package aerrors

import (
	"errors"
	"fmt"

	"golang.org/x/xerrors"
)

func create(msg string) *ErrApp {
	return &ErrApp{
		frame: xerrors.Caller(1),
		// code:  code,
		msg: msg,
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
	if len(m) != 0 {
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
	frame   xerrors.Frame
	msg     string
	code    Code
	infoMsg string
}

func (e *ErrApp) Error() string {
	return e.msg
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
