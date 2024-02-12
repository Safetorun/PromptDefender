package tracer

import "log"

type GenericFuncType func(args ...interface{}) (interface{}, error)

type Tracer interface {
	TraceDecorator(fn GenericFuncType, functionName string) GenericFuncType
}

type EmptyTracer struct {
}

func (t *EmptyTracer) TraceDecorator(fn GenericFuncType, functionName string) GenericFuncType {
	return fn
}

// Wrapper function that conforms to GenericFuncType
func TracerGenericsWrapper[T any, E any](functionToCall func(T) (E, error)) func(...interface{}) (interface{}, error) {

	return func(args ...interface{}) (interface{}, error) {
		log.Default().Printf("Tracing function call, args: %+v\n", args[0])
		arg1, _ := args[0].(T)
		return functionToCall(arg1)
	}
}

func TracerGenericsWrapper2[T any, T1 any, E any](functionToCall func(T, T1) (E, error)) func(...interface{}) (interface{}, error) {
	return func(args ...interface{}) (interface{}, error) {
		log.Default().Printf("Tracing function call, args: %+v and %+v\n", args[0], args[1])
		arg1, _ := args[0].(T)
		arg2, _ := args[1].(T1)

		return functionToCall(arg1, arg2)
	}
}
