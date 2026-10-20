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

type HighScore struct {
    Id int
    Name string
    GameInfo *GameInfo
    Score int64
    ScoreString string
}


func (this *HighScore) initFields() {
    
    
    
    
    
}


func NewHighScore(id int, name string, gameInfo *GameInfo, score int64) *HighScore {
    this := &HighScore{}
    this.initFields()
        this.Id = id
        this.Name = name
        this.GameInfo = gameInfo
        this.Score = score
        this.ScoreString = fmt.Sprint(this.Score)
    return this
}


func (this *HighScore) GetId() int {
        return this.Id
}


func (this *HighScore) GetName() string {
        return this.Name
}


func (this *HighScore) GetScore() int64 {
        return this.Score
}


// throws Exception
func (this *HighScore) GetAsBytes() []int8 {
        var byteArrayOutputStream *strings.Builder = &strings.Builder{}; _ = byteArrayOutputStream
        var outputStream IDataOutputStream = NewDataOutputStream(byteArrayOutputStream); _ = outputStream
        outputStream.WriteUTF(this.GetName())
        outputStream.WriteLong(this.GetScore())
        return byteArrayOutputStream.ToByteArray()
}


func (this *HighScore) GetGameInfo() *GameInfo {
        return this.GameInfo
}


func (this *HighScore) GetScoreString() string {
        return this.ScoreString
}


func (this *HighScore) SetName(name string) {
        this.Name = name
}


func (this *HighScore) ToString() string {
        var commonSeps *CommonSeps = CommonSeps_GetInstance(); _ = commonSeps
        return fmt.Sprint(NewStringMaker().Append(this.Name).Append(commonSeps.COLON).Appendlong(this.Score).Append(commonSeps.FORWARD_SLASH).Append(this.ScoreString))
}

func (this *HighScore) String() string {
    return this.ToString()
}

