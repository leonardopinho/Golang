package auction_test

import (
	"context"
	"fullcycle-auction_go/internal/infra/database/auction"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type AuctionRepositoryMock struct {
	mock.Mock
}

func (m *AuctionRepositoryMock) CloseAuction(ctx context.Context, auctionId string) error {
	args := m.Called(ctx, auctionId)
	return args.Error(0)
}

func TestCloseAuction(t *testing.T) {
	repository := &AuctionRepositoryMock{}
	ctx := context.Background()
	repository.On("CloseAuction", ctx, "123").Return(nil)

	closeTime := time.Second + 5
	go auction.CloseAuctionWatcher(ctx, repository, "123", closeTime)
	time.Sleep(time.Millisecond * 100)
	repository.AssertNumberOfCalls(t, "CloseAuction", 0)

	time.Sleep(time.Millisecond * 1900)
	repository.AssertNumberOfCalls(t, "CloseAuction", 1)
	repository.AssertExpectations(t)
}
