package main

type info struct {
	First string
	Last  string
	POU   string
	SOU   string
	TCID  int
	DATE  string
}

var ui = info{
	fn(),
	ln(),
	pou(),
	sou(),
	0,
	date(),
}
