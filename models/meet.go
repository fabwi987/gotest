package models

import "time"

type Meet struct {
	ID          int       `json:"ID" bson:"ID"`
	Location    string    `json:"Location" bson:"Location"`
	Date        time.Time `json:"Date" bson:"Date"`
	Text        string    `json:"Text" bson:"Text"`
	Created     time.Time `json:"Created" bson:"Created"`
	LastUpdated time.Time `json:"LastUpdated" bson:"LastUpdated"`
	URL         string    `json:"URL" bson:"URL"`
	IDUser      int       `json:"IDuser" bson:"IDuser"`
}

func (db *DB) GetMeets() ([]*Meet, error) {

	rows, err := db.Query("SELECT idmeet, location, date, text, created, lastupdated, url, iduser FROM meet")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	poss := make([]*Meet, 0)
	for rows.Next() {
		ps := new(Meet)
		err = rows.Scan(&ps.ID, &ps.Location, &ps.Date, &ps.Text, &ps.Created, &ps.LastUpdated, &ps.URL, &ps.IDUser)
		if err != nil {
			return nil, err
		}
		poss = append(poss, ps)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return poss, nil
}
