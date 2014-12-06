package injektor

import "sync"

// Injector acts as a container of dependencies and can inject them to Injectable objects.
type Injector interface {
	// Get retrieves and returns an item from the dependency bag assigned to the given key.
	Get(string) interface{}
	// Extract returns an item from the dependency bag and removes it afterwards.
	Extract(string) interface{}
	// Set adds an item to the dependency bag mapped to a given key.
	Set(string, interface{})
	// Remove removes an item from the dependency bag.
	Remove(string)
	// Clear removes all items from the dependency bag.
	Clear()
	// Inject injects dependencies to an injectable object.
	Inject(Injectable)
}

// An Injectable object can be passed to a Injector's Inject method in order to get its dependencies injected.
type Injectable interface {
	SetDependencies(Injector)
}

type injector struct {
	sync.RWMutex
	bag map[string]interface{}
}

var sharedInjector Injector

// Creates a new injector instance.
func NewInjector() Injector {
	return &injector{bag: make(map[string]interface{})}
}

// GetInjector returns the shared injector instance.
func GetInjector() Injector {
	if sharedInjector == nil {
		sharedInjector = NewInjector()
	}

	return sharedInjector
}

// Inject is a shortcut for injecting dependencies to types when using the shared injector.
func Inject(in Injectable) {
	GetInjector().Inject(in)
}

func (i *injector) Get(key string) interface{} {
	i.RLock()
	defer i.RUnlock()
	if v, ok := i.bag[key]; ok {
		return v
	}

	return nil
}

func (i *injector) Extract(key string) interface{} {
	v := i.Get(key)
	i.Remove(key)
	return v
}

func (i *injector) Set(key string, item interface{}) {
	i.Lock()
	i.bag[key] = item
	i.Unlock()
}

func (i *injector) Remove(key string) {
	i.Lock()
	delete(i.bag, key)
	i.Unlock()
}

func (i *injector) Clear() {
	i.Lock()
	i.bag = make(map[string]interface{})
	i.Unlock()
}

func (i *injector) Inject(in Injectable) {
	in.SetDependencies(i)
}
