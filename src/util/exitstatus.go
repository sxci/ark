package util

const (
	EXIT_SUCCESS int = iota
	EXIT_FAILURE_SETUP
	EXIT_FAILURE_PARSE
	EXIT_FAILURE_CONSTRUCTOR
	EXIT_FAILURE_SEMANTIC
	EXIT_FAILURE_CODEGEN
)
