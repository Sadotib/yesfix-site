-- name: InsertUserTest :one 
INSERT INTO users(usid, firstname,lastname,email,phone,city) VALUES($1, $2, $3, $4, $5, $6) RETURNING *;