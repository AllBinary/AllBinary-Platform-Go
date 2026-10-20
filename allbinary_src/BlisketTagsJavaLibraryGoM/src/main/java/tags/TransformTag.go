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

type TransformTag struct {
    TransformInfoTag
    LogUtil *LogUtil
}

// The Java class TransformTag has subclasses so its variables are this interface
type ITransformTag interface {
    ITransformInfoTag
    TransformTag_() *TransformTag
    View() string
    DoStartTag() int
}

func (this *TransformTag) TransformTag_() *TransformTag {
    return this
}


func (this *TransformTag) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}




// Java overloads of new TransformTag
func NewTransformTag(args ...any) *TransformTag {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewTransformTag_1()
        }
        if len(args) == 1 && Java_IsArg[TagHelperFactoryInterface](args[0], pass) {
            return NewTransformTag_2(Java_Arg[TagHelperFactoryInterface](args[0]))
        }
    }
    panic("No Java overload of new TransformTag for the arguments")
}
func NewTransformTag_1() *TransformTag {
    this := &TransformTag{}
        this.TransformInfoTag = *NewTransformInfoTag(NewViewHelperFactory())
        this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        }
    return this
}


func NewTransformTag_2(tagHelperFactoryInterface TagHelperFactoryInterface) *TransformTag {
    this := &TransformTag{}
        this.TransformInfoTag = *NewTransformInfoTag(tagHelperFactoryInterface)
        this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
            this.LogUtil.PutF(this.CommonStrings.START, this, this.CommonStrings.CONSTRUCTOR)
        }
    return this
}


// throws Exception
func (this *TransformTag) View() string {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 string) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAGERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, "view()", e)
            }
            panic(e)
                }
            }()
            var result string = any(reflect.TypeOf(this.GetHelper()).GetMethod("view", nil).Invoke(this.GetHelper(), nil)).(string); _ = result
            tryControl1, tryResult1 = 1, result; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


// throws JspTagException
func (this *TransformTag) DoStartTag() int {
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            AbResponseHandler_SendJspTagLicensingRedirect(this.PageContext, e)
            tryControl1, tryResult1 = 1, TransformInfoTag_SKIP_BODY; return
                }
            }()
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                this.LogUtil.PutF("Creating with: \n" + fmt.Sprint(this.GetPropertiesHashMap()), this, this.TagStrings.DO_START_TAG)
            }
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
                stringBuffer.Append("ViewTag Start For: ")
                stringBuffer.Append(this.GetName())
                stringBuffer.Append("\nViewFile: ")
                stringBuffer.Append(this.GetObjectFile())
                this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.TagStrings.DO_START_TAG)
            }
            this.SetHelper()
            this.PageContext.GetOut().Print(this.View())
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().JSPTAG) {
                this.LogUtil.PutF("Tag End", this, this.TagStrings.DO_START_TAG)
            }
            tryControl1, tryResult1 = 1, TransformInfoTag_SKIP_BODY; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

