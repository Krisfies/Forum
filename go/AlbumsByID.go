package main

// import (
// 	"database/sql"
// 	"fmt"
// )

// // albumByID queries for the album with the specified ID.
// func AlbumByID(id int64) (Album, error) {
// 	// An album to hold data from the returned row.
// 	var alb Album

// 	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
// 	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
// 		if err == sql.ErrNoRows {
// 			return alb, fmt.Errorf("albumsById %d: no such album", id)
// 		}
// 		return alb, fmt.Errorf("albumsById %d: %v", id, err)
// 	}
// 	return alb, nil
// }
