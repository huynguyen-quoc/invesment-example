package rredisapi

// Args is a helper for constructing command arguments from structured values.
type Args []interface{}

// Add returns the result of appending value to args.
func (args *Args) Add(values ...interface{}) *Args {
	*args = append(*args, values...)
	return args
}

// NewArgs makes a new *Args
func NewArgs(args ...interface{}) *Args {
	result := &Args{}
	*result = Args(args)
	return result
}

// Value return Args type the pointer points to
func (args *Args) Value() Args {
	return *args
}