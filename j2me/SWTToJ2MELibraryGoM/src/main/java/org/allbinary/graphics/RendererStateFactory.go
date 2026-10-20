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

type RendererStateFactory struct {
    Loaded bool
    SurfaceChanged bool
}

var RendererStateFactory_Instance *RendererStateFactory = NewRendererStateFactory()

func (this *RendererStateFactory) initFields() {
    this.Loaded = false
    this.SurfaceChanged = false
}

func NewRendererStateFactory() *RendererStateFactory {
    this := &RendererStateFactory{}
    this.initFields()
    return this
}


func RendererStateFactory_GetInstance() *RendererStateFactory {
        return RendererStateFactory_Instance
}


func (this *RendererStateFactory) SetLoaded(loaded bool) {
        this.Loaded = loaded
}


func (this *RendererStateFactory) IsLoaded() bool {
        return this.Loaded
}


func (this *RendererStateFactory) SetSurfaceChanged(initialized bool) {
        this.SurfaceChanged = initialized
}


func (this *RendererStateFactory) IsSurfaceChanged() bool {
        return this.SurfaceChanged
}

