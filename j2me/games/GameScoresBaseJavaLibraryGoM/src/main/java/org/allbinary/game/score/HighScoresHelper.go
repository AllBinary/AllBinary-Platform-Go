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

type HighScoresHelper struct {
    HighScoresHelperBase
    CircularIndexUtil *CircularIndexUtil
    LastIndex int
}


func (this *HighScoresHelper) initFields() {
    this.CircularIndexUtil = CircularIndexUtil_CreateInstanceAt(0, 0)
    this.LastIndex = -1
}

func NewHighScoresHelper() *HighScoresHelper {
    this := &HighScoresHelper{}
    this.HighScoresHelperBase = *NewHighScoresHelperBase()
    this.initFields()
    return this
}


func (this *HighScoresHelper) SetHighScoresArray(highScoresArrayP []IHighScores) {
        this.HighScoresHelperBase.SetHighScoresArray(highScoresArrayP)
        this.CircularIndexUtil.SetSize(len(this.HighScoresArrayP))
}


func (this *HighScoresHelper) GetNextHighScores() IHighScores {
        var highScores IHighScores = this.HighScoresArrayP[this.CircularIndexUtil.GetIndex()]; _ = highScores
        var index int = 0; _ = index
        for (highScores.GetTotal() < 1 || this.LastIndex == this.CircularIndexUtil.GetIndex()) && index < len(this.HighScoresArrayP) {
            this.CircularIndexUtil.Next()
            highScores = this.HighScoresArrayP[this.CircularIndexUtil.GetIndex()]
            index++
        }
        this.LastIndex = this.CircularIndexUtil.GetIndex()
        return highScores
}


func (this *HighScoresHelper) IsAnyHighScores() bool {
        if len(this.HighScoresArrayP) < 1 {
            return false
        }
        var anyScores bool = false; _ = anyScores
        var highScores IHighScores; _ = highScores
        for index := int(len(this.HighScoresArrayP) - 1); index >= 0; index-- {
            highScores = this.HighScoresArrayP[index]
            if highScores.GetTotal() > 0 {
                anyScores = true
            }
        }
        return anyScores
}

