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

type HistoryDbInitInfo struct {
    DbInitInfo
}

var HistoryDbInitInfo_INITFILE string = "historydbinitdata.dat"


func NewHistoryDbInitInfo() *HistoryDbInitInfo {
    this := &HistoryDbInitInfo{}
        this.DbInitInfo = *NewDbInitInfo(HistoryDbInitInfo_INITFILE, true)
    return this
}


// Java overloaded constructor: Go needs a different name
func NewHistoryDbInitInfo(read bool) *HistoryDbInitInfo {
    this := &HistoryDbInitInfo{}
        this.DbInitInfo = *NewDbInitInfo(HistoryDbInitInfo_INITFILE, read)
    return this
}

