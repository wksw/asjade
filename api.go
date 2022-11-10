package asjade

import (
	"asjade/config"
	"asjade/server"
)

// Asjade .
type Asjade struct {
	// 协议列表
	schemas []*Schema
	// 是否已初始化
	inited bool
}

// Schema .
type Schema struct {
	// 协议名称
	name string
	// 协议处理方法
	structPtr interface{}
	options   []server.RegisteOption
}

var asjade *Asjade

func init() {
	asjade = &Asjade{}
}

// RegisteSchema .
// RegisteSchema("rest", &Hello{})
// RegisteSchema("rest", &Account{})
// RegisteSchema("rest", &Login{})
func RegisteSchema(schema string, structPtr interface{}, options ...server.RegisteOption) error {
	return asjade.registeSchema(schema, structPtr, options...)
}

// Init system init
func Init() error {
	// config init
	if err := config.Init(); err != nil {
		return err
	}
	// server init
	if err := server.Init(); err != nil {
		return err
	}
	asjade.inited = true
	return nil
}

// Start  start servers
func Start(options ...server.RunOption) error {
	return asjade.start(options...)
}

func (aj *Asjade) registeSchema(schemaName string, structPtr interface{}, options ...server.RegisteOption) error {
	aj.schemas = append(aj.schemas, &Schema{
		name:      schemaName,
		structPtr: structPtr,
		options:   options,
	})
	return nil
}

func (aj *Asjade) start(options ...server.RunOption) error {
	// system init
	if !aj.inited {
		if err := Init(); err != nil {
			return err
		}
	}
	for _, schema := range aj.schemas {
		s, err := server.GetServerByName(schema.name)
		if err != nil {
			return err
		}
		if err := s.Registe(schema.structPtr, schema.options...); err != nil {
			return err
		}
	}
	return server.StartServers(options...)
}
