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

type WorkAreaDom struct {
    Name string
    CanvasNodeList IBasicArrayList
}

var WorkAreaDom_WORKAREA string = "workArea"
var WorkAreaDom_NAME string = "name"
var WorkAreaDom_FRAMES string = "frames"

func (this *WorkAreaDom) initFields() {
    
    
}


func NewWorkAreaDom(document *Document) *WorkAreaDom {
    this := &WorkAreaDom{}
    this.initFields()
        var workAreaNodeList *NodeList = document.GetElementsByTagName(WorkAreaDom_WORKAREA); _ = workAreaNodeList
        if workAreaNodeList != nil {
            var numberOfworkAreas int = workAreaNodeList.GetLength(); _ = numberOfworkAreas
            if numberOfworkAreas == 1 {
                var numberOfNodes int = workAreaNodeList.GetLength(); _ = numberOfNodes
                var nameNode *Node = DomHelper_GetInstance().SearchNodeList(WorkAreaDom_NAME, workAreaNodeList.Item(0).GetChildNodes()); _ = nameNode
                var nameTextNode *Node = nameNode.GetFirstChild(); _ = nameTextNode
                this.Name = nameTextNode.GetNodeValue()
            } else {
                panic(errors.New(fmt.Sprint("Wrong Number of WorkAreas: " + fmt.Sprint(numberOfworkAreas))))
            }
            this.CanvasNodeList = DomHelper_GetInstance().GetChildrenWithoutTextNodes(WorkAreaDom_FRAMES, workAreaNodeList.Item(0).GetChildNodes())
        } else {
            panic(errors.New(fmt.Sprint("workArea Node Not Found")))
        }
    return this
}


func (this *WorkAreaDom) GetName() string {
        return this.Name
}


func (this *WorkAreaDom) GetCanvasNodes() IBasicArrayList {
        return this.CanvasNodeList
}

