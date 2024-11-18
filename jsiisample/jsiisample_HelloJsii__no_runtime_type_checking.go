//go:build no_runtime_type_checking

// hello, world
package jsiisample

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HelloJsii) validateSayHelloParameters(name *string) error {
	return nil
}

func validateNewHelloJsiiParameters(props *HelloJsiiProps) error {
	return nil
}

