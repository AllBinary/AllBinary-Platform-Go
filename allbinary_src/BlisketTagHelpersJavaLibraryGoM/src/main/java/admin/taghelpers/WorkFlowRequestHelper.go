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

type WorkFlowRequestHelper struct {
    ModifyTable
    LogUtil *LogUtil
    AbeClientInformation AbeClientInformationInterface
    WorkFlowInterface WorkFlowInterface
}


func (this *WorkFlowRequestHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.AbeClientInformation = ServiceClientInformationInterfaceFactory_GetInstance()
    
}


func NewWorkFlowRequestHelper(hashMap *HashMap, pageContext *PageContext) *WorkFlowRequestHelper {
    this := &WorkFlowRequestHelper{}
    this.ModifyTable = *NewModifyTable()
    this.initFields()
        this.WorkFlowInterface = NewWorkFlowFactory_GetInstance().GetInstance(this.AbeClientInformation, hashMap, pageContext)
    return this
}


// throws Exception
func (this *WorkFlowRequestHelper) Delete() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if tryControl2, tryResult2 := func() (tryControl2 int, tryResult2 string) {
                defer func() {
                    if recovered := recover(); recovered != nil {
                        e2 := Java_ToError(recovered)
                        _ = e2
                var error_ string = "Failed to get data from workflowinterface to set error string"; _ = error_
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                    this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "delete()", e2)
                }
                tryControl2, tryResult2 = 1, error_; return
                    }
                }()
                var error_ string = "Failed to remove workflow with " + fmt.Sprint(WorkFlowData_GetInstance().NAME) + "= " + this.WorkFlowInterface.GetName() + " from "; _ = error_
                if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                    this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "delete()", e)
                }
                panic(errors.New(fmt.Sprint(e)))
                return
            }(); tryControl2 == 1 {
                tryControl1, tryResult1 = 1, tryResult2; return
            }
                }
            }()
            var stringUtil *StringUtil = StringUtil_GetInstance(); _ = stringUtil
            WorkFlowEntityFactory_GetInstance().Create2().Delete(this.WorkFlowInterface.GetName(), stringUtil.EMPTY_STRING)
            var success string = "Successfully Removed the workflow with " + fmt.Sprint(WorkFlowData_GetInstance().NAME) + "= " + this.WorkFlowInterface.GetName() + " from "; _ = success
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "delete()")
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *WorkFlowRequestHelper) Insert() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to add User"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "add()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var success string = "New User Successfully added to the Users Table"; _ = success
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "add()")
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *WorkFlowRequestHelper) Update() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to update WorkFlow"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "update()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var success string = "Updated WorkFlow Successfully"; _ = success
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "update()")
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

