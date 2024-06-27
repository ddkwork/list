package main

import (
	"list"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("list", func(w *unison.Window) {
		w.Content().AddChild(list.New().Layout())
	})
}
