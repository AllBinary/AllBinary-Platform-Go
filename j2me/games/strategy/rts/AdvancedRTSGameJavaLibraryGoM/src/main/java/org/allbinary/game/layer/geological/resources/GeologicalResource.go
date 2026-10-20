/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2006 AllBinary 
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

type GeologicalResource struct {
    Total int
    InitTotal int
}


func (this *GeologicalResource) initFields() {
    
    
}


func NewGeologicalResource(initTotal int) *GeologicalResource {
    this := &GeologicalResource{}
    this.initFields()
        this.InitTotal = initTotal
        this.Init()
    return this
}


func (this *GeologicalResource) Init() {
        this.SetTotal(this.InitTotal)
}


func (this *GeologicalResource) Add(value int) {
        var max_ int = Integer_MAX_VALUE; _ = max_
        if this.GetTotal() + value > max_ {
            this.SetTotal(max_)
        } else {
            this.SetTotal(this.GetTotal() + value)
        }
}


func (this *GeologicalResource) Remove(value int) {
        if this.GetTotal() - value < 0 {
            this.SetTotal(0)
        } else {
            this.SetTotal(this.GetTotal() - value)
        }
}


func (this *GeologicalResource) GetTotal() int {
        return this.Total
}


func (this *GeologicalResource) SetTotal(total int) {
        this.Total = total
}

