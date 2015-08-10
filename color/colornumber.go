package color

type ColorNumber uint8

// 1 - 9
const (
	Red ColorNumber = iota + 1
	Yellow
	Green
	Cyan
	Blue
	Magenta
	White
	Grey128
	Grey192
)

// 250 - 255
const (
	Grey51 = iota + 250
	Grey91
	Grey132
	Grey173
	Grey214
	Grey255
)

var (
	ColorRGB = [][]uint8{
		[]uint8{0, 0, 0},       // 0
		[]uint8{255, 0, 0},     // 1
		[]uint8{255, 255, 0},   // 2
		[]uint8{0, 255, 0},     // 3
		[]uint8{0, 255, 255},   // 4
		[]uint8{0, 0, 255},     // 5
		[]uint8{255, 0, 255},   // 6
		[]uint8{255, 255, 255}, // 7
		[]uint8{128, 128, 128}, // 8
		[]uint8{192, 192, 192}, // 9
		[]uint8{255, 0, 0},     // 10
		[]uint8{255, 127, 127}, // 11
		[]uint8{204, 0, 0},     // 12
		[]uint8{204, 102, 102}, // 13
		[]uint8{153, 0, 0},     // 14
		[]uint8{153, 76, 76},   // 15
		[]uint8{127, 0, 0},     // 16
		[]uint8{127, 63, 63},   // 17
		[]uint8{76, 0, 0},      // 18
		[]uint8{76, 38, 38},    // 19
		[]uint8{255, 63, 0},    // 20
		[]uint8{255, 159, 127}, // 21
		[]uint8{204, 51, 0},    // 22
		[]uint8{204, 127, 102}, // 23
		[]uint8{153, 38, 0},    // 24
		[]uint8{153, 95, 76},   // 25
		[]uint8{127, 31, 0},    // 26
		[]uint8{127, 79, 63},   // 27
		[]uint8{76, 19, 0},     // 28
		[]uint8{76, 47, 38},    // 29
		[]uint8{255, 127, 0},   // 30
		[]uint8{255, 191, 127}, // 31
		[]uint8{204, 102, 0},   // 32
		[]uint8{204, 153, 102}, // 33
		[]uint8{153, 76, 0},    // 34
		[]uint8{153, 114, 76},  // 35
		[]uint8{127, 63, 0},    // 36
		[]uint8{127, 95, 63},   // 37
		[]uint8{76, 38, 0},     // 38
		[]uint8{76, 57, 38},    // 39
		[]uint8{255, 191, 0},   // 40
		[]uint8{255, 223, 127}, // 41
		[]uint8{204, 153, 0},   // 42
		[]uint8{204, 178, 102}, // 43
		[]uint8{153, 114, 0},   // 44
		[]uint8{153, 133, 76},  // 45
		[]uint8{127, 95, 0},    // 46
		[]uint8{127, 111, 63},  // 47
		[]uint8{76, 57, 0},     // 48
		[]uint8{76, 66, 38},    // 49
		[]uint8{255, 255, 0},   // 50
		[]uint8{255, 255, 127}, // 51
		[]uint8{204, 204, 0},   // 52
		[]uint8{204, 204, 102}, // 53
		[]uint8{153, 153, 0},   // 54
		[]uint8{153, 153, 76},  // 55
		[]uint8{127, 127, 0},   // 56
		[]uint8{127, 127, 63},  // 57
		[]uint8{76, 76, 0},     // 58
		[]uint8{76, 76, 38},    // 59
		[]uint8{191, 255, 0},   // 60
		[]uint8{223, 255, 127}, // 61
		[]uint8{153, 204, 0},   // 62
		[]uint8{178, 204, 102}, // 63
		[]uint8{114, 153, 0},   // 64
		[]uint8{133, 153, 76},  // 65
		[]uint8{95, 127, 0},    // 66
		[]uint8{111, 127, 63},  // 67
		[]uint8{57, 76, 0},     // 68
		[]uint8{66, 76, 38},    // 69
		[]uint8{127, 255, 0},   // 70
		[]uint8{191, 255, 127}, // 71
		[]uint8{102, 204, 0},   // 72
		[]uint8{153, 204, 102}, // 73
		[]uint8{76, 153, 0},    // 74
		[]uint8{114, 153, 76},  // 75
		[]uint8{63, 127, 0},    // 76
		[]uint8{95, 127, 63},   // 77
		[]uint8{38, 76, 0},     // 78
		[]uint8{57, 76, 38},    // 79
		[]uint8{63, 255, 0},    // 80
		[]uint8{159, 255, 127}, // 81
		[]uint8{51, 204, 0},    // 82
		[]uint8{127, 204, 102}, // 83
		[]uint8{38, 153, 0},    // 84
		[]uint8{95, 153, 76},   // 85
		[]uint8{31, 127, 0},    // 86
		[]uint8{79, 127, 63},   // 87
		[]uint8{19, 76, 0},     // 88
		[]uint8{47, 76, 38},    // 89
		[]uint8{0, 255, 0},     // 90
		[]uint8{127, 255, 127}, // 91
		[]uint8{0, 204, 0},     // 92
		[]uint8{102, 204, 102}, // 93
		[]uint8{0, 153, 0},     // 94
		[]uint8{76, 153, 76},   // 95
		[]uint8{0, 127, 0},     // 96
		[]uint8{63, 127, 63},   // 97
		[]uint8{0, 76, 0},      // 98
		[]uint8{38, 76, 38},    // 99
		[]uint8{0, 255, 63},    // 100
		[]uint8{127, 255, 159}, // 101
		[]uint8{0, 204, 51},    // 102
		[]uint8{102, 204, 127}, // 103
		[]uint8{0, 153, 38},    // 104
		[]uint8{76, 153, 95},   // 105
		[]uint8{0, 127, 31},    // 106
		[]uint8{63, 127, 79},   // 107
		[]uint8{0, 76, 19},     // 108
		[]uint8{38, 76, 47},    // 109
		[]uint8{0, 255, 127},   // 110
		[]uint8{127, 255, 191}, // 111
		[]uint8{0, 204, 102},   // 112
		[]uint8{102, 204, 153}, // 113
		[]uint8{0, 153, 76},    // 114
		[]uint8{76, 153, 114},  // 115
		[]uint8{0, 127, 63},    // 116
		[]uint8{63, 127, 95},   // 117
		[]uint8{0, 76, 38},     // 118
		[]uint8{38, 76, 57},    // 119
		[]uint8{0, 255, 191},   // 120
		[]uint8{127, 255, 223}, // 121
		[]uint8{0, 204, 153},   // 122
		[]uint8{102, 204, 178}, // 123
		[]uint8{0, 153, 114},   // 124
		[]uint8{76, 153, 133},  // 125
		[]uint8{0, 127, 95},    // 126
		[]uint8{63, 127, 111},  // 127
		[]uint8{0, 76, 57},     // 128
		[]uint8{38, 76, 66},    // 129
		[]uint8{0, 255, 255},   // 130
		[]uint8{127, 255, 255}, // 131
		[]uint8{0, 204, 204},   // 132
		[]uint8{102, 204, 204}, // 133
		[]uint8{0, 153, 153},   // 134
		[]uint8{76, 153, 153},  // 135
		[]uint8{0, 127, 127},   // 136
		[]uint8{63, 127, 127},  // 137
		[]uint8{0, 76, 76},     // 138
		[]uint8{38, 76, 76},    // 139
		[]uint8{0, 191, 255},   // 140
		[]uint8{127, 223, 255}, // 141
		[]uint8{0, 153, 204},   // 142
		[]uint8{102, 178, 204}, // 143
		[]uint8{0, 114, 153},   // 144
		[]uint8{76, 133, 153},  // 145
		[]uint8{0, 95, 127},    // 146
		[]uint8{63, 111, 127},  // 147
		[]uint8{0, 57, 76},     // 148
		[]uint8{38, 66, 76},    // 149
		[]uint8{0, 127, 255},   // 150
		[]uint8{127, 191, 255}, // 151
		[]uint8{0, 102, 204},   // 152
		[]uint8{102, 153, 204}, // 153
		[]uint8{0, 76, 153},    // 154
		[]uint8{76, 114, 153},  // 155
		[]uint8{0, 63, 127},    // 156
		[]uint8{63, 95, 127},   // 157
		[]uint8{0, 38, 76},     // 158
		[]uint8{38, 57, 76},    // 159
		[]uint8{0, 63, 255},    // 160
		[]uint8{127, 159, 255}, // 161
		[]uint8{0, 51, 204},    // 162
		[]uint8{102, 127, 204}, // 163
		[]uint8{0, 38, 153},    // 164
		[]uint8{76, 95, 153},   // 165
		[]uint8{0, 31, 127},    // 166
		[]uint8{63, 79, 127},   // 167
		[]uint8{0, 19, 76},     // 168
		[]uint8{38, 47, 76},    // 169
		[]uint8{0, 0, 255},     // 170
		[]uint8{127, 127, 255}, // 171
		[]uint8{0, 0, 204},     // 172
		[]uint8{102, 102, 204}, // 173
		[]uint8{0, 0, 153},     // 174
		[]uint8{76, 76, 153},   // 175
		[]uint8{0, 0, 127},     // 176
		[]uint8{63, 63, 127},   // 177
		[]uint8{0, 0, 76},      // 178
		[]uint8{38, 38, 76},    // 179
		[]uint8{63, 0, 255},    // 180
		[]uint8{159, 127, 255}, // 181
		[]uint8{51, 0, 204},    // 182
		[]uint8{127, 102, 204}, // 183
		[]uint8{38, 0, 153},    // 184
		[]uint8{95, 76, 153},   // 185
		[]uint8{31, 0, 127},    // 186
		[]uint8{79, 63, 127},   // 187
		[]uint8{19, 0, 76},     // 188
		[]uint8{47, 38, 76},    // 189
		[]uint8{127, 0, 255},   // 190
		[]uint8{191, 127, 255}, // 191
		[]uint8{102, 0, 204},   // 192
		[]uint8{153, 102, 204}, // 193
		[]uint8{76, 0, 153},    // 194
		[]uint8{114, 76, 153},  // 195
		[]uint8{63, 0, 127},    // 196
		[]uint8{95, 63, 127},   // 197
		[]uint8{38, 0, 76},     // 198
		[]uint8{57, 38, 76},    // 199
		[]uint8{191, 0, 255},   // 200
		[]uint8{223, 127, 255}, // 201
		[]uint8{153, 0, 204},   // 202
		[]uint8{178, 102, 204}, // 203
		[]uint8{114, 0, 153},   // 204
		[]uint8{133, 76, 153},  // 205
		[]uint8{95, 0, 127},    // 206
		[]uint8{111, 63, 127},  // 207
		[]uint8{57, 0, 76},     // 208
		[]uint8{66, 38, 76},    // 209
		[]uint8{255, 0, 255},   // 210
		[]uint8{255, 127, 255}, // 211
		[]uint8{204, 0, 204},   // 212
		[]uint8{204, 102, 204}, // 213
		[]uint8{153, 0, 153},   // 214
		[]uint8{153, 76, 153},  // 215
		[]uint8{127, 0, 127},   // 216
		[]uint8{127, 63, 127},  // 217
		[]uint8{76, 0, 76},     // 218
		[]uint8{76, 38, 76},    // 219
		[]uint8{255, 0, 191},   // 220
		[]uint8{255, 127, 223}, // 221
		[]uint8{204, 0, 153},   // 222
		[]uint8{204, 102, 178}, // 223
		[]uint8{153, 0, 114},   // 224
		[]uint8{153, 76, 133},  // 225
		[]uint8{127, 0, 95},    // 226
		[]uint8{127, 63, 111},  // 227
		[]uint8{76, 0, 57},     // 228
		[]uint8{76, 38, 66},    // 229
		[]uint8{255, 0, 127},   // 230
		[]uint8{255, 127, 191}, // 231
		[]uint8{204, 0, 102},   // 232
		[]uint8{204, 102, 153}, // 233
		[]uint8{153, 0, 76},    // 234
		[]uint8{153, 76, 114},  // 235
		[]uint8{127, 0, 63},    // 236
		[]uint8{127, 63, 95},   // 237
		[]uint8{76, 0, 38},     // 238
		[]uint8{76, 38, 57},    // 239
		[]uint8{255, 0, 63},    // 240
		[]uint8{255, 127, 159}, // 241
		[]uint8{204, 0, 51},    // 242
		[]uint8{204, 102, 127}, // 243
		[]uint8{153, 0, 38},    // 244
		[]uint8{153, 76, 95},   // 245
		[]uint8{127, 0, 31},    // 246
		[]uint8{127, 63, 79},   // 247
		[]uint8{76, 0, 19},     // 248
		[]uint8{76, 38, 47},    // 249
		[]uint8{51, 51, 51},    // 250
		[]uint8{91, 91, 91},    // 251
		[]uint8{132, 132, 132}, // 252
		[]uint8{173, 173, 173}, // 253
		[]uint8{214, 214, 214}, // 254
		[]uint8{255, 255, 255}, // 255
	}
)
