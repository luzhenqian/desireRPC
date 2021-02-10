package server

import (
	"desire/protocol"
	"errors"
	"fmt"
	"log"
	"net"
	"reflect"
	"runtime"
	"strings"
)

// Server rpc server 接口
type Server interface {
	Register(fn interface{}) error
	call(name string, params ...interface{}) (interface{}, error)
	Run()
}

// Desire rpc server 实现
type Desire struct {
	fns map[string]interface{}
}

// Address rpc 服务端地址
type Address string

// New 创建 rpc 服务实例
func New() *Desire {
	return &Desire{
		fns: make(map[string]interface{}, 1),
	}
}

// Register 注册函数
func (s *Desire) Register(fn interface{}) error {
	if !checkIsFn(fn) { // 检查是否为函数
		return errors.New("register failed, params is not a func")
	}
	name, err := getFnName(fn) // 获取函数名 getFnName
	if err != nil {
		return err
	}
	return addFn(s, name, fn) // 向 fns 添加 fn
}

// Run 启动 rpc 服务
func (s *Desire) Run(addr Address) {
	listener, err := net.Listen("tcp", string(addr))
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept() // 接收 tcp 请求
		if err != nil {
			log.Fatal(err)
		}
		// process(conn) // 处理请求
		dp := protocol.Desire{}
		var i interface{}
		dp.Read(conn, i)
		// TODO
		// conn.Close()
	}
}

func checkIsFn(fn interface{}) bool {
	return reflect.ValueOf(fn).Kind() == reflect.Func
}

func getFnName(fn interface{}) (string, error) {
	fullName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	tmp := strings.Split(fullName, ".")
	if len(tmp) < 2 {
		return "", errors.New("get func name failed")
	}
	name := tmp[len(tmp)-1]
	return name, nil
}

func addFn(s *Desire, name string, fn interface{}) error {
	if _, ok := s.fns[name]; ok {
		return errors.New("func " + name + " is registered")
	}
	s.fns[name] = fn
	return nil
}

// call 调用
func (s *Desire) call(name string, params ...interface{}) (interface{}, error) {
	defer protect()
	if fn, ok := s.fns[name]; ok {
		if checkIsFn(fn) {
			paramsV := paramConversion(params) // 转换参数
			v := reflect.ValueOf(fn)
			rets := v.Call(paramsV)
			retsV := resultConversion(rets) // 转换返回值
			return retsV, nil
		}
		return nil, errors.New(name + "is not a Func")
	}
	return nil, errors.New("Func " + name + " does not exist")
}

func protect() {
	err := recover()
	if err != nil {
		switch err.(type) {
		case runtime.Error:
			fmt.Println("runtime error:", err)
		default:
			fmt.Println("error:", err)
		}
	}
}

func paramConversion(params []interface{}) []reflect.Value {
	paramsV := make([]reflect.Value, len(params))
	for i, p := range params {
		paramsV[i] = reflect.ValueOf(p)
	}
	return paramsV
}

func resultConversion(rets []reflect.Value) []interface{} {
	retsV := make([]interface{}, len(rets))
	for i, ret := range rets {
		retsV[i] = ret.Interface()
	}
	return retsV
}
