package manager

import (
	"gopkg.in/vinxi/vinxi.v0"
)

// JSONInstance represents the Instance entity for JSON serialization.
type JSONInstance struct {
	Info   *vinxi.Metadata `json:"info"`
	Scopes []JSONScope     `json:"scopes"`
}

func createInstance(instance *Instance) JSONInstance {
	return JSONInstance{
		Info:   instance.Metadata(),
		Scopes: createScopes(instance.Scopes()),
	}
}

func createInstances(instances []*Instance) []JSONInstance {
	list := []JSONInstance{}
	for _, instance := range instances {
		list = append(list, createInstance(instance))
	}
	return list
}

// InstancesController represents the rules entity HTTP controller.
type InstancesController struct{}

func (InstancesController) List(ctx *Context) {
	ctx.Send(createInstances(ctx.Manager.Instances()))
}

func (InstancesController) Get(ctx *Context) {
	ctx.Send(createInstance(ctx.Instance))
}

func (InstancesController) Delete(ctx *Context) {
	if ctx.Manager.RemoveInstance(ctx.Instance.ID()) {
		ctx.SendNoContent()
	} else {
		ctx.SendError(500, "Cannot remove instance")
	}
}