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

type PropertiesTag struct {
    CommandTag
    LogUtil *LogUtil
    PropertiesHashMap *HashMap
}

// The Java class PropertiesTag has subclasses so its variables are this interface
type IPropertiesTag interface {
    ICommandTag
    PropertiesTag_() *PropertiesTag
    GetPropertiesHashMap() *HashMap
    DoEndTag() int
}

func (this *PropertiesTag) PropertiesTag_() *PropertiesTag {
    return this
}


func (this *PropertiesTag) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewPropertiesTag() *PropertiesTag {
    this := &PropertiesTag{}
    this.CommandTag = *NewCommandTag()
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        }
        this.PropertiesHashMap = StdUtil_GetInstance().CreateHashMap()
    return this
}


func (this *PropertiesTag) GetPropertiesHashMap() *HashMap {
        return this.PropertiesHashMap
}


func (this *PropertiesTag) DoEndTag() int {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
            this.LogUtil.PutF("Tag Ended", this, "doEndTag")
        }
        this.PropertiesHashMap = StdUtil_GetInstance().CreateHashMap()
        return CommandTag_EVAL_PAGE
}

