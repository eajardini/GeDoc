package main

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Documento struct {
	gorm.Model
	Titulo        string    `json:"titulo"`
	ArqDocumento  []byte    `json:"arqdocumento"`
	DataDocumento time.Time `gorm:"type:date; default:null",json:"datadocumento"`
}

type DocumentoRetorno struct {
	Id            string `json:"id"`
	Titulo        string `json:"titulo"`
	DataDocumento string `json:"datadocumento"`
}

var conn *gorm.DB

func GetDocs(c *fiber.Ctx) {
	db := conn
	var Docs Documento

	IDDocumento, _ := strconv.Atoi(c.FormValue("IDDocumento"))

	db.First(&Docs, IDDocumento)

	c.Set("Content-type", "application/pdf")
	c.SendBytes(Docs.ArqDocumento)

	// c.Send(documento.ArqDocumento)
}

func GetAllDocs(c *fiber.Ctx) {

	var Datas []DocumentoRetorno

	db := conn

	SQL := `
	select to_char(id,'9999') as "Id" , titulo as "Titulo", COALESCE(to_char(data_documento, 'DD/MM/YYYY'),' ') as "DataDocumento" 
	from documentos
	`

	rows, _ := db.Raw(SQL).Rows() // (*sql.Rows, error)
	defer rows.Close()

	for rows.Next() {
		var Data DocumentoRetorno
		// ScanRows scan a row into user
		// db.ScanRows(rows, &Data)
		rows.Scan(&Data.Id, &Data.Titulo, &Data.DataDocumento)
		Data.Id = strings.TrimSpace(Data.Id)
		Datas = append(Datas, Data)
	}

	// fmt.Println("Valor da Documento:", Datas)

	c.JSON(Datas)
}

func SetDocs(c *fiber.Ctx) {
	type MSGRetorno struct {
		Resposta string
	}
	var msgRetorno MSGRetorno

	var doc Documento
	var err error
	var file *multipart.FileHeader

	doc.DataDocumento, _ = time.Parse("02/01/2006", c.FormValue("ValorDaData"))
	file, err = c.FormFile("documento")

	if err != nil {
		fmt.Println("Erro:", err.Error())

	}

	// Check for errors:
	if err == nil {
		// Save file to root directory:
		c.SaveFile(file, fmt.Sprintf("./docDestino/%s", file.Filename))

		fileMultipart, _ := file.Open()
		doc.Titulo = file.Filename
		doc.ArqDocumento, _ = ioutil.ReadAll(fileMultipart)

		conn.Create(&doc)
		msgRetorno.Resposta = "Salvo com sucesso!"
		c.JSON(msgRetorno)

	} else {
		msgRetorno.Resposta = "Não foi possível salvar o documento"
		c.SendStatus(400)
		c.JSON(msgRetorno)
	}
}

func GetStatus(c *fiber.Ctx) {
	type MSGRetorno struct {
		Resposta string
	}
	var resp MSGRetorno
	resp.Resposta = "Servidor OnLine"
	c.JSON(resp)
}

// Routes

func setupRoutes(app *fiber.App) {
	app.Post("/GetDocs", GetDocs)
	app.Get("/GetAllDocs", GetAllDocs)
	app.Get("/GetStatus", GetStatus)
	app.Post("/SetDocs", SetDocs)

}

func initDatabase() {
	var err error
	dbConnectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable ", "172.17.0.1", 5432, "postgres", "postdba", "fibertestes")

	conn, err = gorm.Open("postgres", dbConnectionString)
	if err != nil {
		fmt.Println("Erro ao realizar a conexão com a Base de Dados..: ", err)
	}

	fmt.Println("Connection Opened to Database")
	conn.AutoMigrate(&Documento{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	initDatabase()

	setupRoutes(app)
	app.Listen(3000)

	defer conn.Close()
}
