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





// Java overloads of new TransformInfoDomNodeView
func NewTransformInfoDomNodeView(args ...any) *TransformInfoDomNodeView {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[TransformInfoInterface](args[0], pass) && Java_IsArg[string](args[1], pass) {
            return NewTransformInfoDomNodeView_1(Java_Arg[TransformInfoInterface](args[0]), Java_Arg[string](args[1]))
        }
        if len(args) == 1 && Java_IsArg[TransformInfoInterface](args[0], pass) {
            return NewTransformInfoDomNodeView_2(Java_Arg[TransformInfoInterface](args[0]))
        }
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[string](args[1], pass) {
            return NewTransformInfoDomNodeView_3(Java_Arg[string](args[0]), Java_Arg[string](args[1]))
        }
        if len(args) == 1 && Java_IsArg[string](args[0], pass) {
            return NewTransformInfoDomNodeView_4(Java_Arg[string](args[0]))
        }
    }
    panic("No Java overload of new TransformInfoDomNodeView for the arguments")
}
func NewTransformInfoDomNodeView_1(transformInfoInterface TransformInfoInterface, mapping string) *TransformInfoDomNodeView {
    this := &TransformInfoDomNodeView{}
        this.TransformInfoDomNode = *NewTransformInfoDomNode(transformInfoInterface, mapping)
    return this
}


func NewTransformInfoDomNodeView_2(transformInfoInterface TransformInfoInterface) *TransformInfoDomNodeView {
    this := &TransformInfoDomNodeView{}
        this.TransformInfoDomNode = *NewTransformInfoDomNode(transformInfoInterface)
    return this
}


func NewTransformInfoDomNodeView_3(name string, mappedName string) *TransformInfoDomNodeView {
    this := &TransformInfoDomNodeView{}
        this.TransformInfoDomNode = *NewTransformInfoDomNode(name, mappedName)
    return this
}


func NewTransformInfoDomNodeView_4(name string) *TransformInfoDomNodeView {
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

