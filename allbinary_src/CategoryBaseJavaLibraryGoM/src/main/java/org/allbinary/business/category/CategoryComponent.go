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

type CategoryComponent struct {
    CategoryView
    NO_TYPE int
}

// Implements TransformInterface


func (this *CategoryComponent) initFields() {
    this.NO_TYPE = 0
}


func NewCategoryComponent(categoryInterface CategoryInterface) *CategoryComponent {
    this := &CategoryComponent{}
        this.CategoryView = *NewCategoryView(categoryInterface)
        this.initFields()
    return this
}


func (this *CategoryComponent) GetTypeId() int {
        return this.NO_TYPE
}


func (this *CategoryComponent) GetTransformDocumentInterface() TransformDocumentInterface {
        return nil
}


// throws Exception
func (this *CategoryComponent) GetTransformInfoInterface() TransformInfoInterface {
        return nil
}


func (this *CategoryComponent) SetTransformDocumentInterface(viewDocumentInterface TransformDocumentInterface) {
}


// throws Exception
func (this *CategoryComponent) ToXmlDoc() *Document {
        var document *Document = DomDocumentHelper_Create(); _ = document
        document.AppendChild(NewCategoryView(this.GetCategoryInterface()).ToXmlNode(document))
        return document
}


// throws Exception
func (this *CategoryComponent) View() string {
        return nil
}

