/*
 * 
 * AllBinary Open License Version 1 
 * Copyright (c) 2025 AllBinary 
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

type PersistanceStrings struct {
    SAVING string
    NUMBER_OF_RECORDS string
    LOADING_ID string
    LOAD_ALL string
    DELETING_WITH_ID string
    CLOSING_RECORDSTORE string
    NOT_SAVING string
    ERROR_LOADING string
    ERROR_LOADING_ID string
    ID string
    GAME_ACTION_INPUT string
}

var PersistanceStrings_Instance *PersistanceStrings = NewPersistanceStrings()

func (this *PersistanceStrings) initFields() {
    this.SAVING = "Saving: "
    this.NUMBER_OF_RECORDS = "Number of records: "
    this.LOADING_ID = "Loading data with id: "
    this.LOAD_ALL = "loadAll"
    this.DELETING_WITH_ID = "Deleting data with id: "
    this.CLOSING_RECORDSTORE = "Closing RecordStore"
    this.NOT_SAVING = "Not Saving: "
    this.ERROR_LOADING = "Error Loading gameActionInput: "
    this.ERROR_LOADING_ID = "Error Loading id: "
    this.ID = " id: "
    this.GAME_ACTION_INPUT = " GameActionInput: "
}

func NewPersistanceStrings() *PersistanceStrings {
    this := &PersistanceStrings{}
    this.initFields()
    return this
}


func PersistanceStrings_GetInstance() *PersistanceStrings {
        return PersistanceStrings_Instance
}

