package main

type Product interface {
	GetName() string
}

type Factory interface {
	CreateProduct(model string)
}

type CPU struct {
	model string
}

func (c *CPU) CPU(model string) {
	c.model = model
}
func (c *CPU) GetName() string {
	return "CPU with model: " + c.model
}

type Intel struct {
	CPU
}

func (intl Intel) CreateProduct(model string) CPU {
	return CPU{model}
}

type GPU struct {
	model  string
	opengl bool
}

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
	GPU
}

func (n_vid NVidia) CreateProduct(model string) GPU {
	gpu := GPU{model: model, opengl: true}
	return gpu
}

//func describeProductFromFactory(f *Factory) {
//product := f.CreateProduct()
//println(product.GetName())
//product = Product{}
//}

func main() {
	var intel Intel
	cpu := intel.CreateProduct("Core i7")
	println(cpu.GetName())

	var nvidia NVidia
	gpu := nvidia.CreateProduct("GForce")
	println(gpu.GetName())
}
