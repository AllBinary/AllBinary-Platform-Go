/*
 *
 * AllBinary Open License Version 1
 * Copyright (c) 2011 AllBinary
 * 
 * By agreeing to this license you and any business entity you represent are
 * legally bound to the AllBinary Open License Version 1 legal agreement.
 * 
 * You may obtain the AllBinary Open License Version 1 legal agreement from
 * AllBinary or the root directory of AllBinary's AllBinary Platform repository.
 * 
 * Created By: Travis Berthelot
 *  
 */
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

type LogData struct {
    MAX int64
    Extension string
    BackupFileExt string
    ALLBINARY string
    NAME string
    DESCRIPTION string
}


func (this *LogData) initFields() {
    this.MAX = 5000000
    this.Extension = "log"
    this.BackupFileExt = fmt.Sprint(AbPathData_GetInstance().EXTENSION_SEP) + "bak" + fmt.Sprint(AbPathData_GetInstance().EXTENSION_SEP) + "txt"
    this.ALLBINARY = "allbinary"
    this.NAME = "LOG_NAME"
    this.DESCRIPTION = "LOG_DESCRIPTION"
}

func NewLogData() *LogData {
    this := &LogData{}
    this.initFields()
    return this
}

