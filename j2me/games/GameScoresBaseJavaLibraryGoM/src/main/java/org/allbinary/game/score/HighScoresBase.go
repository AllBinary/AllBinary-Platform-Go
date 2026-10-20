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

type HighScoresBase struct {
    CommonStrings *CommonStrings
}

// Implements HighScoresFactoryInterface

// The Java class HighScoresBase has subclasses so its variables are this interface
type IHighScoresBase interface {
    HighScoresBase_() *HighScoresBase
    FetchHighScores(gameInfo *GameInfo, highScoresResultsListener HighScoresResultsListener)
    FetchHighScoresPreload(gameInfo *GameInfo, highScoresResultsListener HighScoresResultsListener, preload bool)
    CreateHighScoresHelper() HighScoresHelperBaseInterface
}

func (this *HighScoresBase) HighScoresBase_() *HighScoresBase {
    return this
}


func (this *HighScoresBase) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
}


func NewHighScoresBase() *HighScoresBase {
    this := &HighScoresBase{}
    this.initFields()
    return this
}


func (this *HighScoresBase) FetchHighScores(gameInfo *GameInfo, highScoresResultsListener HighScoresResultsListener) {
        panic(errors.New(fmt.Sprint()))
}


func (this *HighScoresBase) FetchHighScoresPreload(gameInfo *GameInfo, highScoresResultsListener HighScoresResultsListener, preload bool) {
        panic(errors.New(fmt.Sprint()))
}


func (this *HighScoresBase) CreateHighScoresHelper() HighScoresHelperBaseInterface {
        panic(errors.New(fmt.Sprint()))
}

