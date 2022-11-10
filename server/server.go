package server

// Server .
type Server interface {
	Registe(structPtr interface{}, options ...RegisteOption) error
	Start(options ...RunOption) error
	Stop() error
}
