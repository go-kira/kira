package kira

import "fmt"

// TODO: Later we can do some stuff with the registered services.
//       For example: Hooks, init methods...

// RegisterService ...
func (app *App) RegisterService(instance Service) error {
	service := instance.Service()
	if service.Name == "" {
		return fmt.Errorf("missing ServiceInstance.Name")
	}
	if service.New == nil {
		return fmt.Errorf("missing ServiceInstance.New")
	}
	if val := service.New(); val == nil {
		return fmt.Errorf("ServiceInstance.New must return a non-nil service instance")
	}

	app.mutex.Lock()
	defer app.mutex.Unlock()

	if _, ok := app.Services[service.Name]; ok {
		return fmt.Errorf("service already registered: %s", service.Name)
	}
	app.Services[service.Name] = service

	return nil
}

// GetService returns service information from its full name.
func (app *App) GetService(name string) (ServiceInstance, error) {
	app.mutex.Lock()
	defer app.mutex.Unlock()
	m, ok := app.Services[name]
	if !ok {
		return ServiceInstance{}, fmt.Errorf("service not registered: %s", name)
	}
	return m, nil
}
