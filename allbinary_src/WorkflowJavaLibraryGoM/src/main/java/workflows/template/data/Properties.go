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

type Properties struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    PropertiesHashMap *HashMap
}

// The Java class Properties has subclasses so its variables are this interface
type IProperties interface {
    Properties_() *Properties
    GetPropertiesHashMap() *HashMap
    DoEndTag() int
}

func (this *Properties) Properties_() *Properties {
    return this
}


func (this *Properties) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
}


func NewProperties() *Properties {
    this := &Properties{}
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
            this.LogUtil.PutF("Start/Tag", this, this.CommonStrings.CONSTRUCTOR)
        }
        this.PropertiesHashMap = StdUtil_GetInstance().CreateHashMap()
    return this
}


func (this *Properties) GetPropertiesHashMap() *HashMap {
        return this.PropertiesHashMap
}


func (this *Properties) DoEndTag() int {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
            this.LogUtil.PutF("Tag Ended", this, "doEndTag")
        }
        this.PropertiesHashMap = StdUtil_GetInstance().CreateHashMap()
        return TagSupport_EVAL_PAGE
}

