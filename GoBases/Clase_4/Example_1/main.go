package main

/*
Crear un programa que cumpla los siguiente puntos:

Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
Tener un slice global de Product llamado Products instanciado con valores.
2 métodos asociados a la estructura Product: Save(), GetAll(). El método Save()

	deberá tomar el slice de Products y añadir el producto desde el cual se llama al método.
	El método GetAll() deberá imprimir todos los productos guardados en el slice Products.

Una función getById() al cual se le deberá pasar un INT como parámetro y retorna el producto correspondiente al parámetro pasado.
Ejecutar al menos una vez cada método y función definido desde main().
*/
import (
	"fmt"
)

// Definir la estructura Product
type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

func (p Product) Save() {
	Products = append(Products, p)
}

func (p Product) GetAll() {
	for _, product := range Products {
		fmt.Printf("ID: %d\nName: %s\nPrice: $%.2f\nDescription: %s\nCategory: %s\n\n", product.ID, product.Name, product.Price, product.Description, product.Category)
	}
}

func getById(id int) *Product {
	for _, product := range Products {
		if product.ID == id {
			return &product
		}
	}
	return nil
}

// Crear un slice global de Product llamado Products
var Products = []Product{
	{ID: 1, Name: "Producto1", Price: 19.99, Description: "Descripción del Producto1", Category: "Electrónicos"},
	{ID: 2, Name: "Producto2", Price: 29.99, Description: "Descripción del Producto2", Category: "Ropa"},
	{ID: 3, Name: "Producto3", Price: 39.99, Description: "Descripción del Producto3", Category: "Hogar"},
}

func main() {
	// Crear un nuevo producto
	newProduct := Product{
		ID:          4,
		Name:        "Nuevo Producto",
		Price:       49.99,
		Description: "Descripción del Nuevo Producto",
		Category:    "Accesorios",
	}

	// Llamar al método Save() para agregar el nuevo producto al slice Products
	newProduct.Save()

	// Llamar al método GetAll() para imprimir todos los productos
	fmt.Println("Lista de Productos:")
	newProduct.GetAll()

	// Llamar a la función getById() para obtener un producto por ID
	idToGet := 2
	productByID := getById(idToGet)
	if productByID != nil {
		fmt.Printf("\nProducto con ID %d encontrado:\n", idToGet)
		fmt.Printf("ID: %d\nName: %s\nPrice: $%.2f\nDescription: %s\nCategory: %s\n\n", productByID.ID, productByID.Name, productByID.Price, productByID.Description, productByID.Category)
	} else {
		fmt.Printf("\nProducto con ID %d no encontrado.\n", idToGet)
	}
}
