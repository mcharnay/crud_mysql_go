package models

import "gorm.io/gorm"

type Empresa struct {
	gorm.Model
	NombreEmpresa string
	RutEmpresa    string
	//Empleados     []Empleado `gorm:"constraint:OnUpdate:CASCADE;"`
	//Empleados []Empleado `gorm:"constraint:OnDelete:CASCADE;"`
	Empleados []Empleado `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
