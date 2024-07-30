package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func GetConnectionPool(context context.Context, config Config) *pgxpool.Pool {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable statement_cache_mode=describe pool_max_conns=%s pool_max_conn_idle_time=%s",
		config.Host,
		config.Port,
		config.UserName,
		config.Password,
		config.DbName,
		config.MaxConnections,
		config.MaxConnectionIdleTime)
	connect, err := pgxpool.Connect(context, connString)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		panic(err)
	}
	return connect
}
