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

type GatewayValidationView struct {
    HttpStoreComponentView
}

// Implements ValidationComponentInterface



func NewGatewayValidationView(transformInfoInterface TransformInfoInterface) *GatewayValidationView {
    this := &GatewayValidationView{}
        this.HttpStoreComponentView = *NewHttpStoreComponentView(transformInfoInterface)
    return this
}


func (this *GatewayValidationView) IsValid() bool {
        return true
}


func (this *GatewayValidationView) ToValidationInfoDoc() *Document {
        return nil
}


func (this *GatewayValidationView) ToValidationInfoNode(document *Document) *Node {
        return nil
}


// throws Exception
func (this *GatewayValidationView) ValidationInfo() string {
        panic(errors.New(fmt.Sprint("validationInfo")))
}


func (this *GatewayValidationView) ToXmlDoc() *Document {
        return nil
}


func (this *GatewayValidationView) ToXmlNode(document *Document) *Node {
        return nil
}


// throws Exception
func (this *GatewayValidationView) View() string {
        return NewGatewayView(this.GetTransformInfoInterface()).View()
}

