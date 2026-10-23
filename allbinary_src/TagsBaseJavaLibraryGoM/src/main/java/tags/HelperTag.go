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

type HelperTag struct {
    PropertiesTag
    LogUtil *LogUtil
    TagHelperFactoryInterface TagHelperFactoryInterface
    AnyType any
}

// The Java class HelperTag has subclasses so its variables are this interface
type IHelperTag interface {
    IPropertiesTag
    HelperTag_() *HelperTag
    SetHelper()
    GetHelper() any
    DoEndTag() int
}

func (this *HelperTag) HelperTag_() *HelperTag {
    return this
}


func (this *HelperTag) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
}


func NewHelperTag(tagHelperFactoryInterface TagHelperFactoryInterface) *HelperTag {
    this := &HelperTag{}
    this.PropertiesTag = *NewPropertiesTag()
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        }
        this.TagHelperFactoryInterface = tagHelperFactoryInterface
    return this
}


// throws Exception
func (this *HelperTag) SetHelper() {
        if this.AnyType == nil {
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                this.LogUtil.PutF("Creating TagHelper with: \n" + fmt.Sprint(this.GetPropertiesHashMap()), this, "doStartTag")
            }
            this.AnyType = this.TagHelperFactoryInterface.GetInstance(this.GetPropertiesHashMap(), this.PageContext)
        }
}


func (this *HelperTag) GetHelper() any {
        return this.AnyType
}


func (this *HelperTag) DoEndTag() int {
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
            this.LogUtil.PutF("Tag Ended", this, "doEndTag")
        }
        this.AnyType = nil
        return this.PropertiesTag.DoEndTag()
}

