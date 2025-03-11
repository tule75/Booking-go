-- name: CreateProperty :execresult
INSERT INTO properties (id, owner_id, name, description, location, price, amenities)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetPropertyByID :one
SELECT 
  properties.id AS property_id,
  properties.owner_id,
  properties.name,
  properties.description,
  properties.location,
  properties.price,
  properties.amenities,
  properties.created_at,
  COALESCE(
    JSON_ARRAYAGG(
      JSON_OBJECT('id', v.id, 'name', v.name, 'price', v.price, 'max_guests', v.max_guests, 'is_available', v.is_available)),
    '[]') AS rooms
FROM properties LEFT JOIN rooms v ON properties.id = rooms.property_id
WHERE properties.id = ? AND properties.deleted_at IS NULL;

-- name: SearchProperties :many
SELECT 
  properties.id AS property_id,
  properties.owner_id,
  properties.name,
  properties.description,
  properties.location,
  properties.price,
  properties.amenities,
  properties.created_at
FROM properties
WHERE properties.deleted_at IS NULL
  AND MATCH(properties.name, properties.location) AGAINST (? IN NATURAL LANGUAGE MODE)
  AND properties.price BETWEEN ? AND ?
ORDER BY properties.price ASC, properties.created_at DESC
LIMIT ? OFFSET ?;

-- name: ListPropertiesByOwner :many
SELECT 
  properties.id AS property_id,
  properties.owner_id,
  properties.name,
  properties.description,
  properties.location,
  properties.price,
  properties.amenities,
  properties.created_at
FROM properties
WHERE properties.owner_id = ? AND properties.deleted_at IS NULL
ORDER BY properties.created_at DESC
LIMIT ? OFFSET ?;

-- name: UpdateProperty :exec
UPDATE properties
SET name = ?, description = ?, location = ?, price = ?, amenities = ?
WHERE id = ? AND deleted_at IS NULL;

-- name: SoftDeleteProperty :exec
UPDATE properties SET deleted_at = NOW() WHERE id = ?;
