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

type CommonStrings struct {
    EMPTY string
    CREATE_IMAGE string
    PLEASE_WAIT string
    PLEASE_WAIT_FOR_SERVER string
    PERIODS string
    LOADING string
    CONSTRUCTOR string
    START string
    END string
    EXCEPTION string
    EXCEPTION_LABEL string
    CLEANUP string
    INIT string
    CLOSE string
    UP string
    DOWN string
    LEFT string
    RIGHT string
    DIRECTION_NAME []string
    CREATE string
    REMOVE string
    UPDATE string
    DELETE string
    Delete string
    LOAD string
    SAVE string
    ADD string
    DROP string
    INSERT string
    VISIT string
    START_METHOD_NAME string
    END_METHOD_NAME string
    GET string
    PROCESS string
    GET_INSTANCE string
    GET_LIST string
    RUN string
    START_RUNNABLE string
    RUNNING string
    END_RUNNABLE string
    DISABLE string
    ENABLE string
    SUCCESS string
    FAILURE string
    NOT_IMPLEMENTED string
    UNKNOWN string
    ON_EVENT string
    IS_VALID string
    TOSTRING string
    ADD_LISTENER string
    REMOVE_LISTENER string
    SIZE string
}

var CommonStrings_Instance *CommonStrings = NewCommonStrings()

func (this *CommonStrings) initFields() {
    this.EMPTY = "Empty"
    this.CREATE_IMAGE = "createImage"
    this.PLEASE_WAIT = "Please Wait"
    this.PLEASE_WAIT_FOR_SERVER = "Please Wait Getting Data From Server..."
    this.PERIODS = "..."
    this.LOADING = "Loading..."
    this.CONSTRUCTOR = "Constructor"
    this.START = "Start"
    this.END = "End"
    this.EXCEPTION = "Exception"
    this.EXCEPTION_LABEL = "Exception: "
    this.CLEANUP = "cleanup"
    this.INIT = "init"
    this.CLOSE = "close"
    this.UP = "up"
    this.DOWN = "down"
    this.LEFT = "left"
    this.RIGHT = "right"
    this.DIRECTION_NAME = []string{this.DOWN, this.LEFT, this.RIGHT, this.UP}
    this.CREATE = "Create"
    this.REMOVE = "Remove"
    this.UPDATE = "Update"
    this.DELETE = "Delete"
    this.Delete = "delete"
    this.LOAD = "load"
    this.SAVE = "save"
    this.ADD = "Add"
    this.DROP = "Drop"
    this.INSERT = "Insert"
    this.VISIT = "visit"
    this.START_METHOD_NAME = this.START
    this.END_METHOD_NAME = "end"
    this.GET = "get"
    this.PROCESS = "process"
    this.GET_INSTANCE = "getInstance"
    this.GET_LIST = "getList"
    this.RUN = "run"
    this.START_RUNNABLE = "Start Runnable"
    this.RUNNING = "Running"
    this.END_RUNNABLE = "End Runnable"
    this.DISABLE = "disable"
    this.ENABLE = "enable"
    this.SUCCESS = "Command Success"
    this.FAILURE = "Command Failed"
    this.NOT_IMPLEMENTED = "Not Implemented"
    this.UNKNOWN = "Unknown"
    this.ON_EVENT = "onEvent"
    this.IS_VALID = "isValid"
    this.TOSTRING = "toString"
    this.ADD_LISTENER = "addListener"
    this.REMOVE_LISTENER = "removeListener"
    this.SIZE = "size"
}

func NewCommonStrings() *CommonStrings {
    this := &CommonStrings{}
    this.initFields()
    return this
}


func CommonStrings_GetInstance() *CommonStrings {
        return CommonStrings_Instance
}

