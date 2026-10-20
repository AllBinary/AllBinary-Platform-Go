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

type TransformInfoData struct {
    PARENT string
    MAPPED string
    NAME string
    DESCRIPTION string
    CATEGORY string
    OBJECTFILENAME string
    OBJECT string
    OBJECTCONFIGFILENAME string
    OBJECTCONFIG string
    TEMPLATEFILENAME string
    TEMPLATE string
    DATAFILENAME string
    DATA string
    LABEL string
    OWNER string
    PARTIAL string
}

var TransformInfoData_Instance *TransformInfoData = NewTransformInfoData()

func (this *TransformInfoData) initFields() {
    this.PARENT = "TRANSFORM_INFO_PARENT"
    this.MAPPED = "TRANSFORM_INFO_MAPPED"
    this.NAME = "TRANSFORM_INFO_NAME"
    this.DESCRIPTION = "TRANSFORM_INFO_DESCRIPTION"
    this.CATEGORY = "TRANSFORM_INFO_CATEGORY"
    this.OBJECTFILENAME = "TRANSFORM_INFO_OBJECTFILENAME"
    this.OBJECT = "TRANSFORM_INFO_OBJECT"
    this.OBJECTCONFIGFILENAME = "TRANSFORM_INFO_OBJECTCONFIGFILENAME"
    this.OBJECTCONFIG = "TRANSFORM_INFO_OBJECTCONFIG"
    this.TEMPLATEFILENAME = "TRANSFORM_INFO_TEMPLATEFILENAME"
    this.TEMPLATE = "TRANSFORM_INFO_TEMPLATE"
    this.DATAFILENAME = "TRANSFORM_INFO_DATAFILENAME"
    this.DATA = "TRANSFORM_INFO_DATA"
    this.LABEL = "TRANSFORM_INFO_LABEL"
    this.OWNER = "OWNER_" + this.NAME
    this.PARTIAL = "PARTIAL_" + this.NAME
}


func TransformInfoData_GetInstance() *TransformInfoData {
        return TransformInfoData_Instance
}


func NewTransformInfoData() *TransformInfoData {
    this := &TransformInfoData{}
    this.initFields()
    return this
}

