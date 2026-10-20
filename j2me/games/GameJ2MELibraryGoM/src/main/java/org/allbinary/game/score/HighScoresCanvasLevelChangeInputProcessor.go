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

type HighScoresCanvasLevelChangeInputProcessor struct {
    HighScoresCanvasInputProcessor
    DisplayInfoSingleton *DisplayInfoSingleton
    INSTRUCTIONS string
    MyFontProcessor IMyFontProcessor
    Anchor int
    FontHeight int
}

// Implements UpdateMyFontInterface


func (this *HighScoresCanvasLevelChangeInputProcessor) initFields() {
    this.DisplayInfoSingleton = DisplayInfoSingleton_GetInstance()
    this.INSTRUCTIONS = "(Right = Next Track, Left = Previous Track)"
    this.MyFontProcessor = NewUpdateMyFontProcessor(this)
    this.Anchor = Anchor_TOP_LEFT
    this.FontHeight = 0
}


func NewHighScoresCanvasLevelChangeInputProcessor(highScoresCanvas *HighScoresCanvas) *HighScoresCanvasLevelChangeInputProcessor {
    this := &HighScoresCanvasLevelChangeInputProcessor{}
        this.HighScoresCanvasInputProcessor = *NewHighScoresCanvasInputProcessor(highScoresCanvas)
        this.initFields()
    return this
}


func (this *HighScoresCanvasLevelChangeInputProcessor) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        this.FontHeight = font.GetHeight()
        this.MyFontProcessor = MyFontProcessor_GetInstance()
}


func (this *HighScoresCanvasLevelChangeInputProcessor) Open() {
        GameKeyEventHandler_GetInstance().AddListener(this)
}


func (this *HighScoresCanvasLevelChangeInputProcessor) Close() {
        GameKeyEventHandler_GetInstance().RemoveListener(this)
}


// synchronized
func (this *HighScoresCanvasLevelChangeInputProcessor) Update() {
        var list IBasicArrayList = this.GetGameKeyEventList(); _ = list
        var size int = list.Size(); _ = size
        for index := int(0); index < size; index++ {
            var anyType any = list.BasicArrayList_().ObjectArray[index]; _ = anyType
            var key int = GameKeyEventUtil_GetKey(anyType); _ = key
            if key == Canvas_RIGHT {
                this.GetHighScoresCanvas().GetGameInfo().NextGameLevel()
            } else if key == Canvas_LEFT {
                this.GetHighScoresCanvas().GetGameInfo().PreviousGameLevel()
            }
            this.GetHighScoresCanvas().ExecuteUpdate()
        }
        list.Clear()
}


func (this *HighScoresCanvasLevelChangeInputProcessor) Paint(graphics IGraphics) {
        this.MyFontProcessor.Process(graphics)
        var width int = this.DisplayInfoSingleton.GetLastWidth(); _ = width
        var topScoresWidth int = (graphics.GetFont().StringWidth(this.INSTRUCTIONS) >> 1); _ = topScoresWidth
        graphics.DrawString(this.INSTRUCTIONS, (width >> 1) - topScoresWidth, this.FontHeight * 2, this.Anchor)
}

