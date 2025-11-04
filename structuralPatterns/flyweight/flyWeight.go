package flyweight

import "fmt"

type DB struct {
	id int
}

func NewDB(id int) *DB {
	return &DB{id: id}
}

func (db DB) Query(str string) {
	fmt.Printf("使用 %d 连接对象 进行查询操作 %s\n", db.id, str)
}

type DBPool struct {
	pool   map[int]*DB
	nextID int
}

func NewDBPool(num int) *DBPool {
	var pool = map[int]*DB{}

	for i := 0; i < num; i++ {
		pool[i] = NewDB(i)
	}

	return &DBPool{pool: pool, nextID: num - 1}
}

// GetDB 从连接池里面获取连接
func (p *DBPool) GetDB() *DB {
	if len(p.pool) > 0 {
		for id, db := range p.pool {
			delete(p.pool, id)
			return db
		}
	}

	p.nextID++
	db := NewDB(p.nextID)
	return db
}

// ReleaseDB 释放连接池
func (p *DBPool) ReleaseDB(db *DB) {
	p.pool[db.id] = db
}
