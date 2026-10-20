/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2002 AllBinary 
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

type Capital struct {
    Value int
}

// The Java class Capital has subclasses so its variables are this interface
type ICapital interface {
    Capital_() *Capital
    GetTotalMoney() int
    RemoveMoney(points int)
    AddMoney(points int)
}

func (this *Capital) Capital_() *Capital {
    return this
}


func (this *Capital) initFields() {
    
}


func NewCapital() *Capital {
    this := &Capital{}
    this.initFields()
    return this
}


func (this *Capital) GetTotalMoney() int {
        return this.Value
}


func (this *Capital) RemoveMoney(points int) {
        this.Value -= points
}


func (this *Capital) AddMoney(points int) {
        this.Value += points
}

