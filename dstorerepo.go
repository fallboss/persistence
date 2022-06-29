package persistence

type DSClient struct{}

func GetDataStoreRepo(cnt *DSClient) DBQuery {
	return &datastoreRepo{cnt}
}

type datastoreRepo struct {
	cnt *DSClient
}

func (r *datastoreRepo) FindColumnsByConditions(tableName string, columns []string, conditions []DBCondition, response interface{}) error {
	panic("implement me")
}

func (r *datastoreRepo) ExistsBy(tableName string, condition DBCondition) (bool, error) {
	panic("implement me")
}

func (r *datastoreRepo) Update(tableName string, columnValues []DBValue, conditions []DBCondition) error {
	panic("implement me")
}

func (r *datastoreRepo) ExecuteQuery(query string, response interface{}) error {
	panic("Feature not implemented for datastore")
}

func (r *datastoreRepo) FindAllBy(tableName string, condition DBCondition, response interface{}) error {
	panic("Feature not implemented for datastore")
}

func (r *datastoreRepo) FindByConditions(tableName string, conditions []DBCondition, response interface{}) error {
	panic("Feature not implemented for datastore")
}

func (r *datastoreRepo) InsertInto(tableName string, columnValues []DBValue) error {
	panic("Feature not implemented for datastore")
}
