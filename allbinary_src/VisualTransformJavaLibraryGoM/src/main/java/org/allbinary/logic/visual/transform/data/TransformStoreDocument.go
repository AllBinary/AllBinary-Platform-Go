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

type TransformStoreDocument struct {
    TransformHttpRequestDocument
    LogUtil *LogUtil
}


func (this *TransformStoreDocument) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}




// Java overloads of new TransformStoreDocument
func NewTransformStoreDocument(args ...any) *TransformStoreDocument {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[*PageContext](args[0], pass) && Java_IsArg[*WeblisketSession](args[1], pass) {
            return NewTransformStoreDocument_1(Java_Arg[*PageContext](args[0]), Java_Arg[*WeblisketSession](args[1]))
        }
        if len(args) == 1 && Java_IsArg[*SearchRequest](args[0], pass) {
            return NewTransformStoreDocument_2(Java_Arg[*SearchRequest](args[0]))
        }
    }
    panic("No Java overload of new TransformStoreDocument for the arguments")
}
func NewTransformStoreDocument_1(pageContext *PageContext, weblisketSession *WeblisketSession) *TransformStoreDocument {
    this := &TransformStoreDocument{}
        this.TransformHttpRequestDocument = *NewTransformHttpRequestDocument(pageContext, weblisketSession)
        this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.CONSTRUCTOR, e)
            }
            panic(e)
                }
            }()
            var storeName string = weblisketSession.GetStoreName(); _ = storeName
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
                this.LogUtil.PutF("Using Session StoreName: " + storeName, this, this.CommonStrings.CONSTRUCTOR)
            }
            this.GetBaseNode().AppendChild(StoreFrontViewFactory_GetInstance(storeName).ToXmlNode(this.GetDoc()))
            this.GetBaseNode().AppendChild(NewRequestParams(any(pageContext.GetRequest()).(*HttpServletRequest)).ToXmlNode(this.GetDoc()))
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


func NewTransformStoreDocument_2(searchRequest *SearchRequest) *TransformStoreDocument {
    this := &TransformStoreDocument{}
        this.TransformHttpRequestDocument = *NewTransformHttpRequestDocument(searchRequest)
        this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEWERROR) {
                this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.CONSTRUCTOR, e)
            }
            panic(e)
                }
            }()
            this.GetBaseNode().AppendChild(StoreFrontViewFactory_GetInstance(searchRequest.GetStoreFront().GetName()).ToXmlNode(this.GetDoc()))
            this.GetBaseNode().AppendChild(searchRequest.GetParams().GetParamsNode(this.GetDoc()))
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}

