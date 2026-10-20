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

type RemoteHighScores struct {
    HighScores
    LogUtil *LogUtil
    EnumerationUtil *EnumerationUtil
    AbeClientInformation AbeClientInformationInterface
    SoftwareInformation *SoftwareInformation
    Ascending bool
    ASCENDING string
}

var RemoteHighScores_HashTable *ABHashtable[any, any] = StdUtil_GetInstance().CreateHashtable()

func (this *RemoteHighScores) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.EnumerationUtil = EnumerationUtil_GetInstance()
    
    
    
    this.ASCENDING = "ASCENDING"
}


func NewRemoteHighScores(abeClientInformation AbeClientInformationInterface, softwareInformation *SoftwareInformation, gameInfo *GameInfo, heading string, columnTwoHeading string, ascending bool, preload bool) *RemoteHighScores {
    this := &RemoteHighScores{}
        this.HighScores = *NewHighScores(fmt.Sprint(gameInfo), heading, columnTwoHeading)
        this.initFields()
        this.AbeClientInformation = abeClientInformation
        this.SoftwareInformation = softwareInformation
        this.SetAscending(ascending)
        if preload {
            RemoteHighScoresProcessorFactory_GetInstance().Process(this, this.AbeClientInformation, gameInfo)
        }
    return this
}


// synchronized
func RemoteHighScores_GetInstance(abeClientInformation AbeClientInformationInterface, softwareInformation *SoftwareInformation, gameInfo *GameInfo, heading string, columnTwoHeading string, isAscending bool) IHighScores {
        return RemoteHighScores_GetInstancePreload(abeClientInformation, softwareInformation, gameInfo, heading, columnTwoHeading, isAscending, true)
}


// synchronized
func RemoteHighScores_GetInstancePreload(abeClientInformation AbeClientInformationInterface, softwareInformation *SoftwareInformation, gameInfo *GameInfo, heading string, columnTwoHeading string, isAscending bool, preload bool) IHighScores {
        var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
        if tryControl1, tryResult1 := func() (tryControl1 int, tryResult1 IHighScores) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            logUtil.Put(commonStrings.EXCEPTION, RemoteErrorHighScoresSingletonFactory_GetInstance(), commonStrings.GET_INSTANCE, e)
            tryControl1, tryResult1 = 1, RemoteErrorHighScoresSingletonFactory_GetInstance(); return
                }
            }()
            var highScores IHighScores = any(RemoteHighScores_HashTable.Get(gameInfo)).(IHighScores); _ = highScores
            if highScores == nil {
                highScores = NewRemoteHighScores(abeClientInformation, softwareInformation, gameInfo, heading, columnTwoHeading, isAscending, preload)
                RemoteHighScores_HashTable.Put(gameInfo, highScores)
            }
            tryControl1, tryResult1 = 1, highScores; return
            return
        }(); tryControl1 == 1 {
            return tryResult1
        }
        panic("unreachable")
}


func (this *RemoteHighScores) AddHighScore(newHighScore *HighScore) {
        RemoteHighScoresSubmissionProcessorFactory_GetInstance().Process(this, this.AbeClientInformation, newHighScore)
}


func (this *RemoteHighScores) Update(hashtable *ABHashtable[any, any]) {
        this.GetList().Clear()
        var vector IBasicArrayList = any(hashtable.Get(any(RemoteHighScoresData_GetInstance().HIGH_SCORES).(any))).(IBasicArrayList); _ = vector
        if vector != nil {
            var size int = vector.Size(); _ = size
            for index := int(0); index < size; index++ {
                var highScoreVector IBasicArrayList = any(vector.Get(index)).(IBasicArrayList); _ = highScoreVector
                var displayName string = any(highScoreVector.Get(0)).(string); _ = displayName
                var score string = any(highScoreVector.Get(1)).(string); _ = score
                var longScore int64 = func() int64 { value, _ := strconv.ParseInt(score, 10, 64); return value }(); _ = longScore
                var highScore *HighScore = NewHighScore(-1, displayName, nil, longScore); _ = highScore
                this.GetList().Add(highScore)
            }
        } else {
            var NEXT_ELEMENT string = fmt.Sprint(NewStringMaker().Append(CommonLabels_GetInstance().NEXT).Append("Element: ")); _ = NEXT_ELEMENT
            var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
            var enumeration Enumeration = hashtable.Elements(); _ = enumeration
            var nextElement any; _ = nextElement
            for this.EnumerationUtil.HasMoreElements(enumeration) {
                nextElement = this.EnumerationUtil.NextElement(enumeration)
                this.LogUtil.PutF(NEXT_ELEMENT + fmt.Sprint(nextElement), this, commonStrings.PROCESS)
            }
        }
}


func (this *RemoteHighScores) SetAscending(ascending bool) {
        this.Ascending = ascending
}


func (this *RemoteHighScores) GetAscending() bool {
        return this.Ascending
}


func (this *RemoteHighScores) GetSoftwareInformation() *SoftwareInformation {
        return this.SoftwareInformation
}

