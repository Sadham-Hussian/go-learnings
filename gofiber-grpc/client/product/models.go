package product

// ProductItem struct to model ProductItem
type ProductItem struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Prize int32  `json:"prize"`
}
