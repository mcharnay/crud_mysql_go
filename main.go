package main

import (
	"fmt"
	"restApi/config"
	"restApi/models"
	"time"

	//"restApi/routes"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDB()
)

func main() {
	defer config.DisconnectDB(db)

	//////////////////////Migraciones///////////////////////////
	db.AutoMigrate(&models.Empresa{})
	db.AutoMigrate(&models.Empleado{})

	//CREAR EMPRESA//
	empresa := models.Empresa{NombreEmpresa: "Wit", RutEmpresa: "8.555.444-4"}
	db.Create(&empresa) // pass pointer of data to Create

	//CREAR EMPLEADO//
	fecha := time.Now()
	empleado := models.Empleado{Nombre: "Steve", Apellido: "Carrel", Rut: "8.555.444-4", FechaNacimiento: fecha, EmpresaID: empresa.ID}
	db.Create(&empleado) // pass pointer of data to Create

	//iNGRESAR UN EMPLEADO A UNA EMPRESA POR ID

	//select de empresa,saca el id de la empresa.
	// empresaSelectAll := []models.Empresa{}
	// querySelectAll := db.Find(&empresaSelectAll)
	// fmt.Println(querySelectAll)
	// fmt.Println(empresaSelectAll)

	//rellenar arreglo de empleados dentro de empresa, muestra todas las empresas con sus empleados dentro.
	// var empresas []models.Empresa
	// err := db.Model(&models.Empresa{}).Preload("Empleados").Find(&empresas).Error //Empleados es el nombre de la variable array, vienen mas de 1
	// if err != nil {
	// 	fmt.Println("Error ", err)
	// }
	// fmt.Println(empresas)

	//mostrar el id recorrido por un for
	// for _, v2 := range empresas {
	// 	fmt.Printf("%v ", v2.ID)
	// }

	//agrgear empleado por id empresa
	// empresaId := models.Empresa{} //este se usará enn hacer el select * where id = es 2, yo lo puse abajo.
	// db.First(&empresaId, 1)
	// empleado = models.Empleado{Nombre: "Stivi", Apellido: "Nix", Rut: "8.555.444-4", FechaNacimiento: fecha, EmpresaID: empresaId.ID}
	// db.Create(&empleado)

	//Updatear empleado
	empleado = models.Empleado{}                                       //este se usará enn hacer el select * where id = es 2, yo lo puse abajo.
	db.First(&empleado, 6)                                             // saco el id = 2
	queryUpdatePorId := db.Model(&empleado).Update("nombre", "Tomate") //dar el update al campo.
	fmt.Println(queryUpdatePorId)

	//Updatear empresa
	// empresa = models.Empresa{}
	// db.First(&empresa, 2)
	// queryUpdatePorId = db.Model(&empresa).Update("nombre_empresa", "Super junior ltda")
	// fmt.Println(queryUpdatePorId)

	//eliminar empleado soft delete funciona
	// empleado = models.Empleado{}
	// db.Delete(&empleado, 20)
	//fmt.Println(queryBorrarPorId)

	//eliminar empresa soft delete no lo sé aún
	// empresa = models.Empresa{}
	// //db.Delete(&empresa, 3)
	// db.Select("Empleado").Delete(&empresa)

	//eliminar empleado permanente si, funciona
	// var empresaa models.Empresa
	// db.Unscoped().Where("id= ?", 2).Delete(empresaa)

	//eliminar empresa permanente no funciona
	// empresa = models.Empresa
	// db.Where("id= ?", 1).Delete(&empresa)

	//eliminar empresa de diucheba
	// var empresas models.Empresa
	// // var empl models.Empleado
	// db.Unscoped().Where("id = ?", 5).Delete(&empresas)

	//buscar find id empresa
	empresa = models.Empresa{} //este se usará enn hacer el select * where id = es 2, yo lo puse abajo.
	//db.First(&empresa, 6)
	//err := e.DB.Where("ID = ?", s).Preload("Empleados").First(&empr).Error
	db.Where("ID = ?", 7).Preload("Empleados").First(&empresa)
	db.Unscoped().Select("Empleados").Delete(&empresa)

	//////////////////////EMPLEADO/////////////////////////// No pescar esto de abajo.
	//esto migra la tabla, lo toma del model. Le paso tantos models como quiero crear.

	//Objeto con datos.
	//Defino el objeto que quiero crear, esto descomentarlo cuando insertaré este objeto.
	//empleado := models.Empleado{Nombre: "Steve", Apellido: "Carrel", Rut: "8.555.444-4"}

	//recordar que para retornar un select all, hay que dejar el modelo como arreglo.
	//empleado2 := models.Empleado{} //este se usará enn hacer el select * where id = 1
	//empleado2 := []models.Empleado{} //este se usará enn hacer el select * findAll, se hace con arreglo.

	//CREATE
	//acá hace el insert., descomentarlo cuando quieres agregar empleado
	//result := db.Create(&empleado) // pass pointer of data to Create

	//muestra las líneas afectadas.
	//fmt.Println(result.RowsAffected)

	//fmt.Println(empleado)

	//SELECT POR ID = 1
	//query que retorna el Empleado con id 1.
	//queryPrimerEmpleado := db.First(&empleado2)

	//SELECT ALL
	//querySelectAll := db.Find(&empleado2)
	//querySelectAll.RowsAffected // returns found records count, equals `len(users)`
	//querySelectAll.Error

	//UPDATE POR ID
	//primero buscar el registro con el id
	/*
		empleado3 := models.Empleado{}                                        //este se usará enn hacer el select * where id = es 2, yo lo puse abajo.
		db.First(&empleado3, 2)                                               // saco el id = 2
		queryUpdatePorId := db.Model(&empleado3).Update("nombre", "Casimiro") //dar el update al campo.
		fmt.Println(queryUpdatePorId)                                         //sólo lo uso para que no de error
	*/

	//BORRAR POR ID SOFT DELETE
	/*
		empleado4 := models.Empleado{}
		db.First(&empleado4, 1)
		db.Delete(&empleado4, 1)
		//fmt.Println(queryBorrarPorId)
	*/

	//BORRAR PERMANENTEMENTE.
	// empleado5 := models.Empleado{}
	// db.Unscoped().Delete(&empleado5, 1)

	fmt.Println("Query primer empleado")
	//fmt.Println(queryPrimerEmpleado)
	fmt.Println("Query Select all")
	//fmt.Println(querySelectAll)
	//fmt.Println(empleado2)
	//fmt.Println(empleado2.ID)
	fmt.Println()

	/////////////////////////////////////////////////////////////////////////////////////////////////////////

}
