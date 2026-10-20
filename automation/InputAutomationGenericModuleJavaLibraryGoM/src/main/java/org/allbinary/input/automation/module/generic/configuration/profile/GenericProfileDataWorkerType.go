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

type GenericProfileDataWorkerType struct {
    Name string
}

// Implements GenericProfileDataWorkerTypeInterface

// The Java class GenericProfileDataWorkerType has subclasses so its variables are this interface
type IGenericProfileDataWorkerType interface {
    GenericProfileDataWorkerType_() *GenericProfileDataWorkerType
    GetName() string
    SetName(name string)
    ToXmlNode(document *Document) *Node
    ToString() string
}

func (this *GenericProfileDataWorkerType) GenericProfileDataWorkerType_() *GenericProfileDataWorkerType {
    return this
}

var GenericProfileDataWorkerType_HashMap *HashMap = StdUtil_GetInstance().CreateHashMap()
var GenericProfileDataWorkerType_SAVED_CAPTURE IGenericProfileDataWorkerType = SavedCaptureGenericProfileDataWorkerType_SAVED_CAPTURE
var GenericProfileDataWorkerType_SCREEN_CAPTURE IGenericProfileDataWorkerType = NewGenericProfileDataWorkerType("Screen " + InputImageType_CAPTURE.GetName())
var GenericProfileDataWorkerType_COMPARISON IGenericProfileDataWorkerType = NewGenericProfileDataWorkerType(InputImageType_COMPARISON.GetName())
var GenericProfileDataWorkerType_MOTION IGenericProfileDataWorkerType = NewGenericProfileDataWorkerType(InputImageType_MOTION.GetName())

func (this *GenericProfileDataWorkerType) initFields() {
    
}


func NewGenericProfileDataWorkerType(name string) *GenericProfileDataWorkerType {
    this := &GenericProfileDataWorkerType{}
    this.initFields()
        this.SetName(name)
        GenericProfileDataWorkerType_HashMap.Put(this.GetName(), this)
    return this
}


func (this *GenericProfileDataWorkerType) GetName() string {
        return this.Name
}


func (this *GenericProfileDataWorkerType) SetName(name string) {
        this.Name = name
}


// throws Exception
func (this *GenericProfileDataWorkerType) ToXmlNode(document *Document) *Node {
        var node *Node = ModDomHelper_CreateTextNode(document, GenericProfileDataWorkerData_NAME, this.Name); _ = node
        return node
}


func (this *GenericProfileDataWorkerType) ToString() string {
        return "GenericProfileDataWorkerType: " + this.GetName()
}

func (this *GenericProfileDataWorkerType) String() string {
    return this.ToString()
}


func GenericProfileDataWorkerType_GetInstance(node *Node) IGenericProfileDataWorkerType {
        var name string = DomNodeHelper_GetTextNodeValue(node); _ = name
        return any(GenericProfileDataWorkerType_HashMap.Get(name)).(IGenericProfileDataWorkerType)
}

