package models

import (
	"context"

	"github.com/google/uuid"
)

type ReqToken struct {
	RefreshToken string `json:"refresh_token"`
}

// ErrorResponse defines the structure of an error response
//
//	@description	This struct represents the response structure for error handling.
//	@example		{ "code": 500, "message": "Internal server error"}
type ErrorResponse struct {
	Message string `json:"message" example:"Internal server error" format:"string"` // Human-readable error message
	Code    int    `json:"code" example:"500" format:"int"`                         // HTTP status code
}

type GeneralService interface {
	ResetMedicineService(ctx context.Context) error
	ResetUserService(ctx context.Context) error
	ResetAddressService(ctx context.Context) error
	GenerateToken(ctx context.Context, refreshToken string) (TokenResponseDTO, error)
	RevokeRefreshToken(ctx context.Context, refreshToken string) error
}

type GeneralRepository interface {
	ResetMedicineRepo(ctx context.Context) error
	ResetUserRepo(ctx context.Context) error
	ResetAddressRepo(ctx context.Context) error
	RevokeToken(ctx context.Context, token string) error
	CreateRefreshToken(ctx context.Context, token string, id uuid.UUID) error
	FindUserFromToken(ctx context.Context, token string) (User, error)
}
