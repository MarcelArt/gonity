package gameobject

import (
	"github.com/MarcelArt/gonity/pkg/arrays"
)

type GameObject struct {
	Name       string
	Tags       []string
	Layers     []string
	Components []Component
}

func (g *GameObject) GetComponent(t string) *Component {
	return arrays.Find(g.Components, func(component Component) bool {
		return component.GetComponentName() == t
	})
}
