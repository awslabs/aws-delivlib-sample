// hello, world
package jsiisample

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/awslabs/aws-delivlib-sample/jsiisample/jsii"
)

// This guy knows how to say hello, in multiple langauges.
type HelloJsii interface {
	// Says goodbye, but not farewell.
	SayGoodbye(times *float64) *string
	// Say hello to someone.
	SayHello(name *string) *string
}

// The jsii proxy struct for HelloJsii
type jsiiProxy_HelloJsii struct {
	_ byte // padding
}

func NewHelloJsii(props *HelloJsiiProps) HelloJsii {
	_init_.Initialize()

	if err := validateNewHelloJsiiParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HelloJsii{}

	_jsii_.Create(
		"jsii-sample.HelloJsii",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewHelloJsii_Override(h HelloJsii, props *HelloJsiiProps) {
	_init_.Initialize()

	_jsii_.Create(
		"jsii-sample.HelloJsii",
		[]interface{}{props},
		h,
	)
}

func (h *jsiiProxy_HelloJsii) SayGoodbye(times *float64) *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"sayGoodbye",
		[]interface{}{times},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HelloJsii) SayHello(name *string) *string {
	if err := h.validateSayHelloParameters(name); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"sayHello",
		[]interface{}{name},
		&returns,
	)

	return returns
}

