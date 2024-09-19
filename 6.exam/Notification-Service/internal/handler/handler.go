package handler

import (
	"context"
	"fmt"
	tpb "notification/protos"
	"notification/storage/mongodb"
)

type NotficationService struct {
	tpb.UnimplementedNotificationServiceServer
	Mongo *mongodb.MongoRepo
}

func NewNotficationService(mongo *mongodb.MongoRepo) *NotficationService {
	return &NotficationService{Mongo: mongo}
}

func (n *NotficationService) GetNotfication(ctx context.Context, in *tpb.Empty) (*tpb.NotifyList, error) {
	user_id := ctx.Value("user_id").(string)
	notifies, err := n.Mongo.GetNotficationFromMongoDB(user_id)
	if err != nil {
		return nil, err
	}

	if len(notifies) == 0 {
		return nil, fmt.Errorf("no unread message")
	}

	var notifyList tpb.NotifyList
	for _, notify := range notifies {
		notifyList.NotifyList = append(notifyList.NotifyList, &tpb.Notify{
			UserId:  user_id,
			Message: notify.Message,
			Report: &tpb.Report{
				Income:     notify.Report.Income,
				Expenses:   notify.Report.Expenses,
				NetSavings: notify.Report.NetSavings,
			},
			Date: notify.Date,
		})
	}
	return &notifyList, nil
}

func (n *NotficationService) GetUnreadNotfications(ctx context.Context, in *tpb.Empty) (*tpb.NotifyList, error) {
	user_id := ctx.Value("user_id").(string)
	notifies, err := n.Mongo.GetUnreadNotfications(user_id)
	if err != nil {
		return nil, err
	}

	if len(notifies) == 0 {
		return nil, fmt.Errorf("no unread message")
	}


	var notifyList tpb.NotifyList
	for _, notify := range notifies {
		notifyList.NotifyList = append(notifyList.NotifyList, &tpb.Notify{
			UserId:  user_id,
			Message: notify.Message,
			Report:	&tpb.Report{
				Income:     notify.Report.Income,
				Expenses:   notify.Report.Expenses,
				NetSavings: notify.Report.NetSavings,
			},
			Date: notify.Date,
		})
	}
	return &notifyList, nil
}
