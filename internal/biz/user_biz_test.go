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
	"github.com/Xwudao/neter-template/internal/data/sqlc"
	"github.com/Xwudao/neter-template/internal/domain/params"
	"github.com/Xwudao/neter-template/internal/domain/payloads"
	"github.com/Xwudao/neter-template/internal/system"
	"github.com/Xwudao/neter-template/pkg/utils/bcrypt"
	"github.com/Xwudao/neter-template/pkg/utils/jwt"
)

func init() {
	// Use a lower bcrypt cost in tests for speed.
	bcrypt.Init(bcrypt.WithCost(4))
}

func newTestUserBiz(t *testing.T, ur biz.UserRepository) *biz.UserBiz {
	t.Helper()
	log := zap.NewNop().Sugar()
	appCtx := system.NewTestAppContext()
	jwtClient := jwt.NewClient(&payloads.JwtConfig{
		Secret: "test-secret-key-for-tests-only",
		Expire: time.Hour,
		Issuer: "test",
	})
	return biz.NewUserBiz(log, ur, jwtClient, appCtx)
}

// TestUserBiz_Login_Success verifies a successful login returns a user and non-empty token.
func TestUserBiz_Login_Success(t *testing.T) {
	ctrl := gomock.NewController(t)

	hashedPwd, err := bcrypt.GeneratePassword("password123")
	require.NoError(t, err)

	mockUser := &sqlc.User{ID: 1, Username: "admin", Password: hashedPwd}

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().
		GetBy(gomock.Any(), &params.GetUserByParams{Username: "admin"}).
		Return(mockUser, nil)

	b := newTestUserBiz(t, mockRepo)
	user, token, err := b.Login(context.Background(), &params.UserLoginParams{
		Username: "admin",
		Password: "password123",
	})

	require.NoError(t, err)
	assert.Equal(t, int64(1), user.ID)
	assert.NotEmpty(t, token)
}

// TestUserBiz_Login_UserNotFound verifies that a repo error propagates correctly.
func TestUserBiz_Login_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)

	repoErr := errors.New("user not found")

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().
		GetBy(gomock.Any(), gomock.Any()).
		Return(nil, repoErr)

	b := newTestUserBiz(t, mockRepo)
	_, _, err := b.Login(context.Background(), &params.UserLoginParams{
		Username: "nobody",
		Password: "anything",
	})

	assert.ErrorIs(t, err, repoErr)
}

// TestUserBiz_Login_WrongPassword verifies that an incorrect password returns an error.
func TestUserBiz_Login_WrongPassword(t *testing.T) {
	ctrl := gomock.NewController(t)

	hashedPwd, err := bcrypt.GeneratePassword("correctPassword")
	require.NoError(t, err)

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().
		GetBy(gomock.Any(), gomock.Any()).
		Return(&sqlc.User{ID: 1, Username: "admin", Password: hashedPwd}, nil)

	b := newTestUserBiz(t, mockRepo)
	_, _, err = b.Login(context.Background(), &params.UserLoginParams{
		Username: "admin",
		Password: "wrongPassword",
	})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "密码错误")
}

// TestUserBiz_Create verifies that CreateUserParams are forwarded and the user is returned.
func TestUserBiz_Create(t *testing.T) {
	ctrl := gomock.NewController(t)

	p := &params.CreateUserParams{Username: "newuser", Password: "secret"}
	want := &sqlc.User{ID: 2, Username: "newuser"}

	mockRepo := mocks.NewMockUserRepository(ctrl)
	// Password is hashed before forwarding, so use Any() for the repo call.
	mockRepo.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(want, nil)

	b := newTestUserBiz(t, mockRepo)
	got, err := b.Create(context.Background(), p)

	require.NoError(t, err)
	assert.Equal(t, want.ID, got.ID)
	assert.Equal(t, "newuser", got.Username)
}

// TestUserBiz_Create_RepoError verifies that a repo error is propagated.
func TestUserBiz_Create_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(nil, errors.New("duplicate username"))

	b := newTestUserBiz(t, mockRepo)
	_, err := b.Create(context.Background(), &params.CreateUserParams{
		Username: "dup",
		Password: "pass",
	})

	assert.Error(t, err)
}

// TestUserBiz_Get verifies GetByID delegates to the repository.
func TestUserBiz_Get(t *testing.T) {
	ctrl := gomock.NewController(t)

	want := &sqlc.User{ID: 5, Username: "alice"}

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().
		GetByID(gomock.Any(), int64(5)).
		Return(want, nil)

	b := newTestUserBiz(t, mockRepo)
	got, err := b.Get(context.Background(), 5)

	require.NoError(t, err)
	assert.Equal(t, want, got)
}

// TestUserBiz_Delete verifies DeleteByID delegates to the repository.
func TestUserBiz_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().
		DeleteByID(gomock.Any(), int64(3)).
		Return(nil)

	b := newTestUserBiz(t, mockRepo)
	err := b.Delete(context.Background(), 3)

	assert.NoError(t, err)
}

// TestUserBiz_GetBy verifies GetBy delegates to the repository.
func TestUserBiz_GetBy(t *testing.T) {
	ctrl := gomock.NewController(t)

	p := &params.GetUserByParams{Username: "bob"}
	want := &sqlc.User{ID: 7, Username: "bob"}

	mockRepo := mocks.NewMockUserRepository(ctrl)
	mockRepo.EXPECT().
		GetBy(gomock.Any(), p).
		Return(want, nil)

	b := newTestUserBiz(t, mockRepo)
	got, err := b.GetBy(context.Background(), p)

	require.NoError(t, err)
	assert.Equal(t, want, got)
}
