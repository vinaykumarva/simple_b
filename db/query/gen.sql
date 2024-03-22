

-- name: GetJSON :many
SELECT '{"bar": "baz", "balance": 7.77, "active":false}'::json AS A,
        NULL::json AS B;
