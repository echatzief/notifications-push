package database

import (
	"fmt"
	"notifications-push/types"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB() *sqlx.DB {
	url := "postgres://user:password@localhost:5432/notifications_push?sslmode=disable"
	db, err := sqlx.Connect("postgres", url)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
	} else {
		fmt.Println("Connected to database")
	}
	return db
}

func GetUnreadNotifications(db *sqlx.DB) []types.Notification {
	notifications := []types.Notification{}
	err := db.Select(&notifications, "SELECT * FROM notifications WHERE status = 'not_sent' AND sent_at <= $1", time.Now())
	if err != nil {
		fmt.Println("Failed to get unread notifications")
		return []types.Notification{}
	}
	return notifications
}

func UpdateNotificationsToSent(db *sqlx.DB, ids []int64) {
	query := "UPDATE notifications SET status = 'sent' WHERE id IN (?)"

	query, args, err := sqlx.In(query, ids)
	if err != nil {
		fmt.Println("Failed to set ids")
	}

	query = db.Rebind(query)
	_, err = db.Exec(query, args...)
	if err != nil {
		fmt.Println("Failed to update notifications")
	} else {
		fmt.Println("Successfully updated notification statuses")
	}
}
