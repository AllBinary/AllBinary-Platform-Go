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

type HighScoresHelperBase struct {
    LogUtil *LogUtil
    HighScoresArrayP []IHighScores
}

// Implements HighScoresHelperBaseInterface

// The Java class HighScoresHelperBase has subclasses so its variables are this interface
type IHighScoresHelperBase interface {
    HighScoresHelperBase_() *HighScoresHelperBase
    SetHighScoresArray(highScoresArrayP []IHighScores)
    GetNextHighScores() IHighScores
    IsAnyHighScores() bool
    SetSelectedHighScores(selectedHighScores IHighScores)
    GetSelectedHighScores() IHighScores
    GetHighScoresArray() []IHighScores
    SelectHighScores()
}

func (this *HighScoresHelperBase) HighScoresHelperBase_() *HighScoresHelperBase {
    return this
}


func (this *HighScoresHelperBase) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.HighScoresArrayP = LastFetchHighScoresFactory_GetInstance().HighScoresArray
}

func NewHighScoresHelperBase() *HighScoresHelperBase {
    this := &HighScoresHelperBase{}
    this.initFields()
    return this
}


func (this *HighScoresHelperBase) SetHighScoresArray(highScoresArrayP []IHighScores) {
        if highScoresArrayP != nil {
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(commonStrings.START).Appendint(len(highScoresArrayP))), this, "setHighScoresArray")
        } else {
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            this.LogUtil.PutF(commonStrings.START, this, "setHighScoresArray")
        }
        this.HighScoresArrayP = highScoresArrayP
}


func (this *HighScoresHelperBase) GetNextHighScores() IHighScores {
        return NullHighScoresSingletonFactory_GetInstance()
}


func (this *HighScoresHelperBase) IsAnyHighScores() bool {
        panic(errors.New(fmt.Sprint()))
}


func (this *HighScoresHelperBase) SetSelectedHighScores(selectedHighScores IHighScores) {
}


func (this *HighScoresHelperBase) GetSelectedHighScores() IHighScores {
        return NullHighScoresSingletonFactory_GetInstance()
}


func (this *HighScoresHelperBase) GetHighScoresArray() []IHighScores {
        return this.HighScoresArrayP
}


func (this *HighScoresHelperBase) SelectHighScores() {
}

