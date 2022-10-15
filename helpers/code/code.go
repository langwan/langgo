package code

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"reflect"
)

func Call(ctx context.Context, service interface{}, methodName string, request string) (response interface{}, code int, err error) {
	tp := reflect.TypeOf(service)
	method, ok := tp.MethodByName(methodName)
	if !ok {
		return "", int(codes.NotFound), status.Errorf(codes.NotFound, "%s not find", methodName)
	}

	method.Type.NumIn()

	parameter := method.Type.In(2)
	req := reflect.New(parameter.Elem()).Interface()
	jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(request), req)

	in := make([]reflect.Value, 0)
	in = append(in, reflect.ValueOf(ctx))
	in = append(in, reflect.ValueOf(req))
	call := reflect.ValueOf(service).MethodByName(methodName).Call(in)
	if call[1].Interface() != nil {
		e := call[1].Interface().(error)
		st, _ := status.FromError(e)
		return "", int(st.Code()), e
	}

	return call[0].Interface(), 0, nil
}
