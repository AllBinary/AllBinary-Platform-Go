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

type BasicWorkFlow struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    AbeClientInformation AbeClientInformationInterface
    ValidationDomNodeInterface ValidationComponentInterface
    PropertiesHashMap *HashMap
    PageContext *PageContext
}

// Implements WorkFlowInterface


func (this *BasicWorkFlow) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.AbeClientInformation = ServiceClientInformationInterfaceFactory_GetInstance()
    
    
    
}


func NewBasicWorkFlow(propertiesHashMap *HashMap, pageContext *PageContext) *BasicWorkFlow {
    this := &BasicWorkFlow{}
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().WORKFLOW) {
            this.LogUtil.PutF("Properties: " + fmt.Sprint(propertiesHashMap), this, "BasicWorkFlow()")
        }
        this.PropertiesHashMap = propertiesHashMap
        this.PageContext = pageContext
        this.ValidationDomNodeInterface = any(TransformFactory_GetInstance().GetInstance(this.AbeClientInformation, propertiesHashMap, pageContext)).(ValidationComponentInterface)
    return this
}


// throws Exception
func (this *BasicWorkFlow) GetName() string {
        return any(this.GetPropertiesHashMap().Get(WorkFlowData_GetInstance().NAME)).(string)
}


func (this *BasicWorkFlow) GetPropertiesHashMap() *HashMap {
        return this.PropertiesHashMap
}


func (this *BasicWorkFlow) GetPageContext() *PageContext {
        return this.PageContext
}


// throws Exception
func (this *BasicWorkFlow) Process() int {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().WORKFLOW) {
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.PROCESS)
        }
        if this.ValidationDomNodeInterface.IsValid() {
            this.PageContext.GetOut().Print(this.ValidationDomNodeInterface.View())
            return fmt.Sprint(TagSupport_EVAL_BODY_INCLUDE)
        } else {
            this.PageContext.GetOut().Print(this.ValidationDomNodeInterface.ValidationInfo())
            return fmt.Sprint(TagSupport_SKIP_BODY)
        }
}

