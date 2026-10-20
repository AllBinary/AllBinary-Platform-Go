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

type MyCanvasEventSource struct {
    Command string
    TreeNode *MutableTreeNode
    Angle float64
}


func (this *MyCanvasEventSource) initFields() {
    
    
    this.Angle = 0.0
}




// Java overloads of new MyCanvasEventSource
func NewMyCanvasEventSource(args ...any) *MyCanvasEventSource {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[*MutableTreeNode](args[1], pass) {
            return NewMyCanvasEventSource_1(Java_Arg[string](args[0]), Java_Arg[*MutableTreeNode](args[1]))
        }
        if len(args) == 3 && Java_IsArg[string](args[0], pass) && Java_IsArg[float64](args[1], pass) && Java_IsArg[*MutableTreeNode](args[2], pass) {
            return NewMyCanvasEventSource_2(Java_Arg[string](args[0]), Java_Arg[float64](args[1]), Java_Arg[*MutableTreeNode](args[2]))
        }
    }
    panic("No Java overload of new MyCanvasEventSource for the arguments")
}
func NewMyCanvasEventSource_1(command string, treeNode *MutableTreeNode) *MyCanvasEventSource {
    this := &MyCanvasEventSource{}
    this.initFields()
        this.Command = command
        this.TreeNode = treeNode
    return this
}


func NewMyCanvasEventSource_2(command string, angle float64, treeNode *MutableTreeNode) *MyCanvasEventSource {
    this := &MyCanvasEventSource{}
    this.initFields()
        this.Command = command
        this.TreeNode = treeNode
        this.SetAngle(angle)
    return this
}


func (this *MyCanvasEventSource) GetCommand() string {
        return this.Command
}


func (this *MyCanvasEventSource) GetTreeNode() *MutableTreeNode {
        return this.TreeNode
}


func (this *MyCanvasEventSource) GetAngle() float64 {
        return this.Angle
}


func (this *MyCanvasEventSource) SetAngle(angle float64) {
        this.Angle = angle
}

