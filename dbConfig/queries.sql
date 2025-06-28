-- name: InsertUserTest :one 
INSERT INTO prebook_data(uid, firstname,lastname,email,phone,city,locality) VALUES($1, $2, $3, $4, $5, $6,$7) RETURNING *;