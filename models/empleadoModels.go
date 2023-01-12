package models

import (
	"time"

	"gorm.io/gorm"
)

type Empleado struct {
	gorm.Model
	Nombre          string
	Apellido        string
	Rut             string
	FechaNacimiento time.Time
	EmpresaID       uint
}
