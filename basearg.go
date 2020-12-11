package wzmodlib

// BaseArg is a base object for arguments. All objects should inherit it.
type BaseArg struct {
	Debug string
}

// Validate the input
func (arg *BaseArg) Validate() error {
	if arg.Debug != "yes" {
		arg.Debug = "no"
	}
	return nil
}
