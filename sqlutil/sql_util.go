package sqlutil

import (
	"strconv"
	"time"

	"github.com/warrior21st/go-utils/timeutil"
)

//string数组转sql条件string
func StringsToSqlCondition(args *[]string) string {
	sqlCon := ""
	for i := range *args {
		sqlCon += "'" + (*args)[i] + "',"
	}
	if sqlCon != "" {
		sqlCon = sqlCon[0 : len(sqlCon)-1]
	}

	return sqlCon
}

//int数组转sql条件string
func IntToSqlCondition(args *[]int) string {
	sqlCon := ""
	for i := range *args {
		sqlCon += strconv.Itoa((*args)[i]) + ","
	}
	if sqlCon != "" {
		sqlCon = sqlCon[0 : len(sqlCon)-1]
	}

	return sqlCon
}

//int32数组转sql条件string
func Int32sToSqlCondition(args *[]int32) string {
	sqlCon := ""
	for i := range *args {
		sqlCon += strconv.FormatInt(int64((*args)[i]), 10) + ","
	}
	if sqlCon != "" {
		sqlCon = sqlCon[0 : len(sqlCon)-1]
	}

	return sqlCon
}

//int64数组转sql条件string
func Int64sToSqlCondition(args *[]int64) string {
	sqlCon := ""
	for i := range *args {
		sqlCon += strconv.FormatInt((*args)[i], 10) + ","
	}
	if sqlCon != "" {
		sqlCon = sqlCon[0 : len(sqlCon)-1]
	}

	return sqlCon
}

//float32数组转sql条件string
func Float32ToSqlCondition(args *[]float32) string {
	sqlCon := ""
	for i := range *args {
		sqlCon += strconv.FormatFloat(float64((*args)[i]), 'f', -1, 64) + ","
	}
	if sqlCon != "" {
		sqlCon = sqlCon[0 : len(sqlCon)-1]
	}

	return sqlCon
}

//float64数组转sql条件string
func Float64sToSqlCondition(args *[]float64) string {
	sqlCon := ""
	for i := range *args {
		sqlCon += strconv.FormatFloat((*args)[i], 'f', -1, 64) + ","
	}
	if sqlCon != "" {
		sqlCon = sqlCon[0 : len(sqlCon)-1]
	}

	return sqlCon
}

//时间字符串转为oracle to_date形式的字符串
func TimeStrToOracleSqlStr(timeStr string) string {
	return "to_date('" + timeutil.StandardTimeString(timeStr) + "', 'yyyy-MM-dd hh24:mi:ss')"
}

//时间转为oracle to_date形式的字符串
func TimeToOracleSqlStr(t *time.Time) string {
	return TimeStrToOracleSqlStr(timeutil.TimeToString(*t))
}
