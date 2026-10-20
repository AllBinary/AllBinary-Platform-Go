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

type NoHighScoresFactory struct {
    HighScoresBase
    LogUtil *LogUtil
    NO_HIGH_SCORES []IHighScores
}

var NoHighScoresFactory_Instance *NoHighScoresFactory = NewNoHighScoresFactory()

func (this *NoHighScoresFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.NO_HIGH_SCORES = make([]IHighScores, 0)
}


func NoHighScoresFactory_GetInstance() *NoHighScoresFactory {
        return NoHighScoresFactory_Instance
}


func NewNoHighScoresFactory() *NoHighScoresFactory {
    this := &NoHighScoresFactory{}
    this.HighScoresBase = *NewHighScoresBase()
    this.initFields()
    return this
}


func (this *NoHighScoresFactory) FetchHighScores(gameInfo *GameInfo, highScoresResultsListener HighScoresResultsListener) {
        this.LogUtil.PutF("Getting No HighScores", this, "fetchHighScores")
        highScoresResultsListener.SetHighScoresArray(this.NO_HIGH_SCORES)
}


func (this *NoHighScoresFactory) CreateHighScoresHelper() HighScoresHelperBaseInterface {
        return NewHighScoresHelperBase()
}

