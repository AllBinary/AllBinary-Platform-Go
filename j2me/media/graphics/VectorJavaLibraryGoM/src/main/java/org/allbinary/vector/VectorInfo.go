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

type VectorInfo struct {
    Width int
    Height int
    Points [][]int
    TotalFrames int
}


func (this *VectorInfo) initFields() {
    
    
    
    
}


func NewVectorInfo(width int, height int, points [][]int, totalFrames int) *VectorInfo {
    this := &VectorInfo{}
    this.initFields()
        this.Width = width
        this.Height = height
        this.Points = points
        this.TotalFrames = totalFrames
    return this
}


func (this *VectorInfo) GetWidth() int {
        return this.Width
}


func (this *VectorInfo) SetWidth(width int) {
        this.Width = width
}


func (this *VectorInfo) GetHeight() int {
        return this.Height
}


func (this *VectorInfo) SetHeight(height int) {
        this.Height = height
}


func (this *VectorInfo) GetPoints() [][]int {
        return this.Points
}


func (this *VectorInfo) SetPoints(points [][]int) {
        this.Points = points
}


func (this *VectorInfo) GetTotalFrames() int {
        return this.TotalFrames
}


func (this *VectorInfo) SetTotalFrames(totalFrames int) {
        this.TotalFrames = totalFrames
}

