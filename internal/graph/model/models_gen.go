// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type CreateProduct struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ThumbnailID string  `json:"thumbnailID"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Object struct {
	ID         string `json:"id"`
	FileName   string `json:"fileName"`
	Type       string `json:"type"`
	SignedURL  string `json:"signedUrl"`
	ExpiredAt  string `json:"expiredAt"`
	IsPublic   bool   `json:"isPublic"`
	UploadedBy string `json:"uploadedBy"`
	CreatedAt  string `json:"createdAt"`
}

type PaginationMeta struct {
	Search  string   `json:"search"`
	Sort    []string `json:"sort"`
	Limit   int      `json:"limit"`
	Page    int      `json:"page"`
	Count   int      `json:"count"`
	MaxPage int      `json:"maxPage"`
}

type PaginationRequest struct {
	Search         string   `json:"search"`
	Sort           []string `json:"sort"`
	Limit          int      `json:"limit"`
	Page           int      `json:"page"`
	IncludeDeleted bool     `json:"includeDeleted"`
}

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ThumbnailID string  `json:"thumbnailID"`
	Thumbnail   *Object `json:"thumbnail"`
	OwnerID     string  `json:"ownerID"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
	DeletedAt   string  `json:"deletedAt"`
}

type ProductPaginationResponse struct {
	Meta  *PaginationMeta `json:"meta"`
	Items []*Product      `json:"items"`
}

type Register struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateProduct struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ThumbnailID string  `json:"thumbnailID"`
}

type User struct {
	ID        string `json:"id"`
	FullName  string `json:"fullName"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
