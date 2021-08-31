package main

type info struct {
	First string
	Last  string
	POU   string
	SOU   string
	TCID  int
}

var ui = info{
	fn(),
	ln(),
	pou(),
	sou(),
	0,
}
