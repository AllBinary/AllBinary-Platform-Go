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

type ChoiceGroupImageUtil struct {
}

var ChoiceGroupImageUtil_Instance *ChoiceGroupImageUtil = NewChoiceGroupImageUtil()

func NewChoiceGroupImageUtil() *ChoiceGroupImageUtil {
    this := &ChoiceGroupImageUtil{}
    return this
}


func ChoiceGroupImageUtil_GetInstance() *ChoiceGroupImageUtil {
        return ChoiceGroupImageUtil_Instance
}


func (this *ChoiceGroupImageUtil) Init() {
        var features *Features = Features_GetInstance(); _ = features
        var choiceGroupFactory *ChoiceGroupImageFactory = ChoiceGroupImageFactory_GetInstance(); _ = choiceGroupFactory
        var imageArray []IImage = choiceGroupFactory.GetImageArray(); _ = imageArray
        if features.IsDefault(OpenGLFeatureFactory_GetInstance().OPENGL) {
            var preResourceImageUtil *PreResourceImageUtil = PreResourceImageUtil_GetInstance(); _ = preResourceImageUtil
            var size int = len(imageArray); _ = size
            for index := int(0); index < size; index++ {
                imageArray[index] = preResourceImageUtil.Encapsulate(imageArray[index])
            }
            choiceGroupFactory.Init(imageArray)
        }
}


func (this *ChoiceGroupImageUtil) Update_1(gl GL) {
        var features *Features = Features_GetInstance(); _ = features
        var choiceGroupFactory *ChoiceGroupImageFactory = ChoiceGroupImageFactory_GetInstance(); _ = choiceGroupFactory
        var imageArray []IImage = choiceGroupFactory.GetImageArray(); _ = imageArray
        if features.IsDefault(OpenGLFeatureFactory_GetInstance().OPENGL) {
            var preResourceImageUtil *PreResourceImageUtil = PreResourceImageUtil_GetInstance(); _ = preResourceImageUtil
            var size int = len(imageArray); _ = size
            for index := int(0); index < size; index++ {
                preResourceImageUtil.UpdateGL(gl, imageArray[index])
            }
        }
        choiceGroupFactory.Init(imageArray)
}

// Java overloads of ChoiceGroupImageUtil.update
func (this *ChoiceGroupImageUtil) Update(args ...any) {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 1 && Java_IsArg[GL](args[0], pass) {
            this.Update_1(Java_Arg[GL](args[0]))
            return
        }
        if len(args) == 1 && Java_IsArg[IGraphics](args[0], pass) {
            this.Update_2(Java_Arg[IGraphics](args[0]))
            return
        }
    }
    panic("No Java overload of ChoiceGroupImageUtil.update for the arguments")
}


func (this *ChoiceGroupImageUtil) Update_2(graphics IGraphics) {
        var features *Features = Features_GetInstance(); _ = features
        var choiceGroupFactory *ChoiceGroupImageFactory = ChoiceGroupImageFactory_GetInstance(); _ = choiceGroupFactory
        var imageArray []IImage = choiceGroupFactory.GetImageArray(); _ = imageArray
        if features.IsDefault(OpenGLFeatureFactory_GetInstance().OPENGL) {
            var preResourceImageUtil *PreResourceImageUtil = PreResourceImageUtil_GetInstance(); _ = preResourceImageUtil
            var size int = len(imageArray); _ = size
            for index := int(0); index < size; index++ {
                preResourceImageUtil.Update(graphics, imageArray[index])
            }
        }
        choiceGroupFactory.Init(imageArray)
}

