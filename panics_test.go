package panics

import (
	"errors"
	"fmt"
	"testing"
)

func TestNow(t *testing.T) {
	willPanic(t, "Panic now", func() {
		Now()
	})
	for _, test := range []struct {
		name       string
		msgAndArgs []interface{}
		output     string
	}{
		{
			"no message",
			nil,
			"Panic now",
		},
		{
			"no message",
			[]interface{}{"foo"},
			"foo",
		},
		{
			"no message",
			[]interface{}{"foo %s %d", "bar", 1},
			"foo bar 1",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			willPanic(t, test.output, func() {
				Now(test.msgAndArgs...)
			})
		})
	}
}

func TestIf(t *testing.T) {
	willPanic(t, "Panic because true", func() {
		If(true)
	})
	willNotPanic(t, func() {
		If(false)
	})
}

func TestIfNot(t *testing.T) {
	willPanic(t, "Panic because false", func() {
		IfNot(false)
	})
	willNotPanic(t, func() {
		IfNot(true)
	})
}

func TestIfError(t *testing.T) {
	willPanic(t, "Panic because error: foo", func() {
		IfError(errors.New("foo"))
	})
	willNotPanic(t, func() {
		IfError(nil)
	})
}

func TestIfNotError(t *testing.T) {
	willPanic(t, "Panic because no error", func() {
		IfNotError(nil)
	})
	willNotPanic(t, func() {
		IfNotError(errors.New("foo"))
	})
}

func TestIfNil(t *testing.T) {
	willPanic(t, "Panic because nil", func() {
		IfNil(nil)
	})
	willNotPanic(t, func() {
		IfNil(false)
	})
}

func TestIfNotNil(t *testing.T) {
	willPanic(t, "Panic because not nil: have 0 (int)", func() {
		IfNotNil(0)
	})
	willNotPanic(t, func() {
		IfNotNil(nil)
	})
}

func TestIfEqual(t *testing.T) {
	for _, test := range []struct {
		name       string
		val        interface{}
		expected   interface{}
		msgAndArgs []interface{}
		willPanic  bool
		output     string
	}{
		{
			"1=1",
			1, 1,
			nil,
			true, "Panic because equal: have 1 (int)",
		},
		{
			"1 != 2",
			1, 2,
			nil,
			false, "",
		},
		{
			"int(1) != int16(1)",
			1, int16(1),
			nil,
			false, "",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			if test.willPanic {
				willPanic(t, test.output, func() {
					IfEqual(test.val, test.expected, test.msgAndArgs...)
				})
			} else {
				willNotPanic(t, func() {
					IfEqual(test.val, test.expected, test.msgAndArgs...)
				})
			}
		})
	}
}

func TestIfNotEqual(t *testing.T) {
	for _, test := range []struct {
		name       string
		val        interface{}
		expected   interface{}
		msgAndArgs []interface{}
		willPanic  bool
		output     string
	}{
		{
			"1=1",
			1, 1,
			nil,
			false, "",
		},
		{
			"1 != 2",
			1, 2,
			nil,
			true, "Panic because not equal: have 2 (int), want 1 (int)",
		},
		{
			"int(1) != int16(1)",
			1, int16(1),
			nil,
			true, "Panic because not equal: have 1 (int16), want 1 (int)",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			if test.willPanic {
				willPanic(t, test.output, func() {
					IfNotEqual(test.val, test.expected, test.msgAndArgs...)
				})
			} else {
				willNotPanic(t, func() {
					IfNotEqual(test.val, test.expected, test.msgAndArgs...)
				})
			}
		})
	}
}

func TestFormatMsgAndArgs(t *testing.T) {
	for _, test := range []struct {
		name       string
		msgAndArgs []interface{}
		expected   string
	}{
		{
			"No message",
			nil,
			"",
		},
		{
			"One message",
			[]interface{}{"foo"},
			"foo",
		},
		{
			"Message format",
			[]interface{}{"foo %s", "bar"},
			"foo bar",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			actual := formatMsgAndArgs(test.msgAndArgs...)
			if actual != test.expected {
				t.Errorf(`Got "%s", want "%s"`, actual, test.expected)
			}
		})
	}
}

func panicResult(fn func()) (interface{}, bool) {
	didPanic := false
	var message interface{}
	func() {
		defer func() {
			if message = recover(); message != nil {
				didPanic = true
			}
		}()
		fn()
	}()
	return message, didPanic
}

func willNotPanic(t *testing.T, fn func()) {
	if message, didPanic := panicResult(fn); didPanic {
		t.Errorf("should not panic\n\tPanic value:\t%#v", message)
	}
}

func willPanic(t *testing.T, msg string, fn func()) {
	message, didPanic := panicResult(fn)
	if !didPanic {
		t.Errorf("should panic")
		return
	}

	got := fmt.Sprintf("%v", message)
	if got != msg {
		t.Errorf("panic should have returned \"%s\", got \"%s\"", msg, got)
	}
}
