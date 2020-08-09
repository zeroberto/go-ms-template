package sqldbdriver

// BeginTransaction is responsible for initiating a transaction in the database
func (driver *SQLDBDriver) BeginTransaction() error {
	return nil
}

// Commit is responsible for persisting the modified information in the database
func (driver *SQLDBDriver) Commit() error {
	return driver.DB.Commit()
}

// EndTransaction is responsible for ending the transaction in the database
func (driver *SQLDBDriver) EndTransaction() error {
	return driver.DB.Rollback()
}

// Rollback is responsible for undoing all modifications made to the database within the current transaction
func (driver *SQLDBDriver) Rollback() error {
	return nil
}
