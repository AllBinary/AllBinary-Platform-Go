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

type SavedCaptureGenericProfileDataWorkerType struct {
    GenericProfileDataWorkerType
    Path string
    LastFrame int
}

var SavedCaptureGenericProfileDataWorkerType_SAVED_CAPTURE *SavedCaptureGenericProfileDataWorkerType = NewSavedCaptureGenericProfileDataWorkerType("Saved " + InputImageType_CAPTURE.GetName(), "c:\\share\\captures\\2\\", 100)

func (this *SavedCaptureGenericProfileDataWorkerType) initFields() {
    
    
}


func NewSavedCaptureGenericProfileDataWorkerType(name string, path string, lastFrame int) *SavedCaptureGenericProfileDataWorkerType {
    this := &SavedCaptureGenericProfileDataWorkerType{}
        this.GenericProfileDataWorkerType = *NewGenericProfileDataWorkerType(name)
        this.initFields()
        this.SetPath(path)
        this.SetLastFrame(lastFrame)
    return this
}


func (this *SavedCaptureGenericProfileDataWorkerType) GetPath() string {
        return this.Path
}


func (this *SavedCaptureGenericProfileDataWorkerType) SetPath(path string) {
        this.Path = path
}


func (this *SavedCaptureGenericProfileDataWorkerType) GetLastFrame() int {
        return this.LastFrame
}


func (this *SavedCaptureGenericProfileDataWorkerType) SetLastFrame(lastFrame int) {
        this.LastFrame = lastFrame
}

