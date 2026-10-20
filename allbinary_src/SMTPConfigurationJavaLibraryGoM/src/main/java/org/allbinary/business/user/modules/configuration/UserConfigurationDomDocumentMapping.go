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

type UserConfigurationDomDocumentMapping struct {
    UserConfigurationInterface UserConfigurationInterface
}

// Implements DomDocumentMappingInterface


func (this *UserConfigurationDomDocumentMapping) initFields() {
    
}


func NewUserConfigurationDomDocumentMapping(userConfigurationInterface UserConfigurationInterface) *UserConfigurationDomDocumentMapping {
    this := &UserConfigurationDomDocumentMapping{}
    this.initFields()
        this.UserConfigurationInterface = userConfigurationInterface
    return this
}


// throws Exception
func (this *UserConfigurationDomDocumentMapping) ToXmlDoc() *Document {
        var document *Document = DomDocumentHelper_Create(); _ = document
        var node *Node = NewUserConfigurationView(this.UserConfigurationInterface).ToXmlNode(document); _ = node
        document.AppendChild(node)
        return document
}


// throws Exception
func (this *UserConfigurationDomDocumentMapping) ToDomDocumentString() string {
        return XmlDocumentHelper_ToString(this.ToXmlDoc())
}

