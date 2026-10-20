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

type InsertThemeValidationView struct {
    ThemeCustomizerView
    LogUtil *LogUtil
    RequestHashMap *HashMap
}

// Implements ValidationComponentInterface


func (this *InsertThemeValidationView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewInsertThemeValidationView(transformInfoInterface TransformInfoInterface) *InsertThemeValidationView {
    this := &InsertThemeValidationView{}
        this.ThemeCustomizerView = *NewThemeCustomizerView(transformInfoInterface)
        this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF(this.CommonStrings.CONSTRUCTOR, this, "started")
        }
        this.RequestHashMap = NewNameSpaceRequestParams(this.GetPageContext()).ToHashMap()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF(this.CommonStrings.CONSTRUCTOR, this, "begin loading theme")
        }
    return this
}


// throws Exception
func (this *InsertThemeValidationView) ToXmlDoc() *Document {
        return nil
}


func (this *InsertThemeValidationView) IsValid() bool {
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
            var isValid bool = true; _ = isValid
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("Started Validation", this, this.CommonStrings.IS_VALID)
            }
            var themeValidation *ThemeValidation = NewThemeValidation(this.GetTransformInfoInterface(), this.RequestHashMap); _ = themeValidation
            if themeValidation.IsValid() == false {
                isValid = false
            }
            this.ValidationInterface = themeValidation.GetCssStyleValidation()
            isValid = this.ValidationInterface.IsValid()
            if isValid == true {
                CustomizerUtil_GetInstance().Insert(this.GetTransformInfoInterface(), any(this.ValidationInterface).(DomNodeInterface))
            }
            tryControl1, tryResult1 = 1, isValid; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *InsertThemeValidationView) ValidationInfo() string {
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
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *InsertThemeValidationView) ToValidationInfoDoc() *Document {
        return nil
}


func (this *InsertThemeValidationView) ToValidationInfoNode(document *Document) *Node {
        return nil
}


// throws Exception
func (this *InsertThemeValidationView) View() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "view()", e)
            }
            panic(e)
                }
            }()
            tryControl1, tryResult1 = 1, StoreCustomizerComponentUtil_GetInstance().Generate(this.AbeClientInformation, this.GetTransformInfoInterface()); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

