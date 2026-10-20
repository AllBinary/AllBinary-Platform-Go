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

type HighScoresCanvasInputProcessor struct {
    PlayerGameInput
    HighScoresCanvas *HighScoresCanvas
}

// The Java class HighScoresCanvasInputProcessor has subclasses so its variables are this interface
type IHighScoresCanvasInputProcessor interface {
    IPlayerGameInput
    HighScoresCanvasInputProcessor_() *HighScoresCanvasInputProcessor
    Open()
    Close()
    OnPressGameKeyEvent(gameKeyEvent *GameKeyEvent)
    OnDownGameKeyEvent(gameKeyEvent *GameKeyEvent)
    OnUpGameKeyEvent(gameKeyEvent *GameKeyEvent)
    Paint(graphics IGraphics)
    GetHighScoresCanvas() *HighScoresCanvas
}

func (this *HighScoresCanvasInputProcessor) HighScoresCanvasInputProcessor_() *HighScoresCanvasInputProcessor {
    return this
}


func (this *HighScoresCanvasInputProcessor) initFields() {
    
}


func NewHighScoresCanvasInputProcessor(highScoresCanvas *HighScoresCanvas) *HighScoresCanvasInputProcessor {
    this := &HighScoresCanvasInputProcessor{}
        this.PlayerGameInput = *NewPlayerGameInput(NewBasicArrayListD(), NewBasicArrayListD(), -1)
        this.initFields()
        this.HighScoresCanvas = highScoresCanvas
    return this
}


func (this *HighScoresCanvasInputProcessor) Open() {
}


func (this *HighScoresCanvasInputProcessor) Close() {
}


// synchronized
func (this *HighScoresCanvasInputProcessor) OnPressGameKeyEvent(gameKeyEvent *GameKeyEvent) {
        this.PlayerGameInput.OnPressGameKeyEvent(gameKeyEvent)
        this.Update()
}


// synchronized
func (this *HighScoresCanvasInputProcessor) OnDownGameKeyEvent(gameKeyEvent *GameKeyEvent) {
        this.PlayerGameInput.OnDownGameKeyEvent(gameKeyEvent)
        this.Update()
}


// synchronized
func (this *HighScoresCanvasInputProcessor) OnUpGameKeyEvent(gameKeyEvent *GameKeyEvent) {
        this.PlayerGameInput.OnUpGameKeyEvent(gameKeyEvent)
        this.Update()
}


func (this *HighScoresCanvasInputProcessor) Paint(graphics IGraphics) {
}


func (this *HighScoresCanvasInputProcessor) GetHighScoresCanvas() *HighScoresCanvas {
        return this.HighScoresCanvas
}

