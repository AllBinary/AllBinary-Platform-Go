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

type NoDecimalTrigTable struct {
    MathUtil *MathUtil
    NoDecimalSin []int64
    NoDecimalCos []int64
    NoDecimalTan []int64
    SCALE int
}

var NoDecimalTrigTable_Instance *NoDecimalTrigTable = NewNoDecimalTrigTable()

func (this *NoDecimalTrigTable) initFields() {
    this.MathUtil = MathUtil_GetInstance()
    
    
    
    this.SCALE = 10000
}


func NoDecimalTrigTable_GetInstance() *NoDecimalTrigTable {
        return NoDecimalTrigTable_Instance
}


func NewNoDecimalTrigTable() *NoDecimalTrigTable {
    this := &NoDecimalTrigTable{}
    this.initFields()
        var TOTAL_ANGLE int = int(AngleFactory_GetInstance().TOTAL_ANGLE); _ = TOTAL_ANGLE
        var MAX_VALUE int64 = (int64(Integer_MAX_VALUE)); _ = MAX_VALUE
        this.NoDecimalSin = make([]int64, TOTAL_ANGLE)
        this.NoDecimalCos = make([]int64, TOTAL_ANGLE)
        this.NoDecimalTan = make([]int64, TOTAL_ANGLE + 1)
        this.NoDecimalSin[0] = 0
        this.NoDecimalSin[1] = 174
        this.NoDecimalSin[2] = 348
        this.NoDecimalSin[3] = 523
        this.NoDecimalSin[4] = 697
        this.NoDecimalSin[5] = 871
        this.NoDecimalSin[6] = 1045
        this.NoDecimalSin[7] = 1218
        this.NoDecimalSin[8] = 1391
        this.NoDecimalSin[9] = 1564
        this.NoDecimalSin[10] = 1736
        this.NoDecimalSin[11] = 1908
        this.NoDecimalSin[12] = 2079
        this.NoDecimalSin[13] = 2249
        this.NoDecimalSin[14] = 2419
        this.NoDecimalSin[15] = 2588
        this.NoDecimalSin[16] = 2756
        this.NoDecimalSin[17] = 2923
        this.NoDecimalSin[18] = 3090
        this.NoDecimalSin[19] = 3255
        this.NoDecimalSin[20] = 3420
        this.NoDecimalSin[21] = 3583
        this.NoDecimalSin[22] = 3746
        this.NoDecimalSin[23] = 3907
        this.NoDecimalSin[24] = 4067
        this.NoDecimalSin[25] = 4226
        this.NoDecimalSin[26] = 4383
        this.NoDecimalSin[27] = 4539
        this.NoDecimalSin[28] = 4694
        this.NoDecimalSin[29] = 4848
        this.NoDecimalSin[30] = 4999
        this.NoDecimalSin[31] = 5150
        this.NoDecimalSin[32] = 5299
        this.NoDecimalSin[33] = 5446
        this.NoDecimalSin[34] = 5591
        this.NoDecimalSin[35] = 5735
        this.NoDecimalSin[36] = 5877
        this.NoDecimalSin[37] = 6018
        this.NoDecimalSin[38] = 6156
        this.NoDecimalSin[39] = 6293
        this.NoDecimalSin[40] = 6427
        this.NoDecimalSin[41] = 6560
        this.NoDecimalSin[42] = 6691
        this.NoDecimalSin[43] = 6819
        this.NoDecimalSin[44] = 6946
        this.NoDecimalSin[45] = 7071
        this.NoDecimalSin[46] = 7193
        this.NoDecimalSin[47] = 7313
        this.NoDecimalSin[48] = 7431
        this.NoDecimalSin[49] = 7547
        this.NoDecimalSin[50] = 7660
        this.NoDecimalSin[51] = 7771
        this.NoDecimalSin[52] = 7880
        this.NoDecimalSin[53] = 7986
        this.NoDecimalSin[54] = 8090
        this.NoDecimalSin[55] = 8191
        this.NoDecimalSin[56] = 8290
        this.NoDecimalSin[57] = 8386
        this.NoDecimalSin[58] = 8480
        this.NoDecimalSin[59] = 8571
        this.NoDecimalSin[60] = 8660
        this.NoDecimalSin[61] = 8746
        this.NoDecimalSin[62] = 8829
        this.NoDecimalSin[63] = 8910
        this.NoDecimalSin[64] = 8987
        this.NoDecimalSin[65] = 9063
        this.NoDecimalSin[66] = 9135
        this.NoDecimalSin[67] = 9205
        this.NoDecimalSin[68] = 9271
        this.NoDecimalSin[69] = 9335
        this.NoDecimalSin[70] = 9396
        this.NoDecimalSin[71] = 9455
        this.NoDecimalSin[72] = 9510
        this.NoDecimalSin[73] = 9563
        this.NoDecimalSin[74] = 9612
        this.NoDecimalSin[75] = 9659
        this.NoDecimalSin[76] = 9702
        this.NoDecimalSin[77] = 9743
        this.NoDecimalSin[78] = 9781
        this.NoDecimalSin[79] = 9816
        this.NoDecimalSin[80] = 9848
        this.NoDecimalSin[81] = 9876
        this.NoDecimalSin[82] = 9902
        this.NoDecimalSin[83] = 9925
        this.NoDecimalSin[84] = 9945
        this.NoDecimalSin[85] = 9961
        this.NoDecimalSin[86] = 9975
        this.NoDecimalSin[87] = 9986
        this.NoDecimalSin[88] = 9993
        this.NoDecimalSin[89] = 9998
        this.NoDecimalSin[90] = 10000
        this.NoDecimalSin[91] = 9998
        this.NoDecimalSin[92] = 9993
        this.NoDecimalSin[93] = 9986
        this.NoDecimalSin[94] = 9975
        this.NoDecimalSin[95] = 9961
        this.NoDecimalSin[96] = 9945
        this.NoDecimalSin[97] = 9925
        this.NoDecimalSin[98] = 9902
        this.NoDecimalSin[99] = 9876
        this.NoDecimalSin[100] = 9848
        this.NoDecimalSin[101] = 9816
        this.NoDecimalSin[102] = 9781
        this.NoDecimalSin[103] = 9743
        this.NoDecimalSin[104] = 9702
        this.NoDecimalSin[105] = 9659
        this.NoDecimalSin[106] = 9612
        this.NoDecimalSin[107] = 9563
        this.NoDecimalSin[108] = 9510
        this.NoDecimalSin[109] = 9455
        this.NoDecimalSin[110] = 9396
        this.NoDecimalSin[111] = 9335
        this.NoDecimalSin[112] = 9271
        this.NoDecimalSin[113] = 9205
        this.NoDecimalSin[114] = 9135
        this.NoDecimalSin[115] = 9063
        this.NoDecimalSin[116] = 8987
        this.NoDecimalSin[117] = 8910
        this.NoDecimalSin[118] = 8829
        this.NoDecimalSin[119] = 8746
        this.NoDecimalSin[120] = 8660
        this.NoDecimalSin[121] = 8571
        this.NoDecimalSin[122] = 8480
        this.NoDecimalSin[123] = 8386
        this.NoDecimalSin[124] = 8290
        this.NoDecimalSin[125] = 8191
        this.NoDecimalSin[126] = 8090
        this.NoDecimalSin[127] = 7986
        this.NoDecimalSin[128] = 7880
        this.NoDecimalSin[129] = 7771
        this.NoDecimalSin[130] = 7660
        this.NoDecimalSin[131] = 7547
        this.NoDecimalSin[132] = 7431
        this.NoDecimalSin[133] = 7313
        this.NoDecimalSin[134] = 7193
        this.NoDecimalSin[135] = 7071
        this.NoDecimalSin[136] = 6946
        this.NoDecimalSin[137] = 6819
        this.NoDecimalSin[138] = 6691
        this.NoDecimalSin[139] = 6560
        this.NoDecimalSin[140] = 6427
        this.NoDecimalSin[141] = 6293
        this.NoDecimalSin[142] = 6156
        this.NoDecimalSin[143] = 6018
        this.NoDecimalSin[144] = 5877
        this.NoDecimalSin[145] = 5735
        this.NoDecimalSin[146] = 5591
        this.NoDecimalSin[147] = 5446
        this.NoDecimalSin[148] = 5299
        this.NoDecimalSin[149] = 5150
        this.NoDecimalSin[150] = 4999
        this.NoDecimalSin[151] = 4848
        this.NoDecimalSin[152] = 4694
        this.NoDecimalSin[153] = 4539
        this.NoDecimalSin[154] = 4383
        this.NoDecimalSin[155] = 4226
        this.NoDecimalSin[156] = 4067
        this.NoDecimalSin[157] = 3907
        this.NoDecimalSin[158] = 3746
        this.NoDecimalSin[159] = 3583
        this.NoDecimalSin[160] = 3420
        this.NoDecimalSin[161] = 3255
        this.NoDecimalSin[162] = 3090
        this.NoDecimalSin[163] = 2923
        this.NoDecimalSin[164] = 2756
        this.NoDecimalSin[165] = 2588
        this.NoDecimalSin[166] = 2419
        this.NoDecimalSin[167] = 2249
        this.NoDecimalSin[168] = 2079
        this.NoDecimalSin[169] = 1908
        this.NoDecimalSin[170] = 1736
        this.NoDecimalSin[171] = 1564
        this.NoDecimalSin[172] = 1391
        this.NoDecimalSin[173] = 1218
        this.NoDecimalSin[174] = 1045
        this.NoDecimalSin[175] = 871
        this.NoDecimalSin[176] = 697
        this.NoDecimalSin[177] = 523
        this.NoDecimalSin[178] = 348
        this.NoDecimalSin[179] = 174
        this.NoDecimalSin[180] = 0
        this.NoDecimalSin[181] = int64(-174)
        this.NoDecimalSin[182] = int64(-348)
        this.NoDecimalSin[183] = int64(-523)
        this.NoDecimalSin[184] = int64(-697)
        this.NoDecimalSin[185] = int64(-871)
        this.NoDecimalSin[186] = int64(-1045)
        this.NoDecimalSin[187] = int64(-1218)
        this.NoDecimalSin[188] = int64(-1391)
        this.NoDecimalSin[189] = int64(-1564)
        this.NoDecimalSin[190] = int64(-1736)
        this.NoDecimalSin[191] = int64(-1908)
        this.NoDecimalSin[192] = int64(-2079)
        this.NoDecimalSin[193] = int64(-2249)
        this.NoDecimalSin[194] = int64(-2419)
        this.NoDecimalSin[195] = int64(-2588)
        this.NoDecimalSin[196] = int64(-2756)
        this.NoDecimalSin[197] = int64(-2923)
        this.NoDecimalSin[198] = int64(-3090)
        this.NoDecimalSin[199] = int64(-3255)
        this.NoDecimalSin[200] = int64(-3420)
        this.NoDecimalSin[201] = int64(-3583)
        this.NoDecimalSin[202] = int64(-3746)
        this.NoDecimalSin[203] = int64(-3907)
        this.NoDecimalSin[204] = int64(-4067)
        this.NoDecimalSin[205] = int64(-4226)
        this.NoDecimalSin[206] = int64(-4383)
        this.NoDecimalSin[207] = int64(-4539)
        this.NoDecimalSin[208] = int64(-4694)
        this.NoDecimalSin[209] = int64(-4848)
        this.NoDecimalSin[210] = int64(-5000)
        this.NoDecimalSin[211] = int64(-5150)
        this.NoDecimalSin[212] = int64(-5299)
        this.NoDecimalSin[213] = int64(-5446)
        this.NoDecimalSin[214] = int64(-5591)
        this.NoDecimalSin[215] = int64(-5735)
        this.NoDecimalSin[216] = int64(-5877)
        this.NoDecimalSin[217] = int64(-6018)
        this.NoDecimalSin[218] = int64(-6156)
        this.NoDecimalSin[219] = int64(-6293)
        this.NoDecimalSin[220] = int64(-6427)
        this.NoDecimalSin[221] = int64(-6560)
        this.NoDecimalSin[222] = int64(-6691)
        this.NoDecimalSin[223] = int64(-6819)
        this.NoDecimalSin[224] = int64(-6946)
        this.NoDecimalSin[225] = int64(-7071)
        this.NoDecimalSin[226] = int64(-7193)
        this.NoDecimalSin[227] = int64(-7313)
        this.NoDecimalSin[228] = int64(-7431)
        this.NoDecimalSin[229] = int64(-7547)
        this.NoDecimalSin[230] = int64(-7660)
        this.NoDecimalSin[231] = int64(-7771)
        this.NoDecimalSin[232] = int64(-7880)
        this.NoDecimalSin[233] = int64(-7986)
        this.NoDecimalSin[234] = int64(-8090)
        this.NoDecimalSin[235] = int64(-8191)
        this.NoDecimalSin[236] = int64(-8290)
        this.NoDecimalSin[237] = int64(-8386)
        this.NoDecimalSin[238] = int64(-8480)
        this.NoDecimalSin[239] = int64(-8571)
        this.NoDecimalSin[240] = int64(-8660)
        this.NoDecimalSin[241] = int64(-8746)
        this.NoDecimalSin[242] = int64(-8829)
        this.NoDecimalSin[243] = int64(-8910)
        this.NoDecimalSin[244] = int64(-8987)
        this.NoDecimalSin[245] = int64(-9063)
        this.NoDecimalSin[246] = int64(-9135)
        this.NoDecimalSin[247] = int64(-9205)
        this.NoDecimalSin[248] = int64(-9271)
        this.NoDecimalSin[249] = int64(-9335)
        this.NoDecimalSin[250] = int64(-9396)
        this.NoDecimalSin[251] = int64(-9455)
        this.NoDecimalSin[252] = int64(-9510)
        this.NoDecimalSin[253] = int64(-9563)
        this.NoDecimalSin[254] = int64(-9612)
        this.NoDecimalSin[255] = int64(-9659)
        this.NoDecimalSin[256] = int64(-9702)
        this.NoDecimalSin[257] = int64(-9743)
        this.NoDecimalSin[258] = int64(-9781)
        this.NoDecimalSin[259] = int64(-9816)
        this.NoDecimalSin[260] = int64(-9848)
        this.NoDecimalSin[261] = int64(-9876)
        this.NoDecimalSin[262] = int64(-9902)
        this.NoDecimalSin[263] = int64(-9925)
        this.NoDecimalSin[264] = int64(-9945)
        this.NoDecimalSin[265] = int64(-9961)
        this.NoDecimalSin[266] = int64(-9975)
        this.NoDecimalSin[267] = int64(-9986)
        this.NoDecimalSin[268] = int64(-9993)
        this.NoDecimalSin[269] = int64(-9998)
        this.NoDecimalSin[270] = int64(-10000)
        this.NoDecimalSin[271] = int64(-9998)
        this.NoDecimalSin[272] = int64(-9993)
        this.NoDecimalSin[273] = int64(-9986)
        this.NoDecimalSin[274] = int64(-9975)
        this.NoDecimalSin[275] = int64(-9961)
        this.NoDecimalSin[276] = int64(-9945)
        this.NoDecimalSin[277] = int64(-9925)
        this.NoDecimalSin[278] = int64(-9902)
        this.NoDecimalSin[279] = int64(-9876)
        this.NoDecimalSin[280] = int64(-9848)
        this.NoDecimalSin[281] = int64(-9816)
        this.NoDecimalSin[282] = int64(-9781)
        this.NoDecimalSin[283] = int64(-9743)
        this.NoDecimalSin[284] = int64(-9702)
        this.NoDecimalSin[285] = int64(-9659)
        this.NoDecimalSin[286] = int64(-9612)
        this.NoDecimalSin[287] = int64(-9563)
        this.NoDecimalSin[288] = int64(-9510)
        this.NoDecimalSin[289] = int64(-9455)
        this.NoDecimalSin[290] = int64(-9396)
        this.NoDecimalSin[291] = int64(-9335)
        this.NoDecimalSin[292] = int64(-9271)
        this.NoDecimalSin[293] = int64(-9205)
        this.NoDecimalSin[294] = int64(-9135)
        this.NoDecimalSin[295] = int64(-9063)
        this.NoDecimalSin[296] = int64(-8987)
        this.NoDecimalSin[297] = int64(-8910)
        this.NoDecimalSin[298] = int64(-8829)
        this.NoDecimalSin[299] = int64(-8746)
        this.NoDecimalSin[300] = int64(-8660)
        this.NoDecimalSin[301] = int64(-8571)
        this.NoDecimalSin[302] = int64(-8480)
        this.NoDecimalSin[303] = int64(-8386)
        this.NoDecimalSin[304] = int64(-8290)
        this.NoDecimalSin[305] = int64(-8191)
        this.NoDecimalSin[306] = int64(-8090)
        this.NoDecimalSin[307] = int64(-7986)
        this.NoDecimalSin[308] = int64(-7880)
        this.NoDecimalSin[309] = int64(-7771)
        this.NoDecimalSin[310] = int64(-7660)
        this.NoDecimalSin[311] = int64(-7547)
        this.NoDecimalSin[312] = int64(-7431)
        this.NoDecimalSin[313] = int64(-7313)
        this.NoDecimalSin[314] = int64(-7193)
        this.NoDecimalSin[315] = int64(-7071)
        this.NoDecimalSin[316] = int64(-6946)
        this.NoDecimalSin[317] = int64(-6819)
        this.NoDecimalSin[318] = int64(-6691)
        this.NoDecimalSin[319] = int64(-6560)
        this.NoDecimalSin[320] = int64(-6427)
        this.NoDecimalSin[321] = int64(-6293)
        this.NoDecimalSin[322] = int64(-6156)
        this.NoDecimalSin[323] = int64(-6018)
        this.NoDecimalSin[324] = int64(-5877)
        this.NoDecimalSin[325] = int64(-5735)
        this.NoDecimalSin[326] = int64(-5591)
        this.NoDecimalSin[327] = int64(-5446)
        this.NoDecimalSin[328] = int64(-5299)
        this.NoDecimalSin[329] = int64(-5150)
        this.NoDecimalSin[330] = int64(-5000)
        this.NoDecimalSin[331] = int64(-4848)
        this.NoDecimalSin[332] = int64(-4694)
        this.NoDecimalSin[333] = int64(-4539)
        this.NoDecimalSin[334] = int64(-4383)
        this.NoDecimalSin[335] = int64(-4226)
        this.NoDecimalSin[336] = int64(-4067)
        this.NoDecimalSin[337] = int64(-3907)
        this.NoDecimalSin[338] = int64(-3746)
        this.NoDecimalSin[339] = int64(-3583)
        this.NoDecimalSin[340] = int64(-3420)
        this.NoDecimalSin[341] = int64(-3255)
        this.NoDecimalSin[342] = int64(-3090)
        this.NoDecimalSin[343] = int64(-2923)
        this.NoDecimalSin[344] = int64(-2756)
        this.NoDecimalSin[345] = int64(-2588)
        this.NoDecimalSin[346] = int64(-2419)
        this.NoDecimalSin[347] = int64(-2249)
        this.NoDecimalSin[348] = int64(-2079)
        this.NoDecimalSin[349] = int64(-1908)
        this.NoDecimalSin[350] = int64(-1736)
        this.NoDecimalSin[351] = int64(-1564)
        this.NoDecimalSin[352] = int64(-1391)
        this.NoDecimalSin[353] = int64(-1218)
        this.NoDecimalSin[354] = int64(-1045)
        this.NoDecimalSin[355] = int64(-871)
        this.NoDecimalSin[356] = int64(-697)
        this.NoDecimalSin[357] = int64(-523)
        this.NoDecimalSin[358] = int64(-348)
        this.NoDecimalSin[359] = int64(-174)
        this.NoDecimalCos[0] = 10000
        this.NoDecimalCos[1] = 9998
        this.NoDecimalCos[2] = 9993
        this.NoDecimalCos[3] = 9986
        this.NoDecimalCos[4] = 9975
        this.NoDecimalCos[5] = 9961
        this.NoDecimalCos[6] = 9945
        this.NoDecimalCos[7] = 9925
        this.NoDecimalCos[8] = 9902
        this.NoDecimalCos[9] = 9876
        this.NoDecimalCos[10] = 9848
        this.NoDecimalCos[11] = 9816
        this.NoDecimalCos[12] = 9781
        this.NoDecimalCos[13] = 9743
        this.NoDecimalCos[14] = 9702
        this.NoDecimalCos[15] = 9659
        this.NoDecimalCos[16] = 9612
        this.NoDecimalCos[17] = 9563
        this.NoDecimalCos[18] = 9510
        this.NoDecimalCos[19] = 9455
        this.NoDecimalCos[20] = 9396
        this.NoDecimalCos[21] = 9335
        this.NoDecimalCos[22] = 9271
        this.NoDecimalCos[23] = 9205
        this.NoDecimalCos[24] = 9135
        this.NoDecimalCos[25] = 9063
        this.NoDecimalCos[26] = 8987
        this.NoDecimalCos[27] = 8910
        this.NoDecimalCos[28] = 8829
        this.NoDecimalCos[29] = 8746
        this.NoDecimalCos[30] = 8660
        this.NoDecimalCos[31] = 8571
        this.NoDecimalCos[32] = 8480
        this.NoDecimalCos[33] = 8386
        this.NoDecimalCos[34] = 8290
        this.NoDecimalCos[35] = 8191
        this.NoDecimalCos[36] = 8090
        this.NoDecimalCos[37] = 7986
        this.NoDecimalCos[38] = 7880
        this.NoDecimalCos[39] = 7771
        this.NoDecimalCos[40] = 7660
        this.NoDecimalCos[41] = 7547
        this.NoDecimalCos[42] = 7431
        this.NoDecimalCos[43] = 7313
        this.NoDecimalCos[44] = 7193
        this.NoDecimalCos[45] = 7071
        this.NoDecimalCos[46] = 6946
        this.NoDecimalCos[47] = 6819
        this.NoDecimalCos[48] = 6691
        this.NoDecimalCos[49] = 6560
        this.NoDecimalCos[50] = 6427
        this.NoDecimalCos[51] = 6293
        this.NoDecimalCos[52] = 6156
        this.NoDecimalCos[53] = 6018
        this.NoDecimalCos[54] = 5877
        this.NoDecimalCos[55] = 5735
        this.NoDecimalCos[56] = 5591
        this.NoDecimalCos[57] = 5446
        this.NoDecimalCos[58] = 5299
        this.NoDecimalCos[59] = 5150
        this.NoDecimalCos[60] = 5000
        this.NoDecimalCos[61] = 4848
        this.NoDecimalCos[62] = 4694
        this.NoDecimalCos[63] = 4539
        this.NoDecimalCos[64] = 4383
        this.NoDecimalCos[65] = 4226
        this.NoDecimalCos[66] = 4067
        this.NoDecimalCos[67] = 3907
        this.NoDecimalCos[68] = 3746
        this.NoDecimalCos[69] = 3583
        this.NoDecimalCos[70] = 3420
        this.NoDecimalCos[71] = 3255
        this.NoDecimalCos[72] = 3090
        this.NoDecimalCos[73] = 2923
        this.NoDecimalCos[74] = 2756
        this.NoDecimalCos[75] = 2588
        this.NoDecimalCos[76] = 2419
        this.NoDecimalCos[77] = 2249
        this.NoDecimalCos[78] = 2079
        this.NoDecimalCos[79] = 1908
        this.NoDecimalCos[80] = 1736
        this.NoDecimalCos[81] = 1564
        this.NoDecimalCos[82] = 1391
        this.NoDecimalCos[83] = 1218
        this.NoDecimalCos[84] = 1045
        this.NoDecimalCos[85] = 871
        this.NoDecimalCos[86] = 697
        this.NoDecimalCos[87] = 523
        this.NoDecimalCos[88] = 348
        this.NoDecimalCos[89] = 174
        this.NoDecimalCos[90] = 0
        this.NoDecimalCos[91] = int64(-174)
        this.NoDecimalCos[92] = int64(-348)
        this.NoDecimalCos[93] = int64(-523)
        this.NoDecimalCos[94] = int64(-697)
        this.NoDecimalCos[95] = int64(-871)
        this.NoDecimalCos[96] = int64(-1045)
        this.NoDecimalCos[97] = int64(-1218)
        this.NoDecimalCos[98] = int64(-1391)
        this.NoDecimalCos[99] = int64(-1564)
        this.NoDecimalCos[100] = int64(-1736)
        this.NoDecimalCos[101] = int64(-1908)
        this.NoDecimalCos[102] = int64(-2079)
        this.NoDecimalCos[103] = int64(-2249)
        this.NoDecimalCos[104] = int64(-2419)
        this.NoDecimalCos[105] = int64(-2588)
        this.NoDecimalCos[106] = int64(-2756)
        this.NoDecimalCos[107] = int64(-2923)
        this.NoDecimalCos[108] = int64(-3090)
        this.NoDecimalCos[109] = int64(-3255)
        this.NoDecimalCos[110] = int64(-3420)
        this.NoDecimalCos[111] = int64(-3583)
        this.NoDecimalCos[112] = int64(-3746)
        this.NoDecimalCos[113] = int64(-3907)
        this.NoDecimalCos[114] = int64(-4067)
        this.NoDecimalCos[115] = int64(-4226)
        this.NoDecimalCos[116] = int64(-4383)
        this.NoDecimalCos[117] = int64(-4539)
        this.NoDecimalCos[118] = int64(-4694)
        this.NoDecimalCos[119] = int64(-4848)
        this.NoDecimalCos[120] = int64(-4999)
        this.NoDecimalCos[121] = int64(-5150)
        this.NoDecimalCos[122] = int64(-5299)
        this.NoDecimalCos[123] = int64(-5446)
        this.NoDecimalCos[124] = int64(-5591)
        this.NoDecimalCos[125] = int64(-5735)
        this.NoDecimalCos[126] = int64(-5877)
        this.NoDecimalCos[127] = int64(-6018)
        this.NoDecimalCos[128] = int64(-6156)
        this.NoDecimalCos[129] = int64(-6293)
        this.NoDecimalCos[130] = int64(-6427)
        this.NoDecimalCos[131] = int64(-6560)
        this.NoDecimalCos[132] = int64(-6691)
        this.NoDecimalCos[133] = int64(-6819)
        this.NoDecimalCos[134] = int64(-6946)
        this.NoDecimalCos[135] = int64(-7071)
        this.NoDecimalCos[136] = int64(-7193)
        this.NoDecimalCos[137] = int64(-7313)
        this.NoDecimalCos[138] = int64(-7431)
        this.NoDecimalCos[139] = int64(-7547)
        this.NoDecimalCos[140] = int64(-7660)
        this.NoDecimalCos[141] = int64(-7771)
        this.NoDecimalCos[142] = int64(-7880)
        this.NoDecimalCos[143] = int64(-7986)
        this.NoDecimalCos[144] = int64(-8090)
        this.NoDecimalCos[145] = int64(-8191)
        this.NoDecimalCos[146] = int64(-8290)
        this.NoDecimalCos[147] = int64(-8386)
        this.NoDecimalCos[148] = int64(-8480)
        this.NoDecimalCos[149] = int64(-8571)
        this.NoDecimalCos[150] = int64(-8660)
        this.NoDecimalCos[151] = int64(-8746)
        this.NoDecimalCos[152] = int64(-8829)
        this.NoDecimalCos[153] = int64(-8910)
        this.NoDecimalCos[154] = int64(-8987)
        this.NoDecimalCos[155] = int64(-9063)
        this.NoDecimalCos[156] = int64(-9135)
        this.NoDecimalCos[157] = int64(-9205)
        this.NoDecimalCos[158] = int64(-9271)
        this.NoDecimalCos[159] = int64(-9335)
        this.NoDecimalCos[160] = int64(-9396)
        this.NoDecimalCos[161] = int64(-9455)
        this.NoDecimalCos[162] = int64(-9510)
        this.NoDecimalCos[163] = int64(-9563)
        this.NoDecimalCos[164] = int64(-9612)
        this.NoDecimalCos[165] = int64(-9659)
        this.NoDecimalCos[166] = int64(-9702)
        this.NoDecimalCos[167] = int64(-9743)
        this.NoDecimalCos[168] = int64(-9781)
        this.NoDecimalCos[169] = int64(-9816)
        this.NoDecimalCos[170] = int64(-9848)
        this.NoDecimalCos[171] = int64(-9876)
        this.NoDecimalCos[172] = int64(-9902)
        this.NoDecimalCos[173] = int64(-9925)
        this.NoDecimalCos[174] = int64(-9945)
        this.NoDecimalCos[175] = int64(-9961)
        this.NoDecimalCos[176] = int64(-9975)
        this.NoDecimalCos[177] = int64(-9986)
        this.NoDecimalCos[178] = int64(-9993)
        this.NoDecimalCos[179] = int64(-9998)
        this.NoDecimalCos[180] = int64(-10000)
        this.NoDecimalCos[181] = int64(-9998)
        this.NoDecimalCos[182] = int64(-9993)
        this.NoDecimalCos[183] = int64(-9986)
        this.NoDecimalCos[184] = int64(-9975)
        this.NoDecimalCos[185] = int64(-9961)
        this.NoDecimalCos[186] = int64(-9945)
        this.NoDecimalCos[187] = int64(-9925)
        this.NoDecimalCos[188] = int64(-9902)
        this.NoDecimalCos[189] = int64(-9876)
        this.NoDecimalCos[190] = int64(-9848)
        this.NoDecimalCos[191] = int64(-9816)
        this.NoDecimalCos[192] = int64(-9781)
        this.NoDecimalCos[193] = int64(-9743)
        this.NoDecimalCos[194] = int64(-9702)
        this.NoDecimalCos[195] = int64(-9659)
        this.NoDecimalCos[196] = int64(-9612)
        this.NoDecimalCos[197] = int64(-9563)
        this.NoDecimalCos[198] = int64(-9510)
        this.NoDecimalCos[199] = int64(-9455)
        this.NoDecimalCos[200] = int64(-9396)
        this.NoDecimalCos[201] = int64(-9335)
        this.NoDecimalCos[202] = int64(-9271)
        this.NoDecimalCos[203] = int64(-9205)
        this.NoDecimalCos[204] = int64(-9135)
        this.NoDecimalCos[205] = int64(-9063)
        this.NoDecimalCos[206] = int64(-8987)
        this.NoDecimalCos[207] = int64(-8910)
        this.NoDecimalCos[208] = int64(-8829)
        this.NoDecimalCos[209] = int64(-8746)
        this.NoDecimalCos[210] = int64(-8660)
        this.NoDecimalCos[211] = int64(-8571)
        this.NoDecimalCos[212] = int64(-8480)
        this.NoDecimalCos[213] = int64(-8386)
        this.NoDecimalCos[214] = int64(-8290)
        this.NoDecimalCos[215] = int64(-8191)
        this.NoDecimalCos[216] = int64(-8090)
        this.NoDecimalCos[217] = int64(-7986)
        this.NoDecimalCos[218] = int64(-7880)
        this.NoDecimalCos[219] = int64(-7771)
        this.NoDecimalCos[220] = int64(-7660)
        this.NoDecimalCos[221] = int64(-7547)
        this.NoDecimalCos[222] = int64(-7431)
        this.NoDecimalCos[223] = int64(-7313)
        this.NoDecimalCos[224] = int64(-7193)
        this.NoDecimalCos[225] = int64(-7071)
        this.NoDecimalCos[226] = int64(-6946)
        this.NoDecimalCos[227] = int64(-6819)
        this.NoDecimalCos[228] = int64(-6691)
        this.NoDecimalCos[229] = int64(-6560)
        this.NoDecimalCos[230] = int64(-6427)
        this.NoDecimalCos[231] = int64(-6293)
        this.NoDecimalCos[232] = int64(-6156)
        this.NoDecimalCos[233] = int64(-6018)
        this.NoDecimalCos[234] = int64(-5877)
        this.NoDecimalCos[235] = int64(-5735)
        this.NoDecimalCos[236] = int64(-5591)
        this.NoDecimalCos[237] = int64(-5446)
        this.NoDecimalCos[238] = int64(-5299)
        this.NoDecimalCos[239] = int64(-5150)
        this.NoDecimalCos[240] = int64(-5000)
        this.NoDecimalCos[241] = int64(-4848)
        this.NoDecimalCos[242] = int64(-4694)
        this.NoDecimalCos[243] = int64(-4539)
        this.NoDecimalCos[244] = int64(-4383)
        this.NoDecimalCos[245] = int64(-4226)
        this.NoDecimalCos[246] = int64(-4067)
        this.NoDecimalCos[247] = int64(-3907)
        this.NoDecimalCos[248] = int64(-3746)
        this.NoDecimalCos[249] = int64(-3583)
        this.NoDecimalCos[250] = int64(-3420)
        this.NoDecimalCos[251] = int64(-3255)
        this.NoDecimalCos[252] = int64(-3090)
        this.NoDecimalCos[253] = int64(-2923)
        this.NoDecimalCos[254] = int64(-2756)
        this.NoDecimalCos[255] = int64(-2588)
        this.NoDecimalCos[256] = int64(-2419)
        this.NoDecimalCos[257] = int64(-2249)
        this.NoDecimalCos[258] = int64(-2079)
        this.NoDecimalCos[259] = int64(-1908)
        this.NoDecimalCos[260] = int64(-1736)
        this.NoDecimalCos[261] = int64(-1564)
        this.NoDecimalCos[262] = int64(-1391)
        this.NoDecimalCos[263] = int64(-1218)
        this.NoDecimalCos[264] = int64(-1045)
        this.NoDecimalCos[265] = int64(-871)
        this.NoDecimalCos[266] = int64(-697)
        this.NoDecimalCos[267] = int64(-523)
        this.NoDecimalCos[268] = int64(-348)
        this.NoDecimalCos[269] = int64(-174)
        this.NoDecimalCos[270] = 0
        this.NoDecimalCos[271] = 174
        this.NoDecimalCos[272] = 348
        this.NoDecimalCos[273] = 523
        this.NoDecimalCos[274] = 697
        this.NoDecimalCos[275] = 871
        this.NoDecimalCos[276] = 1045
        this.NoDecimalCos[277] = 1218
        this.NoDecimalCos[278] = 1391
        this.NoDecimalCos[279] = 1564
        this.NoDecimalCos[280] = 1736
        this.NoDecimalCos[281] = 1908
        this.NoDecimalCos[282] = 2079
        this.NoDecimalCos[283] = 2249
        this.NoDecimalCos[284] = 2419
        this.NoDecimalCos[285] = 2588
        this.NoDecimalCos[286] = 2756
        this.NoDecimalCos[287] = 2923
        this.NoDecimalCos[288] = 3090
        this.NoDecimalCos[289] = 3255
        this.NoDecimalCos[290] = 3420
        this.NoDecimalCos[291] = 3583
        this.NoDecimalCos[292] = 3746
        this.NoDecimalCos[293] = 3907
        this.NoDecimalCos[294] = 4067
        this.NoDecimalCos[295] = 4226
        this.NoDecimalCos[296] = 4383
        this.NoDecimalCos[297] = 4539
        this.NoDecimalCos[298] = 4694
        this.NoDecimalCos[299] = 4848
        this.NoDecimalCos[300] = 5000
        this.NoDecimalCos[301] = 5150
        this.NoDecimalCos[302] = 5299
        this.NoDecimalCos[303] = 5446
        this.NoDecimalCos[304] = 5591
        this.NoDecimalCos[305] = 5735
        this.NoDecimalCos[306] = 5877
        this.NoDecimalCos[307] = 6018
        this.NoDecimalCos[308] = 6156
        this.NoDecimalCos[309] = 6293
        this.NoDecimalCos[310] = 6427
        this.NoDecimalCos[311] = 6560
        this.NoDecimalCos[312] = 6691
        this.NoDecimalCos[313] = 6819
        this.NoDecimalCos[314] = 6946
        this.NoDecimalCos[315] = 7071
        this.NoDecimalCos[316] = 7193
        this.NoDecimalCos[317] = 7313
        this.NoDecimalCos[318] = 7431
        this.NoDecimalCos[319] = 7547
        this.NoDecimalCos[320] = 7660
        this.NoDecimalCos[321] = 7771
        this.NoDecimalCos[322] = 7880
        this.NoDecimalCos[323] = 7986
        this.NoDecimalCos[324] = 8090
        this.NoDecimalCos[325] = 8191
        this.NoDecimalCos[326] = 8290
        this.NoDecimalCos[327] = 8386
        this.NoDecimalCos[328] = 8480
        this.NoDecimalCos[329] = 8571
        this.NoDecimalCos[330] = 8660
        this.NoDecimalCos[331] = 8746
        this.NoDecimalCos[332] = 8829
        this.NoDecimalCos[333] = 8910
        this.NoDecimalCos[334] = 8987
        this.NoDecimalCos[335] = 9063
        this.NoDecimalCos[336] = 9135
        this.NoDecimalCos[337] = 9205
        this.NoDecimalCos[338] = 9271
        this.NoDecimalCos[339] = 9335
        this.NoDecimalCos[340] = 9396
        this.NoDecimalCos[341] = 9455
        this.NoDecimalCos[342] = 9510
        this.NoDecimalCos[343] = 9563
        this.NoDecimalCos[344] = 9612
        this.NoDecimalCos[345] = 9659
        this.NoDecimalCos[346] = 9702
        this.NoDecimalCos[347] = 9743
        this.NoDecimalCos[348] = 9781
        this.NoDecimalCos[349] = 9816
        this.NoDecimalCos[350] = 9848
        this.NoDecimalCos[351] = 9876
        this.NoDecimalCos[352] = 9902
        this.NoDecimalCos[353] = 9925
        this.NoDecimalCos[354] = 9945
        this.NoDecimalCos[355] = 9961
        this.NoDecimalCos[356] = 9975
        this.NoDecimalCos[357] = 9986
        this.NoDecimalCos[358] = 9993
        this.NoDecimalCos[359] = 9998
        this.NoDecimalTan[0] = 0
        this.NoDecimalTan[1] = 174
        this.NoDecimalTan[2] = 349
        this.NoDecimalTan[3] = 524
        this.NoDecimalTan[4] = 699
        this.NoDecimalTan[5] = 874
        this.NoDecimalTan[6] = 1051
        this.NoDecimalTan[7] = 1227
        this.NoDecimalTan[8] = 1405
        this.NoDecimalTan[9] = 1583
        this.NoDecimalTan[10] = 1763
        this.NoDecimalTan[11] = 1943
        this.NoDecimalTan[12] = 2125
        this.NoDecimalTan[13] = 2308
        this.NoDecimalTan[14] = 2493
        this.NoDecimalTan[15] = 2679
        this.NoDecimalTan[16] = 2867
        this.NoDecimalTan[17] = 3057
        this.NoDecimalTan[18] = 3249
        this.NoDecimalTan[19] = 3443
        this.NoDecimalTan[20] = 3639
        this.NoDecimalTan[21] = 3838
        this.NoDecimalTan[22] = 4040
        this.NoDecimalTan[23] = 4244
        this.NoDecimalTan[24] = 4452
        this.NoDecimalTan[25] = 4663
        this.NoDecimalTan[26] = 4877
        this.NoDecimalTan[27] = 5095
        this.NoDecimalTan[28] = 5317
        this.NoDecimalTan[29] = 5543
        this.NoDecimalTan[30] = 5773
        this.NoDecimalTan[31] = 6008
        this.NoDecimalTan[32] = 6248
        this.NoDecimalTan[33] = 6494
        this.NoDecimalTan[34] = 6745
        this.NoDecimalTan[35] = 7002
        this.NoDecimalTan[36] = 7265
        this.NoDecimalTan[37] = 7535
        this.NoDecimalTan[38] = 7812
        this.NoDecimalTan[39] = 8097
        this.NoDecimalTan[40] = 8390
        this.NoDecimalTan[41] = 8692
        this.NoDecimalTan[42] = 9004
        this.NoDecimalTan[43] = 9325
        this.NoDecimalTan[44] = 9656
        this.NoDecimalTan[45] = 9999
        this.NoDecimalTan[46] = 10355
        this.NoDecimalTan[47] = 10723
        this.NoDecimalTan[48] = 11106
        this.NoDecimalTan[49] = 11503
        this.NoDecimalTan[50] = 11917
        this.NoDecimalTan[51] = 12348
        this.NoDecimalTan[52] = 12799
        this.NoDecimalTan[53] = 13270
        this.NoDecimalTan[54] = 13763
        this.NoDecimalTan[55] = 14281
        this.NoDecimalTan[56] = 14825
        this.NoDecimalTan[57] = 15398
        this.NoDecimalTan[58] = 16003
        this.NoDecimalTan[59] = 16642
        this.NoDecimalTan[60] = 17320
        this.NoDecimalTan[61] = 18040
        this.NoDecimalTan[62] = 18807
        this.NoDecimalTan[63] = 19626
        this.NoDecimalTan[64] = 20503
        this.NoDecimalTan[65] = 21445
        this.NoDecimalTan[66] = 22460
        this.NoDecimalTan[67] = 23558
        this.NoDecimalTan[68] = 24750
        this.NoDecimalTan[69] = 26050
        this.NoDecimalTan[70] = 27474
        this.NoDecimalTan[71] = 29042
        this.NoDecimalTan[72] = 30776
        this.NoDecimalTan[73] = 32708
        this.NoDecimalTan[74] = 34874
        this.NoDecimalTan[75] = 37320
        this.NoDecimalTan[76] = 40107
        this.NoDecimalTan[77] = 43314
        this.NoDecimalTan[78] = 47046
        this.NoDecimalTan[79] = 51445
        this.NoDecimalTan[80] = 56712
        this.NoDecimalTan[81] = 63137
        this.NoDecimalTan[82] = 71153
        this.NoDecimalTan[83] = 81443
        this.NoDecimalTan[84] = 95143
        this.NoDecimalTan[85] = 114300
        this.NoDecimalTan[86] = 143006
        this.NoDecimalTan[87] = 190811
        this.NoDecimalTan[88] = 286362
        this.NoDecimalTan[89] = 572899
        this.NoDecimalTan[90] = MAX_VALUE
        this.NoDecimalTan[91] = int64(-572899)
        this.NoDecimalTan[92] = int64(-286362)
        this.NoDecimalTan[93] = int64(-190811)
        this.NoDecimalTan[94] = int64(-143006)
        this.NoDecimalTan[95] = int64(-114300)
        this.NoDecimalTan[96] = int64(-95143)
        this.NoDecimalTan[97] = int64(-81443)
        this.NoDecimalTan[98] = int64(-71153)
        this.NoDecimalTan[99] = int64(-63137)
        this.NoDecimalTan[100] = int64(-56712)
        this.NoDecimalTan[101] = int64(-51445)
        this.NoDecimalTan[102] = int64(-47046)
        this.NoDecimalTan[103] = int64(-43314)
        this.NoDecimalTan[104] = int64(-40107)
        this.NoDecimalTan[105] = int64(-37320)
        this.NoDecimalTan[106] = int64(-34874)
        this.NoDecimalTan[107] = int64(-32708)
        this.NoDecimalTan[108] = int64(-30776)
        this.NoDecimalTan[109] = int64(-29042)
        this.NoDecimalTan[110] = int64(-27474)
        this.NoDecimalTan[111] = int64(-26050)
        this.NoDecimalTan[112] = int64(-24750)
        this.NoDecimalTan[113] = int64(-23558)
        this.NoDecimalTan[114] = int64(-22460)
        this.NoDecimalTan[115] = int64(-21445)
        this.NoDecimalTan[116] = int64(-20503)
        this.NoDecimalTan[117] = int64(-19626)
        this.NoDecimalTan[118] = int64(-18807)
        this.NoDecimalTan[119] = int64(-18040)
        this.NoDecimalTan[120] = int64(-17320)
        this.NoDecimalTan[121] = int64(-16642)
        this.NoDecimalTan[122] = int64(-16003)
        this.NoDecimalTan[123] = int64(-15398)
        this.NoDecimalTan[124] = int64(-14825)
        this.NoDecimalTan[125] = int64(-14281)
        this.NoDecimalTan[126] = int64(-13763)
        this.NoDecimalTan[127] = int64(-13270)
        this.NoDecimalTan[128] = int64(-12799)
        this.NoDecimalTan[129] = int64(-12348)
        this.NoDecimalTan[130] = int64(-11917)
        this.NoDecimalTan[131] = int64(-11503)
        this.NoDecimalTan[132] = int64(-11106)
        this.NoDecimalTan[133] = int64(-10723)
        this.NoDecimalTan[134] = int64(-10355)
        this.NoDecimalTan[135] = int64(-10000)
        this.NoDecimalTan[136] = int64(-9656)
        this.NoDecimalTan[137] = int64(-9325)
        this.NoDecimalTan[138] = int64(-9004)
        this.NoDecimalTan[139] = int64(-8692)
        this.NoDecimalTan[140] = int64(-8390)
        this.NoDecimalTan[141] = int64(-8097)
        this.NoDecimalTan[142] = int64(-7812)
        this.NoDecimalTan[143] = int64(-7535)
        this.NoDecimalTan[144] = int64(-7265)
        this.NoDecimalTan[145] = int64(-7002)
        this.NoDecimalTan[146] = int64(-6745)
        this.NoDecimalTan[147] = int64(-6494)
        this.NoDecimalTan[148] = int64(-6248)
        this.NoDecimalTan[149] = int64(-6008)
        this.NoDecimalTan[150] = int64(-5773)
        this.NoDecimalTan[151] = int64(-5543)
        this.NoDecimalTan[152] = int64(-5317)
        this.NoDecimalTan[153] = int64(-5095)
        this.NoDecimalTan[154] = int64(-4877)
        this.NoDecimalTan[155] = int64(-4663)
        this.NoDecimalTan[156] = int64(-4452)
        this.NoDecimalTan[157] = int64(-4244)
        this.NoDecimalTan[158] = int64(-4040)
        this.NoDecimalTan[159] = int64(-3838)
        this.NoDecimalTan[160] = int64(-3639)
        this.NoDecimalTan[161] = int64(-3443)
        this.NoDecimalTan[162] = int64(-3249)
        this.NoDecimalTan[163] = int64(-3057)
        this.NoDecimalTan[164] = int64(-2867)
        this.NoDecimalTan[165] = int64(-2679)
        this.NoDecimalTan[166] = int64(-2493)
        this.NoDecimalTan[167] = int64(-2308)
        this.NoDecimalTan[168] = int64(-2125)
        this.NoDecimalTan[169] = int64(-1943)
        this.NoDecimalTan[170] = int64(-1763)
        this.NoDecimalTan[171] = int64(-1583)
        this.NoDecimalTan[172] = int64(-1405)
        this.NoDecimalTan[173] = int64(-1227)
        this.NoDecimalTan[174] = int64(-1051)
        this.NoDecimalTan[175] = int64(-874)
        this.NoDecimalTan[176] = int64(-699)
        this.NoDecimalTan[177] = int64(-524)
        this.NoDecimalTan[178] = int64(-349)
        this.NoDecimalTan[179] = int64(-174)
        this.NoDecimalTan[180] = 0
        this.NoDecimalTan[181] = 174
        this.NoDecimalTan[182] = 349
        this.NoDecimalTan[183] = 524
        this.NoDecimalTan[184] = 699
        this.NoDecimalTan[185] = 874
        this.NoDecimalTan[186] = 1051
        this.NoDecimalTan[187] = 1227
        this.NoDecimalTan[188] = 1405
        this.NoDecimalTan[189] = 1583
        this.NoDecimalTan[190] = 1763
        this.NoDecimalTan[191] = 1943
        this.NoDecimalTan[192] = 2125
        this.NoDecimalTan[193] = 2308
        this.NoDecimalTan[194] = 2493
        this.NoDecimalTan[195] = 2679
        this.NoDecimalTan[196] = 2867
        this.NoDecimalTan[197] = 3057
        this.NoDecimalTan[198] = 3249
        this.NoDecimalTan[199] = 3443
        this.NoDecimalTan[200] = 3639
        this.NoDecimalTan[201] = 3838
        this.NoDecimalTan[202] = 4040
        this.NoDecimalTan[203] = 4244
        this.NoDecimalTan[204] = 4452
        this.NoDecimalTan[205] = 4663
        this.NoDecimalTan[206] = 4877
        this.NoDecimalTan[207] = 5095
        this.NoDecimalTan[208] = 5317
        this.NoDecimalTan[209] = 5543
        this.NoDecimalTan[210] = 5773
        this.NoDecimalTan[211] = 6008
        this.NoDecimalTan[212] = 6248
        this.NoDecimalTan[213] = 6494
        this.NoDecimalTan[214] = 6745
        this.NoDecimalTan[215] = 7002
        this.NoDecimalTan[216] = 7265
        this.NoDecimalTan[217] = 7535
        this.NoDecimalTan[218] = 7812
        this.NoDecimalTan[219] = 8097
        this.NoDecimalTan[220] = 8390
        this.NoDecimalTan[221] = 8692
        this.NoDecimalTan[222] = 9004
        this.NoDecimalTan[223] = 9325
        this.NoDecimalTan[224] = 9656
        this.NoDecimalTan[225] = 9999
        this.NoDecimalTan[226] = 10355
        this.NoDecimalTan[227] = 10723
        this.NoDecimalTan[228] = 11106
        this.NoDecimalTan[229] = 11503
        this.NoDecimalTan[230] = 11917
        this.NoDecimalTan[231] = 12348
        this.NoDecimalTan[232] = 12799
        this.NoDecimalTan[233] = 13270
        this.NoDecimalTan[234] = 13763
        this.NoDecimalTan[235] = 14281
        this.NoDecimalTan[236] = 14825
        this.NoDecimalTan[237] = 15398
        this.NoDecimalTan[238] = 16003
        this.NoDecimalTan[239] = 16642
        this.NoDecimalTan[240] = 17320
        this.NoDecimalTan[241] = 18040
        this.NoDecimalTan[242] = 18807
        this.NoDecimalTan[243] = 19626
        this.NoDecimalTan[244] = 20503
        this.NoDecimalTan[245] = 21445
        this.NoDecimalTan[246] = 22460
        this.NoDecimalTan[247] = 23558
        this.NoDecimalTan[248] = 24750
        this.NoDecimalTan[249] = 26050
        this.NoDecimalTan[250] = 27474
        this.NoDecimalTan[251] = 29042
        this.NoDecimalTan[252] = 30776
        this.NoDecimalTan[253] = 32708
        this.NoDecimalTan[254] = 34874
        this.NoDecimalTan[255] = 37320
        this.NoDecimalTan[256] = 40107
        this.NoDecimalTan[257] = 43314
        this.NoDecimalTan[258] = 47046
        this.NoDecimalTan[259] = 51445
        this.NoDecimalTan[260] = 56712
        this.NoDecimalTan[261] = 63137
        this.NoDecimalTan[262] = 71153
        this.NoDecimalTan[263] = 81443
        this.NoDecimalTan[264] = 95143
        this.NoDecimalTan[265] = 114300
        this.NoDecimalTan[266] = 143006
        this.NoDecimalTan[267] = 190811
        this.NoDecimalTan[268] = 286362
        this.NoDecimalTan[269] = 572899
        this.NoDecimalTan[270] = MAX_VALUE
        this.NoDecimalTan[271] = int64(-572899)
        this.NoDecimalTan[272] = int64(-286362)
        this.NoDecimalTan[273] = int64(-190811)
        this.NoDecimalTan[274] = int64(-143006)
        this.NoDecimalTan[275] = int64(-114300)
        this.NoDecimalTan[276] = int64(-95143)
        this.NoDecimalTan[277] = int64(-81443)
        this.NoDecimalTan[278] = int64(-71153)
        this.NoDecimalTan[279] = int64(-63137)
        this.NoDecimalTan[280] = int64(-56712)
        this.NoDecimalTan[281] = int64(-51445)
        this.NoDecimalTan[282] = int64(-47046)
        this.NoDecimalTan[283] = int64(-43314)
        this.NoDecimalTan[284] = int64(-40107)
        this.NoDecimalTan[285] = int64(-37320)
        this.NoDecimalTan[286] = int64(-34874)
        this.NoDecimalTan[287] = int64(-32708)
        this.NoDecimalTan[288] = int64(-30776)
        this.NoDecimalTan[289] = int64(-29042)
        this.NoDecimalTan[290] = int64(-27474)
        this.NoDecimalTan[291] = int64(-26050)
        this.NoDecimalTan[292] = int64(-24750)
        this.NoDecimalTan[293] = int64(-23558)
        this.NoDecimalTan[294] = int64(-22460)
        this.NoDecimalTan[295] = int64(-21445)
        this.NoDecimalTan[296] = int64(-20503)
        this.NoDecimalTan[297] = int64(-19626)
        this.NoDecimalTan[298] = int64(-18807)
        this.NoDecimalTan[299] = int64(-18040)
        this.NoDecimalTan[300] = int64(-17320)
        this.NoDecimalTan[301] = int64(-16642)
        this.NoDecimalTan[302] = int64(-16003)
        this.NoDecimalTan[303] = int64(-15398)
        this.NoDecimalTan[304] = int64(-14825)
        this.NoDecimalTan[305] = int64(-14281)
        this.NoDecimalTan[306] = int64(-13763)
        this.NoDecimalTan[307] = int64(-13270)
        this.NoDecimalTan[308] = int64(-12799)
        this.NoDecimalTan[309] = int64(-12348)
        this.NoDecimalTan[310] = int64(-11917)
        this.NoDecimalTan[311] = int64(-11503)
        this.NoDecimalTan[312] = int64(-11106)
        this.NoDecimalTan[313] = int64(-10723)
        this.NoDecimalTan[314] = int64(-10355)
        this.NoDecimalTan[315] = int64(-10000)
        this.NoDecimalTan[316] = int64(-9656)
        this.NoDecimalTan[317] = int64(-9325)
        this.NoDecimalTan[318] = int64(-9004)
        this.NoDecimalTan[319] = int64(-8692)
        this.NoDecimalTan[320] = int64(-8390)
        this.NoDecimalTan[321] = int64(-8097)
        this.NoDecimalTan[322] = int64(-7812)
        this.NoDecimalTan[323] = int64(-7535)
        this.NoDecimalTan[324] = int64(-7265)
        this.NoDecimalTan[325] = int64(-7002)
        this.NoDecimalTan[326] = int64(-6745)
        this.NoDecimalTan[327] = int64(-6494)
        this.NoDecimalTan[328] = int64(-6248)
        this.NoDecimalTan[329] = int64(-6008)
        this.NoDecimalTan[330] = int64(-5773)
        this.NoDecimalTan[331] = int64(-5543)
        this.NoDecimalTan[332] = int64(-5317)
        this.NoDecimalTan[333] = int64(-5095)
        this.NoDecimalTan[334] = int64(-4877)
        this.NoDecimalTan[335] = int64(-4663)
        this.NoDecimalTan[336] = int64(-4452)
        this.NoDecimalTan[337] = int64(-4244)
        this.NoDecimalTan[338] = int64(-4040)
        this.NoDecimalTan[339] = int64(-3838)
        this.NoDecimalTan[340] = int64(-3639)
        this.NoDecimalTan[341] = int64(-3443)
        this.NoDecimalTan[342] = int64(-3249)
        this.NoDecimalTan[343] = int64(-3057)
        this.NoDecimalTan[344] = int64(-2867)
        this.NoDecimalTan[345] = int64(-2679)
        this.NoDecimalTan[346] = int64(-2493)
        this.NoDecimalTan[347] = int64(-2308)
        this.NoDecimalTan[348] = int64(-2125)
        this.NoDecimalTan[349] = int64(-1943)
        this.NoDecimalTan[350] = int64(-1763)
        this.NoDecimalTan[351] = int64(-1583)
        this.NoDecimalTan[352] = int64(-1405)
        this.NoDecimalTan[353] = int64(-1227)
        this.NoDecimalTan[354] = int64(-1051)
        this.NoDecimalTan[355] = int64(-874)
        this.NoDecimalTan[356] = int64(-699)
        this.NoDecimalTan[357] = int64(-524)
        this.NoDecimalTan[358] = int64(-349)
        this.NoDecimalTan[359] = int64(-174)
        this.NoDecimalTan[0] = int64(-174)
    return this
}


func (this *NoDecimalTrigTable) Sin(angle int) int64 {
        return this.NoDecimalSin[angle]
}


func (this *NoDecimalTrigTable) Cos(angle int) int64 {
        return this.NoDecimalCos[angle]
}


func (this *NoDecimalTrigTable) Tan(angle int) int64 {
        return this.NoDecimalTan[angle]
}


// throws Exception
func (this *NoDecimalTrigTable) AntiTan(dx int, dy int) int16 {
        var MAX_VALUE int64 = (int64(Integer_MAX_VALUE)); _ = MAX_VALUE
        var MIN_VALUE int64 = (int64(Integer_MIN_VALUE)); _ = MIN_VALUE
        var ratio int64 = (MAX_VALUE - 1); _ = ratio
        if dy != 0 {
            var dxl int64 = (int64(dx)); _ = dxl
            ratio = int64(this.SCALE) * dxl / int64(dy)
            if dx <= 0 && dy < 0 {
                ratio = this.MathUtil.Abslong(ratio)
            }
        }
        if (dx >= 0 && dy < 0) {
            for index := int(180); index > 90; index-- {
                if ratio <= this.NoDecimalTan[index] {
                    if ratio > this.NoDecimalTan[index - 1] {
                        return (int16(index))
                    } else if this.NoDecimalTan[index - 1] == MAX_VALUE && ratio > MIN_VALUE {
                        return (int16(index))
                    }
                }
            }
        } else if (dx < 0 && dy <= 0) {
            for index := int(270); index > 180; index-- {
                if ratio < this.NoDecimalTan[index] && ratio >= this.NoDecimalTan[index - 1] {
                    return (int16(index))
                }
            }
        } else if (dx < 0 && dy > 0) {
            for index := int((len(this.NoDecimalTan) - 1)); index > 270; index-- {
                if ratio < this.NoDecimalTan[index] {
                    if ratio >= this.NoDecimalTan[index - 1] {
                        return (int16(index))
                    } else if this.NoDecimalTan[index - 1] == MAX_VALUE && ratio >= MIN_VALUE {
                        return (int16(index))
                    }
                }
            }
        } else if (dx >= 0 && dy >= 0) {
            for index := int(90); index > 0; index-- {
                if ratio < this.NoDecimalTan[index] && ratio >= this.NoDecimalTan[index - 1] {
                    return (int16(index))
                }
            }
        }
        var stringBuffer *StringMaker = NewStringMaker(); _ = stringBuffer
        stringBuffer.Append("Invalid Ratio: ")
        stringBuffer.Appendlong(ratio)
        stringBuffer.Append(CommonSeps_GetInstance().SPACE)
        stringBuffer.Append(PositionStrings_GetInstance().DX_LABEL)
        stringBuffer.Appendint(dx)
        stringBuffer.Append(CommonSeps_GetInstance().SPACE)
        stringBuffer.Append(PositionStrings_GetInstance().DX_LABEL)
        stringBuffer.Appendint(dy)
        panic(errors.New(fmt.Sprint(fmt.Sprint(stringBuffer))))
}


// throws Exception
func (this *NoDecimalTrigTable) AntiTanDebug(screenX int, screenY int, targetX int, targetY int) int16 {
        var METHOD_NAME string = "antiTanDebug"; _ = METHOD_NAME
        var TARGET_ANGLE string = " Targeting angle: "; _ = TARGET_ANGLE
        var stringMaker *StringMaker = NewStringMaker(); _ = stringMaker
        PreLogUtil_Put(fmt.Sprint(stringMaker.Append("screenX: ").Appendint(screenX).Append(" screenY: ").Appendint(screenY).Append(" targetX: ").Appendint(targetX).Append(" targetY: ").Appendint(targetY)), this, METHOD_NAME)
        var targetX2 int = -targetX; _ = targetX2
        var targetY2 int = -targetY; _ = targetY2
        var screenX2 int = -screenX; _ = screenX2
        var screenY2 int = -screenY; _ = screenY2
        var dx []int = make([]int, 4); _ = dx
        dx[0] = (screenX - targetX)
        dx[1] = (screenX - targetX2)
        dx[2] = (screenX2 - targetX)
        dx[3] = (screenX2 - targetX2)
        var dy []int = make([]int, 4); _ = dy
        dy[0] = (screenY - targetY)
        dy[1] = (screenY - targetY2)
        dy[2] = (screenY2 - targetY)
        dy[3] = (screenY2 - targetY2)
        var angleOfTarget int16; _ = angleOfTarget
        var angleFactory *AngleFactory = AngleFactory_GetInstance(); _ = angleFactory
        var angle IAngle; _ = angle
        for index := int(0); index < len(dx); index++ {
            for index2 := int(0); index2 < len(dy); index2++ {
                angleOfTarget = this.AntiTan(dx[index], dy[index2])
                angle = angleFactory.GetAt(int(angleOfTarget))
                stringMaker.Delete(0, stringMaker.Length())
                PreLogUtil_Put(fmt.Sprint(stringMaker.Append(CommonLabels_GetInstance().INDEX_LABEL).Appendint(index).Append(" index2: ").Appendint(index2).Append(PositionStrings_GetInstance().DX_LABEL).Appendint(dx[index]).Append(PositionStrings_GetInstance().DY_LABEL).Appendint(dy[index2]).Append(TARGET_ANGLE).Appendshort(angleOfTarget).Append(CommonSeps_GetInstance().EQUALS).Appendshort(angle.GetValue())), this, METHOD_NAME)
            }
        }
        angleOfTarget = this.AntiTan(dx[0], dy[0])
        return angleOfTarget
}

