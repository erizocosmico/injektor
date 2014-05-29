package injektor

// Injector acts as a container of dependencies and can inject them to Injectable objects.
type Injector interface {
	Get(string) interface{}
	Extract(string) interface{}
	Set(string, interface{})
	Remove(string)
	Clear()
	Inject(Injectable)
}

// An Injectable object can be passed to a Injector's Inject method in order to get its dependencies injected.
type Injectable interface {
	SetDependencies(Injector)
}

type injector struct {
	bag map[string]interface{}
}

var sharedInjector *injector

// Creates a new injector instance.
func NewInjector() Injector {
	return &injector{make(map[string]interface{})}
}

// GetInjector returns the shared injector instance.
func GetInjector() Injector {
	if sharedInjector == nil {
		sharedInjector = &injector{make(map[string]interface{})}
	}

	return sharedInjector
}

// Get retrieves and returns an item from the dependency bag
// assigned to the given key.
func (i *injector) Get(key string) interface{} {
	if v, ok := i.bag[key]; ok {
		return v
	}

	return nil
}

// Extract returns an item from the dependency bag
// and removes it afterwards.
func (i *injector) Extract(key string) interface{} {
	v := i.Get(key)
	i.Remove(key)
	return v
}

// Set adds an item to the dependency bag mapped to a given key
func (i *injector) Set(key string, item interface{}) {
	i.bag[key] = item
}

// Remove removes an item from the dependency bag.
func (i *injector) Remove(key string) {
	delete(i.bag, key)
}

// Clear removes all items from the dependency bag.
func (i *injector) Clear() {
	i.bag = make(map[string]interface{})
}

// Inject injects dependencies to an injectable object.
func (i *injector) Inject(in Injectable) {
	in.SetDependencies(i)
}
