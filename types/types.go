package types

type Notification struct {
	Id      int64  `db:"id"`
	Email   string `db:"email"`
	Name    string `db:"name"`
	Message string `db:"message"`
	Status  string `db:"status"`
	SentAt  string `db:"sent_at"`
}
