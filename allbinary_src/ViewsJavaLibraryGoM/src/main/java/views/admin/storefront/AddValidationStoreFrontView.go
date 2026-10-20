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

type AddValidationStoreFrontView struct {
    HttpContextView
    LogUtil *LogUtil
    NewStoreFrontInterface StoreFrontInterface
}

// Implements ValidationComponentInterface


func (this *AddValidationStoreFrontView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewAddValidationStoreFrontView(transformInfoInterface TransformInfoInterface) *AddValidationStoreFrontView {
    this := &AddValidationStoreFrontView{}
        this.HttpContextView = *NewHttpContextView(transformInfoInterface)
        this.initFields()
        this.NewStoreFrontInterface = any(NewStoreFront(any(this.GetPageContext().GetRequest()).(*HttpServletRequest))).(StoreFrontInterface)
        if this.NewStoreFrontInterface.GetName() == "" {
            this.NewStoreFrontInterface.SetName(this.GetWeblisketSession().GetStoreName())
        }
    return this
}


func (this *AddValidationStoreFrontView) IsValid() bool {
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
            if this.NewStoreFrontInterface.IsValid() == false {
                valid = false
            }
            if StoreFrontFactory_GetInstance(this.NewStoreFrontInterface.GetName()) != nil {
                valid = false
            }
            if AbFile_CreateAbFileFromAbPath(NewAbPath(this.GetStoreViewsPath(), StringUtil_GetInstance().EMPTY_STRING)).IsFile() {
                valid = false
            }
            tryControl1, tryResult1 = 1, valid; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws Exception
func (this *AddValidationStoreFrontView) GetStoreViewsPath() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(HttpContextView_URLGLOBALS.GetMainPath())
        stringBuffer.Append(HttpContextView_FREEBLISKET_PATH_GLOBALS.GetInstance().XSLPATH)
        stringBuffer.Append(this.NewStoreFrontInterface.GetName())
        return fmt.Sprint(stringBuffer)
}


func (this *AddValidationStoreFrontView) ValidationInfo() string {
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
            if this.NewStoreFrontInterface.IsValid() == false {
                stringBuffer.Append(this.NewStoreFrontInterface.ValidationInfo())
            }
            if StoreFrontFactory_GetInstance(this.NewStoreFrontInterface.GetName()) != nil {
                stringBuffer.Append("Store name already used<br/>")
            }
            if AbFile_CreateAbFileFromAbPath(NewAbPath(this.GetStoreViewsPath(), StringUtil_GetInstance().EMPTY_STRING)).IsFile() {
                stringBuffer.Append("Store name clashes with template name<br/>")
            }
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *AddValidationStoreFrontView) ToValidationInfoDoc() *Document {
        return nil
}


func (this *AddValidationStoreFrontView) ToValidationInfoNode(document *Document) *Node {
        return nil
}


// throws Exception
func (this *AddValidationStoreFrontView) View() string {
        return this.Views.ValidationOnlyTempUtil.GetInstance().View(this)
}

