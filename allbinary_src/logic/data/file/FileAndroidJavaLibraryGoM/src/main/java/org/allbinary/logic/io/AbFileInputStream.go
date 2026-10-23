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

type AbFileInputStream struct {
    InputStream
    FileInputStream *FileInputStream
}

// The Java class AbFileInputStream has subclasses so its variables are this interface
type IAbFileInputStream interface {
    IInputStream
    AbFileInputStream_() *AbFileInputStream
    Mark(readlimit int)
    MarkSupported() bool
    Reset()
    Available() int
    Close()
    GetChannel() *FileChannel
    Read(args ...any) int
    Skip(n int64) int64
}

func (this *AbFileInputStream) AbFileInputStream_() *AbFileInputStream {
    return this
}


func (this *AbFileInputStream) initFields() {
    
}


func NewAbFileInputStream(fileInputStream *FileInputStream) *AbFileInputStream {
    this := &AbFileInputStream{}
    this.InputStream = *NewInputStream()
    this.initFields()
        this.FileInputStream = fileInputStream
    return this
}


func (this *AbFileInputStream) Mark(readlimit int) {
        this.FileInputStream.Mark(readlimit)
}


func (this *AbFileInputStream) MarkSupported() bool {
        return this.FileInputStream.MarkSupported()
}


// throws IOException
// synchronized
func (this *AbFileInputStream) Reset() {
        this.FileInputStream.Reset()
}


// throws IOException
func (this *AbFileInputStream) Available() int {
        return this.FileInputStream.Available()
}


// throws IOException
func (this *AbFileInputStream) Close() {
        this.FileInputStream.Close()
}


func (this *AbFileInputStream) GetChannel() *FileChannel {
        return this.FileInputStream.GetChannel()
}


// throws IOException
func (this *AbFileInputStream) Read_1() int {
        return this.FileInputStream.Read()
}

// Java overloads of AbFileInputStream.read
func (this *AbFileInputStream) Read(args ...any) int {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 0 {
            return this.Read_1()
        }
        if len(args) == 1 && Java_IsArg[[]int8](args[0], pass) {
            return this.Read_2(Java_Arg[[]int8](args[0]))
        }
        if len(args) == 3 && Java_IsArg[[]int8](args[0], pass) && Java_IsArg[int](args[1], pass) && Java_IsArg[int](args[2], pass) {
            return this.Read_3(Java_Arg[[]int8](args[0]), Java_Arg[int](args[1]), Java_Arg[int](args[2]))
        }
    }
    panic("No Java overload of AbFileInputStream.read for the arguments")
}


// throws IOException
func (this *AbFileInputStream) Read_2(b []int8) int {
        return this.FileInputStream.Read(b)
}


// throws IOException
func (this *AbFileInputStream) Read_3(buffer []int8, offset int, count int) int {
        return this.FileInputStream.Read(buffer, offset, count)
}


// throws IOException
func (this *AbFileInputStream) Skip(n int64) int64 {
        return this.FileInputStream.Skip(n)
}

