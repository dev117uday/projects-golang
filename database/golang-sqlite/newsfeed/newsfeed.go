package newsfeed

import "database/sql"

// Feed : connect to db
type Feed struct {
	DB *sql.DB
}

// NewFeed : to create a connection to db
func NewFeed(db *sql.DB) *Feed {
	stmnt, _ := db.Prepare(`
		create table if not exists newsfeed
			(
				ID integer not null
					constraint newsfeed_pk
						primary key autoincrement,
				CONTENT text
			);
	`)
	stmnt.Exec()
	return &Feed{
		DB: db,
	}
}

// AddItem : to add item to db
func (feed *Feed) AddItem(item Item) {
	stmnt, _ := feed.DB.Prepare(`
		INSERT INTO newsfeed (content) values (?)
	`)
	stmnt.Exec(item.Content)
}

// GetItem : to get item from db
func (feed *Feed) GetItem() []Item {

	rows, _ := feed.DB.Query(`
		Select * from newsfeed
	`)
	items := []Item{}
	var id int
	var content string
	for rows.Next() {
		rows.Scan(&id, &content)
		payload := Item{
			ID:      id,
			Content: content,
		}
		items = append(items, payload)
	}
	return items
}
