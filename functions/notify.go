package main

import (
	"context"
	"fmt"
	"net/http"
	"notifications-push/config"
	"notifications-push/database"
	"notifications-push/mail"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	db := database.ConnectDB()
	defer db.Close()

	notifications := database.GetUnreadNotifications(db)
	ids := []int64{}
	for _, notification := range notifications {
		fmt.Printf("ID: %d, Email: %s, Status: %s, SentAt: %s\n", notification.Id, notification.Email, notification.Status, notification.SentAt)
		mail.SendEmail(config.SENDER, notification.Email, notification.Name, notification.Message)
		ids = append(ids, notification.Id)
	}

	if len(ids) > 0 {
		database.UpdateNotificationsToSent(db, ids)
	} else {
		fmt.Println("No notifications to update")
	}

	return events.APIGatewayProxyResponse{
		Body:       "All good",
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
