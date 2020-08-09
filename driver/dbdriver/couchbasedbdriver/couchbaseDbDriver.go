package couchbasedriver

// CouchbaseDBDriver is responsible for performing operations on a Couchbase database
type CouchbaseDBDriver struct {
}

// AddDoc is responsible for creating a new document in the Couchbase database
func (driver *CouchbaseDBDriver) AddDoc(doc interface{}, args ...interface{}) (result *interface{}, err error) {
	return nil, nil
}

// ReplaceDoc is responsible for replacing an existing document in the Couchbase database
func (driver *CouchbaseDBDriver) ReplaceDoc(UID string, doc interface{}, args ...interface{}) (result *interface{}, err error) {
	return nil, nil
}

// Command is responsible for executing an command in the Couchbase database
func (driver *CouchbaseDBDriver) Command(command string, args ...interface{}) (result *interface{}, err error) {
	return nil, nil
}

// QueryDocs is responsible for obtaining several documents from the Couchbase database
func (driver *CouchbaseDBDriver) QueryDocs(query string, args ...interface{}) ([]interface{}, error) {
	return nil, nil
}

// QueryDoc is responsible for obtaining a single document from the Couchbase database
func (driver *CouchbaseDBDriver) QueryDoc(query string, args ...interface{}) (result *interface{}, err error) {
	return nil, nil
}
