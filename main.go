package main

import (
	"go_practice/arrayslices"
	"go_practice/booleanconditional"
	"go_practice/format"
	"go_practice/functions"
	"go_practice/helloworld"
	lopps "go_practice/loops"
	"go_practice/maps"
	"go_practice/packages"
	"go_practice/pointers"
	"go_practice/scopes"
	"go_practice/structsreciever"
	"go_practice/variablesstringsnumbers"
)

func main() {
	format.Run()
	booleanconditional.Run()
	arrayslices.RunArray()
	arrayslices.RunSlices()
	functions.Run()
	helloworld.Run()
	lopps.Run()
	packages.Run()
	scopes.Run()
	variablesstringsnumbers.RunFloat()
	variablesstringsnumbers.RunInt()
	variablesstringsnumbers.RunVar()
	maps.Run()
	pointers.Run()
	structsreciever.Run()
}
