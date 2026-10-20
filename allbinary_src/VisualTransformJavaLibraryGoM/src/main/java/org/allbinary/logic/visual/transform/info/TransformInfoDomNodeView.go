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

type TransformInfoDomNodeView struct {
    TransformInfoDomNode
}

// Implements DomNodeInterface



func NewTransformInfoDomNodeView(transformInfoInterface TransformInfoInterface, mapping string) *TransformInfoDomNodeView {
    this := &TransformInfoDomNodeView{}
        this.TransformInfoDomNode = *NewTransformInfoDomNode(transformInfoInterface, mapping)
    return this
}


// Java overloaded constructor: Go needs a different name
func NewTransformInfoDomNodeView(transformInfoInterface TransformInfoInterface) *TransformInfoDomNodeView {
    this := &TransformInfoDomNodeView{}
        this.TransformInfoDomNode = *NewTransformInfoDomNode(transformInfoInterface)
    return this
}


// Java overloaded constructor: Go needs a different name
func NewTransformInfoDomNodeView(name string, mappedName string) *TransformInfoDomNodeView {
    this := &TransformInfoDomNodeView{}
        this.TransformInfoDomNode = *NewTransformInfoDomNode(name, mappedName)
    return this
}


// Java overloaded constructor: Go needs a different name
func NewTransformInfoDomNodeView(name string) *TransformInfoDomNodeView {
    this := &TransformInfoDomNodeView{}
        this.TransformInfoDomNode = *NewTransformInfoDomNode(name, nil)
    return this
}


// throws Exception
func (this *TransformInfoDomNodeView) ToXmlNode(document *Document) *Node {
        var transformInfoData *TransformInfoData = TransformInfoData_GetInstance(); _ = transformInfoData
        var subViewNode *Node = document.CreateElement(transformInfoData.NAME); _ = subViewNode
        var viewNameAttr *Attr = document.CreateAttribute(transformInfoData.NAME); _ = viewNameAttr
        viewNameAttr.SetValue(this.GetTransformInfoInterface().GetName())
        var componentNodeAttributes *NamedNodeMap = subViewNode.GetAttributes(); _ = componentNodeAttributes
        componentNodeAttributes.SetNamedItem(viewNameAttr)
        if this.GetMappedName() != nil {
            var mappedNameAttr *Attr = document.CreateAttribute(transformInfoData.MAPPED); _ = mappedNameAttr
            mappedNameAttr.SetValue(this.GetMappedName())
            componentNodeAttributes.SetNamedItem(mappedNameAttr)
        }
        return subViewNode
}

