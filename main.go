package main

import (
	"ecommerce-poo/cart"
	"ecommerce-poo/order"
	"ecommerce-poo/product"
	"fmt"
	"os"

	"github.com/google/uuid"
)

func main() {

	catalogo := make(map[string]*product.Product)

	p1ID := uuid.New().String()
	p2ID := uuid.New().String()

	mesa := product.NewProduct(p1ID, "Mesa de Comedor Madera", 250.00, 100)
	silla := product.NewProduct(p2ID, "Silla de Comedor Tapizada", 75.00, 100)

	catalogo["mesa"] = &mesa
	catalogo["silla"] = &silla

	myCart := cart.Cart{}
	ecuadorFiscal := order.EcuadorTax{Rate: 0.15}

	//bucle for
	for {
		fmt.Println("\n=========================================")
		fmt.Println("    UIDE DESIGN    ")
		fmt.Println("=========================================")
		fmt.Println("1. Ver Catálogo de Productos y Stock")
		fmt.Println("2. Agregar Muebles al Carrito")
		fmt.Println("3. Procesar Compra (Checkout)")
		fmt.Println("0. Salir del Sistema")
		fmt.Print("👉 Seleccione una opción: ")

		var opcion int
		fmt.Scan(&opcion)

		//estructura switch-case
		switch opcion {
		case 1:
			fmt.Println("\n--- CATÁLOGO EN BODEGA ---")
			for clave, prod := range catalogo {
				fmt.Printf("• Código para comprar: [%s] | %s - Precio: $%.2f | Stock: %d unids.\n",
					clave, prod.Name, prod.Price, prod.GetStock())
			}

		case 2:
			fmt.Println("\n--- AGREGAR PRODUCTOS AL CARRITO ---")
			var codigoProducto string
			var cantidad int

			fmt.Print("Escriba el código del producto (mesa / silla): ")
			fmt.Scan(&codigoProducto)

			prod, existe := catalogo[codigoProducto]
			if !existe {
				fmt.Println("❌ Error: Ese código de producto no existe en el catálogo.")
				continue
			}

			fmt.Printf("¿Cuántas unidades de '%s' desea agregar?: ", prod.Name)
			fmt.Scan(&cantidad)

			if cantidad > 0 {
				myCart.AddItem(*prod, cantidad)
				_ = prod.ReduceStock(cantidad)
				fmt.Println("✅ ¡Producto añadido al carrito con éxito!")
			} else {
				fmt.Println("❌ Error: La cantidad debe ser mayor a 0.")
			}

		case 3:
			fmt.Println("\nProcesando su pedido en el sistema...")
			myOrder := order.Order{}
			orderID := uuid.New().String()

			myOrder.ProcessCheckout(orderID, myCart, ecuadorFiscal)

			fmt.Println("\n======= FACTURACIÓN ELECTRÓNICA  =======")
			fmt.Printf("ID de Orden: %s\n", myOrder.ID)
			fmt.Printf("Fecha:       %s\n", myOrder.CreatedAt.Format("2006-01-02 15:04:05"))
			fmt.Println("-----------------------------------------------------")

			for _, item := range myOrder.CartItems {
				fmt.Printf("- %d x %s (Precio Unitario: $%.2f)\n",
					item.Quantity, item.Product.Name, item.Product.Price)
			}

			fmt.Println("-----------------------------------------------------")
			fmt.Printf("Subtotal:       $%.2f\n", myOrder.Subtotal)
			fmt.Printf("IVA (15%%):      $%.2f\n", myOrder.Tax)
			fmt.Printf("Total a Pagar:  $%.2f\n", myOrder.Total)
			fmt.Println("=====================================================")
			fmt.Println("¡Gracias por su compra! Saliendo del sistema...")
			os.Exit(0)
		case 0:
			fmt.Println("Saliendo del sistema...")
			os.Exit(0)

		default:
			fmt.Println("❌ Opción inválida. Por favor, intente de nuevo.")
		}
	}
}
