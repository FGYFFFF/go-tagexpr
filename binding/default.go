package binding

import "net/http"

var defaultBinding = New(nil)

// Default returns the default binding.
// NOTE:
//  path tag name is 'path';
//  query tag name is 'query';
//  header tag name is 'header';
//  cookie tag name is 'cookie';
//  rawbody tag name is 'rawbody';
//  form tag name is 'form';
//  validator tag name is 'vd';
//  protobuf tag name is 'protobuf';
//  json tag name is 'json'.
func Default() *Binding {
	return defaultBinding
}

// SetErrorFactory customizes the factory of validation error.
// NOTE:
//  If errFactory==nil, the default is used
func SetErrorFactory(bindErrFactory, validatingErrFactory func(failField, msg string) error) {
	defaultBinding.SetErrorFactory(bindErrFactory, validatingErrFactory)
}

// BindAndValidate binds the request parameters and validates them if needed.
func BindAndValidate(structPointer interface{}, req *http.Request, pathParams PathParams) error {
	return defaultBinding.BindAndValidate(structPointer, req, pathParams)
}

// Bind binds the request parameters.
func Bind(structPointer interface{}, req *http.Request, pathParams PathParams) error {
	return defaultBinding.Bind(structPointer, req, pathParams)
}

// Validate validates whether the fields of value is valid.
func Validate(value interface{}) error {
	return defaultBinding.Validate(value)
}