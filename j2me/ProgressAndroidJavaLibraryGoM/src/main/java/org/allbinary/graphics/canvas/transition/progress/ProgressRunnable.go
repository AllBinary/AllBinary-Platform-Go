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

type ProgressRunnable struct {
    ARunnable
    MidletActivity SimpleProgressActivityInterface
    ProgressCanvas IProgressCanvas
}

// The Java class ProgressRunnable has subclasses so its variables are this interface
type IProgressRunnable interface {
    IARunnable
    ProgressRunnable_() *ProgressRunnable
    Run()
}

func (this *ProgressRunnable) ProgressRunnable_() *ProgressRunnable {
    return this
}


func (this *ProgressRunnable) initFields() {
    
    
}


func NewProgressRunnable(midletActivity *Activity, progressCanvas IProgressCanvas) *ProgressRunnable {
    this := &ProgressRunnable{}
    this.ARunnable = *NewARunnable()
    this.initFields()
        var midletActivity2 IMidletActivityBase = any(midletActivity).(IMidletActivityBase); _ = midletActivity2
        this.MidletActivity = any(midletActivity2).(SimpleProgressActivityInterface)
        this.ProgressCanvas = progressCanvas
    return this
}


func (this *ProgressRunnable) Run() {
}

