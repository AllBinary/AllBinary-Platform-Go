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

type GameGraphicsResourceUtil struct {
    LogUtil *LogUtil
    SPRITE string
    QUARTER string
    Name string
}

var GameGraphicsResourceUtil_Instance any = NullUtil_GetInstance().NULL_OBJECT

func (this *GameGraphicsResourceUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.SPRITE = "_sprite"
    this.QUARTER = "_quarter"
    this.Name = StringUtil_GetInstance().EMPTY_STRING
}


func GameGraphicsResourceUtil_GetInstance() *GameGraphicsResourceUtil {
        if GameGraphicsResourceUtil_Instance == NullUtil_GetInstance().NULL_OBJECT {
            GameGraphicsResourceUtil_Instance = NewGameGraphicsResourceUtil()
        }
        return any(GameGraphicsResourceUtil_Instance).(*GameGraphicsResourceUtil)
}


func NewGameGraphicsResourceUtil() *GameGraphicsResourceUtil {
    this := &GameGraphicsResourceUtil{}
    this.initFields()
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.CONSTRUCTOR, e)
                }
            }()
            this.Name = this.GetString()
            return
        }(); tryControl1 == 1 {
            return
        }
    return this
}


// throws Exception
func (this *GameGraphicsResourceUtil) GetString() string {
        var graphicsFeatureFactory *GraphicsFeatureFactory = GraphicsFeatureFactory_GetInstance(); _ = graphicsFeatureFactory
        var graphicsFeature *GraphicsFeature = graphicsFeatureFactory.NONE; _ = graphicsFeature
        var features *Features = Features_GetInstance(); _ = features
        if features.IsFeature(graphicsFeatureFactory.IMAGE_TO_ARRAY_GRAPHICS) {
            graphicsFeature = graphicsFeatureFactory.IMAGE_TO_ARRAY_GRAPHICS
        } else if features.IsFeature(graphicsFeatureFactory.SPRITE_FULL_GRAPHICS) {
            graphicsFeature = graphicsFeatureFactory.SPRITE_FULL_GRAPHICS
        } else if features.IsFeature(graphicsFeatureFactory.SPRITE_QUARTER_ROTATION_GRAPHICS) {
            graphicsFeature = graphicsFeatureFactory.SPRITE_QUARTER_ROTATION_GRAPHICS
        } else if features.IsFeature(graphicsFeatureFactory.IMAGE_ROTATION_ON_THE_FLY) {
            graphicsFeature = graphicsFeatureFactory.IMAGE_ROTATION_ON_THE_FLY
        }
        return this.GetStringForFeature(graphicsFeature)
}


// throws Exception
func (this *GameGraphicsResourceUtil) GetStringForFeature(graphicsFeature *GraphicsFeature) string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var graphicsFeatureFactory *GraphicsFeatureFactory = GraphicsFeatureFactory_GetInstance(); _ = graphicsFeatureFactory
        stringBuffer.Delete(0, stringBuffer.Length())
        var features *Features = Features_GetInstance(); _ = features
        if features.IsFeature(graphicsFeatureFactory.VECTOR_GRAPHICS) {
        } else if features.IsFeature(graphicsFeatureFactory.IMAGE_GRAPHICS) {
            if graphicsFeature == graphicsFeatureFactory.IMAGE_TO_ARRAY_GRAPHICS || graphicsFeature == graphicsFeatureFactory.IMAGE_ROTATION_ON_THE_FLY {
            } else if graphicsFeature == graphicsFeatureFactory.SPRITE_FULL_GRAPHICS {
                stringBuffer.Append(this.SPRITE)
            } else if graphicsFeature == graphicsFeatureFactory.SPRITE_QUARTER_ROTATION_GRAPHICS {
                stringBuffer.Append(this.QUARTER)
                stringBuffer.Append(this.SPRITE)
            } else {
                panic(errors.New(fmt.Sprint("None/Unknown Sub Image Resource Type Specified")))
            }
        } else if OpenGLFeatureUtil_GetInstance().IsAnyThreed() {
            stringBuffer.Append(CanvasStrings_GetInstance().OBJ_MODEL)
        } else {
            panic(errors.New(fmt.Sprint("None/Unknown Main Image Resource Type Specified")))
        }
        return fmt.Sprint(stringBuffer)
}


func (this *GameGraphicsResourceUtil) GetName() string {
        return this.Name
}

