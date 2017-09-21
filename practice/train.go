package main

type Product struct{}

func (p Product) GetName() string {
	str := string("")
	return str
}

type Factory struct{}

func (f *Factory) CreateProduct() Product {return Product{}}

type CPU struct {
	p     Product
	model string
}

func (c *CPU) CPU(model string){
	c.model = model
}
func (c CPU) GetName() string {
	return "CPU with model: " + c.model
}

type Intel struct {
	f Factory
}

func (intl *Intel) CreateProduct() Product {
	return new(CPU).p
}

type GPU struct {
	p      Product
	model  string
	opengl bool
}

func (g *GPU) GPU(model string, supportOpenGL bool){
	g.model = model
	g.opengl = supportOpenGL
}
func (g GPU) GetName() string {
	s := string("GPU with model: " + g.model + "OpenGL?")
	tmp := string("")
	if g.opengl {
		tmp = tmp + "YES"
	} else {
		tmp = tmp + "NO"
	}
	return s + "" + tmp
}

type NVidia struct {
	f Factory
}

func (n *NVidia) CreateProduct() Product {
	return n.f.CreateProduct()
}

func describeProductFromFactory(f *Factory) {
	product := f.CreateProduct()
	println(product.GetName())
	product = Product{}
}
