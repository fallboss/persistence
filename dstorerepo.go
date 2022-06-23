package persistence

type DSClient struct{}

func GetDataStoreRepo(cnt *DSClient) DBQuery {
	return &datastoreRepo{cnt}
}

type datastoreRepo struct {
	cnt *DSClient
}

<<<<<<< HEAD
func (r *datastoreRepo) FindColumnsByConditions(tableName string, columns []string, conditions []DBCondition, response *interface{}) error {
=======
func (r *datastoreRepo) FindColumnsByConditions(tableName string, columns []string, conditions []DBCondition, response *interface{}) {
>>>>>>> 6bb88ef53f641b5d26df89f10cc0f97aeb5f422f
	//TODO implement me
	panic("implement me")
}

func (r *datastoreRepo) ExistsBy(tableName string, condition DBCondition, response *bool) error {
	panic("implement me")
}

func (r *datastoreRepo) Update(tableName string, columnValues []DBValue, conditions []DBCondition, response *interface{}) error {
	panic("implement me")
}

func (r *datastoreRepo) ExecuteQuery(query string, response *interface{}) error {
	panic("Feature not implemented for datastore")
}

func (r *datastoreRepo) FindAllBy(tableName string, condition DBCondition, response *interface{}) error {
	panic("Feature not implemented for datastore")
}

func (r *datastoreRepo) FindByConditions(tableName string, conditions []DBCondition, response *interface{}) error {
	panic("Feature not implemented for datastore")
}

func (r *datastoreRepo) InsertInto(tableName string, columnValues []DBValue, response *interface{}) error {
	panic("Feature not implemented for datastore")
}
