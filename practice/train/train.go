package main

type Product struct {
	virtual interface{}
}
type VirtualProduct interface {
	Name() string
	GetName() string
}

func (p *Product) GetName() string {
	return p.virtual.(VirtualProduct).Name()
}

type Factory struct {
	virtual interface{}
}
type VirtualFactory interface {
	virFac() Product
	CreateProduct() Product
}

func (f *Factory) CreateProduct() Product {
	return f.virtual.(VirtualFactory).virFac()
}

type CPU struct {
	p     Product
	model string
}

func (c *CPU) super() *Product { return &c.p }

func (c *CPU) CPU(model string) {
	c.model = model
}
func (c *CPU) GetName() string {
	return "CPU with model: " + c.model
	//return c.GetName()
}

type Intel struct {
	f Factory
}

func (intl *Intel) super() *Factory { return &intl.f }

func (intl *Intel) CreateProduct() Product {
	cpu := CPU{model: "Core i7"}
	return cpu.p
}

type GPU struct {
	p      Product
	model  string
	opengl bool
}

func (g *GPU) super() *Product { return &g.p }

func (g *GPU) GPU(model string, supportOpenGL bool) {
	g.model = model
	g.opengl = supportOpenGL
}

func (g *GPU) GetName() string {
	s := string("GPU with model: " + g.model + " OpenGL? ")
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

func (n_vid *NVidia) super() *Factory { return &n_vid.f }

func (n_vid *NVidia) CreateProduct() Product {
	gpu := GPU{model: "GeForce GTX 745", opengl: true}
	return gpu.p
}

func describeProductFromFactory(f *Factory) {
	product := f.CreateProduct()
	println(product.GetName())
	product = Product{}
}

func main() {
	var c *CPU = &CPU{Product{nil}, "Core i5"}
	var gpu *GPU = &GPU{Product{nil}, "Nvidia", false}

	println(c.GetName())
	println(gpu.GetName())

	var intl *Intel = &Intel{Factory{Product{nil}}}
	product := intl.f.CreateProduct()
	println(product.GetName())

	var n_vid *NVidia = &NVidia{Factory{Product{nil}}}
	product = n_vid.f.CreateProduct()
	println(product.GetName())
}
