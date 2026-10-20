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

type StoreTagWorkFlowHelper struct {
    TagHelper
    LogUtil *LogUtil
    AbeClientInformation AbeClientInformationInterface
    StoreWorkFlowInterface StoreWorkFlowInterface
}


func (this *StoreTagWorkFlowHelper) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.AbeClientInformation = ServiceClientInformationInterfaceFactory_GetInstance()
    
}


func NewStoreTagWorkFlowHelper(hashMap *HashMap, pageContext *PageContext) *StoreTagWorkFlowHelper {
    this := &StoreTagWorkFlowHelper{}
    this.TagHelper = *NewTagHelper()
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPER) {
            this.LogUtil.PutF("Properties: " + fmt.Sprint(hashMap), this, "StoreTagWorkFlowHelper()")
        }
        this.StoreWorkFlowInterface = StoreTagWorkFlowFactory_GetInstance().GetInstance(this.AbeClientInformation, hashMap, pageContext)
    return this
}


// throws Exception
func (this *StoreTagWorkFlowHelper) Process() int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var error_ string = "Failed to process workflow: "; _ = error_
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().SQLTAGSERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "view()", e)
            }
            panic(e)
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPER) {
                this.LogUtil.PutF("Process", this, this.CommonStrings.PROCESS)
            }
            tryControl1, tryResult1 = 1, this.StoreWorkFlowInterface.Process(); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

