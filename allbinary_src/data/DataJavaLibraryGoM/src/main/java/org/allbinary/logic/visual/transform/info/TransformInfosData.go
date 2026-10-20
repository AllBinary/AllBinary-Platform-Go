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

type TransformInfosData struct {
    NAME string
    GROUP string
    ALL string
    PREVIEW string
    SMALL_PREVIEW string
}

var TransformInfosData_Instance *TransformInfosData = NewTransformInfosData()

func (this *TransformInfosData) initFields() {
    this.NAME = "TRANSFORM_INFOS_NAME"
    this.GROUP = "TRANSFORM_INFOS_GROUP"
    this.ALL = "TRANSFORM_INFOS_ALL"
    this.PREVIEW = "TRANSFORM_INFOS_PREVIEW"
    this.SMALL_PREVIEW = "TRANSFORM_INFOS_SMALL_PREVIEW"
}


func NewTransformInfosData() *TransformInfosData {
    this := &TransformInfosData{}
    this.initFields()
    return this
}


func TransformInfosData_GetInstance() *TransformInfosData {
        return TransformInfosData_Instance
}

