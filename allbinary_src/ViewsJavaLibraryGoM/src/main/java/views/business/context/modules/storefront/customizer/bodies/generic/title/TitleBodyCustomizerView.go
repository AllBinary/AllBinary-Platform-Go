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

type TitleBodyCustomizerView struct {
    HttpStoreComponentView
    LogUtil *LogUtil
    TitleBody *TitleBodyValidation
}

// Implements DomNodeInterface

var TitleBodyCustomizerView_NAME string = "None"

func (this *TitleBodyCustomizerView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewTitleBodyCustomizerView(transformInfoInterface TransformInfoInterface) *TitleBodyCustomizerView {
    this := &TitleBodyCustomizerView{}
        this.HttpStoreComponentView = *NewHttpStoreComponentView(transformInfoInterface)
        this.initFields()
    return this
}


// throws Exception
func (this *TitleBodyCustomizerView) AddDomNodeInterfaces() {
        this.AddDomNodeInterface(any(this.TitleBody).(DomNodeInterface))
        this.AddDomNodeInterface(any(this).(DomNodeInterface))
}


// throws Exception
func (this *TitleBodyCustomizerView) ToXmlNode(document *Document) *Node {
        var pageName string = TransformTemplateCustomizerUtil_GetInstance().GetPageNameHack(this.GetTransformInfoInterface().GetName(), this.GetWeblisketSession().GetStoreName()); _ = pageName
        if StringValidationUtil_GetInstance().IsEmpty(pageName) {
            pageName = TitleBodyCustomizerView_NAME
        }
        return ModDomHelper_CreateNameValueNodes(document, PageData_GetInstance().NAME, pageName)
}


// throws Exception
func (this *TitleBodyCustomizerView) View() string {
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
            var success string = XmlDocumentHelper_ToString(this.GetDoc()); _ = success
            var result string = NewStoreTransformer(this.AbeClientInformation, this.GetTransformInfoInterface()).Translate(success); _ = result
            tryControl1, tryResult1 = 1, StringEscapeUtils_UnescapeHtml3(result); return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}

