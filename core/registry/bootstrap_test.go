package registry_test

import (
	"github.com/go-chassis/go-archaius"
	"github.com/go-chassis/go-chassis/core/config"
	"github.com/go-chassis/go-chassis/core/config/model"
	"github.com/go-chassis/go-chassis/core/config/schema"
	"github.com/go-chassis/go-chassis/core/lager"
	"github.com/go-chassis/go-chassis/core/registry"
	"github.com/go-chassis/go-chassis/core/registry/mock"
	"github.com/go-chassis/go-chassis/pkg/runtime"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	lager.Init(&lager.Options{
		LoggerLevel: "INFO",
	})
	archaius.Init(archaius.WithMemorySource())
	archaius.Set("cse.service.registry.address", "http://127.0.0.1:30100")
	archaius.Set("service_description.name", "Client")
	runtime.HostName = "localhost"
	config.MicroserviceDefinition = &model.MicroserviceCfg{}
	archaius.UnmarshalConfig(config.MicroserviceDefinition)
	config.ReadGlobalConfigFromArchaius()
}

func TestRegisterService(t *testing.T) {
	runtime.Init()

	config.MicroserviceDefinition.ServiceDescription.Schemas = []string{"schemaId2", "schemaId3", "schemaId4"}

	testRegistryObj := new(mock.RegistratorMock)
	registry.DefaultRegistrator = testRegistryObj
	testRegistryObj.On("UnRegisterMicroServiceInstance", "microServiceID", "microServiceInstanceID").Return(nil)

	m := make(map[string]string, 0)
	m["id1"] = "schemaInfo1"
	m["id2"] = "schemaInfo2"
	m["id3"] = "schemaInfo3"

	// 	case schemaIDs is empty
	registry.RegisterService()
	registry.RegisterServiceInstances()
	err := schema.SetSchemaInfoByMap(m)
	assert.NoError(t, err)

	// 	case schemaIDs is empty
	registry.RegisterService()
	registry.RegisterServiceInstances()
	assert.NoError(t, err)

}
