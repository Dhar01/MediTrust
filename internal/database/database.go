package database

import (
	"medicine-app/internal/database/cart/cartDB"
	"medicine-app/internal/database/general/genDB"
	"medicine-app/internal/database/medicine/pgMedicineDB"
	"medicine-app/internal/database/user/userDB"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	User     userDB.Queries
	Medicine pgMedicineDB.Queries
	Cart     cartDB.Queries
	Helper   genDB.Queries
}

func New(pool *pgxpool.Pool) *DB {
	return &DB{
		User:     *userDB.New(pool),
		Medicine: *pgMedicineDB.New(pool),
		Cart:     *cartDB.New(pool),
		Helper:   *genDB.New(pool),
	}
}
