package persistence

import (
	"fmt"
	"github.com/fallboss/persistence/formatter"
	"github.com/georgysavva/scany/pgxscan"
	"strings"
	"time"
)

func GetPgRepo(cnt *PgClient) DBQuery {
	return &pgRepo{cnt}
}

type pgRepo struct {
	cnt *PgClient
}

const queryExists = "SELECT EXISTS (SELECT 1 FROM %s WHERE %s %s '%v')"

func (r *pgRepo) FindColumnsByConditions(tableName string, columns []string, conditions []DBCondition, response interface{}) error {
	var conditionsQuery = ""
	for _, condition := range conditions {
		if !strings.EqualFold(conditionsQuery, "") {
			conditionsQuery = conditionsQuery + " AND "
		}
		op, err := parseOperator(condition.Operator)
		if err != nil {
			return err
		}
		conditionsQuery = conditionsQuery + fmt.Sprintf("(%s %s '%v')", condition.FieldName, op, condition.Value)
	}
	var query = fmt.Sprintf("SELECT %s FROM %s WHERE %s", strings.Join(columns, ","), tableName, conditionsQuery)
	row, err := r.cnt.RunQueryRows(query)
	if err != nil {
		return err
	}
	return (*row).Scan(&response)
}

func (r *pgRepo) ExecuteQuery(query string, response interface{}) error {
	row, err := r.cnt.RunQueryRows(query)
	if err != nil {
		return err
	}
	return pgxscan.ScanOne(response, *row)
}

func (r *pgRepo) FindAllBy(tableName string, condition DBCondition, response interface{}) error {
	op, err := parseOperator(condition.Operator)
	if err != nil {
		return err
	}
	var query = fmt.Sprintf("SELECT * FROM %s WHERE %s %s '%v'", tableName, condition.FieldName, op, condition.Value)
	row, err := r.cnt.RunQueryRows(query)
	if err != nil {
		return err
	}
	return pgxscan.ScanOne(response, *row)
}

func (r *pgRepo) FindByConditions(tableName string, conditions []DBCondition, response interface{}) error {
	var conditionsQuery = ""
	for _, condition := range conditions {
		if !strings.EqualFold(conditionsQuery, "") {
			conditionsQuery = conditionsQuery + " AND "
		}
		op, err := parseOperator(condition.Operator)
		if err != nil {
			return err
		}
		conditionsQuery = conditionsQuery + fmt.Sprintf("(%s %s '%v')", condition.FieldName, op, condition.Value)
	}
	var query = fmt.Sprintf("SELECT * FROM %s WHERE %s", tableName, conditionsQuery)
	row, err := r.cnt.RunQueryRows(query)
	if err != nil {
		return err
	}
	return pgxscan.ScanOne(response, *row)
}

func (r *pgRepo) ExistsBy(tableName string, condition DBCondition) (bool, error) {
	op, err := parseOperator(condition.Operator)
	if err != nil {
		return false, err
	}
	var query = fmt.Sprintf(queryExists, tableName, condition.FieldName, op, condition.Value)
	rows, err := r.cnt.RunQueryRows(query)
	if err != nil {
		return false, err
	}
	var exists = struct {
		Exists bool
	}{}
	err = pgxscan.ScanOne(&exists, *rows)
	return exists.Exists, err
}

func (r *pgRepo) InsertInto(tableName string, columnValues []DBValue, response interface{}) error {
	var columnNames []string
	var values []string
	for _, column := range columnValues {
		columnNames = append(columnNames, column.FieldName)
		if value, ok := column.Value.(time.Time); ok {
			values = append(values, fmt.Sprintf("'%v'", value.UTC().Format(formatter.ParseFormatUtc)))
		} else {
			values = append(values, fmt.Sprintf("'%v'", column.Value))
		}
	}
	var query = fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING * ", tableName, strings.Join(columnNames, ","), strings.Join(values, ","))
	row, err := r.cnt.RunQueryRows(query)
	if err != nil {
		return err
	}
	return pgxscan.ScanOne(response, *row)
}

func (r *pgRepo) Update(tableName string, columnValues []DBValue, conditions []DBCondition, response interface{}) error {
	var columnQuery = ""
	for _, column := range columnValues {
		if !strings.EqualFold(columnQuery, "") {
			columnQuery = columnQuery + " , "
		}
		if value, ok := column.Value.(time.Time); ok {
			date := value.UTC().Format(formatter.ParseFormatUtc)
			columnQuery = columnQuery + fmt.Sprintf("%s = '%v'", column.FieldName, date)
		} else {
			columnQuery = columnQuery + fmt.Sprintf("%s = '%v'", column.FieldName, column.Value)
		}
	}
	var conditionsQuery = ""
	for _, condition := range conditions {
		if !strings.EqualFold(conditionsQuery, "") {
			conditionsQuery = conditionsQuery + " AND "
		}
		op, err := parseOperator(condition.Operator)
		if err != nil {
			return err
		}
		conditionsQuery = conditionsQuery + fmt.Sprintf("%s %s %v", condition.FieldName, op, condition.Value)
	}
	var query = fmt.Sprintf("UPDATE %s SET %s WHERE %s ", tableName, columnQuery, conditionsQuery)
	row, err := r.cnt.RunQueryRows(query)
	if err != nil {
		return err
	}
	return pgxscan.ScanOne(response, *row)
}

func parseOperator(op DBOperator) (string, error) {
	switch op {
	case EQUAL:
		return "=", nil
	case NOTEQUAL:
		return "<>", nil
	case GREATER:
		return ">", nil
	case LESS:
		return "<", nil
	case GREATERE:
		return ">=", nil
	case LESSE:
		return "<=", nil
	case LIKE:
		return "LIKE", nil
	case NOTLIKE:
		return "NOT LIKE", nil
	case IN:
		return "IN", nil
	case NOTIN:
		return "NOT IN", nil
	case ISNULL:
		return "IS NULL", nil
	case ISNOTNULL:
		return "IS NOT NULL", nil
	default:
		return "", fmt.Errorf("operator '%s' not suported", op)
	}
}
