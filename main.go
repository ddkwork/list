package main

import (
	"list"

	"github.com/ddkwork/unison"
	"github.com/ddkwork/unison/app"
)

func main() {
	app.Run("list", func(w *unison.Window) {
		w.Content().AddChild(list.New().Layout())
	})
}
