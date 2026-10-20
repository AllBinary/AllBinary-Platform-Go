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

type SqlTypeStrings struct {
    CHAR_COLUMN string
    INT_COLUMN string
    MAX_SIZE string
    SIXTY_SIZE string
    NOT_NULL string
    ONE_SIZE string
    TWO_SIZE string
    THREE_SIZE string
    SIX_SIZE string
    ELEVEN_SIZE string
    TWELVE_SIZE string
    ONE_KB_CHAR_COLUMN string
    MAX_CHAR_COLUMN string
    MAX_CHAR_COLUMN_NOT_NULL string
    SIXTY_CHAR_COLUMN_NOT_NULL string
    ONE_CHAR_COLUMN_NOT_NULL string
    SIX_CHAR_COLUMN_NOT_NULL string
    TWELVE_CHAR_COLUMN_NOT_NULL string
    TWO_INT_NOT_NULL string
    THREE_INT_NOT_NULL string
    FOUR_INT_NOT_NULL string
    MAX_INT_NOT_NULL string
    MAX_INT_UNSIGNED string
    MAX_INT_UNSIGNED_NOT_NULL string
    MAX_BIG_INT_NOT_NULL string
    MAX_BIG_INT_UNSIGNED string
    MAX_BIG_INT_UNSIGNED_NOT_NULL string
    MAX_BIG_INT_UNSIGNED_AUTO_INCREMENT_NOT_NULL string
    LONG_BLOB string
    BLOB string
    BLOB_NOT_NULL string
}

var SqlTypeStrings_Instance *SqlTypeStrings = NewSqlTypeStrings()

func (this *SqlTypeStrings) initFields() {
    this.CHAR_COLUMN = "VARCHAR"
    this.INT_COLUMN = "VARCHAR"
    this.MAX_SIZE = "(255)"
    this.SIXTY_SIZE = "(60)"
    this.NOT_NULL = "NOT NULL"
    this.ONE_SIZE = "(1)"
    this.TWO_SIZE = "(2)"
    this.THREE_SIZE = "(3)"
    this.SIX_SIZE = "(6)"
    this.ELEVEN_SIZE = "(11)"
    this.TWELVE_SIZE = "(12)"
    this.ONE_KB_CHAR_COLUMN = fmt.Sprint(NewStringMaker().Append(CommonSeps_GetInstance().SPACE).Append(this.CHAR_COLUMN).Append("(1024)").Append(CommonSeps_GetInstance().COMMA_SEP))
    this.MAX_CHAR_COLUMN = fmt.Sprint(NewStringMaker().Append(CommonSeps_GetInstance().SPACE).Append(this.CHAR_COLUMN).Append(this.MAX_SIZE).Append(CommonSeps_GetInstance().COMMA_SEP))
    this.MAX_CHAR_COLUMN_NOT_NULL = fmt.Sprint(NewStringMaker().Append(CommonSeps_GetInstance().SPACE).Append(this.CHAR_COLUMN).Append(this.MAX_SIZE).Append(CommonSeps_GetInstance().SPACE).Append(this.NOT_NULL).Append(CommonSeps_GetInstance().COMMA_SEP))
    this.SIXTY_CHAR_COLUMN_NOT_NULL = fmt.Sprint(NewStringMaker().Append(CommonSeps_GetInstance().SPACE).Append(this.CHAR_COLUMN).Append(this.SIXTY_SIZE).Append(this.NOT_NULL).Append(CommonSeps_GetInstance().COMMA_SEP))
    this.ONE_CHAR_COLUMN_NOT_NULL = fmt.Sprint(NewStringMaker().Append(CommonSeps_GetInstance().SPACE).Append(this.CHAR_COLUMN).Append(this.ONE_SIZE).Append(this.NOT_NULL).Append(CommonSeps_GetInstance().COMMA_SEP))
    this.SIX_CHAR_COLUMN_NOT_NULL = fmt.Sprint(NewStringMaker().Append(CommonSeps_GetInstance().SPACE).Append(this.CHAR_COLUMN).Append(this.SIX_SIZE).Append(this.NOT_NULL).Append(CommonSeps_GetInstance().COMMA_SEP))
    this.TWELVE_CHAR_COLUMN_NOT_NULL = fmt.Sprint(NewStringMaker().Append(CommonSeps_GetInstance().SPACE).Append(this.CHAR_COLUMN).Append(this.TWELVE_SIZE).Append(this.NOT_NULL).Append(CommonSeps_GetInstance().COMMA_SEP))
    this.TWO_INT_NOT_NULL = " INT(2) NOT NULL, "
    this.THREE_INT_NOT_NULL = " INT(3) NOT NULL, "
    this.FOUR_INT_NOT_NULL = " INT(4) NOT NULL, "
    this.MAX_INT_NOT_NULL = " INT(11) NOT NULL, "
    this.MAX_INT_UNSIGNED = " INT(11) UNSIGNED, "
    this.MAX_INT_UNSIGNED_NOT_NULL = " INT(11) UNSIGNED NOT NULL, "
    this.MAX_BIG_INT_NOT_NULL = " BIGINT(19) NOT NULL, "
    this.MAX_BIG_INT_UNSIGNED = " BIGINT(19) UNSIGNED, "
    this.MAX_BIG_INT_UNSIGNED_NOT_NULL = " BIGINT(19) UNSIGNED NOT NULL, "
    this.MAX_BIG_INT_UNSIGNED_AUTO_INCREMENT_NOT_NULL = " BIGINT(19) UNSIGNED AUTO_INCREMENT NOT NULL,"
    this.LONG_BLOB = " LONGBLOB, "
    this.BLOB = " BLOB, "
    this.BLOB_NOT_NULL = " BLOB NOT NULL, "
}

func NewSqlTypeStrings() *SqlTypeStrings {
    this := &SqlTypeStrings{}
    this.initFields()
    return this
}


func SqlTypeStrings_GetInstance() *SqlTypeStrings {
        return SqlTypeStrings_Instance
}

