package routes

import (
	"crudProject/database"
	"crudProject/models"
	"errors"
	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func createResponseProduct(productModel models.Product) Product {
	return Product{
		ID:           productModel.ID,
		Name:         productModel.Name,
		SerialNumber: productModel.SerialNumber,
	}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&product)
	responseProduct := createResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}

func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}
	database.Database.Db.Find(&products)
	responseProducts := []Product{}
	for _, product := range products {
		responseProduct := createResponseProduct(product)
		responseProducts = append(responseProducts, responseProduct)
	}
	return c.Status(200).JSON(responseProducts)
}

func FindProduct(id int, product *models.Product) error {
	database.Database.Db.Find("id = ?", id)
	if product.ID == 0 {
		return errors.New("Product not found")
	}
	return nil
}

func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product models.Product

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := FindProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseProduct := createResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product models.Product

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := FindProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateProduct struct {
		Name         string `json:"name" gorm:"not null"`
		SerialNumber string `json:"serial_number" gorm:"not null"`
	}

	var updateData UpdateProduct

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	product.Name = updateData.Name
	product.SerialNumber = updateData.SerialNumber
	database.Database.Db.Save(&product)
	responseProduct := createResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product models.Product

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := FindProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).JSON("Product successfully deleted")
}
