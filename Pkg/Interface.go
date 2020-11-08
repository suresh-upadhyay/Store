package Pkg
import (
"github.com/dunzoit/Store/Model"
)

type Repo interface {
	ProductRepo
	StoreRepo
}

type ProductRepo interface {
	InsertProduct(Model.Product) (error)
	GetProduct(ProductId int) ([]Model.Product, error)
	UpdateProduct(int) error
	DeleteProduct(ProductId int) error
	GetProducts(Model.Product) (Model.Product, error)
}

type StoreRepo interface {
	GetStoreProduct(StoreId int) ([]Model.Product, error)
	InsertTodoList(Product Model.Product) (error)
}
