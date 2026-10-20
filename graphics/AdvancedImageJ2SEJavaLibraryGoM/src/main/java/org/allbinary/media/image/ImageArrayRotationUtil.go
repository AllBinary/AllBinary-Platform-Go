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

type ImageArrayRotationUtil struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    TOTAL_ANGLE string
}

var ImageArrayRotationUtil_Instance *ImageArrayRotationUtil = NewImageArrayRotationUtil()

func (this *ImageArrayRotationUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.TOTAL_ANGLE = "totalAngle: "
}

func NewImageArrayRotationUtil() *ImageArrayRotationUtil {
    this := &ImageArrayRotationUtil{}
    this.initFields()
    return this
}


func ImageArrayRotationUtil_GetInstance() *ImageArrayRotationUtil {
        return ImageArrayRotationUtil_Instance
}


// throws IOException
func (this *ImageArrayRotationUtil) Process(imageProcessorInput *ImageProcessorInput, input string, visitor ImageProcessedVisitor) {
        var totalAngle int; _ = totalAngle
        if input == this.CommonStrings.UP {
            totalAngle = fmt.Sprint(-90)
        } else if input == this.CommonStrings.DOWN {
            totalAngle = fmt.Sprint(90)
        } else {
            totalAngle = fmt.Sprint(input)
        }
        var generatedBufferedImage *BufferedImage; _ = generatedBufferedImage
        var bufferedImageArray []*BufferedImage = imageProcessorInput.GetBufferedImageArray(); _ = bufferedImageArray
        for index := int(0); index < len(bufferedImageArray); index++ {
            this.LogUtil.PutF(this.TOTAL_ANGLE + fmt.Sprint(totalAngle), this, this.CommonStrings.RUN)
            generatedBufferedImage = ImageJ2SERotationUtil_GetInstance().GetRotatedImage(bufferedImageArray[index], int(totalAngle))
            visitor.Visit(generatedBufferedImage, input, index)
        }
}

