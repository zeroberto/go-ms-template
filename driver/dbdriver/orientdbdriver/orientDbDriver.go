package orientdbdriver

// OrientDBDriver is responsible for performing operations on OrientDB database
type OrientDBDriver struct {
}

// AddElement is responsible for creating a new element on the graph database
func (driver *OrientDBDriver) AddElement(element interface{}, elementType interface{}, args ...interface{}) (result *interface{}, err error) {
	return nil, nil
}

// Command is responsible for executing an command in the graph database
func (driver *OrientDBDriver) Command(command string, args ...interface{}) (result *interface{}, err error) {
	return nil, nil
}

// QueryElements is responsible for obtaining several elements from the graph database
func (driver *OrientDBDriver) QueryElements(query string, args ...interface{}) ([]interface{}, error) {
	return nil, nil
}

// QueryElement is responsible for obtaining a single element from the graph database
func (driver *OrientDBDriver) QueryElement(query string, args ...interface{}) (result *interface{}, err error) {
	return nil, nil
}
