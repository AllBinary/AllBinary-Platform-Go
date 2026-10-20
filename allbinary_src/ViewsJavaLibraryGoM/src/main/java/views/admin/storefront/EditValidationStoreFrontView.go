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

type EditValidationStoreFrontView struct {
    StoreFrontComponent
    LogUtil *LogUtil
}

// Implements ValidationComponentInterface


func (this *EditValidationStoreFrontView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewEditValidationStoreFrontView(transformInfoInterface TransformInfoInterface) *EditValidationStoreFrontView {
    this := &EditValidationStoreFrontView{}
        this.StoreFrontComponent = *NewStoreFrontComponent(transformInfoInterface)
        this.initFields()
    return this
}


func (this *EditValidationStoreFrontView) IsValid() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put("Failed to validate form", this, this.CommonStrings.IS_VALID, e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            var valid bool = true; _ = valid
            if this.NewStoreFrontInterface.IsNameValid() == false {
                valid = false
            }
            if StoreFrontFactory_GetInstance(this.NewStoreFrontInterface.GetName()) == nil {
                valid = false
            }
            this.NewStoreFrontInterface = StoreFrontFactory_GetInstance(this.NewStoreFrontInterface.GetName())
            tryControl1, tryResult1 = 1, valid; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *EditValidationStoreFrontView) ValidationInfo() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put("Failed to generate validation error info", this, "validationInfo()", e)
            }
            tryControl1, tryResult1 = 1, "Error Validating Form"; return
                }
            }()
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            if this.NewStoreFrontInterface.IsNameValid() == false {
                stringBuffer.Append(this.NewStoreFrontInterface.NameValidationInfo())
            }
            if StoreFrontFactory_GetInstance(this.NewStoreFrontInterface.GetName()) == nil {
                stringBuffer.Append(fmt.Sprint(this.NewStoreFrontInterface) + " does not exist.<br>")
            }
            tryControl1, tryResult1 = 1, "Unkown Error"; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *EditValidationStoreFrontView) ToValidationInfoDoc() *Document {
        return nil
}


func (this *EditValidationStoreFrontView) ToValidationInfoNode(document *Document) *Node {
        return nil
}

