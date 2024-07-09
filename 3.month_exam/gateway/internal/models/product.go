package models

type ProductInfo struct {
    ID        string  `json:"id,omitempty"`
    Name      string  `json:"name"`
    Category  string  `json:"category"`
    Quantity  int32   `json:"quantity"`
    Price     float64 `json:"price"`
    CreatedAt string  `json:"created_at,omitempty"`
    ExpiredAt string  `json:"expired_at,omitempty"`
}

type CategoryRequest struct{
    Category string `json:"category"`     
}