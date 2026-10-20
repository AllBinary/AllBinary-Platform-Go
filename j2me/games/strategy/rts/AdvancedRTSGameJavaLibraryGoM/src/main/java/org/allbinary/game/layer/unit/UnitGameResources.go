/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2006 AllBinary 
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

type UnitGameResources struct {
    BasicGameResources
    NAME string
    RESOURCE_ICON_ANIMATION string
    RESOURCE_ICON_BUILD string
    RESOURCE_ICON string
    ICON_SIZE string
}

var UnitGameResources_ICON string = "_icon"
var UnitGameResources_BUILD string = "_build"
var UnitGameResources_DESTROY string = "_destroy"
var UnitGameResources_ICON_BUILD string = "_icon" + fmt.Sprint(UnitGameResources_BUILD)

func (this *UnitGameResources) initFields() {
    this.NAME = StringUtil_GetInstance().EMPTY_STRING
    this.RESOURCE_ICON_ANIMATION = StringUtil_GetInstance().EMPTY_STRING
    this.RESOURCE_ICON_BUILD = StringUtil_GetInstance().EMPTY_STRING
    this.RESOURCE_ICON = StringUtil_GetInstance().EMPTY_STRING
    this.ICON_SIZE = "_64_by_64.png"
}


func NewUnitGameResources() *UnitGameResources {
    this := &UnitGameResources{}
    this.BasicGameResources = *NewBasicGameResources()
    this.initFields()
    return this
}


// throws Exception
func (this *UnitGameResources) Append(ROOT string, sizeString string) {
        this.BasicGameResources.Append(ROOT, sizeString)
        var string_ string = GameGraphicsResourceUtil_GetInstance().GetName(); _ = string_
        BasicGameResources_StringBuffer.Delete(0, BasicGameResources_StringBuffer.Length())
        BasicGameResources_StringBuffer.Append(ROOT)
        BasicGameResources_StringBuffer.Append(UnitGameResources_DESTROY)
        BasicGameResources_StringBuffer.Append(string_)
        BasicGameResources_StringBuffer.Append(sizeString)
        this.RESOURCE_DESTROY = fmt.Sprint(BasicGameResources_StringBuffer)
        BasicGameResources_StringBuffer.Delete(0, BasicGameResources_StringBuffer.Length())
        BasicGameResources_StringBuffer.Append(UnitGameResources_ICON)
        BasicGameResources_StringBuffer.Append(ROOT)
        this.RESOURCE_ICON_ANIMATION = fmt.Sprint(BasicGameResources_StringBuffer)
        BasicGameResources_StringBuffer.Delete(0, BasicGameResources_StringBuffer.Length())
        BasicGameResources_StringBuffer.Append(ROOT)
        BasicGameResources_StringBuffer.Append(UnitGameResources_ICON)
        BasicGameResources_StringBuffer.Append(string_)
        BasicGameResources_StringBuffer.Append(this.ICON_SIZE)
        this.RESOURCE_ICON = fmt.Sprint(BasicGameResources_StringBuffer)
        BasicGameResources_StringBuffer.Delete(0, BasicGameResources_StringBuffer.Length())
        BasicGameResources_StringBuffer.Append(ROOT)
        BasicGameResources_StringBuffer.Append(UnitGameResources_ICON_BUILD)
        BasicGameResources_StringBuffer.Append(string_)
        BasicGameResources_StringBuffer.Append(this.ICON_SIZE)
        this.RESOURCE_ICON_BUILD = fmt.Sprint(BasicGameResources_StringBuffer)
}


// throws Exception
func (this *UnitGameResources) GetString() string {
        if AndroidUtil_IsAndroid() {
            return GameGraphicsResourceUtil_GetInstance().GetStringForFeature(GraphicsFeatureFactory_GetInstance().SPRITE_QUARTER_ROTATION_GRAPHICS)
        } else {
            return this.BasicGameResources.GetString()
        }
}

