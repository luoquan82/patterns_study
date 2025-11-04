package flyweight

import "testing"

func Test_FlyWeight(t *testing.T) {
	pool := NewDBPool(1)

	db1 := pool.GetDB()
	db1.Query("select * from XXX")
	pool.ReleaseDB(db1)

	db2 := pool.GetDB()
	db2.Query("select * from XXX")
	pool.ReleaseDB(db2)

	db3 := pool.GetDB()
	db1.Query("select * from XXX")
	pool.ReleaseDB(db3)
}
