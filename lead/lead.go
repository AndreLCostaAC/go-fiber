package lead

import (
	"github.com/AndreLCostaAC/go-fiber/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Lead struct{
	gorm.Model
	Name 	string		`json: "name"`
	Company string		`json: "company"`
	Email 	string		`json: "email"`
	Phone 	int			`json: "phone"`
}

func GetLeads(c *fiber.Ctx){
	db := database.DbConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

func GetLead(c *fiber.Ctx){
	id := c.Params("id")
	db := database.DbConn 
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead) //responds using json
}

func DeleteLead(c *fiber.Ctx){
	id := c.Params("id")
	db := database.DbConn
	var lead Lead
	db.First(&lead, id)

	if lead.Name == ""{
		c.Status(500).Send("No lead found with ID")
		return
	}
	db.Delete(&lead)
	c.Send("Lead with ID %v deleted", id)
}

func NewLead(c *fiber.Ctx){
	db := database.DbConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil{
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}

