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

type HighScoresPaintable struct {
    Paintable
    DisplayInfoSingleton *DisplayInfoSingleton
    MyFontProcessor IMyFontProcessor
    BasicColor *BasicColor
    HighScores IHighScores
    Anchor int
    CharHeight int
}

// Implements ColorChangeListener, UpdateMyFontInterface


func (this *HighScoresPaintable) initFields() {
    this.DisplayInfoSingleton = DisplayInfoSingleton_GetInstance()
    this.MyFontProcessor = NewUpdateMyFontProcessor(this)
    this.BasicColor = BasicColorFactory_GetInstance().WHITE
    this.HighScores = NullHighScoresSingletonFactory_GetInstance()
    this.Anchor = Anchor_TOP_LEFT
    
}


func NewHighScoresPaintable() *HighScoresPaintable {
    this := &HighScoresPaintable{}
    this.Paintable = *NewPaintable()
    this.initFields()
    return this
}


func (this *HighScoresPaintable) UpdateMeasurement(graphics IGraphics) {
        var font Font = graphics.GetFont(); _ = font
        this.CharHeight = font.GetHeight()
        this.MyFontProcessor = MyFontProcessor_GetInstance()
}


func (this *HighScoresPaintable) OnEvent(eventObject IAllBinaryEventObject) {
        var colorChangeEvent *ColorChangeEvent = any(eventObject).(*ColorChangeEvent); _ = colorChangeEvent
        this.BasicColor = colorChangeEvent.GetBasicColorP()
}


func (this *HighScoresPaintable) Paint(graphics IGraphics) {
        this.MyFontProcessor.Process(graphics)
        var width int = this.DisplayInfoSingleton.GetLastWidth(); _ = width
        var height int = this.DisplayInfoSingleton.GetLastHeight(); _ = height
        graphics.SetColor(this.GetBasicColorP().IntValue())
        var heading string = this.HighScores.GetHeading(); _ = heading
        var topScoresWidth int = (graphics.GetFont().StringWidth(heading) >> 1); _ = topScoresWidth
        var charHeight int = this.CharHeight; _ = charHeight
        graphics.DrawString(heading, (width >> 1) - topScoresWidth, charHeight, this.Anchor)
        graphics.DrawString(this.HighScores.GetColumnOneHeading(), 10, charHeight * 3, this.Anchor)
        var columnTwoHeading string = this.HighScores.GetColumnTwoHeading(); _ = columnTwoHeading
        var columnTwoHeadingWidth int = graphics.GetFont().StringWidth(columnTwoHeading); _ = columnTwoHeadingWidth
        graphics.DrawString(columnTwoHeading, width - 10 - columnTwoHeadingWidth, charHeight * 3, this.Anchor)
        var index int = 4; _ = index
        var largestSecondColumnWidth int = columnTwoHeadingWidth; _ = largestSecondColumnWidth
        var list IBasicArrayList = this.HighScores.GetList(); _ = list
        var size int = list.Size(); _ = size
        var vectorIndex int = 0; _ = vectorIndex
        var highScore *HighScore; _ = highScore
        for vectorIndex < size && charHeight * index < height - (charHeight * 2) {
            highScore = any(list.BasicArrayList_().ObjectArray[vectorIndex]).(*HighScore)
            var nextScoreWidth int = graphics.GetFont().StringWidth(highScore.GetScoreString()); _ = nextScoreWidth
            if nextScoreWidth > largestSecondColumnWidth {
                largestSecondColumnWidth = nextScoreWidth
            }
            vectorIndex++
        }
        vectorIndex = 0
        for vectorIndex < size && charHeight * index < height - (charHeight * 2) {
            highScore = any(list.BasicArrayList_().ObjectArray[vectorIndex]).(*HighScore)
            graphics.DrawString(highScore.GetName(), 10, charHeight * index, this.Anchor)
            graphics.DrawString(highScore.GetScoreString(), width - 10 - largestSecondColumnWidth, charHeight * index, this.Anchor)
            index++
            vectorIndex++
        }
}


func (this *HighScoresPaintable) SetBasicColorP(basicColor *BasicColor) {
        this.BasicColor = basicColor
}


func (this *HighScoresPaintable) GetBasicColorP() *BasicColor {
        return this.BasicColor
}


func (this *HighScoresPaintable) SetHighScores(highScores IHighScores) {
        this.HighScores = highScores
}

