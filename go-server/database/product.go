package database

type Product struct {
	ID          int     `json:"id"` // renames the json key
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"imageUrl"`

	// camelCase makes these variables private, other packages cannot access --> postman too
}

var ProductList []Product // slice of products

func init() {
	prod1 := Product{
		ID:          1,
		Title:       "Product 1",
		Description: "This is a product 1",
		Price:       10.00,
		ImageUrl:    "https://via.placeholder.com/150",
	}

	prod2 := Product{
		ID:          2,
		Title:       "Product 2",
		Description: "This is a product 2",
		Price:       20.00,
		ImageUrl:    "https://via.placeholder.com/150",
	}

	prod3 := Product{
		ID:          3,
		Title:       "Product 3",
		Description: "This is a product 3",
		Price:       30.00,
		ImageUrl:    "https://via.placeholder.com/150",
	}

	ProductList = append(ProductList, prod1, prod2, prod3)
}
