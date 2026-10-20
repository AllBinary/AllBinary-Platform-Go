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

type GameConfigurationCentral struct {
    LogUtil *LogUtil
    SCALE IGameConfiguration
    ORIENTATION IGameConfiguration
    SENSOR_UPDATE_RATE IGameConfiguration
    VIBRATION IGameConfiguration
    CHALLENGE_LEVEL IGameConfiguration
    COLLIDE_DAMAGE IGameConfiguration
    DURABILITY_CHALLENGE_LEVEL IGameConfiguration
    SPEED_CHALLENGE_LEVEL IGameConfiguration
    ATTACK_CHALLENGE_LEVEL IGameConfiguration
    CONTROL_LEVEL IGameConfiguration
    PLAYER_INPUT_WAIT IGameConfiguration
    SPEED IGameConfiguration
    SOUND_VOLUME IGameConfiguration
    MAX_GAME_OBJECTS IGameConfiguration
    MAX_LAYERS IGameConfiguration
    GameControlFidelity int
}

var GameConfigurationCentral_SINGLETON any = NullUtil_GetInstance().NULL_OBJECT

func (this *GameConfigurationCentral) initFields() {
    this.LogUtil = LogUtil_GetInstance()
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    this.GameControlFidelity = 36
}


func GameConfigurationCentral_GetInstance() *GameConfigurationCentral {
        if GameConfigurationCentral_SINGLETON == NullUtil_GetInstance().NULL_OBJECT {
            GameConfigurationCentral_SINGLETON = NewGameConfigurationCentral()
        }
        return any(GameConfigurationCentral_SINGLETON).(*GameConfigurationCentral)
}


func NewGameConfigurationCentral() *GameConfigurationCentral {
    this := &GameConfigurationCentral{}
    this.initFields()
        var smallIntegerSingletonFactory *SmallIntegerSingletonFactory = SmallIntegerSingletonFactory_GetInstance(); _ = smallIntegerSingletonFactory
        this.SCALE = NewGameConfiguration("Scale", smallIntegerSingletonFactory.GetAt(2), smallIntegerSingletonFactory.GetAt(2), smallIntegerSingletonFactory.GetAt(3))
        this.ORIENTATION = NewGameConfiguration("Orientation", smallIntegerSingletonFactory.GetAt(0), smallIntegerSingletonFactory.GetAt(0), smallIntegerSingletonFactory.GetAt(8))
        this.SENSOR_UPDATE_RATE = NewGameConfiguration("Sensor Update Rate", smallIntegerSingletonFactory.GetAt(0), smallIntegerSingletonFactory.GetAt(0), smallIntegerSingletonFactory.GetAt(3))
        this.VIBRATION = NewGameConfiguration("Vibration", smallIntegerSingletonFactory.GetAt(0), smallIntegerSingletonFactory.GetAt(0), smallIntegerSingletonFactory.GetAt(3))
        this.CHALLENGE_LEVEL = NewGameConfiguration("Global Challenge Level", smallIntegerSingletonFactory.GetAt(3), smallIntegerSingletonFactory.GetAt(1), smallIntegerSingletonFactory.GetAt(10))
        this.COLLIDE_DAMAGE = NewGameConfiguration("Collide Damage Challenge Level", smallIntegerSingletonFactory.GetAt(3), smallIntegerSingletonFactory.GetAt(1), smallIntegerSingletonFactory.GetAt(10))
        this.DURABILITY_CHALLENGE_LEVEL = NewGameConfiguration("Enemy Durability Challenge Level", smallIntegerSingletonFactory.GetAt(3), smallIntegerSingletonFactory.GetAt(1), smallIntegerSingletonFactory.GetAt(10))
        this.SPEED_CHALLENGE_LEVEL = NewGameConfiguration("Enemy Speed Challenge Level", smallIntegerSingletonFactory.GetAt(3), smallIntegerSingletonFactory.GetAt(1), smallIntegerSingletonFactory.GetAt(10))
        this.ATTACK_CHALLENGE_LEVEL = NewGameConfiguration("Enemy Attack Challenge Level", smallIntegerSingletonFactory.GetAt(3), smallIntegerSingletonFactory.GetAt(1), smallIntegerSingletonFactory.GetAt(10))
        this.CONTROL_LEVEL = NewGameConfiguration("Control Fidelity", smallIntegerSingletonFactory.GetAt(5), smallIntegerSingletonFactory.GetAt(1), smallIntegerSingletonFactory.GetAt(10))
        this.PLAYER_INPUT_WAIT = NewGameConfiguration("Input Wait", smallIntegerSingletonFactory.GetAt(5), smallIntegerSingletonFactory.GetAt(1), smallIntegerSingletonFactory.GetAt(10))
        this.SPEED = NewGameConfiguration("Game Speed (Device Specific)", smallIntegerSingletonFactory.GetAt(5), smallIntegerSingletonFactory.GetAt(1), smallIntegerSingletonFactory.GetAt(10))
        this.SOUND_VOLUME = NewGameConfiguration("Sound Volume", smallIntegerSingletonFactory.GetAt(6), smallIntegerSingletonFactory.GetAt(0), smallIntegerSingletonFactory.GetAt(10))
        this.MAX_GAME_OBJECTS = NewGameConfiguration("Max Game Objects", smallIntegerSingletonFactory.GetAt(20), smallIntegerSingletonFactory.GetAt(1), smallIntegerSingletonFactory.GetAt(50))
        this.MAX_LAYERS = NewGameConfiguration("Max Layers", smallIntegerSingletonFactory.GetAt(20), smallIntegerSingletonFactory.GetAt(1), smallIntegerSingletonFactory.GetAt(50))
    return this
}


func (this *GameConfigurationCentral) Load(abeClientInformation AbeClientInformationInterface) {
        var commonStrings *CommonStrings = CommonStrings_GetInstance(); _ = commonStrings
        if tryControl1 := func() (tryControl1 int) {
            defer func() {
                if recovered := recover(); recovered != nil {
                    e := Java_ToError(recovered)
                    _ = e
            this.LogUtil.Put(commonStrings.EXCEPTION, this, commonStrings.LOAD, e)
                }
            }()
            var smallIntegerSingletonFactory *SmallIntegerSingletonFactory = SmallIntegerSingletonFactory_GetInstance(); _ = smallIntegerSingletonFactory
            var keyValuePersistance IKeyValuePersistance = GameConfigurationPersistanceSingleton_GetInstance(); _ = keyValuePersistance
            keyValuePersistance.Clear()
            keyValuePersistance.LoadAll(abeClientInformation)
            var list IBasicArrayList = keyValuePersistance.GetIds(); _ = list
            if list.Size() > 0 {
                var hashtable *ABHashtable[any, any] = keyValuePersistance.Get(0); _ = hashtable
                var anyType any = any(this.SCALE.GetName()).(any); _ = anyType
                var value string = any(hashtable.Get(anyType)).(string); _ = value
                this.SCALE.SetValue(smallIntegerSingletonFactory.GetAt(int(fmt.Sprint(value))))
            } else {
                this.LogUtil.PutF("No Game Configuration To Load", this, commonStrings.LOAD)
            }
            return
        }(); tryControl1 == 1 {
            return
        }
}


func (this *GameConfigurationCentral) SetGameControlFidelity(gameControlFidelity int) {
        this.GameControlFidelity = gameControlFidelity
}


func (this *GameConfigurationCentral) GetGameControlFidelity() int {
        return this.GameControlFidelity
}


func (this *GameConfigurationCentral) ToString() string {
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        var commonStrings *CommonSeps = CommonSeps_GetInstance(); _ = commonStrings
        stringBuffer.Append(commonStrings.COMMA_SEP)
        stringBuffer.Append(this.SCALE.GetName())
        stringBuffer.Append(commonStrings.EQUALS)
        stringBuffer.Appendint(int(this.SCALE.GetValue()))
        stringBuffer.Append(commonStrings.COMMA_SEP)
        stringBuffer.Append(this.ORIENTATION.GetName())
        stringBuffer.Append(commonStrings.EQUALS)
        stringBuffer.Appendint(int(this.ORIENTATION.GetValue()))
        stringBuffer.Append(commonStrings.COMMA_SEP)
        stringBuffer.Append(this.SENSOR_UPDATE_RATE.GetName())
        stringBuffer.Append(commonStrings.EQUALS)
        stringBuffer.Appendint(int(this.SENSOR_UPDATE_RATE.GetValue()))
        stringBuffer.Append(commonStrings.COMMA_SEP)
        stringBuffer.Append(this.VIBRATION.GetName())
        stringBuffer.Append(commonStrings.EQUALS)
        stringBuffer.Appendint(int(this.VIBRATION.GetValue()))
        stringBuffer.Append(commonStrings.COMMA_SEP)
        stringBuffer.Append(this.CHALLENGE_LEVEL.GetName())
        stringBuffer.Append(commonStrings.EQUALS)
        stringBuffer.Appendint(int(this.CHALLENGE_LEVEL.GetValue()))
        stringBuffer.Append(commonStrings.COMMA_SEP)
        stringBuffer.Append(this.COLLIDE_DAMAGE.GetName())
        stringBuffer.Append(commonStrings.EQUALS)
        stringBuffer.Appendint(int(this.COLLIDE_DAMAGE.GetValue()))
        stringBuffer.Append(commonStrings.COMMA_SEP)
        stringBuffer.Append(this.DURABILITY_CHALLENGE_LEVEL.GetName())
        stringBuffer.Append(commonStrings.EQUALS)
        stringBuffer.Appendint(int(this.DURABILITY_CHALLENGE_LEVEL.GetValue()))
        stringBuffer.Append(commonStrings.COMMA_SEP)
        stringBuffer.Append(this.SPEED_CHALLENGE_LEVEL.GetName())
        stringBuffer.Append(commonStrings.EQUALS)
        stringBuffer.Appendint(int(this.SPEED_CHALLENGE_LEVEL.GetValue()))
        stringBuffer.Append(commonStrings.COMMA_SEP)
        stringBuffer.Append(this.ATTACK_CHALLENGE_LEVEL.GetName())
        stringBuffer.Append(commonStrings.EQUALS)
        stringBuffer.Appendint(int(this.ATTACK_CHALLENGE_LEVEL.GetValue()))
        stringBuffer.Append(commonStrings.COMMA_SEP)
        stringBuffer.Append(this.CONTROL_LEVEL.GetName())
        stringBuffer.Append(commonStrings.EQUALS)
        stringBuffer.Appendint(int(this.CONTROL_LEVEL.GetValue()))
        stringBuffer.Append(commonStrings.COMMA_SEP)
        stringBuffer.Append(this.SPEED.GetName())
        stringBuffer.Append(commonStrings.EQUALS)
        stringBuffer.Appendint(int(this.SPEED.GetValue()))
        stringBuffer.Append(commonStrings.COMMA_SEP)
        stringBuffer.Append(this.MAX_GAME_OBJECTS.GetName())
        stringBuffer.Append(commonStrings.EQUALS)
        stringBuffer.Appendint(int(this.MAX_GAME_OBJECTS.GetValue()))
        stringBuffer.Append(commonStrings.COMMA_SEP)
        stringBuffer.Append(this.MAX_LAYERS.GetName())
        stringBuffer.Append(commonStrings.EQUALS)
        stringBuffer.Appendint(int(this.MAX_LAYERS.GetValue()))
        return fmt.Sprint(stringBuffer)
}

func (this *GameConfigurationCentral) String() string {
    return this.ToString()
}

