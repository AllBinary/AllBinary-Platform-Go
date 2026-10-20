/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2026 AllBinary 
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

type ShapeTypeFactory struct {
    KEY_LINE string
    KEY_CIRCLE string
    KEY_RECTANGLE string
    KEY_RECT string
    LINE *ShapeType
    CIRCLE *ShapeType
    RECTANGLE *ShapeType
}

var ShapeTypeFactory_Instance *ShapeTypeFactory = NewShapeTypeFactory()

func (this *ShapeTypeFactory) initFields() {
    this.KEY_LINE = "line"
    this.KEY_CIRCLE = "circle"
    this.KEY_RECTANGLE = "rectangle"
    this.KEY_RECT = "rect"
    this.LINE = NewShapeType(this.KEY_LINE)
    this.CIRCLE = NewShapeType(this.KEY_CIRCLE)
    this.RECTANGLE = NewShapeType(this.KEY_RECTANGLE)
}

func NewShapeTypeFactory() *ShapeTypeFactory {
    this := &ShapeTypeFactory{}
    this.initFields()
    return this
}


func ShapeTypeFactory_GetInstance() *ShapeTypeFactory {
        return ShapeTypeFactory_Instance
}

