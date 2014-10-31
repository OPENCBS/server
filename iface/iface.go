package iface

import (
	"net/http"
	"database/sql"
	"github.com/OPENCBS/server/model"
)

// Repositories
type UserRepo interface {
	FindAll(db *sql.DB, offset int, limit int) ([]*model.User, error)
	FindById(db *sql.DB, id int) (*model.User, error)
	FindByUsernameAndPassword(db *sql.DB, username string, password string) (*model.User, error)
}

type ClientRepo interface {
	FindAll(db *sql.DB, offset int, limit int) ([]*model.Client, error)
}

var GetDb func(r *http.Request) (*sql.DB, error)

var NewUserRepo func() UserRepo 
var NewClientRepo func() ClientRepo

