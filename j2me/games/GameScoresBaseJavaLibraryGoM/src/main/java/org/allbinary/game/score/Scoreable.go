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

type Scoreable struct {
    Points int
}

// Implements ScoreableInterface

// The Java class Scoreable has subclasses so its variables are this interface
type IScoreable interface {
    Scoreable_() *Scoreable
    AddPoints(points int)
    RemovePoints(points int)
    SetPoints(points int)
    GetPoints() int
}

func (this *Scoreable) Scoreable_() *Scoreable {
    return this
}

var Scoreable_ID int = SmallIntegerSingletonFactory_GetInstance().GetAt(3)

func (this *Scoreable) initFields() {
    
}


func NewScoreable(points int) *Scoreable {
    this := &Scoreable{}
    this.initFields()
        this.Points = points
    return this
}


func (this *Scoreable) AddPoints(points int) {
        this.Points += points
}


func (this *Scoreable) RemovePoints(points int) {
        this.Points -= points
}


func (this *Scoreable) SetPoints(points int) {
        this.Points = points
}


func (this *Scoreable) GetPoints() int {
        return this.Points
}

