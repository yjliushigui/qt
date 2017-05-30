package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quick"

	"github.com/therecipe/qt/internal/examples/qml/extending/components/test_dir/component"
)

func main() {
	gui.NewQGuiApplication(len(os.Args), os.Args)

	view := quick.NewQQuickView(nil)
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.RootContext().SetContextProperty("factory", component.NewPieChartFactory(nil))
	view.SetSource(core.NewQUrl3("qrc:///app.qml", 0))

	view.Show()

	gui.QGuiApplication_Exec()
}
