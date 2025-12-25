package repository

import (
	"context"
	"github.com/practice-cloud-run-functions/myapp/common"
)

func SelectUsers() ([]map[string]any, error) {
	adapter := common.NewAlloyDbAdapter()
    pool, err := adapter.CreatePool() 
	if err != nil {
		return nil, err
	}
	defer pool.Close()

	rows, err := pool.Query(context.Background(), "SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []map[string]any
	for rows.Next() {
		var id int
		var name, email string
		if err := rows.Scan(&id, &name, &email); err != nil {
			return nil, err
		}
		users = append(users, map[string]any{
			"id":    id,
			"name":  name,
			"email": email,
		})
	}

	return users, nil
}
