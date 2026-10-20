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

type ScreenScavangerRobot struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
}


func (this *ScreenScavangerRobot) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
}


func NewScreenScavangerRobot() *ScreenScavangerRobot {
    this := &ScreenScavangerRobot{}
    this.initFields()
    return this
}


// throws Exception
func (this *ScreenScavangerRobot) GetScreenAsBufferedImages() []*BufferedImage {
        this.LogUtil.PutF(this.CommonStrings.START, this, "getScreenAsBufferedImages")
        var dimScreenSize *Dimension = Toolkit_GetDefaultToolkit().GetScreenSize(); _ = dimScreenSize
        var rectScreenSize *Rectangle = NewRectangle(dimScreenSize); _ = rectScreenSize
        var robotHashtable *ABHashtable[any, any] = InputRobotFactory_GetInstance().Get().(*ABHashtable); _ = robotHashtable
        var bufferedImageArray []*BufferedImage = make([]*BufferedImage, 1); _ = bufferedImageArray
        var index int = 0; _ = index
        var inputRobotInterface InputRobotInterface = any(robotHashtable.Get("Java Robot")).(InputRobotInterface); _ = inputRobotInterface
        bufferedImageArray[index] = inputRobotInterface.CreateScreenCapture(rectScreenSize)
        this.LogUtil.PutF("Finish", this, "getScreenAsBufferedImages")
        return bufferedImageArray
}

