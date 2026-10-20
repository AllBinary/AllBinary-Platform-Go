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

type DefaultListModelHelper struct {
    DefaultListModel *DefaultListModel
    Vector IBasicArrayList
}


func (this *DefaultListModelHelper) initFields() {
    
    
}


func NewDefaultListModelHelper() *DefaultListModelHelper {
    this := &DefaultListModelHelper{}
    this.initFields()
        this.Vector = NewBasicArrayListD()
        this.DefaultListModel = NewDefaultListModel()
    return this
}


func (this *DefaultListModelHelper) InitDefaultModelList() {
        this.DefaultListModel = NewDefaultListModel()
        var size int = this.Vector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var moduleName string = any(this.Vector.Get(index)).(string); _ = moduleName
            this.DefaultListModel.AddElement(moduleName)
        }
}


func (this *DefaultListModelHelper) GetListModel() *ListModel {
        return any(this.DefaultListModel).(*ListModel)
}


func (this *DefaultListModelHelper) Add(string_ string) {
        this.Vector.Add(string_)
}


func (this *DefaultListModelHelper) Remove(string_ string) {
        this.Vector.Remove(string_)
}

