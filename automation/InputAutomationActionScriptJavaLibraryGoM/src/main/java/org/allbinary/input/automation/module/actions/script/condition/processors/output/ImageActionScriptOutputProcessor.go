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

type ImageActionScriptOutputProcessor struct {
    LogUtil *LogUtil
}


func (this *ImageActionScriptOutputProcessor) initFields() {
    this.LogUtil = LogUtil_GetInstance()
}


func NewImageActionScriptOutputProcessor() *ImageActionScriptOutputProcessor {
    this := &ImageActionScriptOutputProcessor{}
    this.initFields()
    return this
}


// throws Exception
func ImageActionScriptOutputProcessor_Process(imageActionScriptOutputInterface ImageActionScriptOutputInterface, frame int64) {
        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        logUtil.PutF("Start - Processing at: " + fmt.Sprint(imageActionScriptOutputInterface), "ImageActionScriptOutputProcessor", commonStrings.PROCESS)
        var imageTypes *ImageTypes = imageActionScriptOutputInterface.GetImageTypes(); _ = imageTypes
        var vector IBasicArrayList = imageTypes.GetVector(); _ = vector
        var size int = vector.Size(); _ = size
        for index := int(0); index < size; index++ {
            var imageType *InputImageType = any(vector.Get(index)).(*InputImageType); _ = imageType
            var imageIOInterface ImageIOInterface = InputImageIOInterfaceFactory_GetInstance(imageType); _ = imageIOInterface
            imageIOInterface.Save(frame)
        }
}

