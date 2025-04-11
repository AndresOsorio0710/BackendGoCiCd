package dbcontext

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/AndresOsorio0710/BackendGoCiCd/internals/config"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type DbContext struct {
	DB *sql.DB
}

func NewDbContext(cfg config.PostgresConfig) (*DbContext, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error al abrir conexión: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error al conectar con la base de datos: %w", err)
	}

	log.Println("Conexión a la base de datos PostgreSQL establecida.")
	return &DbContext{DB: db}, nil
}

func (ctx *DbContext) OpenConnection() (*sql.Conn, error) {
	conn, err := ctx.DB.Conn(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error obteniendo conexión: %w", err)
	}
	return conn, nil
}

func (ctx *DbContext) Close() {
	if ctx.DB != nil {
		err := ctx.DB.Close()
		if err != nil {
			log.Println("Error al cerrar la conexión con la base de datos:", err)
		} else {
			log.Println("Conexión con la base de datos cerrada.")
		}
	}
}
