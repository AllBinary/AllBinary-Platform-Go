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

type EntryData struct {
    ID string
    ENCRYPTION string
    TIMECREATED string
    LASTMODIFIED string
    DEFAULT string
    SPECIAL string
    ENABLE string
    YES string
    NO string
}

var EntryData_Instance *EntryData = NewEntryData()

func (this *EntryData) initFields() {
    this.ID = "ENTRY_ID"
    this.ENCRYPTION = "ENTRY_ENCRYPTION"
    this.TIMECREATED = "ENTRY_TIMECREATED"
    this.LASTMODIFIED = "ENTRY_LASTMODIFIED"
    this.DEFAULT = "ENTRY_DEFAULT"
    this.SPECIAL = "ENTRY_SPECIAL"
    this.ENABLE = "ENTRY_ENABLE"
    this.YES = "Yes"
    this.NO = "No"
}


func EntryData_GetInstance() *EntryData {
        return EntryData_Instance
}


func NewEntryData() *EntryData {
    this := &EntryData{}
    this.initFields()
    return this
}

