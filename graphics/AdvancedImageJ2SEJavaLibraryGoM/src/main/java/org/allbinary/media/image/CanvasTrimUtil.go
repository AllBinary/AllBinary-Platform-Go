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

type CanvasTrimUtil struct {
    LogUtil *LogUtil
    StringUtil *StringUtil
}

var CanvasTrimUtil_Instance *CanvasTrimUtil = NewCanvasTrimUtil()

func (this *CanvasTrimUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.StringUtil = StringUtil_GetInstance()
}

func NewCanvasTrimUtil() *CanvasTrimUtil {
    this := &CanvasTrimUtil{}
    this.initFields()
    return this
}


func CanvasTrimUtil_GetInstance() *CanvasTrimUtil {
        return CanvasTrimUtil_Instance
}


// throws IOException
func (this *CanvasTrimUtil) Process(imageProcessorInput *ImageProcessorInput, cutCanvasFromLeftReduction int, cutCanvasFromTopReduction int, cutCanvasFromRightReduction int, cutCanvasFromBottomReduction int, visitor ImageProcessedVisitor) {
        var bufferedImageArray []*BufferedImage = imageProcessorInput.GetBufferedImageArray(); _ = bufferedImageArray
        var subBufferedImageArray []*BufferedImage = make([]*BufferedImage, len(bufferedImageArray)); _ = subBufferedImageArray
        var size int = len(bufferedImageArray); _ = size
        var bufferedImage *BufferedImage; _ = bufferedImage
        for index := int(0); index < size; index++ {
            bufferedImage = bufferedImageArray[index]
            subBufferedImageArray[index] = bufferedImage.GetSubimage(cutCanvasFromLeftReduction, cutCanvasFromTopReduction, bufferedImage.GetWidth() - cutCanvasFromLeftReduction - cutCanvasFromRightReduction, bufferedImage.GetHeight() - cutCanvasFromTopReduction - cutCanvasFromBottomReduction)
            visitor.Visit(subBufferedImageArray[index], this.StringUtil.EMPTY_STRING, index)
        }
}

