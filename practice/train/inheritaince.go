package main

import (
	"fmt"
)

type Product interface {
	getProduct() string
}

type CPU struct {
	brand string
	model string
}

type Intel struct {
	CPU
}

func NewIntel(model string) *Intel {
	return &Intel{
		CPU{
			brand: "Intel",
			model: model,
		},
	}
}

func (i *Intel) getProduct() string {
	return i.brand + " " + i.model
}

type GPU struct {
	brand  string
	model  string
	opengl bool
}

type Nvidia struct {
	GPU
}

func NewNvidia(model string, opengl bool) *Nvidia {
	return &Nvidia{
		GPU{
			brand:  "Nvidia",
			model:  model,
			opengl: opengl,
		},
	}
}

func (n *Nvidia) getProduct() string {
	if n.opengl {
		return n.brand + " " + n.model + " supports OpenGL"
	} else {
		return n.brand + " " + n.model
	}
}

func main() {
	products := []Product{
		NewIntel("Core i7"),
		NewNvidia("GeForce X720", true),
		NewNvidia("GeForce X640", false),
	}

	for _, v := range products {
		fmt.Println(v.getProduct())
	}
}