package helperGrpc

import (
	"context"
	"encoding/json"
	"errors"
	jsoniter "github.com/json-iterator/go"
	"github.com/langwan/langgo/helpers/grpc/pb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"reflect"
	"testing"
)

type TestService struct {
}

func (t TestService) Hello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Msg: request.GetMsg()}, nil
}

func (t TestService) Hello2(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	i := 0
	for i < 100000 {
		i++
	}
	return &pb.HelloResponse{Msg: request.GetMsg()}, nil
}

func Test_jsoniter(t *testing.T) {
	ts := TestService{}
	req := pb.HelloRequest{Msg: "hello"}
	marshal, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(req)
	if err != nil {
		return
	}
	response, code, err := Call(ts, "Hello", string(marshal), nil)
	assert.NoError(t, err)
	assert.Equal(t, code, 0)
	t.Log(response)
}

func Benchmark_jsoniter(b *testing.B) {
	s := TestService{}
	req := pb.HelloRequest{Msg: "hello"}
	marshal, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(req)
	if err != nil {
		return
	}
	reqString := string(marshal)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Call(s, "Hello", reqString, nil)
	}
}

func Benchmark_json(b *testing.B) {
	s := TestService{}
	req := pb.HelloRequest{Msg: "hello"}
	marshal, err := json.Marshal(req)
	if err != nil {
		return
	}
	reqString := string(marshal)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Call2(s, "Hello", reqString, nil)
	}
}

func Benchmark_call(b *testing.B) {
	s := TestService{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Hello(context.Background(), &pb.HelloRequest{Msg: "hello"})

	}
}

func Benchmark_sleep10_jsoniter(b *testing.B) {
	s := TestService{}
	req := pb.HelloRequest{Msg: "hello"}
	marshal, err := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(req)
	if err != nil {
		return
	}
	reqString := string(marshal)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Call(s, "Hello2", reqString, nil)
	}
}

func Benchmark_sleep10_call(b *testing.B) {
	s := TestService{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Hello2(context.Background(), &pb.HelloRequest{Msg: "hello"})
	}
}

func Call2(service interface{}, methodName string, request string, header interface{}) (response string, code int, err error) {
	tp := reflect.TypeOf(service)
	method, ok := tp.MethodByName(methodName)
	if !ok {
		return "", int(codes.NotFound), status.Errorf(codes.NotFound, "%s not find", methodName)
	}

	method.Type.NumIn()

	parameter := method.Type.In(2)
	req := reflect.New(parameter.Elem()).Interface()
	json.Unmarshal([]byte(request), req)

	in := make([]reflect.Value, 0)
	ctx := context.Background()
	in = append(in, reflect.ValueOf(ctx))
	in = append(in, reflect.ValueOf(req))
	call := reflect.ValueOf(service).MethodByName(methodName).Call(in)
	if call[1].Interface() != nil {
		e := call[1].Interface().(error)
		st, _ := status.FromError(e)
		return "", int(st.Code()), e
	}
	marshal, err := json.Marshal(call[0].Interface())
	if err != nil {
		return "", int(codes.Aborted), errors.New("json marshal error")
	}
	return string(marshal), 0, nil
}
