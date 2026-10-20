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

type HighScoreCommandsBaseFactory struct {
    LogUtil *LogUtil
    HIGH_SCORE_COMMANDS []ICommand
}

// The Java class HighScoreCommandsBaseFactory has subclasses so its variables are this interface
type IHighScoreCommandsBaseFactory interface {
    HighScoreCommandsBaseFactory_() *HighScoreCommandsBaseFactory
    IsHighScoreCommand(command ICommand) bool
    GetIndex(command ICommand) int
}

func (this *HighScoreCommandsBaseFactory) HighScoreCommandsBaseFactory_() *HighScoreCommandsBaseFactory {
    return this
}


func (this *HighScoreCommandsBaseFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
}


func NewHighScoreCommandsBaseFactory(HIGH_SCORE_COMMANDS []ICommand) *HighScoreCommandsBaseFactory {
    this := &HighScoreCommandsBaseFactory{}
    this.initFields()
        this.HIGH_SCORE_COMMANDS = HIGH_SCORE_COMMANDS
    return this
}


func (this *HighScoreCommandsBaseFactory) IsHighScoreCommand(command ICommand) bool {
        for index := int(len(this.HIGH_SCORE_COMMANDS) - 1); index >= 0; index-- {
            if command == this.HIGH_SCORE_COMMANDS[index] {
                return true
            }
        }
        return false
}


// throws Exception
func (this *HighScoreCommandsBaseFactory) GetIndex(command ICommand) int {
        for index := int(len(this.HIGH_SCORE_COMMANDS) - 1); index >= 0; index-- {
            if command == this.HIGH_SCORE_COMMANDS[index] {
                return index
            }
        }
        panic(errors.New(fmt.Sprint("No Such Command")))
}

