package builder

type (
	Brand string
	Color string
	Size string
)

const (
	BMW Brand = "BMW"
	MB  	  = "Mercedes-Benz"

	Blue  Color = "Blue"
	Green       = "Green"
	Red         = "Red"

	Big Size = "Big"
	Small    = "Small"
)

type Builder interface {
	Brand(Brand) *Builder
	Color(Color) *Builder
	Size(Size) *Builder
	Build() Product
}

type Product interface {
	Drive() error
	Stop() error
}
