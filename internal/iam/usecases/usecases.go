package usecases

type usecases struct {
}

func NewUsecases() UsecasesInterface {
	return &usecases{}
}

func (u *usecases) Create() {
	panic("not implemented") // TODO: Implement
}

func (u *usecases) Get() {
	panic("not implemented") // TODO: Implement
}

func (u *usecases) Update() {
	panic("not implemented") // TODO: Implement
}

func (u *usecases) Delete() {
	panic("not implemented") // TODO: Implement
}
