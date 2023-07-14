package session

//session用于实现和数据库交互
import (
	"database/sql"
	"gee-orm/clause"
	"gee-orm/dialect"
	"gee-orm/log"
	"gee-orm/schema"
	"strings"
)

// db 连接数据库成功后返回的指针
// 第二三个变量用来拼接SQL语句和SQL语句中占位符的对应值
type Session struct {
	db      *sql.DB
	sql     strings.Builder
	sqlVars []interface{}

	dialect  dialect.Dialect
	refTable *schema.Schema

	clause clause.Clause

	tx *sql.Tx //支持事务
}
type CommonDB interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// _ CommonDB表示声明了一个类型为CommonDB的变量，但使用了下划线 _来忽略它的值。
// (*sql.DB)(nil)和(*sql.Tx)(nil)是将nil转换为对应类型的指针,这样的赋值表达式实际上是为了确保*sql.DB和*sql.Tx类型都实现了CommonDB接口。
var _ CommonDB = (*sql.DB)(nil)
var _ CommonDB = (*sql.Tx)(nil)

func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{db: db,
		dialect: dialect,
	}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
	s.clause = clause.Clause{}
}

func (s *Session) DB() CommonDB {
	if s.tx != nil {
		return s.tx
	}
	return s.db
}

func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}
