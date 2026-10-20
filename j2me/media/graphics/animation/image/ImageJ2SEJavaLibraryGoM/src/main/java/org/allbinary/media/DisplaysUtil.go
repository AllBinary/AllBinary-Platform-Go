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

type DisplaysUtil struct {
}

var DisplaysUtil_Instance *DisplaysUtil = NewDisplaysUtil()

func NewDisplaysUtil() *DisplaysUtil {
    this := &DisplaysUtil{}
    return this
}


func DisplaysUtil_GetInstance() *DisplaysUtil {
        return DisplaysUtil_Instance
}


func (this *DisplaysUtil) GetPrimaryScreenSize() *Rectangle {
        var dimension *Dimension = Toolkit_GetDefaultToolkit().GetScreenSize(); _ = dimension
        return NewRectangle(PointFactory_GetInstance().ZERO_ZERO, dimension.Width, dimension.Height)
}


func (this *DisplaysUtil) GetScreenSizesAsRectangleArray() []*Rectangle {
        var graphicsEnvironment *GraphicsEnvironment = GraphicsEnvironment_GetLocalGraphicsEnvironment(); _ = graphicsEnvironment
        var graphicsDeviceArray []*GraphicsDevice = graphicsEnvironment.GetScreenDevices(); _ = graphicsDeviceArray
        var graphicsDevice *GraphicsDevice; _ = graphicsDevice
        var displayMode *DisplayMode; _ = displayMode
        var rectangle *Rectangle; _ = rectangle
        var size int = len(graphicsDeviceArray); _ = size
        var rectangleArray []*Rectangle = make([]*Rectangle, size); _ = rectangleArray
        for index := int(0); index < size; index++ {
            graphicsDevice = graphicsDeviceArray[index]
            displayMode = graphicsDevice.GetDisplayMode()
            rectangle = NewRectangle(PointFactory_GetInstance().ZERO_ZERO, displayMode.GetWidth(), displayMode.GetHeight())
            rectangleArray[index] = rectangle
        }
        return rectangleArray
}

