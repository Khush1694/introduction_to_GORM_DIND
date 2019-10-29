package structs

// IDs request struct is used to send all the Ids in a request
// for example DeleteByIDs request at services/deleteByIDs.go
type IDs struct {
	IDs []string `validate:"required" json:"ids"`
}
