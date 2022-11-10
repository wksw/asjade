package server

// RunOptions run options
type RunOptions struct{}

// RegisteOptions registe options
type RegisteOptions struct{}

// RegisteOption .
type RegisteOption func(*RegisteOptions)

// RunOption .
type RunOption func(*RunOptions)

// Option .
type Option func(*Options)

// Options start options
type Options struct{}
