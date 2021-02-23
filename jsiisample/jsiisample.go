// hello, world
package jsiisample

import (
	_jsii_ "github.com/aws/jsii-runtime-go"
	_init_ "github.com/awslabs/aws-delivlib-sample/jsiisample/jsii"
	"reflect"
)

// Class interface
type HelloJsiiIface interface {
	SayGoodbye(times float64) string
	SayHello(name string) string
}

// This guy knows how to say hello, in multiple langauges.
// Struct proxy
type HelloJsii struct {
}

func NewHelloJsii(props HelloJsiiPropsIface) HelloJsiiIface {
	_init_.Initialize()
	self := HelloJsii{}
	_jsii_.Create(
		"jsii-sample.HelloJsii",
		[]interface{}{props},
		[]_jsii_.FQN{},
		[]_jsii_.Override{},
		&self,
	)
	return &self
}

func (h *HelloJsii) SayGoodbye(times float64) string {
	var returns string
	_jsii_.Invoke(
		h,
		"sayGoodbye",
		[]interface{}{times},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

func (h *HelloJsii) SayHello(name string) string {
	var returns string
	_jsii_.Invoke(
		h,
		"sayHello",
		[]interface{}{name},
		true,
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}

// HelloJsiiPropsIface is the public interface for the custom type HelloJsiiProps
type HelloJsiiPropsIface interface {
	GetGoodbyeMessage() string
}

// Struct proxy
type HelloJsiiProps struct {
	// The message to emit when saying goodbye.
	GoodbyeMessage string `json:"goodbyeMessage"`
}

func (h *HelloJsiiProps) GetGoodbyeMessage() string {
	var returns string
	_jsii_.Get(
		h,
		"goodbyeMessage",
		&returns,
		map[reflect.Type]reflect.Type{},
	)
	return returns
}


