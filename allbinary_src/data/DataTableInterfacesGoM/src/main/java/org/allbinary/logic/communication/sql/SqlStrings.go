// Generated Code Do Not Modify

package allbinary

import (
    "errors"
    "fmt"
    "math"
    "math/rand"
    "os"
    "reflect"
    "runtime"
    "strconv"
    "strings"
    "time"
    "unicode"
)

var _ = errors.New
var _ = fmt.Sprint
var _ = math.Abs
var _ = rand.Float64
var _ = os.Exit
var _ = reflect.TypeOf
var _ = runtime.GC
var _ = strconv.Itoa
var _ = strings.Contains
var _ = time.Now
var _ = unicode.IsDigit

type SqlStrings struct {
    CREATE_DATABASE string
    DROP_TABLE string
    CREATE_TABLE string
    START string
    SELECT string
    FROM string
    SELECT_ALL string
    WHERE string
    AND string
    UPDATE string
    SET string
    SELECT_ALL_FROM string
    DELETE string
    LIKE_QUOTE string
    INSERT_INTO string
    VALUES string
    ORDER_BY string
    COUNT string
    ASC string
    DESC string
    PRIMARY_KEY string
    END string
    EQUAL_QUOTE string
    CLOSE_QUOTE string
    EQUAL_QUOTE_NO_SPACE string
    ESCAPE string
    DOUBLE_ESCAPE string
    MORE_THAN_QUOTE string
    LESS_THAN_QUOTE string
    SINGLE_QUOTE_COMMA_SEP string
    ID string
    SQL_STATEMENT_LABEL string
    COLUMN_VALUE string
    CREATE_RETURN string
    FIELD_VALUE string
}

var SqlStrings_Instance *SqlStrings = NewSqlStrings()

func (this *SqlStrings) initFields() {
    this.CREATE_DATABASE = "CREATE DATABASE "
    this.DROP_TABLE = "DROP TABLE "
    this.CREATE_TABLE = "CREATE TABLE "
    this.START = " ("
    this.SELECT = "SELECT "
    this.FROM = " FROM "
    this.SELECT_ALL = "SELECT *"
    this.WHERE = " WHERE "
    this.AND = " AND "
    this.UPDATE = "UPDATE "
    this.SET = " SET "
    this.SELECT_ALL_FROM = this.SELECT_ALL + this.FROM
    this.DELETE = "DELETE"
    this.LIKE_QUOTE = " LIKE \""
    this.INSERT_INTO = "INSERT INTO "
    this.VALUES = " VALUES ('"
    this.ORDER_BY = "ORDER BY"
    this.COUNT = "COUNT("
    this.ASC = "ASC"
    this.DESC = "DESC"
    this.PRIMARY_KEY = "PRIMARY KEY("
    this.END = ") )"
    this.EQUAL_QUOTE = " = \""
    this.CLOSE_QUOTE = CommonSeps_GetInstance().QUOTE
    this.EQUAL_QUOTE_NO_SPACE = "=\""
    this.ESCAPE = "\\"
    this.DOUBLE_ESCAPE = "\\\\"
    this.MORE_THAN_QUOTE = " > \""
    this.LESS_THAN_QUOTE = " < \""
    this.SINGLE_QUOTE_COMMA_SEP = "','"
    this.ID = "ID"
    this.SQL_STATEMENT_LABEL = "SQL Statement: "
    this.COLUMN_VALUE = "\nColumn Value: "
    this.CREATE_RETURN = " Created Successfully"
    this.FIELD_VALUE = "\nField Value: "
}

func NewSqlStrings() *SqlStrings {
    this := &SqlStrings{}
    this.initFields()
    return this
}


func SqlStrings_GetInstance() *SqlStrings {
        return SqlStrings_Instance
}

