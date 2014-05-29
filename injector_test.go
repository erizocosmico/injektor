package injektor

import "testing"

type Pens struct {
	Number int
	Colors []string
}

type Pencils struct {
	Number int
}

type PencilCase struct {
	pencils *Pencils
	pens    *Pens
}

func NewPencilCase() *PencilCase {
	p := &PencilCase{}
	GetInjector().Inject(p)

	return p
}

func (p *PencilCase) SetDependencies(in Injector) {
	var pencils, pens interface{}

	if pencils = in.Get("pencils"); pencils != nil {
		p.pencils = pencils.(*Pencils)
	}

	if pens = in.Get("pens"); pens != nil {
		p.pens = pens.(*Pens)
	}
}

func TestSet(t *testing.T) {
	s := "hello"
	i := GetInjector()
	i.Set("a", 1)
	i.Set("b", 2)
	i.Set("c", 3)
	i.Set("d", &s)

	j := 0
	for _, _ = range sharedInjector.bag {
		j++
	}

	if j != 4 {
		t.Errorf("expected 4 items in the dependency bag, %d found", j)
	}
}

func TestGet(t *testing.T) {
	i := GetInjector()
	a := i.Get("a").(int)
	b := i.Get("b").(int)
	c := i.Get("c").(int)
	d := i.Get("d").(*string)

	if a != 1 {
		t.Errorf("expected item a to be 1, %d found", a)
	}

	if b != 2 {
		t.Errorf("expected item b to be 2, %d found", a)
	}

	if c != 3 {
		t.Errorf("expected item c to be 3, %d found", a)
	}

	if *d != "hello" {
		t.Errorf("expected item a to be 'hello', %s found", d)
	}
}

func TestExtract(t *testing.T) {
	i := GetInjector()
	a := i.Extract("a").(int)
	aNil := i.Get("a")

	if a != 1 {
		t.Errorf("expected item a to be 1, %d found", a)
	}

	if aNil != nil {
		t.Error("expected item a to be nil")
	}
}

func TestRemove(t *testing.T) {
	i := GetInjector()
	i.Remove("b")
	b := i.Get("b")

	if b != nil {
		t.Error("expected item b to be nil")
	}
}

func TestClear(t *testing.T) {
	i := GetInjector()
	i.Clear()

	j := 0
	for _, _ = range sharedInjector.bag {
		j++
	}

	if j != 0 {
		t.Errorf("expected 0 items in the dependency bag, %d found", j)
	}
}

func TestGetInjector(t *testing.T) {
	if sharedInjector != GetInjector() {
		t.Error("injectors don't match")
	}
}

func TestInject(t *testing.T) {
	pencils := &Pencils{4}
	pens := &Pens{3, []string{"red", "blue", "black"}}

	i := GetInjector()
	i.Set("pencils", pencils)
	i.Set("pens", pens)

	p := NewPencilCase()

	if p.pencils == nil {
		t.Error("nil object received from the injector, *Pencils expected")
	} else if p.pencils.Number != 4 {
		t.Errorf("expecting number of pencils to be 4, %d found", p.pencils.Number)
	}

	if p.pens == nil {
		t.Error("nil object received from the injector, *Pens expected")
	} else if p.pens.Number != 3 {
		t.Errorf("expecting number of pens to be 3, %d found", p.pens.Number)
	} else if p.pens.Colors[1] != "blue" {
		t.Errorf("expecting color of second pen to be blue, %s found", p.pens.Colors[1])
	}
}
