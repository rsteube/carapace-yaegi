package main

import (
	"github.com/mitchellh/go-homedir"
	"github.com/rsteube/carapace"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

func main() {
    _ = carapace.CallbackValue
    gopath,_ := homedir.Expand("~/go") // TODO handle error, user vendor
    i := interp.New(interp.Options{GoPath: gopath})

	i.Use(stdlib.Symbols)

    _, err := i.EvalPath("./completers/root.go")
	if err != nil {
		panic(err)
	}
}
