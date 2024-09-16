package openapi

import (
	"context"
	"net/http"
)

type APIRouter interface {
	CheckServer(http.ResponseWriter, *http.Request)
	CreateBid(http.ResponseWriter, *http.Request)
	CreateTender(http.ResponseWriter, *http.Request)
	EditBid(http.ResponseWriter, *http.Request)
	EditTender(http.ResponseWriter, *http.Request)
	GetBidReviews(http.ResponseWriter, *http.Request)
	GetBidStatus(http.ResponseWriter, *http.Request)
	GetBidsForTender(http.ResponseWriter, *http.Request)
	GetTenderStatus(http.ResponseWriter, *http.Request)
	GetTenders(http.ResponseWriter, *http.Request)
	GetUserBids(http.ResponseWriter, *http.Request)
	GetUserTenders(http.ResponseWriter, *http.Request)
	RollbackBid(http.ResponseWriter, *http.Request)
	RollbackTender(http.ResponseWriter, *http.Request)
	SubmitBidDecision(http.ResponseWriter, *http.Request)
	SubmitBidFeedback(http.ResponseWriter, *http.Request)
	UpdateBidStatus(http.ResponseWriter, *http.Request)
	UpdateTenderStatus(http.ResponseWriter, *http.Request)
}

type APIServicer interface {
	CheckServer(context.Context) (ImplResponse, error)
	CreateBid(context.Context, CreateBidRequest) (ImplResponse, error)
	CreateTender(context.Context, CreateTenderRequest) (ImplResponse, error)
	EditBid(context.Context, string, string, EditBidRequest) (ImplResponse, error)
	EditTender(context.Context, string, string, EditTenderRequest) (ImplResponse, error)
	GetBidReviews(context.Context, string, string, string, int32, int32) (ImplResponse, error)
	GetBidStatus(context.Context, string, string) (ImplResponse, error)
	GetBidsForTender(context.Context, string, string, int32, int32) (ImplResponse, error)
	GetTenderStatus(context.Context, string, string) (ImplResponse, error)
	GetTenders(context.Context, int32, int32, []TenderServiceType) (ImplResponse, error)
	GetUserBids(context.Context, int32, int32, string) (ImplResponse, error)
	GetUserTenders(context.Context, int32, int32, string) (ImplResponse, error)
	RollbackBid(context.Context, string, int32, string) (ImplResponse, error)
	RollbackTender(context.Context, string, int32, string) (ImplResponse, error)
	SubmitBidDecision(context.Context, string, BidDecision, string) (ImplResponse, error)
	SubmitBidFeedback(context.Context, string, string, string) (ImplResponse, error)
	UpdateBidStatus(context.Context, string, BidStatus, string) (ImplResponse, error)
	UpdateTenderStatus(context.Context, string, TenderStatus, string) (ImplResponse, error)
}
