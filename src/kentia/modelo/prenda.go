package modelo

import "kentia/log"

//Prenda define los datos importantes para una prenda.
type Prenda struct {
	ID           int `gorm:"primary_key"`
	Brillo       int `form:"brillo"`
	Foto         string
	ColorID      int
	ClimaID      int
	TipoPrendaID int
	OcasioniD    int
	UsuarioID    int
}

const coleccionPrenda = "prenda"

//Registrar inserta la prenda en BD.
func (p *Prenda) Registrar() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.Create(p).Error
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}

//Modificar datos de una prenda en la base
func (p *Prenda) Modificar() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.First(p).Save(p).Error
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}

//ConsultarPrendas regresa un catálogo de prendas.
func ConsultarPrendas() (prendas []Prenda) {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.Model(&Prenda{}).Find(&prendas).Error
	if err != nil {
		log.RegistrarError(err)
	}
	return prendas
}

//BuscarPorID busca en la BD un color que coincida con el ID dado.
func (p *Prenda) BuscarPorID() bool {
	conn := conectar()
	defer conn.desconectar()
	err := conn.db.Find(p).First(p).Error
	if err != nil {
		log.RegistrarError(err)
		return false
	}
	return true
}

//BuscarPorBrilloTono busca en la BD un color que coincida con el tono y el brillo.
func (p *Prenda) BuscarPorBrilloTono(prendas []Prenda) bool {
	colorOriginal := Color{ID: p.ColorID}
	colorOriginal.BuscarPorID()
	tipoOriginal := TipoPrenda{ID: p.TipoPrendaID}
	tipoOriginal.BuscarPorID()
	for _, prenda := range prendas {
		colorBusqueda := Color{ID: prenda.ColorID}
		colorBusqueda.BuscarPorID()
		tipoBusqueda := TipoPrenda{ID: prenda.TipoPrendaID}
		tipoBusqueda.BuscarPorID()
		if tipoOriginal.Nombre == tipoBusqueda.Nombre && colorOriginal.Tono == colorBusqueda.Tono && p.Brillo == prenda.Brillo {
			*p = prenda
			return true
		}
	}
	return false
}
