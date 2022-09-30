package dao

type Dao struct{}

var dao *Dao

func New() *Dao {
	return &Dao{}
}
