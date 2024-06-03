package pgparser

import (
	"strings"

	pgQuery "github.com/pganalyze/pg_query_go/v5"
)

// smallint 			2 bytes 	small-range integer					-32768 to +32767
// integer 				4 bytes 	typical choice for integer			-2147483648 to +2147483647
// bigint 				8 bytes 	large-range integer					-9223372036854775808 to +9223372036854775807
// decimal 				variable	user-specified precision, exact 	up to 131072 digits before the decimal point; up to 16383 digits after the decimal point
// numeric 				variable	user-specified precision, exact 	up to 131072 digits before the decimal point; up to 16383 digits after the decimal point
// real 				4 bytes 	variable-precision, inexact 		6 decimal digits precision
// double precision 	8 bytes 	variable-precision, inexact 		15 decimal digits precision
// smallserial 			2 bytes 	small autoincrementing integer 		1 to 32767
// serial 				4 bytes 	autoincrementing integer 			1 to 2147483647
// bigserial 			8 bytes 	large autoincrementing integer 		1 to 9223372036854775807
func convertPostgresType(pgType string, isArray bool) PgDataType {
	lowerType := strings.ToLower(pgType)

	switch lowerType {
	case "int2", "smallint", "serial2":
		return typeOrArray(Smallint, SmallintArray, isArray)

	case "int4", "int", "serial", "serial4":
		return typeOrArray(Int, IntArray, isArray)

	case "bigint", "bigserial", "serial8":
		return typeOrArray(Bigint, BigintArray, isArray)

	case "float", "float8", "double precision":
		return typeOrArray(Double, DoubleArray, isArray)

	case "real", "float4":
		return typeOrArray(Float, FloatArray, isArray)

	case "varchar", "text", "character varying":
		return typeOrArray(Varchar, VarcharArray, isArray)

	case "boolean", "bool":
		return typeOrArray(Boolean, BooleanArray, isArray)

	case "timestamp", "timestamptz",
		"timestamp with time zone",
		"timestamp without time zone":
		return typeOrArray(Timestamp, TimestampArray, isArray)
	case "date":
		return typeOrArray(Date, DateArray, isArray)
	default:
		return typeOrArray(Any, AnyArray, isArray)
	}
}

func typeOrArray(ty, tyArray PgDataType, isArray bool) PgDataType {
	if isArray {
		return tyArray
	}

	return ty
}

func Cast[T *pgQuery.Node, E any](in T) E {

	// if result, ok := in.(E); ok {
	// 	return result
	// }

	var none E
	return none
}
