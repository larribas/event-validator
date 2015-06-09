package domain

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

type GroundFormatChecker struct{}

func (f *GroundFormatChecker) Check(rules []byte) (err error) {
    return
}