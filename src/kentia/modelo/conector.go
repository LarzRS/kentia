package modelo

import (
	//required for mysql conn
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//Conector Go -> Mysql que permite obtener sesiones en una bd.
type conector struct {
	db gorm.DB
}

const dburl = "root:@/kentia?charset=utf8&parseTime=True&loc=Local"

// Realiza una conexión al servidor server y selecciona la BD de kentia
func conectar() (c *conector) {
	bd, err := gorm.Open("mysql", dburl)
	if err != nil {
		return nil
	}
	return &conector{db: bd}
}

// Finaliza la conexión al servidor
func (c *conector) desconectar() {
	if c != nil {
		c.db.Close()
	}
}
