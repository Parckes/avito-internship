package openapi

import (
	"context"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
	"time"
)

// DefaultAPIService is a service that implements the logic for the DefaultAPIServicer
// This service should implement the business logic for every endpoint for the DefaultAPI API.
// Include any external packages or services that will be required by this service.
type DefaultAPIService struct {
	pg      *Postgres
	log     *slog.Logger
	builder squirrel.StatementBuilderType
}

// NewDefaultAPIService creates a default api service
func NewDefaultAPIService(pg *Postgres, log *slog.Logger) *DefaultAPIService {
	return &DefaultAPIService{
		pg:      pg,
		log:     log,
		builder: pg.Builder,
	}
}

// CheckServer - Проверка доступности сервера (good)
func (s *DefaultAPIService) CheckServer(ctx context.Context) (ImplResponse, error) {
	if err := s.pg.Pool.Ping(ctx); err != nil {
		return Response(http.StatusInternalServerError, nil), nil
	}
	return Response(http.StatusOK, "ok"), nil
}

// CreateBid - Создание нового предложения (good)
func (s *DefaultAPIService) CreateBid(ctx context.Context, createBidRequest CreateBidRequest) (ImplResponse, error) {
	const op = "CreateBid"
	log := s.log.With(slog.String("op", op))

	// Проверяем, существует ли тендер с таким ID
	tenderId, _ := s.ConvertIntoUUID(createBidRequest.TenderId)
	authorId, _ := s.ConvertIntoUUID(createBidRequest.AuthorId)

	_, err := s.getTenderById(ctx, tenderId)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return Response(http.StatusNotFound, ErrorResponse{Reason: "Тендер не найден"}), nil
		}
		return Response(http.StatusInternalServerError, ErrorResponse{Reason: err.Error()}), nil
	}

	if createBidRequest.AuthorType == USER {
		_, err := s.getUserById(ctx, authorId)
		if err != nil {
			if errors.Is(err, ErrNoUser) {
				return Response(http.StatusUnauthorized, ErrorResponse{Reason: "Пользователь не существует или некорректен"}), nil
			}
			return Response(http.StatusInternalServerError, ErrorResponse{Reason: err.Error()}), err
		}
		err1, err2 := s.userHasRights(ctx, authorId, tenderId)
		if err1 != nil {
			if errors.Is(err1, ErrUserNoRightsTender) {
				return Response(http.StatusForbidden, ErrorResponse{Reason: "Недостаточно прав для выполнения действия"}), nil
			}
			return Response(http.StatusInternalServerError, ErrorResponse{Reason: err2.Error()}), err2
		}

	} else if createBidRequest.AuthorType == ORGANIZATION {
		_, err := s.getOrganizationById(ctx, authorId)
		if err != nil {
			if errors.Is(err, ErrNoOrganization) {
				return Response(http.StatusUnauthorized, ErrorResponse{Reason: "Организация не существует или некорректна"}), nil
			}
			return Response(http.StatusInternalServerError, ErrorResponse{Reason: err.Error()}), err
		}
		if err := s.organizationHasRights(ctx, authorId, tenderId); err != nil {
			if errors.Is(err, ErrOrgNoRightsTender) {
				return Response(http.StatusForbidden, ErrorResponse{Reason: "Недостаточно прав для выполнения действия"}), nil
			}
			return Response(http.StatusInternalServerError, ErrorResponse{Reason: err.Error()}), err
		}

	}

	currentTime := time.Now()
	rfc3339Time := currentTime.Format(time.RFC3339)

	sql, args, err := s.builder.
		Insert("bids").
		Columns("name", "description", "status", "tender_id", "author_type", "author_id", "created_at").
		Values(createBidRequest.Name, createBidRequest.Description, CREATED, createBidRequest.TenderId, createBidRequest.AuthorType, createBidRequest.AuthorId, rfc3339Time).
		Suffix("RETURNING bid_id, created_at").
		ToSql()

	if err != nil {
		log.Error("SQL generation failed", slog.Any("error", err))
		return Response(http.StatusInternalServerError, ErrorResponse{Reason: "Internal server error"}), nil
	}

	var newBidID uuid.UUID
	var createdAt time.Time
	err = s.pg.Pool.QueryRow(ctx, sql, args...).Scan(&newBidID, &createdAt)
	if err != nil {
		log.Error("Database execution failed", slog.Any("error", err))
		return Response(http.StatusInternalServerError, ErrorResponse{Reason: "Internal server error"}), nil
	}

	bidResponse:= Bid{
		Id:          s.ConvertFromUUID(newBidID),
		Name:        createBidRequest.Name,
		Status:      CREATED_BID,
		AuthorType:  createBidRequest.AuthorType,
		AuthorId:    createBidRequest.AuthorId,
		Version:     1,
		CreatedAt:   rfc3339Time,
	}

	return Response(http.StatusOK, bidResponse), nil
}