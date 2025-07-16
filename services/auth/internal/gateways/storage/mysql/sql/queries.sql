
-- name: GetByOAuthAccId :one
SELECT * FROM user WHERE oauth_acc_id = ?;
