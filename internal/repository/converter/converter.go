package converter

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func ToPgUUIDFromService(id string) pgtype.UUID {
	return pgtype.UUID{Bytes: uuid.MustParse(id), Valid: true}
}
