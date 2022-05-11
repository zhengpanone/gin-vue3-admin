package utils

var (
	RegisterVerify = Rules{"UserName": {NotEmpty()}, "Password": {NotEmpty()}}
)
