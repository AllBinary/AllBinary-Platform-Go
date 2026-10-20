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

type BodyData struct {
    NAME string
    VIEWNAMEKEY string
    MIN int
}

var BodyData_Instance *BodyData = NewBodyData()

func (this *BodyData) initFields() {
    this.NAME = "BODY_NAME"
    this.VIEWNAMEKEY = "Body"
    this.MIN = 0
}


func NewBodyData() *BodyData {
    this := &BodyData{}
    this.initFields()
    return this
}


func BodyData_GetInstance() *BodyData {
        return BodyData_Instance
}

