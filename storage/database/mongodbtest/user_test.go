package mongodbtest

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/ppzxc/chattools/storage/database/model"
	"github.com/ppzxc/chattools/types"
	"github.com/ppzxc/chattools/utils/mono"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
	"time"
)

func userCreate(t *testing.T, ctx context.Context, count int) {
	mono.Init()
	for i := 0; i < count; i++ {
		now := time.Now()
		expire := time.Now().Add(365 * 24 * time.Hour)

		user := model.User{
			State:     types.StateUserActive,
			StatedAt:  &now,
			CreatedAt: &now,
			UpdatedAt: &now,

			Authentication: &model.Authentication{
				UserName:  mono.GetMONOID(),
				Email:     mono.GetMONOID() + "@email.co.kr",
				Password:  mono.GetMONOID(),
				AuthType:  types.StateAuthTypeId,
				AuthLevel: types.StateAuthLevelUser,
				Secret:    mono.GetMONOID(),
				Expires:   &expire,
				CreatedAt: &now,
				UpdatedAt: &now,
			},
			Device: []*model.Device{
				{
					DeviceId:        "1",
					BrowserId:       "2",
					UserAgent:       "3",
					OperationSystem: "4",
					Platform:        "5",
					CreatedAt:       &now,
					UpdatedAt:       &now,
				},
			},
			Profile: &model.Profile{
				Description: "create " + string(rune(i)),
				CreatedAt:   &now,
				UpdatedAt:   &now,
			},
		}

		queryCtx, cancel := context.WithCancel(ctx)
		userId, err := service.UserInsert(queryCtx, user)
		cancel()
		require.NoError(t, err)
		assert.Greater(t, userId, int64(0))
		user.Id = userId
	}
}

func timeEquals(t *testing.T, time1 *time.Time, time2 *time.Time) {
	assert.Equal(t, time1.Year(), time2.Year())
	assert.Equal(t, time1.Month(), time2.Month())
	assert.Equal(t, time1.Day(), time2.Day())
	assert.Equal(t, time1.Hour(), time2.Hour())
	assert.Equal(t, time1.Minute(), time2.Minute())
	assert.Equal(t, time1.Second(), time2.Second())
	assert.Equal(t, time1.Location(), time2.Location())
}

func userEquals(t *testing.T, user1 *model.User, user2 *model.User) {
	assert.Equal(t, user1.Id, user2.Id)
	assert.Equal(t, user1.State, user2.State)
	timeEquals(t, user1.StatedAt, user2.StatedAt)
	timeEquals(t, user1.CreatedAt, user2.CreatedAt)
	timeEquals(t, user1.UpdatedAt, user2.UpdatedAt)

	assert.Equal(t, user1.Authentication.Id, user2.Authentication.Id)
	assert.Equal(t, user1.Authentication.UserId, user2.Authentication.UserId)
	assert.Equal(t, user1.Authentication.UserName, user2.Authentication.UserName)
	assert.Equal(t, user1.Authentication.Email, user2.Authentication.Email)
	assert.Equal(t, user1.Authentication.Password, user2.Authentication.Password)
	assert.Equal(t, user1.Authentication.AuthType, user2.Authentication.AuthType)
	assert.Equal(t, user1.Authentication.AuthLevel, user2.Authentication.AuthLevel)
	assert.Equal(t, user1.Authentication.Secret, user2.Authentication.Secret)
	timeEquals(t, user1.Authentication.Expires, user2.Authentication.Expires)
	timeEquals(t, user1.Authentication.CreatedAt, user2.Authentication.CreatedAt)
	timeEquals(t, user1.Authentication.UpdatedAt, user2.Authentication.UpdatedAt)

	assert.Equal(t, user1.Device[0].DeviceId, user2.Device[0].DeviceId)
	assert.Equal(t, user1.Device[0].BrowserId, user2.Device[0].BrowserId)
	assert.Equal(t, user1.Device[0].UserAgent, user2.Device[0].UserAgent)
	assert.Equal(t, user1.Device[0].OperationSystem, user2.Device[0].OperationSystem)
	assert.Equal(t, user1.Device[0].Platform, user2.Device[0].Platform)
	assert.Equal(t, user1.Device[0].IsLogin, user2.Device[0].IsLogin)
	timeEquals(t, user1.Device[0].CreatedAt, user2.Device[0].CreatedAt)
	timeEquals(t, user1.Device[0].UpdatedAt, user2.Device[0].UpdatedAt)

	assert.Equal(t, user1.Profile.Description, user2.Profile.Description)
	timeEquals(t, user1.Profile.CreatedAt, user2.Profile.CreatedAt)
	timeEquals(t, user1.Profile.UpdatedAt, user2.Profile.UpdatedAt)
}

func TestAnonymousUserInsert(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_ = service.InitializeTable(ctx, true, true, false)

	now := time.Now()
	expire := time.Now().Add(365 * 24 * time.Hour)

	user := model.User{
		State:     types.StateUserActive,
		StatedAt:  &now,
		CreatedAt: &now,
		UpdatedAt: &now,

		Authentication: &model.Authentication{
			UserName:  "anonymous",
			AuthType:  types.UsingAnonymous,
			AuthLevel: types.StateAuthLevelAnonymous,
			Secret:    "asdf",
			Expires:   &expire,
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		Device: []*model.Device{
			{
				DeviceId:        "1",
				BrowserId:       "2",
				UserAgent:       "3",
				OperationSystem: "4",
				Platform:        "5",
				IsLogin:         true,
				CreatedAt:       &now,
				UpdatedAt:       &now,
			},
		},
		Profile: &model.Profile{
			Description: "anonymous user",
			CreatedAt:   &now,
			UpdatedAt:   &now,
		},
	}

	queryCtx, cancel := context.WithCancel(ctx)
	userId, err := service.UserInsert(queryCtx, user)
	cancel()
	require.NoError(t, err)
	assert.Greater(t, userId, int64(0))
	user.Id = userId

	queryCtx, cancel = context.WithCancel(ctx)
	findUser, err := service.UserFindOneById(queryCtx, userId)
	cancel()
	require.NoError(t, err)

	userEquals(t, &user, &findUser)
}

func TestEmailUserInsert(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_ = service.InitializeTable(ctx, true, true, false)

	now := time.Now()
	expire := time.Now().Add(365 * 24 * time.Hour)

	user := model.User{
		State:     types.StateUserActive,
		StatedAt:  &now,
		CreatedAt: &now,
		UpdatedAt: &now,

		Authentication: &model.Authentication{
			UserName:  "anonymous",
			Email:     "anonymous@email.com",
			Password:  "anonymousasdf",
			AuthType:  types.StateAuthTypeId,
			AuthLevel: types.StateAuthLevelAnonymous,
			Secret:    "asdvasdf",
			Expires:   &expire,
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		Device: []*model.Device{
			{
				DeviceId:        "1",
				BrowserId:       "2",
				UserAgent:       "3",
				OperationSystem: "4",
				Platform:        "5",
				CreatedAt:       &now,
				UpdatedAt:       &now,
			},
		},
		Profile: &model.Profile{
			Description: "profile",
			CreatedAt:   &now,
			UpdatedAt:   &now,
		},
	}

	queryCtx, cancel := context.WithCancel(ctx)
	userId, err := service.UserInsert(queryCtx, user)
	cancel()
	require.NoError(t, err)
	assert.Greater(t, userId, int64(0))
	user.Id = userId

	queryCtx, cancel = context.WithCancel(ctx)
	findUser, err := service.UserFindOneById(queryCtx, userId)
	cancel()
	require.NoError(t, err)

	userEquals(t, &user, &findUser)

	queryCtx, cancel = context.WithCancel(ctx)
	findUser, err = service.UserFindOneByEmail(queryCtx, "anonymous@email.com")
	cancel()
	require.NoError(t, err)

	userEquals(t, &user, &findUser)
}

func TestUserDelete(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_ = service.InitializeTable(ctx, true, true, false)

	now := time.Now()
	expire := time.Now().Add(365 * 24 * time.Hour)

	user := model.User{
		State:     types.StateUserActive,
		StatedAt:  &now,
		CreatedAt: &now,
		UpdatedAt: &now,

		Authentication: &model.Authentication{
			UserName:  "anonymous",
			Email:     "anonymous@email.com",
			Password:  "anonymousasdf",
			AuthType:  types.StateAuthTypeId,
			AuthLevel: types.StateAuthLevelAnonymous,
			Secret:    "asdvasdf",
			Expires:   &expire,
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		Device: []*model.Device{
			{
				DeviceId:        "1",
				BrowserId:       "2",
				UserAgent:       "3",
				OperationSystem: "4",
				Platform:        "5",
				CreatedAt:       &now,
				UpdatedAt:       &now,
			},
		},
		Profile: &model.Profile{
			Description: "profile",
			CreatedAt:   &now,
			UpdatedAt:   &now,
		},
	}

	queryCtx, cancel := context.WithCancel(ctx)
	userId, err := service.UserInsert(queryCtx, user)
	cancel()
	require.NoError(t, err)
	assert.Greater(t, userId, int64(0))
	user.Id = userId

	queryCtx, cancel = context.WithCancel(ctx)
	findUser, err := service.UserFindOneById(queryCtx, userId)
	cancel()
	require.NoError(t, err)

	userEquals(t, &user, &findUser)

	queryCtx, cancel = context.WithCancel(ctx)
	err = service.UserDeleteByUserId(queryCtx, userId)
	cancel()
	require.NoError(t, err)

	queryCtx, cancel = context.WithCancel(ctx)
	findUser, err = service.UserFindOneById(queryCtx, userId)
	cancel()
	assert.Equal(t, err, mongo.ErrNoDocuments)
}

func TestEmailUserLogout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_ = service.InitializeTable(ctx, true, true, false)

	now := time.Now()
	expire := time.Now().Add(365 * 24 * time.Hour)

	user := model.User{
		State:     types.StateUserActive,
		StatedAt:  &now,
		CreatedAt: &now,
		UpdatedAt: &now,

		Authentication: &model.Authentication{
			UserName:  "anonymous",
			Email:     "anonymous@email.com",
			Password:  "anonymousasdf",
			AuthType:  types.StateAuthTypeId,
			AuthLevel: types.StateAuthLevelAnonymous,
			Secret:    "asdvasdf",
			Expires:   &expire,
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		Device: []*model.Device{
			{
				DeviceId:        "1",
				BrowserId:       "2",
				UserAgent:       "3",
				OperationSystem: "4",
				Platform:        "5",
				CreatedAt:       &now,
				UpdatedAt:       &now,
			},
		},
		Profile: &model.Profile{
			Description: "profile",
			CreatedAt:   &now,
			UpdatedAt:   &now,
		},
	}

	queryCtx, cancel := context.WithCancel(ctx)
	userId, err := service.UserInsert(queryCtx, user)
	cancel()
	require.NoError(t, err)
	assert.Greater(t, userId, int64(0))
	user.Id = userId

	queryCtx, cancel = context.WithCancel(ctx)
	findUser, err := service.UserFindOneById(queryCtx, userId)
	cancel()
	require.NoError(t, err)

	userEquals(t, &user, &findUser)

	queryCtx, cancel = context.WithCancel(ctx)
	err = service.UserLogout(queryCtx, userId)
	cancel()
	require.NoError(t, err)

	queryCtx, cancel = context.WithCancel(ctx)
	findUser, err = service.UserFindOneById(queryCtx, userId)
	cancel()
	require.NoError(t, err)

	assert.Equal(t, findUser.State, "INACTIVE")
}

func TestAnonymousUserLogout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_ = service.InitializeTable(ctx, true, true, false)

	now := time.Now()
	expire := time.Now().Add(365 * 24 * time.Hour)

	user := model.User{
		State:     types.StateUserActive,
		StatedAt:  &now,
		CreatedAt: &now,
		UpdatedAt: &now,

		Authentication: &model.Authentication{
			UserName:  "anonymous",
			AuthType:  types.UsingAnonymous,
			AuthLevel: types.StateAuthLevelAnonymous,
			Secret:    "asdf",
			Expires:   &expire,
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		Device: []*model.Device{
			{
				DeviceId:        "1",
				BrowserId:       "2",
				UserAgent:       "3",
				OperationSystem: "4",
				Platform:        "5",
				IsLogin:         true,
				CreatedAt:       &now,
				UpdatedAt:       &now,
			},
		},
		Profile: &model.Profile{
			Description: "anonymous user",
			CreatedAt:   &now,
			UpdatedAt:   &now,
		},
	}

	queryCtx, cancel := context.WithCancel(ctx)
	userId, err := service.UserInsert(queryCtx, user)
	cancel()
	require.NoError(t, err)
	assert.Greater(t, userId, int64(0))
	user.Id = userId

	queryCtx, cancel = context.WithCancel(ctx)
	findUser, err := service.UserFindOneById(queryCtx, userId)
	cancel()
	require.NoError(t, err)

	userEquals(t, &user, &findUser)

	queryCtx, cancel = context.WithCancel(ctx)
	err = service.UserLogout(queryCtx, userId)
	cancel()
	require.NoError(t, err)

	queryCtx, cancel = context.WithCancel(ctx)
	findUser, err = service.UserFindOneById(queryCtx, userId)
	cancel()
	assert.Equal(t, err, mongo.ErrNoDocuments)
}

func TestFindAll(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_ = service.InitializeTable(ctx, true, true, false)

	userCreate(t, ctx, 100)

	queryCtx, cancel := context.WithCancel(ctx)
	findUsers, err := service.UserFindAllByPaging(queryCtx, model.Paging{})
	cancel()
	require.NoError(t, err)

	assert.Equal(t, len(findUsers), 100)
}

func TestFindAllByPagingIdOrderByDesc1(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_ = service.InitializeTable(ctx, true, true, false)

	userCreate(t, ctx, 100)

	queryCtx, cancel := context.WithCancel(ctx)
	findUsers, err := service.UserFindAllByPaging(queryCtx, model.Paging{Id: 51})
	cancel()
	require.NoError(t, err)

	assert.Equal(t, len(findUsers), 50)

	var expectIds []int64
	for i := 50; i > 0; i-- {
		expectIds = append(expectIds, int64(i))
	}

	var actualIds []int64
	for _, user := range findUsers {
		actualIds = append(actualIds, user.Id)
	}

	assert.EqualValues(t, expectIds, actualIds)
}

func TestFindAllByPagingIdOrderByDesc2(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_ = service.InitializeTable(ctx, true, true, false)

	userCreate(t, ctx, 200)

	queryCtx, cancel := context.WithCancel(ctx)
	findUsers, err := service.UserFindAllByPaging(queryCtx, model.Paging{Id: 151})
	cancel()
	require.NoError(t, err)

	assert.Equal(t, len(findUsers), 100)

	var expectIds []int64
	for i := 150; i > 50; i-- {
		expectIds = append(expectIds, int64(i))
	}

	var actualIds []int64
	for _, user := range findUsers {
		actualIds = append(actualIds, user.Id)
	}

	assert.EqualValues(t, expectIds, actualIds)
}

func TestFindAllByPagingIdOrderByDesc3(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_ = service.InitializeTable(ctx, true, true, false)

	userCreate(t, ctx, 200)

	queryCtx, cancel := context.WithCancel(ctx)
	findUsers, err := service.UserFindAllByPaging(queryCtx, model.Paging{Id: 201})
	cancel()
	require.NoError(t, err)

	assert.Equal(t, len(findUsers), 100)

	var expectIds []int64
	for i := 200; i > 100; i-- {
		expectIds = append(expectIds, int64(i))
	}

	var actualIds []int64
	for _, user := range findUsers {
		actualIds = append(actualIds, user.Id)
	}

	assert.EqualValues(t, expectIds, actualIds)
}

func TestUserUpdateAllFields(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_ = service.InitializeTable(ctx, true, true, false)

	userCreate(t, ctx, 200)

	queryCtx, cancel := context.WithCancel(ctx)
	user, err := service.UserFindOneById(queryCtx, 100)
	cancel()
	require.NoError(t, err)

	marshal, err := jsoniter.Marshal(user)
	require.NoError(t, err)

	var deepCopyUser model.User
	err = jsoniter.Unmarshal(marshal, &deepCopyUser)
	require.NoError(t, err)

	userEquals(t, &user, &deepCopyUser)

	now := time.Now()
	user.State = "INACTIVE"
	user.StatedAt = &now

	queryCtx, cancel = context.WithCancel(ctx)
	err = service.UserUpdate(queryCtx, user)
	cancel()
	require.NoError(t, err)

	queryCtx, cancel = context.WithCancel(ctx)
	updateUser, err := service.UserFindOneById(queryCtx, 100)
	cancel()
	require.NoError(t, err)

	userEquals(t, &user, &updateUser)
}

func TestUserUpdateEachFields1(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_ = service.InitializeTable(ctx, true, true, false)

	userCreate(t, ctx, 200)

	queryCtx, cancel := context.WithCancel(ctx)
	user, err := service.UserFindOneById(queryCtx, 100)
	cancel()
	require.NoError(t, err)

	marshal, err := jsoniter.Marshal(user)
	require.NoError(t, err)

	var deepCopyUser model.User
	err = jsoniter.Unmarshal(marshal, &deepCopyUser)
	require.NoError(t, err)

	userEquals(t, &user, &deepCopyUser)

	now := time.Now()
	uu := model.User{
		Id:       user.Id,
		State:    types.StateUserInactive,
		StatedAt: &now,
	}

	queryCtx, cancel = context.WithCancel(ctx)
	err = service.UserUpdate(queryCtx, uu)
	cancel()
	require.NoError(t, err)

	queryCtx, cancel = context.WithCancel(ctx)
	updateUser, err := service.UserFindOneById(queryCtx, 100)
	cancel()
	require.NoError(t, err)

	user.State = types.StateUserInactive
	user.StatedAt = &now
	userEquals(t, &user, &updateUser)
}
