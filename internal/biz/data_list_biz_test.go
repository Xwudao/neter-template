package biz_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/biz/mocks"
	"github.com/Xwudao/neter-template/internal/data/ent"
	"github.com/Xwudao/neter-template/internal/domain/params"
	"github.com/Xwudao/neter-template/internal/system"
)

func newTestDataListBiz(t *testing.T, dlr biz.DataListRepository) *biz.DataListBiz {
	t.Helper()
	log := zap.NewNop().Sugar()
	appCtx := system.NewTestAppContext()
	return biz.NewDataListBiz(log, dlr, appCtx)
}

func newDataList(id int64, kind, key, value string) *ent.DataList {
	return &ent.DataList{
		ID:         id,
		Kind:       kind,
		Key:        key,
		Value:      value,
		Label:      "label-" + key,
		ItemOrder:  1,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
}

// TestDataListBiz_Create_Success verifies a normal create call sets ItemOrder to 1 and delegates.
func TestDataListBiz_Create_Success(t *testing.T) {
	ctrl := gomock.NewController(t)

	p := &params.CreateDataListParams{
		Label: "友情链接",
		Key:   "link-1",
		Kind:  "friend_link",
		Value: `{"name":"Example","link":"https://example.com"}`,
		// ItemOrder intentionally left as 0; biz should set it to 1.
	}
	want := newDataList(1, "friend_link", "link-1", p.Value)

	mockRepo := mocks.NewMockDataListRepository(ctrl)
	mockRepo.EXPECT().
		Create(gomock.Any(), gomock.AssignableToTypeOf(&params.CreateDataListParams{})).
		DoAndReturn(func(_ context.Context, got *params.CreateDataListParams) (*ent.DataList, error) {
			assert.Equal(t, 1, got.ItemOrder, "Optimize() should set ItemOrder to 1")
			return want, nil
		})

	b := newTestDataListBiz(t, mockRepo)
	got, err := b.Create(context.Background(), p)

	require.NoError(t, err)
	assert.Equal(t, want.ID, got.ID)
}

// TestDataListBiz_Create_RepoError verifies that a repo error propagates.
func TestDataListBiz_Create_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockDataListRepository(ctrl)
	mockRepo.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(nil, errors.New("duplicate key"))

	b := newTestDataListBiz(t, mockRepo)
	_, err := b.Create(context.Background(), &params.CreateDataListParams{
		Label: "l", Key: "k", Kind: "kind", Value: "v",
	})

	assert.Error(t, err)
}

// TestDataListBiz_Get verifies GetByID delegates to the repository.
func TestDataListBiz_Get(t *testing.T) {
	ctrl := gomock.NewController(t)

	want := newDataList(42, "kind", "k", "v")

	mockRepo := mocks.NewMockDataListRepository(ctrl)
	mockRepo.EXPECT().
		GetByID(gomock.Any(), int64(42)).
		Return(want, nil)

	b := newTestDataListBiz(t, mockRepo)
	got, err := b.Get(context.Background(), 42)

	require.NoError(t, err)
	assert.Equal(t, want, got)
}

// TestDataListBiz_Delete verifies DeleteByID delegates to the repository.
func TestDataListBiz_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockDataListRepository(ctrl)
	mockRepo.EXPECT().
		DeleteByID(gomock.Any(), int64(9)).
		Return(nil)

	b := newTestDataListBiz(t, mockRepo)
	err := b.Delete(context.Background(), 9)
	assert.NoError(t, err)
}

// TestDataListBiz_Update verifies Update delegates to the repository.
func TestDataListBiz_Update(t *testing.T) {
	ctrl := gomock.NewController(t)

	order := 2
	p := &params.UpdateDataListParams{ID: 1, Key: "k", Value: "new-value", ItemOrder: &order}
	want := newDataList(1, "kind", "k", "new-value")

	mockRepo := mocks.NewMockDataListRepository(ctrl)
	mockRepo.EXPECT().
		Update(gomock.Any(), p).
		Return(want, nil)

	b := newTestDataListBiz(t, mockRepo)
	got, err := b.Update(context.Background(), p)

	require.NoError(t, err)
	assert.Equal(t, "new-value", got.Value)
}

// TestDataListBiz_GetAll verifies GetAll delegates to the repository.
func TestDataListBiz_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)

	items := []*ent.DataList{
		newDataList(1, "a", "k1", "v1"),
		newDataList(2, "b", "k2", "v2"),
	}

	mockRepo := mocks.NewMockDataListRepository(ctrl)
	mockRepo.EXPECT().
		GetAll(gomock.Any()).
		Return(items, nil)

	b := newTestDataListBiz(t, mockRepo)
	got, err := b.GetAll(context.Background())

	require.NoError(t, err)
	assert.Len(t, got, 2)
}

// TestDataListBiz_ListByKind verifies ListByKind sets correct offset and delegates.
func TestDataListBiz_ListByKind(t *testing.T) {
	ctrl := gomock.NewController(t)

	p := &params.ListDataByKindParams{Kind: "friend_link", Page: 2, Size: 10}
	wantItems := []*ent.DataList{newDataList(1, "friend_link", "k", "v")}

	mockRepo := mocks.NewMockDataListRepository(ctrl)
	mockRepo.EXPECT().
		ListByKind(gomock.Any(), gomock.AssignableToTypeOf(&params.ListDataByKindParams{})).
		DoAndReturn(func(_ context.Context, got *params.ListDataByKindParams) ([]*ent.DataList, int, error) {
			assert.Equal(t, 10, got.Offset, "Optimize() should compute Offset = (Page-1)*Size")
			return wantItems, 1, nil
		})

	b := newTestDataListBiz(t, mockRepo)
	items, total, err := b.ListByKind(context.Background(), p)

	require.NoError(t, err)
	assert.Equal(t, 1, total)
	assert.Len(t, items, 1)
}

// TestDataListBiz_GetAllByKinds verifies GetAllByKinds delegates to the repository.
func TestDataListBiz_GetAllByKinds(t *testing.T) {
	ctrl := gomock.NewController(t)

	p := &params.GetAllDataListByKindsParams{Kinds: []string{"friend_link", "nav"}}
	items := []*ent.DataList{
		newDataList(1, "friend_link", "k1", "v1"),
		newDataList(2, "nav", "k2", "v2"),
	}

	mockRepo := mocks.NewMockDataListRepository(ctrl)
	mockRepo.EXPECT().
		GetAllByKinds(gomock.Any(), p).
		Return(items, nil)

	b := newTestDataListBiz(t, mockRepo)
	got, err := b.GetAllByKinds(context.Background(), p)

	require.NoError(t, err)
	assert.Len(t, got, 2)
}

// TestDataListBiz_GetSortData verifies GetSortData delegates to the repository.
func TestDataListBiz_GetSortData(t *testing.T) {
	ctrl := gomock.NewController(t)

	p := &params.GetDataListSortDataParams{Kind: "friend_link"}
	items := []*ent.DataList{
		newDataList(1, "friend_link", "k1", "v1"),
		newDataList(2, "friend_link", "k2", "v2"),
	}

	mockRepo := mocks.NewMockDataListRepository(ctrl)
	mockRepo.EXPECT().
		GetSortData(gomock.Any(), p).
		Return(items, nil)

	b := newTestDataListBiz(t, mockRepo)
	got, err := b.GetSortData(context.Background(), p)

	require.NoError(t, err)
	assert.Len(t, got, 2)
}

// TestDataListBiz_UpdateOrder_Success verifies that UpdateOrder delegates when IDs and Orders match.
func TestDataListBiz_UpdateOrder_Success(t *testing.T) {
	ctrl := gomock.NewController(t)

	p := &params.ItemOrderParams{
		IDs:    []int64{1, 2, 3},
		Orders: []int{10, 20, 30},
	}

	mockRepo := mocks.NewMockDataListRepository(ctrl)
	mockRepo.EXPECT().
		UpdateOrder(gomock.Any(), p).
		Return(nil)

	b := newTestDataListBiz(t, mockRepo)
	err := b.UpdateOrder(context.Background(), p)
	assert.NoError(t, err)
}

// TestDataListBiz_UpdateOrder_Mismatch verifies that UpdateOrder rejects mismatched IDs/Orders.
func TestDataListBiz_UpdateOrder_Mismatch(t *testing.T) {
	ctrl := gomock.NewController(t)

	p := &params.ItemOrderParams{
		IDs:    []int64{1, 2},
		Orders: []int{10}, // length mismatch
	}

	// Repo must NOT be called when validation fails.
	mockRepo := mocks.NewMockDataListRepository(ctrl)

	b := newTestDataListBiz(t, mockRepo)
	err := b.UpdateOrder(context.Background(), p)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "不一致")
}
