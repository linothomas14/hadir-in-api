package param

type AddProduct struct {
	Name  string `json:"name" validate:"required"`
	Price uint32 `json:"price" validate:"required"`
}
