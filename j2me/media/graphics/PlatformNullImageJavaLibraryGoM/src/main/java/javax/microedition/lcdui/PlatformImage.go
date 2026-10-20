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

//PlatformNullImageJavaLibrary
type PlatformImage struct {
}

// The Java class PlatformImage has subclasses so its variables are this interface
type IPlatformImage interface {
    PlatformImage_() *PlatformImage
    GetImage(graphicsSurface2 any) any
    GetSurfaceWidth(graphicsSurface2 any, width int) int
    GetSurfaceHeight(graphicsSurface2 any, height int) int
    GetWidth(image IImage, width int) int
    GetHeight(image IImage, height int) int
    GetGraphics(graphicsSurface any, width int, height int, image IImage) IGraphics
    DrawImage(image IImage, x int, y int, anchor int, g2 IGraphics, g any)
    DrawRegion(image IImage, x_src int, y_src int, width int, height int, transform int, x_dst int, y_dst int, anchor int, g IGraphics)
    GetRGB(argb []int, offset int, scanlength int, x int, y int, width int, height int, image IImage)
    SetRGB(argb []int, offset int, scanlength int, x int, y int, width int, height int, image IImage)
    GetData(image IImage) []int
}

func (this *PlatformImage) PlatformImage_() *PlatformImage {
    return this
}


func NewPlatformImage() *PlatformImage {
    this := &PlatformImage{}
    return this
}


func (this *PlatformImage) GetImage(graphicsSurface2 any) any {
        return NullImage_NULL_IMAGE
}


func (this *PlatformImage) GetSurfaceWidth(graphicsSurface2 any, width int) int {
        return width
}


func (this *PlatformImage) GetSurfaceHeight(graphicsSurface2 any, height int) int {
        return height
}


func (this *PlatformImage) GetWidth(image IImage, width int) int {
        return width
}


func (this *PlatformImage) GetHeight(image IImage, height int) int {
        return height
}


func (this *PlatformImage) GetGraphics(graphicsSurface any, width int, height int, image IImage) IGraphics {
        panic(errors.New(fmt.Sprint()))
}


func (this *PlatformImage) DrawImage(image IImage, x int, y int, anchor int, g2 IGraphics, g any) {
}


func (this *PlatformImage) DrawRegion(image IImage, x_src int, y_src int, width int, height int, transform int, x_dst int, y_dst int, anchor int, g IGraphics) {
}


func (this *PlatformImage) GetRGB(argb []int, offset int, scanlength int, x int, y int, width int, height int, image IImage) {
}


func (this *PlatformImage) SetRGB(argb []int, offset int, scanlength int, x int, y int, width int, height int, image IImage) {
}


func (this *PlatformImage) GetData(image IImage) []int {
        return NullUtil_GetInstance().NULL_INT_ARRAY
}

