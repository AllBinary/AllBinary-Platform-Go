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

type RemoteHighScoresData struct {
    HIGH_SCORES string
    CUSTOMER_USER_NAME string
    DISPLAY_NAME string
    SCORE string
    GAME_INFO string
    GAME_CONFIGURATION string
}

var RemoteHighScoresData_SINGLETON *RemoteHighScoresData = NewRemoteHighScoresData()

func (this *RemoteHighScoresData) initFields() {
    this.HIGH_SCORES = "HIGH_SCORES"
    this.CUSTOMER_USER_NAME = "CUSTOMER_USER_NAME"
    this.DISPLAY_NAME = "DISPLAY_NAME"
    this.SCORE = "SCORE"
    this.GAME_INFO = "GAME_INFO"
    this.GAME_CONFIGURATION = "GAME_CONFIGURATION"
}

func NewRemoteHighScoresData() *RemoteHighScoresData {
    this := &RemoteHighScoresData{}
    this.initFields()
    return this
}


func RemoteHighScoresData_GetInstance() *RemoteHighScoresData {
        return RemoteHighScoresData_SINGLETON
}

