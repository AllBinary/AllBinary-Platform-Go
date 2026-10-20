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

type DeleteValidationView struct {
    WorkFlowView
    LogUtil *LogUtil
    WorkFlowName string
}

// Implements ValidationComponentInterface


func (this *DeleteValidationView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewDeleteValidationView(transformInfoInterface TransformInfoInterface) *DeleteValidationView {
    this := &DeleteValidationView{}
        this.WorkFlowView = *NewWorkFlowView(transformInfoInterface)
        this.initFields()
        this.WorkFlowName = this.GetPageContext().GetRequest().GetParameter(WorkFlowData_GetInstance().NAME)
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Name: " + this.WorkFlowName, this, this.CommonStrings.CONSTRUCTOR)
        }
    return this
}


func (this *DeleteValidationView) IsValid() bool {
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
            if this.WorkFlowName == "" {
                tryControl1, tryResult1 = 1, false; return
            }
            if WorkFlowEntityFactory_GetInstance().Create2().Get(this.WorkFlowName, this.GetWeblisketSession().GetStoreName()) == nil {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    this.LogUtil.PutF("WorkFlow does not exist.", this, this.CommonStrings.IS_VALID)
                }
                tryControl1, tryResult1 = 1, false; return
            }
            tryControl1, tryResult1 = 1, true; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *DeleteValidationView) ValidationInfo() string {
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
            if this.WorkFlowName == "" {
                stringBuffer.Append("WorkFlow name is invalid<br />")
            }
            if WorkFlowEntityFactory_GetInstance().Create2().Get(this.WorkFlowName, this.GetWeblisketSession().GetStoreName()) == nil {
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                    this.LogUtil.PutF("WorkFlow does not exist.", this, "validationInfo()")
                }
                stringBuffer.Append("WorkFlow does not exist<br />")
            }
            tryControl1, tryResult1 = 1, fmt.Sprint(stringBuffer); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *DeleteValidationView) ToValidationInfoDoc() *Document {
        return nil
}


func (this *DeleteValidationView) ToValidationInfoNode(document *Document) *Node {
        return nil
}

