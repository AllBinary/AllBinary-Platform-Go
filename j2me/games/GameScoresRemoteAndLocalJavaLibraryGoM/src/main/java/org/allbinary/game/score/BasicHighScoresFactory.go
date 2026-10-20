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

type BasicHighScoresFactory struct {
    HighScoresBase
    LogUtil *LogUtil
    AbeClientInformation AbeClientInformationInterface
    SoftwareInformation *SoftwareInformation
    HighScoresArray []IHighScores
    TOP string
    SCORES string
    PERSONAL_HIGH_SCORES string
    WORLD_TOP_SCORES string
    FETCH string
}


func (this *BasicHighScoresFactory) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    this.HighScoresArray = make([]IHighScores, 2)
    this.TOP = "Top"
    this.SCORES = "Scores"
    this.PERSONAL_HIGH_SCORES = "Personal Top Scores"
    this.WORLD_TOP_SCORES = "World Top Scores"
    this.FETCH = "fetchHighScores"
}


func NewBasicHighScoresFactory(abeClientInformation AbeClientInformationInterface, softwareInformation *SoftwareInformation) *BasicHighScoresFactory {
    this := &BasicHighScoresFactory{}
    this.HighScoresBase = *NewHighScoresBase()
    this.initFields()
        this.AbeClientInformation = abeClientInformation
        this.SoftwareInformation = softwareInformation
    return this
}


func (this *BasicHighScoresFactory) FetchHighScores(gameInfo *GameInfo, highScoresResultsListener HighScoresResultsListener) {
        this.LogUtil.PutF("Getting Remote/Local HighScores", this, this.FETCH)
        this.FetchHighScoresPreload(gameInfo, highScoresResultsListener, true)
}


func (this *BasicHighScoresFactory) FetchHighScoresPreload(gameInfo *GameInfo, highScoresResultsListener HighScoresResultsListener, preload bool) {
        SecondaryThreadPool_GetInstance().RunTask(/* Java anonymous class body: implement it as a named type
            

        func (this *BasicHighScoresFactory) Run() {
                var logUtil *LogUtil = LogUtil_GetInstance(); _ = logUtil
                if tryControl1 := func() (tryControl1 int) {
                    defer func() {
                        if recovered := recover(); recovered != nil {
                            e := Java_ToError(recovered)
                            _ = e
                    logUtil.Put(this.CommonStrings.EXCEPTION, this, this.FETCH, e)
                        }
                    }()
                    logUtil.PutF(this.CommonStrings.START, this, this.FETCH)
                    this.HighScoresArray[0] = RecordStoreHighScores_GetInstance(this.AbeClientInformation, gameInfo, this.TOP, this.PERSONAL_HIGH_SCORES, this.SCORES, NewScoreComparator(true))
                    var gameType *GameType = gameInfo.GetGameType(); _ = gameType
                    var gameTypeFactory *GameTypeFactory = GameTypeFactory_GetInstance(); _ = gameTypeFactory
                    var gameInfo2 *GameInfo = gameInfo; _ = gameInfo2
                    if gameType == gameTypeFactory.SINGLE_PLAYER || gameType == gameTypeFactory.BOT {
                        gameInfo2 = BasicHighScoresGameInfoFactory_GetInstance().SINGLE_PLAYER_SERVER
                    } else if gameType == gameTypeFactory.MULTI_PLAYER {
                        gameInfo2 = BasicHighScoresGameInfoFactory_GetInstance().MULTI_PLAYER_SERVER
                    }
                    this.HighScoresArray[1] = RemoteHighScores_GetInstancePreload(this.AbeClientInformation, this.SoftwareInformation, gameInfo2, this.WORLD_TOP_SCORES, this.SCORES, BooleanFactory_GetInstance().FALSE, preload)
                    logUtil.PutF(this.CommonStrings.END, this, this.FETCH)
                    LastFetchHighScoresFactory_GetInstance().HighScoresArray = this.HighScoresArray
                    highScoresResultsListener.SetHighScoresArray(this.HighScoresArray)
                    return
                }(); tryControl1 == 1 {
                    return
                }
        }
*/ NewARunnable())
}


func (this *BasicHighScoresFactory) CreateHighScoresHelper() HighScoresHelperBaseInterface {
        return NewHighScoresHelper2()
}


func BasicHighScoresFactory_Loaded(index2 int) bool {
        if index2 >= 0 {
            return true
        }
        return false
}

