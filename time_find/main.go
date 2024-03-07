package main

import (
	"io"
	"os"

	parser "github.com/accDB/funcs"
)

func main() {
	var in parser.TimeFindIn
	var out parser.TimeFindOut
	err := in.From(os.Args)
	if err != nil {
		writeErr(err.Error())
	}
	out.LogTime, out.Log = compileFunc(in.Log)
	data, err := out.To()
	if err != nil {
		writeErr(err.Error())
	}
	_, _ = os.Stdout.Write(data)
}

func writeErr(s string) {
	_, _ = io.WriteString(os.Stderr, s)
}
