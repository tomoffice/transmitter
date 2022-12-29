package sql

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var testDB Mysql

func init() {
	db := &Mysql{Config: Config{
		Username:  "",
		Password:  "",
		Protocol:  "tcp",
		Host:      "",
		Port:      3306,
		Database:  "",
		Chartset:  "utf8",
		Location:  "Asia%2FTaipei",
		Parsetime: true,
	}}
	db.Conn()
	testDB.Config = db.Config
	testDB.DB = db.DB
}
func TestMysql_Conn(t *testing.T) {
	tests := []struct {
		name    string
		m       *Mysql
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "connect test",
			m: &Mysql{
				Config: testDB.Config,
				DB:     &sql.DB{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.Conn(); (err != nil) != tt.wantErr {
				t.Errorf("Mysql.Conn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMysql_Select(t *testing.T) {
	//init connect for this testing
	var dd *sql.DB
	func(*sql.DB) {
		db := &Mysql{Config: Config{}}
		db.Conn()
		dd = db.DB
	}(dd)

	type args struct {
		q string
	}
	tests := []struct {
		name    string
		m       Mysql
		args    args
		want    []map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "mysql select test",
			m: Mysql{
				Config: Config{},
				DB:     testDB.DB,
			},
			args: args{
				q: "SELECT * FROM record LIMIT 10",
			},
			//want:    []map[string]interface{}{},
			//we don't know db will output something else so block the return result
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.Select(tt.args.q)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mysql.Select() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got) //just print check value is correct or not
			/*if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mysql.Select() = %v, want %v", got, tt.want)
			}*/
		})
	}
}
