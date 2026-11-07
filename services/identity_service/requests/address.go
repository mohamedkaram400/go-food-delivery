package requests



type CreateAddressesRequest struct {
    Addresses []string `json:"addresses" validate:"required,min=1,max=5,dive,required"`
}
