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

type HighScoreUtil struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    HighScoresFactoryInterface HighScoresFactoryInterface
    HighScoresArray []IHighScores
    HighScore *HighScore
    AbeClientInformation AbeClientInformationInterface
    GameInfo *GameInfo
    HighScoresHelper HighScoresHelperBaseInterface
    FirstTime bool
}

// Implements HighScoresResultsListener

var HighScoreUtil_SUBMIT_TEXTBOX_COMMAND ICommand = NewCommand("Submit", StringUtil_GetInstance().EMPTY_STRING, Command_SCREEN, 1)

func (this *HighScoreUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    
    
    
    
    
    
    this.FirstTime = true
}


func NewHighScoreUtil(highScoresFactoryInterface HighScoresFactoryInterface, highScoresHelper HighScoresHelperBaseInterface, abeClientInformation AbeClientInformationInterface, gameInfo *GameInfo, cmdListener CommandListener, name string, highScore *HighScore) *HighScoreUtil {
    this := &HighScoreUtil{}
    this.initFields()
        this.HighScoresFactoryInterface = highScoresFactoryInterface
        this.HighScoresHelper = highScoresHelper
        this.HighScoresArray = this.HighScoresHelper.GetHighScoresArray()
        this.HighScore = highScore
        this.AbeClientInformation = abeClientInformation
        this.GameInfo = gameInfo
    return this
}


func (this *HighScoreUtil) SetHighScoresArray(highScoresArray []IHighScores) {
        this.HighScoresArray = highScoresArray
        this.FirstTime = false
        this.SaveHighScore()
}


func (this *HighScoreUtil) Update(name string) {
        HighScoreNamePersistanceSingleton_GetInstance().Save(this.AbeClientInformation, this.GameInfo, name)
        this.HighScore.SetName(name)
}


func (this *HighScoreUtil) SaveHighScore() {
        this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append(this.CommonStrings.START).Append(StringUtil_GetInstance().ToString(this.HighScore))), this, "saveHighScore")
        var size int = len(this.HighScoresArray); _ = size
        if this.FirstTime && size == 0 {
            this.LogUtil.PutF("Games canvas did not give us any HighScores", this, "saveHighScore")
            this.HighScoresFactoryInterface.FetchHighScores(this.GameInfo, this)
            return
        }
        var highScores IHighScores; _ = highScores
        var highScoresAsString string; _ = highScoresAsString
        for index := int(0); index < size; index++ {
            highScores = this.HighScoresArray[index]
            highScores.AddHighScore(this.HighScore)
            highScoresAsString = fmt.Sprint(highScores)
            this.LogUtil.PutF(fmt.Sprint(NewStringMaker().Append("Added/Adding Score: ").Append(highScoresAsString)), this, "saveHighScore")
        }
        this.HighScoresHelper.SetHighScoresArray(this.HighScoresArray)
}


func (this *HighScoreUtil) Submit(myCanvas IMyCanvas) {
        var commandListener CommandListener = myCanvas.GetCustomCommandListener(); _ = commandListener
        commandListener.CommandAction(HighScoreUtil_SUBMIT_TEXTBOX_COMMAND, myCanvas)
}

