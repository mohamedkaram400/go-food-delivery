package requests



type CreateAddressesRequest struct {
    Addresses     []string `json:"name" validate:"required,min=2"`
}
