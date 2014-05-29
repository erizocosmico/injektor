# Injektor

Injektor is a tiny implementation of dependency injection for the Go programming language.

[GoDoc documentation](http://godoc.org/github.com/mvader/injektor).
MIT Licensed.

## Shared injector

Injektor has a shared dependency container (injector) available for all packages. It can be accessed using the ```GetInjector()``` function. You can also create new injector instances if you need to with ```NewInjector()```.

The shared injector will not be instantiated until the first time it is requested so if you don't need it it won't be instantiated.

## Designing injectable types

For a type to be Injectable you need to add a ```SetDependencies(injektor.Injector)``` method. Then in the constructor you call the injector's method ```Inject(injektor.Injectable)``` and the dependencies will injected.

*Example:*

```go
import "github.com/mvader/injektor"

type Pens struct {
	Number int
	Colors []string
}

type Pencils struct {
	Number int
}

// This will be our Injectable struct
type PencilCase struct {
	pencils *Pencils
	pens    *Pens
}

// In the constructor we call Inject
func NewPencilCase() *PencilCase {
	p := &PencilCase{}
	injektor.GetInjector().Inject(p)

	return p
}

// The SetDependencies method receives the injector and we setup the dependencies
func (p *PencilCase) SetDependencies(in injektor.Injector) {
	var pencils, pens interface{}

	if pencils = in.Get("pencils"); pencils != nil {
		p.pencils = pencils.(*Pencils)
	}

	if pens = in.Get("pens"); pens != nil {
		p.pens = pens.(*Pens)
	}
}
```
