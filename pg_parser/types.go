package pgparser

type ConstrType int32

const (
	CONSTR_TYPE_UNDEFINED ConstrType = iota
	CONSTR_NULL
	CONSTR_NOTNULL
	CONSTR_DEFAULT
	CONSTR_IDENTITY
	CONSTR_GENERATED
	CONSTR_CHECK
	CONSTR_PRIMARY
	CONSTR_UNIQUE
	CONSTR_EXCLUSION
	CONSTR_FOREIGN
	CONSTR_ATTR_DEFERRABLE
	CONSTR_ATTR_NOT_DEFERRABLE
	CONSTR_ATTR_DEFERRED
	CONSTR_ATTR_IMMEDIATE
)

type PgDataType uint32

const (
	Smallint PgDataType = iota
	Int
	Bigint
	Double
	Float
	Varchar
	Boolean
	Timestamp
	Date
	Any

	SmallintArray
	IntArray
	BigintArray
	DoubleArray
	FloatArray
	VarcharArray
	BooleanArray
	TimestampArray
	DateArray
	AnyArray
)

var goDataType = [...]string{
	"int16",
	"int",
	"int64",
	"float64",
	"float32",
	"string",
	"bool",
	"time.Time",
	"time.Time",
	"any",

	"[]int16",
	"[]int",
	"[]int64",
	"[]float64",
	"[]float32",
	"[]string",
	"[]bool",
	"[]time.Time",
	"[]time.Time",
	"[]any",
}

func (p PgDataType) Go() string {
	return goDataType[p]
}

//TODO:other language
// func (p PgDataType) Java() string {
// 	return javaDataType[p]
// }
