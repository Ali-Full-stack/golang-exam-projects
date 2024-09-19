package mongodb

import (
	"context"
	"fmt"
	"notification/internal/model"
	"notification/protos"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepo struct {
	*mongo.Client
}

func NewMongoRepo(mongo_url string) (*MongoRepo, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongo_url))
	if err != nil {
		return nil, fmt.Errorf("error: failed mongoDB connection: %v", err)
	}
	return &MongoRepo{client}, nil
}

func (m *MongoRepo) AddNotifyIntoMongo(user_id, message string, report *protos.Report)error{
	collection := m.Database("notification").Collection("notification")

	var notify model.Notify

	notify.Seen = false
	notify.Report = &model.Report{
		Income:     report.Income,
		Expenses:   report.Expenses,
		NetSavings: report.NetSavings,
	}
	notify.Message=message
	notify.UserID=user_id
	notify.Date=time.Now().Format(time.ANSIC)

	_, err :=collection.InsertOne(context.Background(), notify)
	if err != nil {
		return fmt.Errorf("cannot insert notification into mongoDB :%v",err)
	}
	return nil
}

func (m *MongoRepo) GetNotficationFromMongoDB(user_id string) ([]*model.Notify, error) {
	collection := m.Database("notification").Collection("notification")

	var notifies []*model.Notify
	cursor, err := collection.Find(context.Background(), bson.M{"user_id":user_id})
	if err != nil {
		return nil, fmt.Errorf("cannot find notification from mongoDB :%v",err)
	}
	defer cursor.Close(context.Background())
	cursor.All(context.Background(), &notifies)

	result, err :=collection.UpdateMany(context.Background(), bson.M{"user_id":user_id, "seen":false}, bson.M{"$set": bson.M{"seen":true}})
	if err != nil {
		return nil, fmt.Errorf("cannot update notification from mongoDB :%v",err)
	}
	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("no unread notification")
	}

	return notifies, nil
}

func (m *MongoRepo) GetUnreadNotfications(user_id string) ([]*model.Notify, error) {
	collection := m.Database("notification").Collection("notification")
	var notifies []*model.Notify
	cursor, err := collection.Find(context.Background(), bson.M{"user_id":user_id, "seen":false})
	if err != nil {
		return nil, fmt.Errorf("cannot find notification from mongoDB :%v",err)
	}
	defer cursor.Close(context.Background())
	cursor.All(context.Background(), &notifies)

	result, err :=collection.UpdateMany(context.Background(), bson.M{"user_id":user_id, "seen":false}, bson.M{"$set": bson.M{"seen":true}})
	if err != nil {
		return nil, fmt.Errorf("cannot update notification from mongoDB :%v",err)
	}
	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("no unread notification")
	}
	return notifies, nil
}