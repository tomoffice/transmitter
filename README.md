#transmitter
#### Mysql or Postgres Usage:
```go
import ( 
	Pkgsql "transmitter/sql"
)
var mysqler PkgSql.Database = PkgSql.NewMysql()
Mysqler.ParseConfig(map[string]interface{}{
	"dbusername": "XXX",
	"dbpassword": "XXX,
	"protocol":   "tcp",
	"host":       "xxx.xxx.xxx.xxx",
	"port":       3306,
	"database":   "XXX",
	"charset":    "utf8",
	"location":   "Asia%2FTaipei",
	"parsetime":  true,
})
err := mysqler.Conn()
if Pkgtools.NewHandle().Retry(err, 5, time.Second*1, mysqler.Conn, func() {
	fmt.Println("database alive")
	}, func() {
		fmt.Println("database dead already")
	}) {
		return
	}
for _, v := range w.wSensors.Sensor {
	if v.Value != nil {
	query:=fmt.Sprintf("INSERT INTO record VALUES (NULL,%d,%.2f,NOW())", v.DbID, v.Value)
	err := mysqler.Insert(query)
		if err != nil {
			continue
		}
    }
}
defer mysqler.Close()
log.Println("Write mysqler success")
```
#### Make ModbusRtu with CRC-16:
```go
var rtu PkgModbus.Modbus = PkgModbus.NewRtuCmd()
rtu.SetCmd(func(cmd string) []byte {
	data := strings.ReplaceAll(cmd, " ", "")
	output, err := hex.DecodeString(data)
	if err != nil {
		fmt.Println(err)
	}
	return output
}("01 03 00 00 00 02"))
rtu.GetCmd()
```
