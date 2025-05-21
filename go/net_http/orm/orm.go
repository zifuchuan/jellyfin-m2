package orm

import "github.com/xormplus/xorm"

func foo() {
	// TODO

	engine, err := xorm.NewEngine("", "")
	engine.Insert()
	_ = engine
	_ = err
}
