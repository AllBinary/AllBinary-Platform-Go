/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

type GraphicsStrings struct {
    ANIMATION string
    ANGLE string
    MOVEMENT_ANGLE string
    ROTATION string
    OPACITY string
    HTML string
    WIDTH string
    HEIGHT string
}

var GraphicsStrings_Instance *GraphicsStrings = NewGraphicsStrings()

func (this *GraphicsStrings) initFields() {
    this.ANIMATION = "animation"
    this.ANGLE = "angle"
    this.MOVEMENT_ANGLE = "movement_angle"
    this.ROTATION = "rotation"
    this.OPACITY = "opacity"
    this.HTML = "HTML"
    this.WIDTH = "width"
    this.HEIGHT = "height"
}

func NewGraphicsStrings() *GraphicsStrings {
    this := &GraphicsStrings{}
    this.initFields()
    return this
}


func GraphicsStrings_GetInstance() *GraphicsStrings {
        return GraphicsStrings_Instance
}

