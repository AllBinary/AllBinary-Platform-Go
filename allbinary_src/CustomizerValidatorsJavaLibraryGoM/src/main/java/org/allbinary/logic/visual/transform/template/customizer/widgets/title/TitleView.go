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

type TitleView struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    Title string
}

// Implements DomNodeInterface

// The Java class TitleView has subclasses so its variables are this interface
type ITitleView interface {
    TitleView_() *TitleView
    GetTitle() string
    GetFormData(hashMap *HashMap)
    ToHashMap() *HashMap
    ToXmlNode(document *Document) *Node
}

func (this *TitleView) TitleView_() *TitleView {
    return this
}

var TitleView_NAME string = "None"

func (this *TitleView) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
}




// Java overloads of new TitleView
func NewTitleView(args ...any) *TitleView {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return NewTitleView_1()
        }
        if len(args) == 1 && Java_IsArg[*Node](args[0], pass) {
            return NewTitleView_2(Java_Arg[*Node](args[0]))
        }
        if len(args) == 1 && Java_IsArg[*HashMap](args[0], pass) {
            return NewTitleView_3(Java_Arg[*HashMap](args[0]))
        }
    }
    panic("No Java overload of new TitleView for the arguments")
}
func NewTitleView_1() *TitleView {
    this := &TitleView{}
    this.initFields()
        this.Title = TitleView_NAME
    return this
}


func NewTitleView_2(node *Node) *TitleView {
    this := &TitleView{}
    this.initFields()
        var titleTextNode *Node = DomSearchHelper_GetNode(TitleData_GetInstance().TEXT, node.GetChildNodes()); _ = titleTextNode
        var valueNode *Node = DomSearchHelper_GetNode(DomData_VALUE, titleTextNode.GetChildNodes()); _ = valueNode
        this.Title = DomNodeHelper_GetTextNodeValue(valueNode)
        if StringValidationUtil_GetInstance().IsEmpty(this.Title) {
            this.Title = TitleView_NAME
        }
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("Title: " + fmt.Sprint(this.GetTitle()), this, "Consructor()")
        }
    return this
}


func NewTitleView_3(hashMap *HashMap) *TitleView {
    this := &TitleView{}
    this.initFields()
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("HashMap: " + fmt.Sprint(hashMap), this, this.CommonStrings.CONSTRUCTOR)
        }
        this.GetFormData(hashMap)
    return this
}


func (this *TitleView) GetTitle() string {
        return this.Title
}


func (this *TitleView) GetFormData(hashMap *HashMap) {
        this.Title = any(hashMap.Get(TitleData_GetInstance().TEXT)).(string)
        if StringValidationUtil_GetInstance().IsEmpty(this.Title) {
            this.Title = TitleView_NAME
        }
}


func (this *TitleView) ToHashMap() *HashMap {
        var hashMap *HashMap = StdUtil_GetInstance().CreateHashMap(); _ = hashMap
        hashMap.Put(TitleData_GetInstance().TEXT, this.GetTitle())
        return hashMap
}


// throws Exception
func (this *TitleView) ToXmlNode(document *Document) *Node {
        var hashMap *HashMap = this.ToHashMap(); _ = hashMap
        if org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypes.LOGGING.Contains(org_Allbinary.Logic.Communication.Log.Config.Type.LogConfigTypeFactory.GetInstance().VIEW) {
            this.LogUtil.PutF("TitleView HashMap: " + fmt.Sprint(hashMap), this, "toXmlNode()")
        }
        return ModDomHelper_CreateNameValueNodes(document, TitleData_GetInstance().NAME, hashMap)
}

