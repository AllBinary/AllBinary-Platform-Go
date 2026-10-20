/*
 * 
 * To change this license header, choose License Headers in Project Properties. 
 * To change this template file, choose Tools | Templates 
 * and open the template in the editor.  
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

type RemoteInfo struct {
}

var RemoteInfo_NEW_PLAYER_ID int = 0
var RemoteInfo_REMOTE_INFO *RemoteInfo = NewRemoteInfo(StringUtil_GetInstance().EMPTY_STRING, -1, -1, RemoteInfo_NEW_PLAYER_ID)


func NewRemoteInfo(username string, actorSessionId int, playerInputId int, id int) *RemoteInfo {
    this := &RemoteInfo{}
    return this
}

