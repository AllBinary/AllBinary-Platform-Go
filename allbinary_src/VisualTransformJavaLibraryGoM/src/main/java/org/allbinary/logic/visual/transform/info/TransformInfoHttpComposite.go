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

type TransformInfoHttpComposite struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    TransformInfoInterface TransformInfoHttpInterface
}


func (this *TransformInfoHttpComposite) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
}


func NewTransformInfoHttpComposite(transformInfoInterface TransformInfoInterface) *TransformInfoHttpComposite {
    this := &TransformInfoHttpComposite{}
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("View Name: " + transformInfoInterface.GetName(), this, this.CommonStrings.CONSTRUCTOR)
        }
        this.TransformInfoInterface = any(transformInfoInterface).(TransformInfoHttpInterface)
    return this
}


// throws Exception
func (this *TransformInfoHttpComposite) GetTransformInfoInterface() TransformInfoInterface {
        return this.TransformInfoInterface
}


func (this *TransformInfoHttpComposite) GetWeblisketSession() *WeblisketSession {
        return NewWeblisketSession(this.GetPropertiesHashMap(), this.GetPageContext())
}


func (this *TransformInfoHttpComposite) GetPageContext() *PageContext {
        return this.TransformInfoInterface.GetPageContext()
}


func (this *TransformInfoHttpComposite) GetPropertiesHashMap() *HashMap {
        return this.TransformInfoInterface.GetPropertiesHashMap()
}

