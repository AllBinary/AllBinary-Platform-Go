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

type GameConfigurationUtil struct {
    LogUtil *LogUtil
    CommonStrings *CommonStrings
    GAUGE_UPDATE string
    FROM string
    TO string
    COMPETITION_VALUE int
}

var GameConfigurationUtil_Instance *GameConfigurationUtil = NewGameConfigurationUtil()

func (this *GameConfigurationUtil) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    this.CommonStrings = CommonStrings_GetInstance()
    this.GAUGE_UPDATE = "Gauge Update: "
    this.FROM = " from: "
    this.TO = " to: "
    this.COMPETITION_VALUE = 1
}

func NewGameConfigurationUtil() *GameConfigurationUtil {
    this := &GameConfigurationUtil{}
    this.initFields()
    return this
}


func GameConfigurationUtil_GetInstance() *GameConfigurationUtil {
        return GameConfigurationUtil_Instance
}


// throws Exception
func (this *GameConfigurationUtil) Change(gameOptionsForm *GameOptionsForm, gauge *GameConfigurationGauge) {
        this.Update(gauge)
        var gameConfigurationSingleton *GameConfigurationSingleton = GameConfigurationSingleton_GetInstance(); _ = gameConfigurationSingleton
        var gameConfiguration IGameConfiguration = gameConfigurationSingleton.GetInstanceByName(gauge.GetLabel()); _ = gameConfiguration
        this.UpdateChallange(gameOptionsForm, gameConfiguration)
}


// throws Exception
func (this *GameConfigurationUtil) Update(gauge *GameConfigurationGauge) {
        var gameConfigurationSingleton *GameConfigurationSingleton = GameConfigurationSingleton_GetInstance(); _ = gameConfigurationSingleton
        var gameConfiguration IGameConfiguration = gameConfigurationSingleton.GetInstanceByName(gauge.GetLabel()); _ = gameConfiguration
        var smallIntegerSingletonFactory *SmallIntegerSingletonFactory = SmallIntegerSingletonFactory_GetInstance(); _ = smallIntegerSingletonFactory
        var value int = smallIntegerSingletonFactory.GetAt(gauge.GetValue() + int(gameConfiguration.GetMinValue())); _ = value
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append(this.GAUGE_UPDATE)
        stringBuffer.Append(gameConfiguration.GetName())
        stringBuffer.Append(this.FROM)
        stringBuffer.Appendint(int(gameConfiguration.GetValue()))
        stringBuffer.Append(this.TO)
        stringBuffer.Appendint(int(value))
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, this.CommonStrings.UPDATE)
        gameConfiguration.SetValue(value)
}


// throws Exception
func (this *GameConfigurationUtil) SetDefault(gauge *GameConfigurationGauge) {
        var gameConfigurationSingleton *GameConfigurationSingleton = GameConfigurationSingleton_GetInstance(); _ = gameConfigurationSingleton
        var gameConfiguration IGameConfiguration = gameConfigurationSingleton.GetInstanceByName(gauge.GetLabel()); _ = gameConfiguration
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("Gauge Default: ")
        stringBuffer.Append(gameConfiguration.GetName())
        stringBuffer.Append(this.TO)
        stringBuffer.Appendint(int(gameConfiguration.GetDefaultValue()))
        this.LogUtil.PutF(fmt.Sprint(stringBuffer), this, "setDefault")
        gauge.SetValue(int(gameConfiguration.GetDefaultValue()) - int(gameConfiguration.GetMinValue()))
        gameConfiguration.SetValue(gameConfiguration.GetDefaultValue())
}


// throws Exception
func (this *GameConfigurationUtil) UpdateChallange(gameOptionsForm *GameOptionsForm, gameConfiguration IGameConfiguration) {
        var gameConfigurationSingleton *GameConfigurationSingleton = GameConfigurationSingleton_GetInstance(); _ = gameConfigurationSingleton
        var gameConfigurationCentral *GameConfigurationCentral = GameConfigurationCentral_GetInstance(); _ = gameConfigurationCentral
        if gameConfiguration == gameConfigurationCentral.CHALLENGE_LEVEL {
            this.LogUtil.PutF(this.CommonStrings.START, this, "updateChallange")
            gameConfigurationCentral.COLLIDE_DAMAGE.SetValue(gameConfiguration.GetValue())
            gameConfigurationCentral.ATTACK_CHALLENGE_LEVEL.SetValue(gameConfiguration.GetValue())
            gameConfigurationCentral.DURABILITY_CHALLENGE_LEVEL.SetValue(gameConfiguration.GetValue())
            gameConfigurationCentral.SPEED_CHALLENGE_LEVEL.SetValue(gameConfiguration.GetValue())
            var size int = gameOptionsForm.Size(); _ = size
            for index := int(0); index < size; index++ {
                var item IItem = gameOptionsForm.Get(index); _ = item
                if func() bool { _, ok := any(item).(*GameConfigurationGauge); return ok }() {
                    var gauge *GameConfigurationGauge = any(item).(*GameConfigurationGauge); _ = gauge
                    var name string = item.GetLabel(); _ = name
                    var nextGameConfiguration IGameConfiguration = gameConfigurationSingleton.GetInstanceByName(name); _ = nextGameConfiguration
                    if nextGameConfiguration == gameConfigurationCentral.COLLIDE_DAMAGE {
                        gauge.SetValue(int(gameConfigurationCentral.COLLIDE_DAMAGE.GetValue()) - int(gameConfiguration.GetMinValue()))
                    } else if nextGameConfiguration == gameConfigurationCentral.DURABILITY_CHALLENGE_LEVEL {
                        gauge.SetValue(int(gameConfigurationCentral.DURABILITY_CHALLENGE_LEVEL.GetValue()) - int(gameConfiguration.GetMinValue()))
                    } else if nextGameConfiguration == gameConfigurationCentral.ATTACK_CHALLENGE_LEVEL {
                        gauge.SetValue(int(gameConfigurationCentral.ATTACK_CHALLENGE_LEVEL.GetValue()) - int(gameConfiguration.GetMinValue()))
                    } else if nextGameConfiguration == gameConfigurationCentral.SPEED_CHALLENGE_LEVEL {
                        gauge.SetValue(int(gameConfigurationCentral.SPEED_CHALLENGE_LEVEL.GetValue()) - int(gameConfiguration.GetMinValue()))
                    }
                }
            }
        }
}


func (this *GameConfigurationUtil) UpdateCompetitionValue() {
        this.LogUtil.PutF(this.CommonStrings.START, this, "updateCompetitionValue")
        var gameConfigurationCentral *GameConfigurationCentral = GameConfigurationCentral_GetInstance(); _ = gameConfigurationCentral
        this.COMPETITION_VALUE = int(gameConfigurationCentral.COLLIDE_DAMAGE.GetValue()) + int(gameConfigurationCentral.ATTACK_CHALLENGE_LEVEL.GetValue()) + int(gameConfigurationCentral.DURABILITY_CHALLENGE_LEVEL.GetValue()) + int(gameConfigurationCentral.SPEED_CHALLENGE_LEVEL.GetValue())
        this.COMPETITION_VALUE = (this.COMPETITION_VALUE >> 2)
}


func (this *GameConfigurationUtil) GetCompetitionValue() int {
        return this.COMPETITION_VALUE
}

