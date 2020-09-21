package panics

import (
	"fmt"
	"reflect"
)

// Now will trigger a panic
func Now(msgAndArgs ...interface{}) {
	panicf("Panic now", msgAndArgs...)
}

// If will trigger a panic only if the condition is true
func If(condition bool, msgAndArgs ...interface{}) {
	if condition {
		panicf("Panic because true", msgAndArgs...)
	}
}

// IfNot will trigger a panic only if the condition is false
func IfNot(condition bool, msgAndArgs ...interface{}) {
	if !condition {
		panicf("Panic because false", msgAndArgs...)
	}
}

// IfError will trigger a panic only if the error is not nil
func IfError(err error, msgAndArgs ...interface{}) {
	if err != nil {
		panicf(fmt.Sprintf("Panic because error: %s", err), msgAndArgs...)
	}
}

// IfNotError will trigger a panic only if the error is nil
func IfNotError(err error, msgAndArgs ...interface{}) {
	if err == nil {
		panicf("Panic because no error", msgAndArgs...)
	}
}

// IfNil will trigger a panic only if the value is nil
func IfNil(v interface{}, msgAndArgs ...interface{}) {
	if v != nil || reflect.ValueOf(v).Kind() == reflect.Ptr && !reflect.ValueOf(v).IsNil() {
		return
	}
	panicf("Panic because nil", msgAndArgs...)
}

// IfNotNil will trigger a panic if the value is not nil
func IfNotNil(v interface{}, msgAndArgs ...interface{}) {
	if v == nil || reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil() {
		return
	}
	panicf(fmt.Sprintf("Panic because not nil: have %#v (%T)", v, v), msgAndArgs...)
}

// IfEqual will trigger a panic only if the expected and actual are equal
func IfEqual(expected, actual interface{}, msgAndArgs ...interface{}) {
	if expected != actual {
		return
	}
	panicf(fmt.Sprintf("Panic because equal: have %#v (%T)", actual, actual), msgAndArgs...)
}

// IfNotEqual will trigger a panic only if the expected and actual are not equal
func IfNotEqual(expected, actual interface{}, msgAndArgs ...interface{}) {
	if expected == actual || reflect.DeepEqual(expected, actual) {
		return
	}
	panicf(fmt.Sprintf("Panic because not equal: have %#v (%T), want %#v (%T)", actual, actual, expected, expected), msgAndArgs...)
}

func panicf(defaultMessage string, msgAndArgs ...interface{}) {
	if len(msgAndArgs) > 0 {
		panic(formatMsgAndArgs(msgAndArgs...))
	}
	panic(defaultMessage)
}

func formatMsgAndArgs(msgAndArgs ...interface{}) string {
	if len(msgAndArgs) == 0 {
		return ""
	}
	format := fmt.Sprintf("%v", msgAndArgs[0])
	if len(msgAndArgs) == 1 {
		return format
	}
	return fmt.Sprintf(format, msgAndArgs[1:]...)
}
