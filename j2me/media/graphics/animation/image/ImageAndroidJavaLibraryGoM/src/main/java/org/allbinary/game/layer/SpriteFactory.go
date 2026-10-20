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

type SpriteFactory struct {
    LogUtil *LogUtil
}

// Implements SpriteFactoryInterface

var SpriteFactory_Instance *SpriteFactory = NewSpriteFactory()

func (this *SpriteFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}

func NewSpriteFactory() *SpriteFactory {
    this := &SpriteFactory{}
    this.initFields()
    return this
}


func SpriteFactory_GetInstance() *SpriteFactory {
        return SpriteFactory_Instance
}


func (this *SpriteFactory) CreateSprite(image IImage, frameWidth int, frameHeight int) *Sprite {
        return NewSprite(image, frameWidth, frameHeight, image.GetWidth() / frameWidth, image.GetHeight() / frameHeight)
}

