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

type AllBinaryRendererBase3 struct {
    List IBasicArrayList
}

// The Java class AllBinaryRendererBase3 has subclasses so its variables are this interface
type IAllBinaryRendererBase3 interface {
    AllBinaryRendererBase3_() *AllBinaryRendererBase3
    Update(gl GL10)
    Add(image IImage)
}

func (this *AllBinaryRendererBase3) AllBinaryRendererBase3_() *AllBinaryRendererBase3 {
    return this
}


func (this *AllBinaryRendererBase3) initFields() {
    this.List = NewBasicArrayListD()
}

func NewAllBinaryRendererBase3() *AllBinaryRendererBase3 {
    this := &AllBinaryRendererBase3{}
    this.initFields()
    return this
}


// throws Exception
func (this *AllBinaryRendererBase3) Update(gl GL10) {
        for this.List.Size() > 0 {
            var openGLESImage IOpenGLESImage = any(this.List.RemoveAt(0)).(IOpenGLESImage); _ = openGLESImage
            if openGLESImage != nil {
                openGLESImage.Set(gl)
            }
        }
}


func (this *AllBinaryRendererBase3) Add(image IImage) {
        this.List.Add(image)
}

