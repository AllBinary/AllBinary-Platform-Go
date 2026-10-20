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

type HighScores struct {
    OrderedHighScoresList IBasicArrayList
    Name string
    Heading string
    ColumnOneHeading string
    ColumnTwoHeading string
}

// Implements HighScoresInterface

// The Java class HighScores has subclasses so its variables are this interface
type IHighScores interface {
    HighScores_() *HighScores
    SetList(orderedHighScoresList IBasicArrayList)
    GetList() IBasicArrayList
    GetName() string
    GetTotal() int
    AddHighScore(newHighScore *HighScore)
    GetHeading() string
    GetColumnOneHeading() string
    GetColumnTwoHeading() string
    IsBestScore(newHighScore *HighScore) bool
    ToString() string
}

func (this *HighScores) HighScores_() *HighScores {
    return this
}


func (this *HighScores) initFields() {
    this.OrderedHighScoresList = NewBasicArrayListD()
    
    
    this.ColumnOneHeading = "Name"
    
}


func NewHighScores(name string, heading string, columnTwoHeading string) *HighScores {
    this := &HighScores{}
    this.initFields()
        this.Name = name
        this.Heading = heading
        this.ColumnTwoHeading = columnTwoHeading
    return this
}


func (this *HighScores) SetList(orderedHighScoresList IBasicArrayList) {
        this.OrderedHighScoresList = orderedHighScoresList
}


func (this *HighScores) GetList() IBasicArrayList {
        return this.OrderedHighScoresList
}


func (this *HighScores) GetName() string {
        return this.Name
}


func (this *HighScores) GetTotal() int {
        return this.GetList().Size()
}


func (this *HighScores) AddHighScore(newHighScore *HighScore) {
        this.OrderedHighScoresList.Add(newHighScore)
}


func (this *HighScores) GetHeading() string {
        return this.Heading
}


func (this *HighScores) GetColumnOneHeading() string {
        return this.ColumnOneHeading
}


func (this *HighScores) GetColumnTwoHeading() string {
        return this.ColumnTwoHeading
}


// throws Exception
func (this *HighScores) IsBestScore(newHighScore *HighScore) bool {
        return false
}


func (this *HighScores) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("HighScores: ")
        stringBuffer.Append(this.Heading)
        stringBuffer.Append(CommonSeps_GetInstance().SPACE)
        stringBuffer.Append(CommonLabels_GetInstance().TOTAL_LABEL)
        stringBuffer.Appendint(this.OrderedHighScoresList.Size())
        return fmt.Sprint(stringBuffer)
}

func (this *HighScores) String() string {
    return this.ToString()
}

