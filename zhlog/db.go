package zhlog

import (
	"fmt"

	"github.com/go-xorm/xorm"
)

func getSession() (session *xorm.Session) {
	orm := cloudProjectEngine()
	orm.ShowSQL(false)
	session = orm.NewSession()
	defer session.Close()
	err := session.Begin()
	Assert(err)
	return
}

func cloudProjectEngine() *xorm.Engine {
	orm, err := mysqlEngine()
	Assert(err)
	return orm
}

func mysqlEngine() (*xorm.Engine, error) {
	Host := "10.0.0.4"
	Port := "3306"
	Name := "cloudproject"
	User := "root"
	Password := "Connext@0101"
	dburl := User + ":" + Password + "@tcp(" + Host + ":" + Port + ")/" + Name + "?charset=utf8"
	fmt.Println("dburl:", dburl)
	return xorm.NewEngine("mysql", dburl)
}
