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

type RemoteHighScoresSubmissionProcessor struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    NoCrypt *NoCrypt
}

// Implements RemoteHighScoresSubmissionProcessorInterface


func (this *RemoteHighScoresSubmissionProcessor) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.NoCrypt = NewNoCrypt()
}


func NewRemoteHighScoresSubmissionProcessor() *RemoteHighScoresSubmissionProcessor {
    this := &RemoteHighScoresSubmissionProcessor{}
    this.initFields()
    return this
}


// synchronized
func (this *RemoteHighScoresSubmissionProcessor) Process(remoteHighScores *RemoteHighScores, abeClientInformation AbeClientInformationInterface, highScore *HighScore) {
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(this.CommonStrings.EXCEPTION, this, this.CommonStrings.PROCESS, e)
                }
            }()
            this.LogUtil.PutF("Begin Remote HighScores Submission", this, this.CommonStrings.PROCESS)
            var gameInfoData *GameInfoData = GameInfoData_GetInstance(); _ = gameInfoData
            var hashtable *ABHashtable[any, any] = abeClientInformation.ToHashtable(); _ = hashtable
            HashtableUtil_GetInstance().PutAll(highScore.GetGameInfo().ToHashtable(), hashtable)
            hashtable.Put(RemoteHighScoresData_GetInstance().CUSTOMER_USER_NAME, "None")
            hashtable.Put(RemoteHighScoresData_GetInstance().DISPLAY_NAME, highScore.GetName())
            hashtable.Put(gameInfoData.SOFTWARE_INFORMATION, fmt.Sprint(remoteHighScores.GetSoftwareInformation()))
            hashtable.Put(remoteHighScores.ASCENDING, fmt.Sprint(remoteHighScores.GetAscending()))
            var displayInfoSingleton *DisplayInfoSingleton = DisplayInfoSingleton_GetInstance(); _ = displayInfoSingleton
            hashtable.Put(displayInfoSingleton.ORIENTATION, BooleanFactory_GetInstance().ToStringb(displayInfoSingleton.IsPortrait()))
            hashtable.Put(RemoteHighScoresData_GetInstance().GAME_CONFIGURATION, fmt.Sprint(GameConfigurationCentral_GetInstance()))
            hashtable.Put(RemoteHighScoresData_GetInstance().SCORE, fmt.Sprint(highScore.GetScore()))
            if XmlRpcAbeClient_IsOnline {
                var resultHashtable *ABHashtable[any, any] = any(NewXmlRpcRemoteHighScoresClient(abeClientInformation, "highscoresubmissionservicessl.php", "HighScoreSubmissionService.process").Get(hashtable, this.NoCrypt)).(*ABHashtable[any, any]); _ = resultHashtable
                remoteHighScores.Update(resultHashtable)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}

