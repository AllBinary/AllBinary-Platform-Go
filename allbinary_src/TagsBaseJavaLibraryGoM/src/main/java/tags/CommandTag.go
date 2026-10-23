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

type CommandTag struct {
    CustomTagSupport
    TagStrings *TagStrings
    Command string
}

// The Java class CommandTag has subclasses so its variables are this interface
type ICommandTag interface {
    ICustomTagSupport
    CommandTag_() *CommandTag
    SetCommand(command string)
    GetCommand() string
}

func (this *CommandTag) CommandTag_() *CommandTag {
    return this
}


func (this *CommandTag) initFields() {
    this.TagStrings = TagStrings_GetInstance()
    
}


func NewCommandTag() *CommandTag {
    this := &CommandTag{}
    this.CustomTagSupport = *NewCustomTagSupport()
    this.initFields()
    return this
}


func (this *CommandTag) SetCommand(command string) {
        this.Command = command
}


func (this *CommandTag) GetCommand() string {
        return this.Command
}

