package sql

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type Config struct {
	Username  string
	Password  string
	Protocol  string
	Host      string
	Port      int
	Database  string
	Chartset  string
	Location  string
	Parsetime bool
}
type Database interface { //抽象產品
	ParseConfig(Map map[string]interface{})
	Conn() error
	Select(string) ([]map[string]interface{}, error)
	Insert(string) error
	Update(string) error
	Delete(string) error
	Close() error
}

// ----------------------------------------------------------------mysql

type Mysql struct { //具體產品
	Config
	DB *sql.DB
}

// set map[string]interface{}
//
//	map[string]interface{}{
//		"dbusername": w.wDB.DBpassword,
//		"dbpassword": w.wDB.DBpassword,
//		"protocol":   w.wDB.Protocol,
//		"host":       w.wDB.Host,
//		"port":       w.wDB.Port,
//		"database":   w.wDB.Database,
//		"charset":    w.wDB.Charset,
//		"location":   w.wDB.Location,
//		"parsetime":  w.wDB.Parsetime,
//	}
func (m *Mysql) ParseConfig(Map map[string]interface{}) {
	config := &Config{
		Username:  Map["dbusername"].(string),
		Password:  Map["dbpassword"].(string),
		Protocol:  Map["protocol"].(string),
		Host:      Map["host"].(string),
		Port:      Map["port"].(int),
		Database:  Map["database"].(string),
		Chartset:  Map["charset"].(string),
		Location:  Map["location"].(string),
		Parsetime: Map["parsetime"].(bool),
	}
	m.Config = *config
}

func (m *Mysql) Conn() error {
	url := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=%s&loc=%s&parseTime=%s", m.Config.Username, m.Config.Password, m.Config.Protocol, m.Config.Host, m.Config.Port, m.Config.Database, m.Config.Chartset, m.Config.Location, strconv.FormatBool(m.Config.Parsetime))
	db, err := sql.Open("mysql", url)
	if err != nil {
		//log.Println("開啟 MySQL 連線發生錯誤，原因為：", err)
		return err
	}
	if err := db.Ping(); err != nil {
		//log.Println("資料庫連線錯誤，原因為：", err.Error())
		return err
	}
	m.DB = db
	return nil
}
func (m Mysql) Select(q string) ([]map[string]interface{}, error) {

	rows, err := m.DB.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	colNames, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	colTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	colNamesNum := len(colNames)
	tableStruct := make([]interface{}, colNamesNum)
	for i, v := range colTypes {
		t := v.ScanType()
		v := reflect.New(t).Interface()
		tableStruct[i] = v
		//fmt.Println(colNames[i], t)
	}
	var list []map[string]interface{}
	for rows.Next() {

		err := rows.Scan(tableStruct...)
		if err != nil {
			return nil, err
		}
		items := make(map[string]interface{}, colNamesNum)
		for i, v := range tableStruct {
			value := v
			switch valueType := value.(type) {
			case *int32:
				items[colNames[i]] = *valueType
			case *sql.NullString:
				items[colNames[i]] = *valueType
			case *sql.NullBool:
				items[colNames[i]] = *valueType
			case *sql.NullFloat64:
				items[colNames[i]] = *valueType
			case *sql.NullInt64:
				items[colNames[i]] = *valueType
			case *sql.RawBytes:
				items[colNames[i]] = string(*valueType)
			case *sql.NullTime:
				items[colNames[i]] = *valueType
			case *sql.NullByte:
				items[colNames[i]] = *valueType
			case *sql.NullInt16:
				items[colNames[i]] = *valueType
			case *sql.NullInt32:
				items[colNames[i]] = *valueType
			default:
				items[colNames[i]] = valueType
				panic("unknow type")
			}
		}
		list = append(list, items)
	}
	if len(list) == 0 {
		return nil, errors.New("query no result")
	}
	return list, nil
}
func (m Mysql) Insert(q string) error {
	_, err := m.DB.Exec(q)
	return err
}
func (m Mysql) Update(q string) error {
	_, err := m.DB.Exec(q)
	return err
}
func (m Mysql) Delete(q string) error {
	_, err := m.DB.Exec(q)
	return err
}
func (m Mysql) Close() error {
	err := m.DB.Close()
	return err
}
func NewMysql() *Mysql {
	return &Mysql{}
}

// ----------------------------------------------------------------postgre
type Psql struct { //具體產品
	Config
	DB *sql.DB
}

func (p *Psql) ParseConfig(Map map[string]interface{}) {

	config := &Config{
		Username:  Map["dbusername"].(string),
		Password:  Map["dbpassword"].(string),
		Protocol:  Map["protocol"].(string),
		Host:      Map["host"].(string),
		Port:      Map["port"].(int),
		Database:  Map["database"].(string),
		Chartset:  Map["charset"].(string),
		Location:  Map["location"].(string),
		Parsetime: Map["parsetime"].(bool),
	}
	p.Config = *config
}
func (p *Psql) Conn(c Config) error {
	url := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.Username, c.Password, c.Database)
	db, err := sql.Open("postgres", url)
	if err != nil {
		fmt.Println("開啟 MySQL 連線發生錯誤，原因為：", err)
		return err
	}

	if err := db.Ping(); err != nil {
		fmt.Println("資料庫連線錯誤，原因為：", err.Error())
		return err
	}
	p.DB = db
	return nil
}
func (p Psql) Select(q string) ([]map[string]interface{}, error) {

	rows, err := p.DB.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	colNames, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	colTypes, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}
	colNamesNum := len(colNames)
	tableStruct := make([]interface{}, colNamesNum)
	for i, v := range colTypes {
		t := v.ScanType()
		v := reflect.New(t).Interface()
		tableStruct[i] = v
		//fmt.Println(colNames[i], t)
	}
	var list []map[string]interface{}
	for rows.Next() {

		err := rows.Scan(tableStruct...)
		if err != nil {
			return nil, err
		}
		items := make(map[string]interface{}, colNamesNum)
		for i, value := range tableStruct {
			switch v := value.(type) {
			case *sql.NullString:
				items[colNames[i]] = *v
			case *sql.NullBool:
				items[colNames[i]] = *v
			case *sql.NullFloat64:
				items[colNames[i]] = *v
			case *sql.NullInt64:
				items[colNames[i]] = *v
			case *sql.RawBytes:
				items[colNames[i]] = string(*v)
			case *sql.NullTime:
				items[colNames[i]] = *v
			case *sql.NullByte:
				items[colNames[i]] = *v
			case *sql.NullInt16:
				items[colNames[i]] = *v
			case *sql.NullInt32:
				items[colNames[i]] = *v
			case *int32:
				items[colNames[i]] = *v
			case *time.Time:
				items[colNames[i]] = *v
			default:
				items[colNames[i]] = v
				fmt.Printf("unknow type %s=>%s\n", colNames[i], reflect.TypeOf(value))
				panic("unknow type")
			}
		}
		list = append(list, items)
	}
	if len(list) == 0 {
		return nil, errors.New("query no result")
	}
	return list, nil
}
func (p Psql) Insert(q string) error {
	return nil
}
func (p Psql) Update(q string) error {
	return nil
}
func (p Psql) Delete(q string) error {
	return nil
}
func (p Psql) Close() error {
	return nil
}
