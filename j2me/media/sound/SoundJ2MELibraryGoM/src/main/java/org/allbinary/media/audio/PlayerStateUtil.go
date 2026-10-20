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

type PlayerStateUtil struct {
    CommonStrings *CommonStrings
    UNREALIZED string
    REALIZED string
    PREFETCHED string
    STARTED string
    CLOSED string
}

var PlayerStateUtil_Instance *PlayerStateUtil = NewPlayerStateUtil()

func (this *PlayerStateUtil) initFields() {
    this.CommonStrings = CommonStrings_GetInstance()
    this.UNREALIZED = "UNREALIZED"
    this.REALIZED = "REALIZED"
    this.PREFETCHED = "PREFETCHED"
    this.STARTED = "STARTED"
    this.CLOSED = "CLOSED"
}

func NewPlayerStateUtil() *PlayerStateUtil {
    this := &PlayerStateUtil{}
    this.initFields()
    return this
}


func PlayerStateUtil_GetInstance() *PlayerStateUtil {
        return PlayerStateUtil_Instance
}


func (this *PlayerStateUtil) Convert(state int) string {
        if state == Player_UNREALIZED {
            return this.UNREALIZED
        } else if state == Player_REALIZED {
            return this.REALIZED
        } else if state == Player_PREFETCHED {
            return this.PREFETCHED
        } else if state == Player_STARTED {
            return this.STARTED
        } else if state == Player_CLOSED {
            return this.CLOSED
        }
        return this.CommonStrings.UNKNOWN
}

