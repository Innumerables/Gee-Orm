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
}

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

func (s *Session) DB() *sql.DB {
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
