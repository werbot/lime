package errors

import "errors"

const (
	MsgNotFound      = "not found"
	MsgWrongPassword = "wrong password"

	MsgCustomerNotFound = "customer not found"
	MsgSettingNotFound  = "setting not found"
)

var (
	ErrNotFound      = errors.New(MsgNotFound)
	ErrWrongPassword = errors.New(MsgWrongPassword)

	ErrCustomerNotFound = errors.New(MsgCustomerNotFound)
	ErrSettingNotFound  = errors.New(MsgSettingNotFound)
)
