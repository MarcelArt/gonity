package gameobject

type Component interface {
	Start()
	Update()
	GetComponentName() string
}
