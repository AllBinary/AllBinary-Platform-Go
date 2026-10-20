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

type AdminConfigurationHelper struct {
    BasicTable
    LogUtil *LogUtil
}


func (this *AdminConfigurationHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewAdminConfigurationHelper(hashMap *HashMap, pageContext *PageContext) *AdminConfigurationHelper {
    this := &AdminConfigurationHelper{}
    this.BasicTable = *NewBasicTable()
    this.initFields()
    return this
}


func (this *AdminConfigurationHelper) Create() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to create table"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "create()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var success string = "No Imp"; _ = success
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "create()")
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *AdminConfigurationHelper) Drop() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to drop table"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.DROP, e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var success string = "No Imp"; _ = success
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, this.CommonStrings.DROP)
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *AdminConfigurationHelper) Restore() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to restore backup"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "restore()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var success string = "No Imp"; _ = success
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *AdminConfigurationHelper) Backup() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to make backup"; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "backup()", e)
            }
            tryControl1, tryResult1 = 1, error_; return
                }
            }()
            var success string = "Backup Successful"; _ = success
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGS) {
                this.LogUtil.PutF(success, this, "backup()")
            }
            tryControl1, tryResult1 = 1, success; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

