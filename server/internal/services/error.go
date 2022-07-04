package services

import "github.com/Tauhoo/adon-desktop/internal/errors"

var (
	BuildPluginFailCode     errors.Code = "BUILD_PLUGIN_FAIL"
	LoadPluginFailCode      errors.Code = "LOAD_PLUGIN_FAIL"
	PluginNotFoundCode      errors.Code = "PLUGIN_NOT_FOUND"
	FunctionNotFoundCode    errors.Code = "FUNCTION_NOT_FOUND"
	FunctionArgsInvalidCode errors.Code = "FUNCTION_ARGS_INVALID"
)
