package main

type Product interface{}

func (p Product) GetName() string {
	
}

type Factory interface{}

func (f *Factory) CreateProduct() Product {}

type CPU struct {
	p     Product
	model string
}

//func (c CPU) CPU(return func model(model){}){}
func (c CPU) GetName() string {
	return "CPU with model: " + c.model
}

type Intel struct {
	f Factory
}

func (intl *Intel) CreateProduct() Product {
	return CPU("Core i7")
}

type GPU struct {
	p      Product
	model  string
	opengl bool
}

//func (g GPU) GPU(model string, supportOpenGL bool)
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
	return
}

func describeProductFromFactory(f *Factory) {
	p := f.CreateProduct()
	println(p.GetName())
	p = nil
}
