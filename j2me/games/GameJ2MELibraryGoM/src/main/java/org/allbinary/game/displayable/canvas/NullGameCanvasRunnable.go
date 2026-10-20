/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

type NullGameCanvasRunnable struct {
    GameStateFactory *GameStateFactory
}

// Implements GameCanvasRunnableInterface

var NullGameCanvasRunnable_NULL_GAME_CANVAS_RUNNABLE *NullGameCanvasRunnable = NewNullGameCanvasRunnable()

func (this *NullGameCanvasRunnable) initFields() {
    this.GameStateFactory = GameStateFactory_GetInstance()
}

func NewNullGameCanvasRunnable() *NullGameCanvasRunnable {
    this := &NullGameCanvasRunnable{}
    this.initFields()
    return this
}


func (this *NullGameCanvasRunnable) GetTitle() string {
        return StringUtil_GetInstance().EMPTY_STRING
}


func (this *NullGameCanvasRunnable) IsInitialized() bool {
        return false
}


func (this *NullGameCanvasRunnable) IsPaused() bool {
        return false
}


func (this *NullGameCanvasRunnable) UnPause() {
}


func (this *NullGameCanvasRunnable) Pause() {
}


func (this *NullGameCanvasRunnable) IsHighScoreSubmitted() bool {
        return false
}


// throws Exception
func (this *NullGameCanvasRunnable) SetHighScoreSubmitted(highScoreSubmitted bool) {
}


// throws Exception
func (this *NullGameCanvasRunnable) ProcessGameOver() {
}


func (this *NullGameCanvasRunnable) IsGameOver() bool {
        return false
}


func (this *NullGameCanvasRunnable) InitCommands(cmdListener CommandListener) {
}


// throws Exception
func (this *NullGameCanvasRunnable) GetCurrentStateHashtable() *ABHashtable[any, any] {
        return StdUtil_GetInstance().NULL_TABLE
}


// throws Exception
func (this *NullGameCanvasRunnable) GetLoadStateHashtable() *ABHashtable[any, any] {
        return StdUtil_GetInstance().NULL_TABLE
}


// throws Exception
func (this *NullGameCanvasRunnable) SetLoadStateHashtable(hashtable *ABHashtable[any, any]) {
}


func (this *NullGameCanvasRunnable) GetType() int {
        return -1
}


// throws Exception
func (this *NullGameCanvasRunnable) SetGameState(gameState *GameState) {
}


func (this *NullGameCanvasRunnable) GetGameState() *GameState {
        return this.GameStateFactory.NO_GAME_STATE
}


func (this *NullGameCanvasRunnable) ItemStateChanged(item IItem) {
}


func (this *NullGameCanvasRunnable) Paint(graphics IGraphics) {
}


func (this *NullGameCanvasRunnable) PaintThreed(graphics IGraphics) {
}


func (this *NullGameCanvasRunnable) IsRunning() bool {
        return false
}


func (this *NullGameCanvasRunnable) SetRunning(isRunning bool) {
}


// throws Exception
func (this *NullGameCanvasRunnable) SetThread(thread IThread) {
}


func (this *NullGameCanvasRunnable) Run() {
}

