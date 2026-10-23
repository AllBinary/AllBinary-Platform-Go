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

//Hack class should not exist
type TransformInfoBasic struct {
    AbContext
    StoreFrontInterface StoreFrontInterface
}

// Implements TransformInfoHttpInterface


func (this *TransformInfoBasic) initFields() {
    
}


func NewTransformInfoBasic(storeFrontInterface StoreFrontInterface, propertiesHashMap *HashMap, pageContext *PageContext) *TransformInfoBasic {
    this := &TransformInfoBasic{}
        this.AbContext = *NewAbContext(propertiesHashMap, pageContext)
        this.initFields()
        this.StoreFrontInterface = storeFrontInterface
    return this
}


func (this *TransformInfoBasic) GetStoreName() string {
        return this.StoreFrontInterface.GetName()
}


func (this *TransformInfoBasic) Log() string {
        return this.CommonStrings.NOT_IMPLEMENTED
}


func (this *TransformInfoBasic) Override(hashMap *HashMap) {
}


func (this *TransformInfoBasic) GetName() string {
        return nil
}


func (this *TransformInfoBasic) GetObjectFile() string {
        return nil
}


func (this *TransformInfoBasic) GetObject() any {
        return nil
}


func (this *TransformInfoBasic) GetObjectConfigInterface() TransformInfoObjectConfigInterface {
        return nil
}


func (this *TransformInfoBasic) SetObjectConfigInterface(transformInfoObjectConfigInterface TransformInfoObjectConfigInterface) {
}


func (this *TransformInfoBasic) GetTemplateFilePath() IAbPath {
        return nil
}


func (this *TransformInfoBasic) GetTemplateFile() string {
        return nil
}


func (this *TransformInfoBasic) GetTemplate() string {
        return nil
}


func (this *TransformInfoBasic) GetDataFilePath() IAbPath {
        return nil
}


func (this *TransformInfoBasic) GetDataFile() string {
        return nil
}


// throws Exception
func (this *TransformInfoBasic) GetDataDocument() *Document {
        panic(errors.New(fmt.Sprint(this.CommonStrings.NOT_IMPLEMENTED)))
}


func (this *TransformInfoBasic) SetName(value string) {
}


func (this *TransformInfoBasic) SetStoreName(value string) {
}


func (this *TransformInfoBasic) SetObjectFile(value string) {
}


func (this *TransformInfoBasic) SetObject(anyType any) {
}


func (this *TransformInfoBasic) SetObjectConfigFile(value string) {
}


func (this *TransformInfoBasic) SetTemplateFile(value string) {
}


func (this *TransformInfoBasic) SetTemplate(value string) {
}


func (this *TransformInfoBasic) SetDataFile(value string) {
}


func (this *TransformInfoBasic) SetData(value string) {
}


func (this *TransformInfoBasic) SetChild() {
}


func (this *TransformInfoBasic) IsChild() bool {
        return true
}


func (this *TransformInfoBasic) GetImportUriPath() string {
        return nil
}


func (this *TransformInfoBasic) SetImportUriPath(importUriPath string) {
}


// throws Exception
func (this *TransformInfoBasic) GetKey() any {
        return nil
}


// throws Exception
func (this *TransformInfoBasic) ToVector() IBasicArrayList {
        return nil
}


// throws Exception
func (this *TransformInfoBasic) ToHashMap() *HashMap {
        return nil
}

