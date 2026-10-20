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

type BufferedGameCanvasPaintHelper struct {
    ProcessPaintable
    GameCanvas ICanvas
    OffScreenImage []IImage
    PreviousImage IImage
    CircularIndexUtil *CircularIndexUtil
    DrawCircularIndexUtil *CircularIndexUtil
    Anchor int
}

var BufferedGameCanvasPaintHelper_MAX_IMAGES int = 4

func (this *BufferedGameCanvasPaintHelper) initFields() {
    this.GameCanvas = NullCanvas_NULL_CANVAS
    this.OffScreenImage = make([]IImage, BufferedGameCanvasPaintHelper_MAX_IMAGES)
    this.PreviousImage = NullImage_NULL_IMAGE
    this.CircularIndexUtil = CircularIndexUtil_CreateInstance(BufferedGameCanvasPaintHelper_MAX_IMAGES)
    this.DrawCircularIndexUtil = CircularIndexUtil_CreateInstanceAt(BufferedGameCanvasPaintHelper_MAX_IMAGES - 1, BufferedGameCanvasPaintHelper_MAX_IMAGES)
    this.Anchor = Anchor_TOP_LEFT
}


func NewBufferedGameCanvasPaintHelper(gameCanvas IAllBinaryGameCanvas) *BufferedGameCanvasPaintHelper {
    this := &BufferedGameCanvasPaintHelper{}
    this.ProcessPaintable = *NewProcessPaintable()
    this.initFields()
        panic(errors.New(fmt.Sprint("No Longer Used")))
    return this
}


func (this *BufferedGameCanvasPaintHelper) Process() {
        var myCanvas IMyCanvas = (any(this.GameCanvas).(IMyCanvas)); _ = myCanvas
        myCanvas.Draw(this.OffScreenImage[this.CircularIndexUtil.GetIndex()].GetGraphics())
        this.CircularIndexUtil.Next()
        this.DrawCircularIndexUtil.Next()
}


func (this *BufferedGameCanvasPaintHelper) Paint(graphics IGraphics) {
        this.PreviousImage = this.OffScreenImage[this.DrawCircularIndexUtil.GetIndex()]
        graphics.DrawImage(this.PreviousImage, 0, 0, this.Anchor)
}

