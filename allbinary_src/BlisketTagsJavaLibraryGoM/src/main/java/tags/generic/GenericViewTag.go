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

type GenericViewTag struct {
    TransformTag
    LogUtil *LogUtil
    Output string
    File string
}


func (this *GenericViewTag) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
}


func NewGenericViewTag() *GenericViewTag {
    this := &GenericViewTag{}
    this.TransformTag = *NewTransformTag()
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        }
    return this
}


func (this *GenericViewTag) SetOutput(value string) {
        this.Output = value
        this.GetPropertiesHashMap().Put(InputOutputTypeData_GetInstance().NAME, this.Output)
}


func (this *GenericViewTag) SetFile(value string) {
        this.File = value
        this.GetPropertiesHashMap().Put(InputOutputTypeData_GetInstance().FILE, this.File)
}

