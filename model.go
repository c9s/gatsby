package gatsby

import "gatsby/sqlfragments"
import "gatsby/sqlutils"
import "strings"
import "fmt"

const (
	MODE_SELECT = iota
	MODE_DELETE
	MODE_UPDATE
	MODE_INSERT
)


type ArgMap map[string]interface{}

type Model struct {
	tableName     string
	mode          int
	selectColumns []string
	whereMap      *ArgMap
	insertMap     *ArgMap
	fragments     sqlfragments.SQLFragments

	limit  int
	offset int

	arguments []interface{}
}

func NewModel(tableName string) *Model {
	model := new(Model)
	model.tableName = tableName
	return model
}

func (m *Model) Select(columns ...string) *Model {
	m.mode = MODE_SELECT
	m.selectColumns = columns
	return m
}

func (m *Model) Insert(argMap ArgMap) *Model {
	m.mode = MODE_INSERT
	m.insertMap = &argMap
	return m
}

func (m *Model) WhereFromMap(argMap ArgMap) *Model {
	m.whereMap = &argMap
	return m
}

func (m *Model) Limit(offset, limit int) *Model {
	m.offset = offset
	m.limit = limit
	return m
}

func (m *Model) String() string {
	// build for select
	switch m.mode {
	case MODE_SELECT:
		var sql string = "SELECT " + strings.Join(m.selectColumns, ", ") + " FROM " + m.tableName
		if m.whereMap != nil {
			whereSql, args := sqlutils.BuildWhereInnerClause(*m.whereMap, "AND")
			sql += " WHERE " + whereSql
			m.arguments = append(m.arguments, args...)
		}
		if m.limit > 0 {
			sql += fmt.Sprintf(" LIMIT %d ", m.limit)
		}
		if m.offset > 0 {
			sql += fmt.Sprintf(" OFFSET %d", m.offset)
		}
		return sql
	case MODE_INSERT:
		var sql string = "INSERT INTO " + m.tableName
		var insertSql, args = sqlutils.BuildInsertColumnsFromMap(*m.insertMap)
		sql += " " + insertSql
		m.arguments = append(m.arguments, args...)
		return sql
	default:
		panic("Unsupported mode")
	}
	return ""
}
