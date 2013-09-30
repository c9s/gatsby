package gatsby

import "github.com/c9s/gatsby/sqlfragments"
import "github.com/c9s/gatsby/sqlutils"
import "strings"
import "fmt"

const (
	MODE_SELECT = iota
	MODE_DELETE
	MODE_UPDATE
	MODE_INSERT
)

type ArgMap map[string]interface{}

type Query struct {
	tableName     string
	holderType    int
	mode          int
	selectColumns []string
	whereMap      *ArgMap
	insertMap     *ArgMap
	updateMap     *ArgMap
	fragments     sqlfragments.SQLFragments

	limit  int
	offset int

	arguments []interface{}
}

func NewFragment() *sqlfragments.SQLFragments {
	return new(sqlfragments.SQLFragments)
}

func NewQuery(tableName string) *Query {
	query := new(Query)
	query.tableName = tableName
	query.holderType = sqlutils.QMARK_HOLDER
	return query
}

func (m *Query) Select(columns ...string) *Query {
	m.mode = MODE_SELECT
	m.selectColumns = columns
	return m
}

func (m *Query) Insert(argMap ArgMap) *Query {
	m.mode = MODE_INSERT
	m.insertMap = &argMap
	return m
}

func (m *Query) Update(argMap ArgMap) *Query {
	m.mode = MODE_UPDATE
	m.updateMap = &argMap
	return m
}

func (m *Query) WhereFromMap(argMap ArgMap) *Query {
	m.whereMap = &argMap
	return m
}

func (m *Query) Limit(offset, limit int) *Query {
	m.offset = offset
	m.limit = limit
	return m
}

func (m *Query) Args() []interface{} {
	return m.arguments
}

func (m *Query) String() string {
	// build for select
	switch m.mode {
	case MODE_SELECT:

		var sql string = "SELECT " + strings.Join(m.selectColumns, ", ") + " FROM " + m.tableName
		if m.whereMap != nil {
			whereSql, args := sqlutils.BuildWhereInnerClause(*m.whereMap, "AND", m.holderType)
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

	case MODE_DELETE:

		var sql = "DELETE FROM " + m.tableName

		if m.whereMap != nil {
			whereSql, args := sqlutils.BuildWhereInnerClause(*m.whereMap, "AND", m.holderType)
			sql += " WHERE " + whereSql
			m.arguments = append(m.arguments, args...)
		}
		return sql

	case MODE_UPDATE:
		var sql = "UPDATE " + m.tableName + " SET "
		var updateSql, args = sqlutils.BuildUpdateColumnsFromMap(*m.updateMap)
		sql += updateSql
		m.arguments = append(m.arguments, args...)
		if m.whereMap != nil {
			whereSql, args := sqlutils.BuildWhereInnerClause(*m.whereMap, "AND", m.holderType)
			sql += " WHERE " + whereSql
			m.arguments = append(m.arguments, args...)
		}
		return sql

	case MODE_INSERT:

		var sql string = "INSERT INTO " + m.tableName
		var insertSql, args = sqlutils.BuildInsertColumnsFromMap(*m.insertMap, m.holderType)
		sql += " " + insertSql
		m.arguments = append(m.arguments, args...)
		return sql

	default:
		panic("Unsupported mode")
	}
	return ""
}
