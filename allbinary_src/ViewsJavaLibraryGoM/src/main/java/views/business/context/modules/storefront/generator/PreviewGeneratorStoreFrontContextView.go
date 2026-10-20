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

type PreviewGeneratorStoreFrontContextView struct {
    GeneratorStoreFrontContextView
    LogUtil *LogUtil
}


func (this *PreviewGeneratorStoreFrontContextView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewPreviewGeneratorStoreFrontContextView(transformInfoInterface TransformInfoInterface) *PreviewGeneratorStoreFrontContextView {
    this := &PreviewGeneratorStoreFrontContextView{}
        this.GeneratorStoreFrontContextView = *NewGeneratorStoreFrontContextView(transformInfoInterface)
        this.initFields()
    return this
}


// throws Exception
func (this *PreviewGeneratorStoreFrontContextView) View() string {
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
            var httpTransformInfoInterface TransformInfoHttpInterface = any(this.GetTransformInfoInterface()).(TransformInfoHttpInterface); _ = httpTransformInfoInterface
            httpTransformInfoInterface.GetPropertiesHashMap().Put(TransformInfoData_GetInstance().PARTIAL, CommonSeps_GetInstance().SPACE + TransformInfosData_GetInstance().PREVIEW)
            tryControl1, tryResult1 = 1, this.View(TransformInfosData_GetInstance().PREVIEW); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

