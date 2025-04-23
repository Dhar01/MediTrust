package database

import (
	"medicine-app/internal/database/cart/cartDB"
	"medicine-app/internal/database/general/genDB"
	"medicine-app/internal/database/medicine/medDB"
	"medicine-app/internal/database/user/userDB"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	User     userDB.Queries
	Medicine medDB.Queries
	Cart     cartDB.Queries
	Helper   genDB.Queries
}

func New(pool *pgxpool.Pool) *DB {
	return &DB{
		User:     *userDB.New(pool),
		Medicine: *medDB.New(pool),
		Cart:     *cartDB.New(pool),
		Helper:   *genDB.New(pool),
	}
}
