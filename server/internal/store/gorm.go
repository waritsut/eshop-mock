package store

import (
	"database/sql"
	"errors"
	"eshop-mock-api/configs"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)
 
var Db Repo

type Repo interface {
	Joins(query string, args ...interface{}) *DbRepo
	Count(count *int64) *DbRepo
	Offset(offset int) *DbRepo
	Limit(limit int) *DbRepo
	Order(value interface{}) *DbRepo
	Clauses(conds ...clause.Expression) *DbRepo
	Attrs(attrs ...interface{}) *DbRepo
	FirstOrInit(dest interface{}, conds ...interface{}) *DbRepo
	FirstOrCreate(dest interface{}, conds ...interface{}) *DbRepo
	SavePoint(name string) *DbRepo
	RollbackTo(name string) *DbRepo
	Create(value interface{}) *DbRepo
	Table(name string, args ...interface{}) *DbRepo
	Model(value interface{}) *DbRepo
	Select(query interface{}, args ...interface{}) *DbRepo
	Omit(columns ...string) *DbRepo
	Save(value interface{}) *DbRepo
	Raw(sql string, values ...interface{}) *DbRepo
	Begin(opts ...*sql.TxOptions) *DbRepo
	Rollback() *DbRepo
	Exec(sql string, values ...interface{}) *DbRepo
	First(dest interface{}, conds ...interface{}) *DbRepo
	Find(dest interface{}, conds ...interface{}) *DbRepo
	Where(query interface{}, args ...interface{}) *DbRepo
	Preload(query string, args ...interface{}) *DbRepo
	Delete(value interface{}, conds ...interface{}) *DbRepo
	Session(config *gorm.Session) *DbRepo
	Scan(dest interface{}) *DbRepo
	Unscoped() *DbRepo
}

type DbRepo struct {
	*gorm.DB
}

func NewMyStore(db *gorm.DB) *DbRepo {
	return &DbRepo{DB: db}
}

func (d *DbRepo) Scan(dest interface{}) *DbRepo {
	return NewMyStore(d.DB.Scan(dest))
}

func (d *DbRepo) Session(config *gorm.Session) *DbRepo {
	return NewMyStore(d.DB.Session(config))
}

func (d *DbRepo) Delete(value interface{}, conds ...interface{}) *DbRepo {
	return NewMyStore(d.DB.Delete(value, conds...))
}

func (d *DbRepo) Joins(query string, args ...interface{}) *DbRepo {
	return NewMyStore(d.DB.Joins(query, args...))
}
func (d *DbRepo) Count(count *int64) *DbRepo {
	return NewMyStore(d.DB.Count(count))
}

func (d *DbRepo) Offset(offset int) *DbRepo {
	return NewMyStore(d.DB.Offset(offset))
}

func (d *DbRepo) Limit(limit int) *DbRepo {
	return NewMyStore(d.DB.Limit(limit))
}

func (d *DbRepo) Order(value interface{}) *DbRepo {
	return NewMyStore(d.DB.Order(value))
}

func (d *DbRepo) Clauses(conds ...clause.Expression) *DbRepo {
	return NewMyStore(d.DB.Clauses(conds...))
}

func (d *DbRepo) Attrs(attrs ...interface{}) *DbRepo {
	return NewMyStore(d.DB.Attrs(attrs...))
}

func (d *DbRepo) FirstOrInit(dest interface{}, conds ...interface{}) *DbRepo {
	return NewMyStore(d.DB.FirstOrInit(dest, conds...))
}

func (d *DbRepo) FirstOrCreate(dest interface{}, conds ...interface{}) *DbRepo {
	return NewMyStore(d.DB.FirstOrCreate(dest, conds...))
}

func (d *DbRepo) SavePoint(name string) *DbRepo {
	return NewMyStore(d.DB.SavePoint(name))
}

func (d *DbRepo) RollbackTo(name string) *DbRepo {
	return NewMyStore(d.DB.RollbackTo(name))

}

func (d *DbRepo) Create(value interface{}) *DbRepo {
	return NewMyStore(d.DB.Create(value))
}

func (d *DbRepo) Table(name string, args ...interface{}) *DbRepo {
	return NewMyStore(d.DB.Table(name, args...))
}

func (d *DbRepo) Model(value interface{}) *DbRepo {
	return NewMyStore(d.DB.Model(value))
}

func (d *DbRepo) Select(query interface{}, args ...interface{}) *DbRepo {
	return NewMyStore(d.DB.Select(query, args...))
}

func (d *DbRepo) Omit(columns ...string) *DbRepo {
	return NewMyStore(d.DB.Omit(columns...))
}

func (d *DbRepo) Save(value interface{}) *DbRepo {
	return NewMyStore(d.DB.Save(value))
}

func (d *DbRepo) Raw(sql string, values ...interface{}) *DbRepo {
	return NewMyStore(d.DB.Raw(sql, values...))
}

func (d *DbRepo) Begin(opts ...*sql.TxOptions) *DbRepo {
	return NewMyStore(d.DB.Begin(opts...))
}

func (d *DbRepo) Rollback() *DbRepo {
	return NewMyStore(d.DB.Rollback())
}

func (d *DbRepo) Exec(sql string, values ...interface{}) *DbRepo {
	return NewMyStore(d.DB.Exec(sql, values...))
}

func (d *DbRepo) First(dest interface{}, conds ...interface{}) *DbRepo {
	return NewMyStore(d.DB.First(dest, conds...))
}

func (d *DbRepo) Find(dest interface{}, conds ...interface{}) *DbRepo {
	return NewMyStore(d.DB.Find(dest, conds...))
}

func (d *DbRepo) Preload(query string, args ...interface{}) *DbRepo {
	return NewMyStore(d.DB.Preload(query, args...))
}

func (d *DbRepo) Where(query interface{}, args ...interface{}) *DbRepo {
	return NewMyStore(d.DB.Where(query, args...))
}

func (d *DbRepo) Unscoped() *DbRepo {
	return NewMyStore(d.DB.Unscoped())
}

func New() (Repo, error) {
	db, err := dbConnect(configs.Get().DbUser, configs.Get().DbPassword,
		configs.Get().DbHost, configs.Get().DbName, configs.Get().DbPort)

	if err != nil {
		return nil, err
	}
	return &DbRepo{DB: db}, nil
}

func dbConnect(user, pass, addr, dbName, port string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", addr, user, pass, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, errors.New(fmt.Sprintf("[db connection failed] Database name: %s", dbName))
	}

	return db, nil
}
