package query

var GetPreview = `SELECT * FROM preview WHERE id = $1`

var AddPreview = `INSERT INTO preview(url)
VALUES($1);`

var DeletePreview = `DELETE FROM preview WHERE id = $1`

var GetVideo = `SELECT * FROM video WHERE id = $1`

var AddVideo = `INSERT INTO video(url)
VALUES($1);`

var DeleteVideo = `DELETE FROM video WHERE id = $1`
