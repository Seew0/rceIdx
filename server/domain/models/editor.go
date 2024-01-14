package models

type Resp struct {
	UUID   string
	STDOUT string
	STDERR string
}

type Req struct {
	UUID  string
	Code  string
	STDIN string
}

type RespStore struct {
	STDOUT string
	STDERR string
}