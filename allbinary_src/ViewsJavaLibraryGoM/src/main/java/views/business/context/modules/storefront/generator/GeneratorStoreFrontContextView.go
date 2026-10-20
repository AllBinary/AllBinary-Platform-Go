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

type GeneratorStoreFrontContextView struct {
    HttpStoreComponentView
    LogUtil *LogUtil
}


func (this *GeneratorStoreFrontContextView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewGeneratorStoreFrontContextView(transformInfoInterface TransformInfoInterface) *GeneratorStoreFrontContextView {
    this := &GeneratorStoreFrontContextView{}
        this.HttpStoreComponentView = *NewHttpStoreComponentView(transformInfoInterface)
        this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("View Name: " + transformInfoInterface.GetName(), this, this.CommonStrings.CONSTRUCTOR)
        }
    return this
}


// throws Exception
func (this *GeneratorStoreFrontContextView) View(group string) string {
        var result string = TransformsGeneratorUtil_GetInstance().GenerateComponentsFromObjectConfig(this.AbeClientInformation, this.GetTransformInfoInterface(), group); _ = result
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPEXTRAOUTPUT) {
            return result
        } else {
            return StringUtil_GetInstance().EMPTY_STRING
        }
}


// throws Exception
// Java overloaded method: Go needs a different name
func (this *GeneratorStoreFrontContextView) View() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().TAGHELPERERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "view()", e)
            }
            panic(e)
                }
            }()
            tryControl1, tryResult1 = 1, this.View(TransformInfosData_GetInstance().ALL); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

