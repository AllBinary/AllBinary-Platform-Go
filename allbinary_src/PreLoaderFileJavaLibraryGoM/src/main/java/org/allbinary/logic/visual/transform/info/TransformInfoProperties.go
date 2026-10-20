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

type TransformInfoProperties struct {
    Name string
    Label string
    Description string
    ObjectFileName string
    ObjectConfigFileName string
    TemplateFileName string
}


func (this *TransformInfoProperties) initFields() {
    
    
    
    
    
    
}


func NewTransformInfoProperties(name string, label string, description string, objectFileName string, objectConfigFileName string, templateFileName string) *TransformInfoProperties {
    this := &TransformInfoProperties{}
    this.initFields()
        this.Name = name
        this.Label = label
        this.Description = description
        this.ObjectFileName = objectFileName
        this.TemplateFileName = templateFileName
        this.ObjectConfigFileName = objectConfigFileName
    return this
}


func (this *TransformInfoProperties) SetName(value string) {
        this.Name = value
}


func (this *TransformInfoProperties) SetDescription(value string) {
        this.Description = value
}


func (this *TransformInfoProperties) SetViewFile(value string) {
        this.ObjectFileName = value
}


func (this *TransformInfoProperties) SetTemplateFile(value string) {
        this.TemplateFileName = value
}


func (this *TransformInfoProperties) SetObjectConfigFile(value string) {
        this.ObjectConfigFileName = value
}


func (this *TransformInfoProperties) GetName() string {
        return this.Name
}


func (this *TransformInfoProperties) GetLabel() string {
        return this.Label
}


func (this *TransformInfoProperties) GetDescription() string {
        return this.Description
}


func (this *TransformInfoProperties) GetViewFile() string {
        return this.ObjectFileName
}


func (this *TransformInfoProperties) GetTemplateFile() string {
        return this.TemplateFileName
}


func (this *TransformInfoProperties) GetObjectConfigFile() string {
        return this.ObjectConfigFileName
}

