/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2022 AllBinary 
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

type ABRunnable struct {
    Running bool
}

// Implements RunnableInterface

// The Java class ABRunnable has subclasses so its variables are this interface
type IABRunnable interface {
    ABRunnable_() *ABRunnable
    SetThread(thread IThread)
    IsRunning() bool
    SetRunning(running bool)
    Run()
    GetType() int
}

func (this *ABRunnable) ABRunnable_() *ABRunnable {
    return this
}


func (this *ABRunnable) initFields() {
    this.Running = false
}

func NewABRunnable() *ABRunnable {
    this := &ABRunnable{}
    this.initFields()
    return this
}


func (this *ABRunnable) SetThread(thread IThread) {
}


// synchronized
func (this *ABRunnable) IsRunning() bool {
        return this.Running
}


// synchronized
func (this *ABRunnable) SetRunning(running bool) {
        this.Running = running
}


func (this *ABRunnable) Run() {
}


func (this *ABRunnable) GetType() int {
        return -1
}
