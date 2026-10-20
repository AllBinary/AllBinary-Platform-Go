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

type EditValidationUserView struct {
    UserView
    LogUtil *LogUtil
    UserName string
}

// Implements ValidationComponentInterface


func (this *EditValidationUserView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewEditValidationUserView(transformInfoInterface TransformInfoInterface) *EditValidationUserView {
    this := &EditValidationUserView{}
        this.UserView = *NewUserView(transformInfoInterface)
        this.initFields()
        this.UserName = this.GetWeblisketSession().GetUserName()
    return this
}


func (this *EditValidationUserView) IsValid() bool {
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
            if UserName_GetInstance().IsValid(this.UserName) == true {
                this.User = UserEntityFactory_GetInstance().GetUser(this.UserName)
                if this.User == nil {
                    tryControl1, tryResult1 = 1, false; return
                }
            } else {
                tryControl1, tryResult1 = 1, false; return
            }
            tryControl1, tryResult1 = 1, this.User.IsValid(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *EditValidationUserView) ValidationInfo() string {
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
            if UserName_GetInstance().IsValid(this.UserName) {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    this.LogUtil.PutF("User Name is valid", this, "validationInfo()")
                }
                this.User = UserEntityFactory_GetInstance().GetUser(this.UserName)
                if this.User == nil {
                    stringBuffer.Append("User does not exist.<br />")
                    tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
                }
            } else {
                stringBuffer.Append("User Name is not valid.<br />")
                tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            }
            if this.User.IsValid() == false {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    this.LogUtil.PutF("User exists but is invalid - Probably manually modified", this, "validationInfo()")
                }
                stringBuffer.Append("User data is not valid - Please call administrator.<br />")
                stringBuffer.Append(this.User.ValidationInfo())
            }
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *EditValidationUserView) ToValidationInfoDoc() *Document {
        return nil
}


func (this *EditValidationUserView) ToValidationInfoNode(document *Document) *Node {
        return nil
}

