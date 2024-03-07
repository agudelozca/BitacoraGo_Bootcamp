package main

import "fmt"

type Product interface {
	Price() float64 //metodo que se encarga de calcular el precio del producto y retorna el total
}

//Definicion de tipos de productos (Small, Medium, Large)

type Small struct {
	price float64
}

func (s *Small) Price() float64 {
	return s.price
}

type Medium struct {
	price float64
}

func (m *Medium) Price() float64 {
	return m.price * 1.03
}

type Large struct {
	price float64
}

func (l *Large) Price() float64 {
	return l.price*1.06 + 2500
}

func CreateProduct(productType string, price float64) Product {
	switch productType {
	case "Small":
		return &Small{price}
	case "Medium":
		return &Medium{price}
	case "Large":
		return &Large{price}
	default:
		return nil
	}
}

func main() {
	// Crear un producto pequeño y calcular su precio
	small := CreateProduct("Small", 100)
	medium := CreateProduct("Medium", 1500)

	// Obtener y mostrar el precio total del producto
	fmt.Printf("Precio Total del Producto Small: $%.2f\n", small.Price())
	fmt.Printf("Precio Total del Producto Medium: $%.2f\n", medium.Price())

}
