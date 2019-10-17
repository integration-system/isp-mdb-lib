package structure

type ExecuteByIdRequest struct {
	Id  string `valid:"required~Required"`
	Arg interface{}
}

type ExecuteRequest struct {
	Script string `valid:"required~Required"`
	Arg    interface{}
}

const (
	ErrorTypeCompile = "Compilation"
	ErrorTypeRunTime = "Runtime"
)

type ScriptResponse struct {
	Result interface{}
	Error  *ScriptResponseError
}

type ScriptResponseError struct {
	Type        string
	Description string
}