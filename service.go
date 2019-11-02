package kira

// Service ...
type Service interface {
	Service() ServiceInstance
}

// ServiceInstance ...
type ServiceInstance struct {
	Name string
	New  func() Service
}

func (si ServiceInstance) String() string { return si.Name }

// TODO: Add other methods like namespace...
