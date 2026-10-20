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

type AbIOSystem struct {
    Type string
}

var AbIOSystem_Instance *AbIOSystem = NewAbIOSystem()

func (this *AbIOSystem) initFields() {
    this.Type = "java.io"
}

func NewAbIOSystem() *AbIOSystem {
    this := &AbIOSystem{}
    this.initFields()
    return this
}


func AbIOSystem_GetInstance() *AbIOSystem {
        return AbIOSystem_Instance
}


func (this *AbIOSystem) GetType() string {
        return this.Type
}


func (this *AbIOSystem) IsType(type_ string) bool {
        if strings.Compare(this.Type, type_) == 0 {
            return true
        } else {
            return false
        }
}

