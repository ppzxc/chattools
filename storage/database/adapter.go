package database

import (
	"context"
	model2 "github.com/ppzxc/chattools/storage/database/model"
)

type Service interface {
	InitializeTable(ctx context.Context, dropTableOnStart bool, createTableOnStart bool, testUserInsertOnStart bool) error
	Close(context.Context) error

	UserInsert(ctx context.Context, user model2.User) (id int64, err error)
	UserFindAllByPaging(ctx context.Context, paging model2.Paging) (users []model2.User, err error)
	UserFindAllByTopicId(ctx context.Context, topicId int64) ([]model2.User, error)
	UserFindOneById(ctx context.Context, id int64) (user model2.User, err error)
	UserFindOneByEmail(ctx context.Context, email string) (user model2.User, err error)
	UserUpdate(ctx context.Context, user model2.User) (err error)
	UserDeleteByUserId(ctx context.Context, userId int64) (err error)
	UserLogout(ctx context.Context, userId int64) (err error)

	TopicInsert(ctx context.Context, topic model2.Topic) (topicId int64, err error)
	TopicFindAll(ctx context.Context, paging model2.Paging) (topics []model2.Topic, err error)
	TopicFindAllByUserId(ctx context.Context, topicId int64) (topics []model2.Topic, err error)
	TopicFindOneById(ctx context.Context, topicId int64) (topic model2.Topic, err error)
	TopicDelete(ctx context.Context, topicId int64) (err error)

	SubscriptionInsert(ctx context.Context, subscription model2.Subscription) (subscriptionId int64, err error)
	SubscriptionFindAllByTopicId(ctx context.Context, topicId int64) (subscriptions []model2.Subscription, err error)
	SubscriptionFindOneByUserIdAndTopicId(ctx context.Context, userId int64, topicId int64) (subscription model2.Subscription, err error)
	SubscriptionExistsByUserIdAndTopicId(ctx context.Context, userId int64, topicId int64) (subscriptionId int64, err error)
	SubscriptionUpdateAck(ctx context.Context, subscription model2.Subscription) (err error)
	SubscriptionDeleteAllByUserId(ctx context.Context, userId int64) (err error)
	SubscriptionDeleteByTopicIdAndUserId(ctx context.Context, topicId int64, userId int64) (err error)

	MessageInsert(ctx context.Context, message model2.Message) (sequenceId int64, err error)
	MessageFindAllByTopicIdAndMoreThanSequenceId(ctx context.Context, topicId int64, sequenceId int64) ([]model2.Message, error)
	MessageFindMaxSequenceIdByTopicId(ctx context.Context, topicId int64) (int64, error)

	FileFindOneById(ctx context.Context, fileId int64) (model2.File, error)
	FileInsert(ctx context.Context, file model2.File) (fileId int64, err error)

	ProfileImageUpdate(ctx context.Context, profile model2.File) (fileId int64, err error)
	ProfileUpdateByUserId(ctx context.Context, profile model2.Profile) (err error)

	NotifyInsertMany(ctx context.Context, notify []*model2.Notify) (err error)
	NotifyUpdate(ctx context.Context, notify *model2.Notify) (err error)
	NotifyFindAllByReceiveUserId(ctx context.Context, receiveUserId int64) ([]*model2.Notify, error)
	NotifyFindOneReceiveUserIdById(ctx context.Context, notifyId int64) (int64, error)
}
