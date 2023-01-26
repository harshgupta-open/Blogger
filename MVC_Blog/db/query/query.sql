-- name: InsertPost :exec
INSERT INTO posts (userName,title,posDescription,created) values($1,$2,$3,$4);

-- name: QueryGetAllPost :many
SELECT * FROM posts;

-- name: QueryGetCommentById :many
SELECT * FROM "comments" where commentPostId = $1; 

-- name: QueryGetPostById :one
SELECT * FROM "posts" where postId = $1;

-- name: InsertComment :exec
insert into "comments" (commentPostId,commentDesc,commentUser,commentCreated) values ($1,$2,$3,$4);

