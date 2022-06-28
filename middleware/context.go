package middleware

type ContextKey string

const InputBodyKey ContextKey = "input_data"
const InputParamsKey ContextKey = "input_params"
const OutputDataKey ContextKey = "output_data"
const ErrorKey ContextKey = "error"
const TokenKey ContextKey = "token"