package database

import "log/slog"

func (postgres *Postgres) CreateAllTables(tables []interface{}) error {
	if err := postgres.Db.AutoMigrate(tables...); err != nil {
		return err
	}
	slog.Info("All tables migrated successfully")
	return nil
}
