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

type MyGraphicItemEventSource struct {
    Command string
    TreeNode *MutableTreeNode
    Angle float64
}


func (this *MyGraphicItemEventSource) initFields() {
    
    
    this.Angle = 0.0
}




// Java overloads of new MyGraphicItemEventSource
func NewMyGraphicItemEventSource(args ...any) *MyGraphicItemEventSource {
    for pass := 0; pass < 2; pass++ {
        if len(args) == 2 && Java_IsArg[string](args[0], pass) && Java_IsArg[*MutableTreeNode](args[1], pass) {
            return NewMyGraphicItemEventSource_1(Java_Arg[string](args[0]), Java_Arg[*MutableTreeNode](args[1]))
        }
        if len(args) == 3 && Java_IsArg[string](args[0], pass) && Java_IsArg[float64](args[1], pass) && Java_IsArg[*MutableTreeNode](args[2], pass) {
            return NewMyGraphicItemEventSource_2(Java_Arg[string](args[0]), Java_Arg[float64](args[1]), Java_Arg[*MutableTreeNode](args[2]))
        }
    }
    panic("No Java overload of new MyGraphicItemEventSource for the arguments")
}
func NewMyGraphicItemEventSource_1(command string, treeNode *MutableTreeNode) *MyGraphicItemEventSource {
    this := &MyGraphicItemEventSource{}
    this.initFields()
        this.Command = command
        this.TreeNode = treeNode
    return this
}


func NewMyGraphicItemEventSource_2(command string, angle float64, treeNode *MutableTreeNode) *MyGraphicItemEventSource {
    this := &MyGraphicItemEventSource{}
    this.initFields()
        this.Command = command
        this.TreeNode = treeNode
        this.SetAngle(angle)
    return this
}


func (this *MyGraphicItemEventSource) GetCommand() string {
        return this.Command
}


func (this *MyGraphicItemEventSource) GetTreeNode() *MutableTreeNode {
        return this.TreeNode
}


func (this *MyGraphicItemEventSource) GetAngle() float64 {
        return this.Angle
}


func (this *MyGraphicItemEventSource) SetAngle(angle float64) {
        this.Angle = angle
}

