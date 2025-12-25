package common

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// 環境変数用構造体
type AlloyConfig struct {
	Host     string
	Port     string
	DB_name  string
	User     string
	Password string
}

// コンストラクタ
func NewAlloyDbAdapter() AlloyConfig {
	config := AlloyConfig{}
	config.getEnvVariable()
	return config
}

// 環境変数の取得
func (ac *AlloyConfig) getEnvVariable() {
	ac.Host = os.Getenv("DB_HOST")
	ac.Port = os.Getenv("DB_PORT")
	ac.DB_name = os.Getenv("DB_NAME")
	ac.User = os.Getenv("DB_USER")
	ac.Password = os.Getenv("DB_PASSWORD")
}

// DBプールの作成
func (ac *AlloyConfig) CreatePool() (*pgxpool.Pool, error) {
	// DSNの作成
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		ac.User,
		ac.Password,
		ac.Host,
		ac.Port,
		ac.DB_name,
	)
	// 接続タイムアウト設定
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	// DB接続
	pool, err := pgxpool.New(ctx, dsn)

	return pool, err
}
