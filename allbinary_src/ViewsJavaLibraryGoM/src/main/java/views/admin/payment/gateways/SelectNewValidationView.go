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

type SelectNewValidationView struct {
    HttpStoreComponentView
}

// Implements ValidationComponentInterface



func NewSelectNewValidationView(transformInfoInterface TransformInfoInterface) *SelectNewValidationView {
    this := &SelectNewValidationView{}
        this.HttpStoreComponentView = *NewHttpStoreComponentView(transformInfoInterface)
    return this
}


func (this *SelectNewValidationView) IsValid() bool {
        return true
}


func (this *SelectNewValidationView) ToValidationInfoDoc() *Document {
        return nil
}


func (this *SelectNewValidationView) ToValidationInfoNode(document *Document) *Node {
        return nil
}


// throws Exception
func (this *SelectNewValidationView) ValidationInfo() string {
        panic(errors.New(fmt.Sprint("validationInfo")))
}


func (this *SelectNewValidationView) ToXmlDoc() *Document {
        return nil
}


func (this *SelectNewValidationView) ToXmlNode(document *Document) *Node {
        return nil
}


// throws Exception
func (this *SelectNewValidationView) View() string {
        return NewSelectNewView(this.GetTransformInfoInterface()).View()
}

