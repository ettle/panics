
# Go Panics

[![Go Report Card](https://goreportcard.com/badge/github.com/ettle/panics)](https://goreportcard.com/report/github.com/ettle/panics)
[![GoDoc](https://godoc.org/github.com/ettle/panics?status.svg)](https://pkg.go.dev/github.com/ettle/panics)

`import "github.com/ettle/panics"`

## <a name="pkg-overview">Overview</a>
Package panics implements some helper methods around panicing

While you should handle and return errors instead of throwing a panic
(<a href="https://blog.golang.org/error-handling-and-go">https://blog.golang.org/error-handling-and-go</a>), there are times where
having a runtime assertion is useful.




## <a name="pkg-index">Index</a>
* [func If(condition bool, msgAndArgs ...interface{})](#If)
* [func IfEqual(expected, actual interface{}, msgAndArgs ...interface{})](#IfEqual)
* [func IfError(err error, msgAndArgs ...interface{})](#IfError)
* [func IfNil(v interface{}, msgAndArgs ...interface{})](#IfNil)
* [func IfNot(condition bool, msgAndArgs ...interface{})](#IfNot)
* [func IfNotEqual(expected, actual interface{}, msgAndArgs ...interface{})](#IfNotEqual)
* [func IfNotError(err error, msgAndArgs ...interface{})](#IfNotError)
* [func IfNotNil(v interface{}, msgAndArgs ...interface{})](#IfNotNil)
* [func Now(msgAndArgs ...interface{})](#Now)





## <a name="If">func</a> [If](./panics.go#L14)
``` go
func If(condition bool, msgAndArgs ...interface{})
```
If will trigger a panic only if the condition is true



## <a name="IfEqual">func</a> [IfEqual](./panics.go#L58)
``` go
func IfEqual(expected, actual interface{}, msgAndArgs ...interface{})
```
IfEqual will trigger a panic only if the expected and actual are equal



## <a name="IfError">func</a> [IfError](./panics.go#L28)
``` go
func IfError(err error, msgAndArgs ...interface{})
```
IfError will trigger a panic only if the error is not nil



## <a name="IfNil">func</a> [IfNil](./panics.go#L42)
``` go
func IfNil(v interface{}, msgAndArgs ...interface{})
```
IfNil will trigger a panic only if the value is nil



## <a name="IfNot">func</a> [IfNot](./panics.go#L21)
``` go
func IfNot(condition bool, msgAndArgs ...interface{})
```
IfNot will trigger a panic only if the condition is false



## <a name="IfNotEqual">func</a> [IfNotEqual](./panics.go#L66)
``` go
func IfNotEqual(expected, actual interface{}, msgAndArgs ...interface{})
```
IfNotEqual will trigger a panic only if the expected and actual are not equal



## <a name="IfNotError">func</a> [IfNotError](./panics.go#L35)
``` go
func IfNotError(err error, msgAndArgs ...interface{})
```
IfNotError will trigger a panic only if the error is nil



## <a name="IfNotNil">func</a> [IfNotNil](./panics.go#L50)
``` go
func IfNotNil(v interface{}, msgAndArgs ...interface{})
```
IfNotNil will trigger a panic if the value is not nil



## <a name="Now">func</a> [Now](./panics.go#L9)
``` go
func Now(msgAndArgs ...interface{})
```
Now will trigger a panic








