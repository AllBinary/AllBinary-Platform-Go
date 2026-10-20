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

type ChoiceGroupImageFactory struct {
    MultiOff []int8
    MultiOn []int8
    RadioOff []int8
    RadioOn []int8
    ImgMultiOffO IImage
    ImgMultiOnO IImage
    ImgRadioOffO IImage
    ImgRadioOnO IImage
    ImgMultiOff IImage
    ImgMultiOn IImage
    ImgRadioOff IImage
    ImgRadioOn IImage
    ImageArray []IImage
}

var ChoiceGroupImageFactory_Instance *ChoiceGroupImageFactory = NewChoiceGroupImageFactory()

func (this *ChoiceGroupImageFactory) initFields() {
    this.MultiOff = []int8{int8(-119), 80, 78, 71, 13, 10, 26, 10, 0, 0, 0, 13, 73, 72, 68, 82, 0, 0, 0, 10, 0, 0, 0, 11, 2, 3, 0, 0, 0, 59, 0, int8(-12), int8(-117), 0, 0, 0, 6, 80, 76, 84, 69, int8(-1), int8(-1), int8(-1), int8(-69), int8(-69), int8(-69), int8(-57), 75, int8(-33), int8(-8), 0, 0, 0, 30, 73, 68, 65, 84, 120, int8(-38), 99, 96, 96, 96, 96, 12, 101, int8(-8), int8(-51), int8(-32), int8(-64), 32, int8(-64), int8(-60), int8(-64), int8(-64), int8(-128), 11, 51, int8(-122), 50, int8(-4), 6, 0, 63, 116, 3, 1, 53, int8(-108), 39, int8(-26), 0, 0, 0, 0, 73, 69, 78, 68, int8(-82), 66, 96, int8(-126)}
    this.MultiOn = []int8{int8(-119), 80, 78, 71, 13, 10, 26, 10, 0, 0, 0, 13, 73, 72, 68, 82, 0, 0, 0, 10, 0, 0, 0, 11, 2, 3, 0, 0, 0, 59, 0, int8(-12), int8(-117), 0, 0, 0, 12, 80, 76, 84, 69, int8(-1), int8(-1), int8(-1), int8(-69), int8(-69), int8(-69), 106, 106, 106, 2, 2, 2, 106, int8(-103), 14, int8(-47), 0, 0, 0, 53, 73, 68, 65, 84, 120, int8(-38), 99, 96, 96, 124, int8(-64), int8(-16), int8(-1), int8(-77), 3, int8(-45), 65, int8(-111), 15, 76, 12, 108, 12, 76, 12, int8(-4), 12, 76, 12, 18, 12, 76, int8(-68), 127, 24, int8(-104), 126, 45, 96, 96, int8(-7), int8(-11), int8(-109), int8(-127), int8(-23), int8(-65), 3, 3, int8(-29), 127, int8(-122), int8(-113), 0, 5, 37, 12, int8(-34), 1, int8(-99), int8(-83), 100, 0, 0, 0, 0, 73, 69, 78, 68, int8(-82), 66, 96, int8(-126)}
    this.RadioOff = []int8{int8(-119), 80, 78, 71, 13, 10, 26, 10, 0, 0, 0, 13, 73, 72, 68, 82, 0, 0, 0, 11, 0, 0, 0, 11, 2, 3, 0, 0, 0, int8(-44), int8(-62), int8(-97), int8(-75), 0, 0, 0, 9, 80, 76, 84, 69, int8(-1), int8(-1), int8(-1), int8(-69), int8(-69), int8(-69), 106, 106, 106, int8(-44), 13, int8(-1), int8(-24), 0, 0, 0, 42, 73, 68, 65, 84, 120, int8(-38), 99, 96, 90, int8(-59), int8(-64), 32, int8(-63), 48, int8(-127), 65, int8(-127), 65, int8(-127), 41, int8(-127), int8(-31), 5, 19, 3, 3, 3, 50, 102, 80, 96, 80, 96, int8(-6), int8(-63), 80, int8(-64), int8(-64), int8(-76), int8(-118), 1, 0, 113, 24, 5, 61, 73, int8(-68), int8(-100), 98, 0, 0, 0, 0, 73, 69, 78, 68, int8(-82), 66, 96, int8(-126)}
    this.RadioOn = []int8{int8(-119), 80, 78, 71, 13, 10, 26, 10, 0, 0, 0, 13, 73, 72, 68, 82, 0, 0, 0, 11, 0, 0, 0, 11, 2, 3, 0, 0, 0, int8(-44), int8(-62), int8(-97), int8(-75), 0, 0, 0, 12, 80, 76, 84, 69, int8(-1), int8(-1), int8(-1), int8(-69), int8(-69), int8(-69), 106, 106, 106, 2, 2, 2, 106, int8(-103), 14, int8(-47), 0, 0, 0, 50, 73, 68, 65, 84, 120, int8(-38), 5, int8(-63), 65, 13, 0, 32, 12, 4, int8(-63), int8(-19), int8(-11), int8(-117), 1, 18, 68, int8(-100), 10, 52, 19, 94, 72, 64, 17, 101, int8(-122), 44, int8(-44), int8(-29), 98, int8(-52), 89, 77, int8(-102), 40, 2, 85, int8(-95), int8(-73), int8(-63), int8(-104), int8(-63), 37, int8(-117), 15, int8(-40), 119, 10, 41, 78, 26, int8(-79), 59, 0, 0, 0, 0, 73, 69, 78, 68, int8(-82), 66, 96, int8(-126)}
    this.ImgMultiOffO = Image_CreateImage(this.MultiOff, 0, len(this.MultiOff))
    this.ImgMultiOnO = Image_CreateImage(this.MultiOn, 0, len(this.MultiOn))
    this.ImgRadioOffO = Image_CreateImage(this.RadioOff, 0, len(this.RadioOff))
    this.ImgRadioOnO = Image_CreateImage(this.RadioOn, 0, len(this.RadioOn))
    this.ImgMultiOff = this.ImgMultiOffO
    this.ImgMultiOn = this.ImgMultiOnO
    this.ImgRadioOff = this.ImgRadioOffO
    this.ImgRadioOn = this.ImgRadioOnO
    this.ImageArray = []IImage{this.ImgMultiOff, this.ImgMultiOn, this.ImgRadioOff, this.ImgRadioOn}
}

func NewChoiceGroupImageFactory() *ChoiceGroupImageFactory {
    this := &ChoiceGroupImageFactory{}
    this.initFields()
    return this
}


func ChoiceGroupImageFactory_GetInstance() *ChoiceGroupImageFactory {
        return ChoiceGroupImageFactory_Instance
}


func (this *ChoiceGroupImageFactory) GetImageArray() []IImage {
        return this.ImageArray
}


func (this *ChoiceGroupImageFactory) Init(imageArray []IImage) {
        this.ImgMultiOff = imageArray[0]
        this.ImageArray[0] = this.ImgMultiOff
        this.ImgMultiOn = imageArray[1]
        this.ImageArray[1] = this.ImgMultiOn
        this.ImgRadioOff = imageArray[2]
        this.ImageArray[2] = this.ImgRadioOff
        this.ImgRadioOn = imageArray[3]
        this.ImageArray[3] = this.ImgRadioOn
}

