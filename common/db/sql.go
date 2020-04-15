package db

import (
	"bytes"
	"fmt"
	"strings"

	gorp "gopkg.in/gorp.v1"
)

const (
	sqlDefaultSep    = ","
	placeholderLimit = 35000
)

// 创建 sql in 查询条件
func SqlPlaceholds(count int) string {
	return appendDuplicateString("?", sqlDefaultSep, count)
}

// sql 多字段查询／插入条件
func SqlMultiColumnPlaceholders(columnCount int, count int) string {
	if columnCount <= 0 || count <= 0 {
		return ""
	}
	one := fmt.Sprintf("(%s)", appendDuplicateString("?", sqlDefaultSep, columnCount))
	return appendDuplicateString(one, sqlDefaultSep, count)
}

// 创建Sql 多个or 查询条件
func SqlMultiCondition(singleCondition string, count int) string {
	return appendDuplicateString(singleCondition, " or ", count)
}

// 创建 sql 批量插入条件
func SqlInMultiInsertValues(columnCount int, count int) string {
	return SqlMultiColumnPlaceholders(columnCount, count)
}

func appendDuplicateString(character, separator string, count int) string {
	if count <= 0 {
		return ""
	}
	var b bytes.Buffer
	for i := 0; i < count; i++ {
		if i > 0 {
			b.WriteString(separator)
		}
		b.WriteString(character)
	}
	return b.String()
}

func BuildSqlArgs(args ...interface{}) ([]interface{}, error) {
	newArgs := make([]interface{}, 0)
	addEleFun := func(ele interface{}) {
		newArgs = append(newArgs, ele)
		return
	}
	for _, arg := range args {
		switch v := arg.(type) {
		case string, int, int32, int64, bool, *string, *int, *int32, *int64:
			addEleFun(v)
		case []string:
			for _, e := range v {
				addEleFun(e)
			}
		case []int:
			for _, e := range v {
				addEleFun(e)
			}
		case []int32:
			for _, e := range v {
				addEleFun(e)
			}
		case []int64:
			for _, e := range v {
				addEleFun(e)
			}
		case []*string:
			for _, e := range v {
				addEleFun(e)
			}
		default:
			return nil, fmt.Errorf("Invalid Arg Type: %+v", arg)
		}
	}
	return newArgs, nil
}

func SqlColumns(columns ...string) string {
	var b bytes.Buffer
	for i, column := range columns {
		if i > 0 {
			b.WriteString(", ")
		}
		shouldEscape := !strings.ContainsAny(column, "`.() ")
		if shouldEscape {
			b.WriteByte('`')
		}
		b.WriteString(column)
		if shouldEscape {
			b.WriteByte('`')
		}
	}
	return b.String()
}

type BuildArgsByIndexFunc func(index int) []interface{}

// 分片批量插入，提高效率的同时避免 MySQL 的 placeholder 不能超过 65535 个的问题
func SqlBatchInsert(
	tx gorp.SqlExecutor,
	table string,
	columns []string,
	ignore bool,
	count int,
	buildArgs BuildArgsByIndexFunc) error {
	cc := len(columns)
	if count <= 0 || cc == 0 {
		return nil
	}
	limit := placeholderLimit / cc

	var ignoreStmt string
	if ignore {
		ignoreStmt = "IGNORE"
	}
	columnStmt := SqlColumns(columns...)
	tpl := "INSERT %s INTO `%s`(%s) VALUES %s;"
	tpl = fmt.Sprintf(tpl, ignoreStmt, table, columnStmt, "%s")

	for i := 0; i < count; i += limit {
		sliceCount := limit
		if i+sliceCount > count {
			sliceCount = count - i
		}
		query := fmt.Sprintf(tpl, SqlMultiColumnPlaceholders(cc, sliceCount))

		args := make([]interface{}, 0, cc*sliceCount)
		maxIndex := i + sliceCount
		for index := i; index < maxIndex; index++ {
			rowargs := buildArgs(index)
			if len(rowargs) != cc {
				return fmt.Errorf("argument count mismatch: expected %d, got %d", cc, len(rowargs))
			}
			args = append(args, rowargs...)
		}

		_, err := tx.Exec(query, args...)
		if err != nil {
			return err
		}
	}
	return nil
}
