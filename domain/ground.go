package domain

/*
IMPORTANT NOTE:
The implementations described in this file have no real functionality.
Ground implementations are the minimal amount of logic capable of creating a make-believe environment in which most of
the domain's functionality can be executed.
*/

// GroundRepository is a nil implementation of the Repository interface
type GroundRepository struct{}

func (r *GroundRepository) Create(validator *Validator) (version int) {
	return
}

func (r *GroundRepository) GetNextVersion(_type string) (version int) {
	return
}

func (r *GroundRepository) Inspect(_type string, version int) (v *Validator, err error) {
	return
}

// GroundFormatChecker is a nil implementation of the FormatChecker interface
type GroundFormatChecker struct{}

func (f *GroundFormatChecker) Check(rules []byte) (err error) {
	return
}
