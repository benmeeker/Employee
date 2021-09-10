package main

type info struct {
	First       string
	Last        string
	POU         string
	SOU         string
	TCID        int
	TWOST       int
	EXT         int
	DATE        string
	ALIAS       string
	GROUPS      []string
	EMAILGROUPS []string
}

var ui = info{
	fn(),
	ln(),
	pou(),
	sou(),
	mantcid(),
	twost(),
	ext(),
	date(),
	"",
	groups(),
	nil,
}
