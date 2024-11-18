package jsiisample

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"jsii-sample.HelloJsii",
		reflect.TypeOf((*HelloJsii)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "sayGoodbye", GoMethod: "SayGoodbye"},
			_jsii_.MemberMethod{JsiiMethod: "sayHello", GoMethod: "SayHello"},
		},
		func() interface{} {
			return &jsiiProxy_HelloJsii{}
		},
	)
	_jsii_.RegisterStruct(
		"jsii-sample.HelloJsiiProps",
		reflect.TypeOf((*HelloJsiiProps)(nil)).Elem(),
	)
}
