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

type NewTitleBodyValidationView struct {
    TitleBodyCustomizerView
    LogUtil *LogUtil
}

// Implements ValidationComponentInterface


func (this *NewTitleBodyValidationView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewNewTitleBodyValidationView(transformInfoInterface TransformInfoInterface) *NewTitleBodyValidationView {
    this := &NewTitleBodyValidationView{}
        this.TitleBodyCustomizerView = *NewTitleBodyCustomizerView(transformInfoInterface)
        this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("View Name: " + transformInfoInterface.GetName(), this, "NewTitleBodyValidationView()")
        }
        this.TitleBody = NewTitleBodyValidation()
    return this
}


// throws Exception
func (this *NewTitleBodyValidationView) ToXmlDoc() *Document {
        return nil
}


func (this *NewTitleBodyValidationView) IsValid() bool {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 bool) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put("Failed to validate", this, this.CommonStrings.IS_VALID, e)
            }
            tryControl1, tryResult1 = 1, false; return
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("View Name: " + this.GetTransformInfoInterface().GetName(), this, this.CommonStrings.IS_VALID)
            }
            {
                CustomizerUtil_GetInstance().Insert(this.GetTransformInfoInterface(), any(this.TitleBody).(DomNodeInterface))
            }
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *NewTitleBodyValidationView) ValidationInfo() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put("Failed to generate validation error info", this, "validationInfo()", e)
            }
            tryControl1, tryResult1 = 1, "Error Getting Validation Info"; return
                }
            }()
            var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
            stringBuffer.Append(this.TitleBody.ValidationInfo())
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *NewTitleBodyValidationView) ToValidationInfoDoc() *Document {
        return nil
}


func (this *NewTitleBodyValidationView) ToValidationInfoNode(document *Document) *Node {
        return nil
}


// throws Exception
func (this *NewTitleBodyValidationView) View() string {
        return this.Views.ValidationOnlyTempUtil.GetInstance().View(this)
}

