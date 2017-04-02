// This file is auto-generated using the
// program found in 'cmd/fitgen/main.go'
// DO NOT EDIT.
// SDK Version: 16.20

package fit

import (
	"reflect"

	"github.com/tormoder/fit/internal/types"
)

var knownMsgNums = map[MesgNum]bool{
	MesgNumFileId:                  true,
	MesgNumCapabilities:            true,
	MesgNumDeviceSettings:          true,
	MesgNumUserProfile:             true,
	MesgNumHrmProfile:              true,
	MesgNumSdmProfile:              true,
	MesgNumBikeProfile:             true,
	MesgNumZonesTarget:             true,
	MesgNumHrZone:                  true,
	MesgNumPowerZone:               true,
	MesgNumMetZone:                 true,
	MesgNumSport:                   true,
	MesgNumGoal:                    true,
	MesgNumSession:                 true,
	MesgNumLap:                     true,
	MesgNumRecord:                  true,
	MesgNumEvent:                   true,
	MesgNumDeviceInfo:              true,
	MesgNumWorkout:                 true,
	MesgNumWorkoutStep:             true,
	MesgNumSchedule:                true,
	MesgNumWeightScale:             true,
	MesgNumCourse:                  true,
	MesgNumCoursePoint:             true,
	MesgNumTotals:                  true,
	MesgNumActivity:                true,
	MesgNumSoftware:                true,
	MesgNumFileCapabilities:        true,
	MesgNumMesgCapabilities:        true,
	MesgNumFieldCapabilities:       true,
	MesgNumFileCreator:             true,
	MesgNumBloodPressure:           true,
	MesgNumSpeedZone:               true,
	MesgNumMonitoring:              true,
	MesgNumTrainingFile:            true,
	MesgNumHrv:                     true,
	MesgNumLength:                  true,
	MesgNumMonitoringInfo:          true,
	MesgNumSlaveDevice:             true,
	MesgNumCadenceZone:             true,
	MesgNumSegmentLap:              true,
	MesgNumMemoGlob:                true,
	MesgNumSegmentId:               true,
	MesgNumSegmentLeaderboardEntry: true,
	MesgNumSegmentPoint:            true,
	MesgNumSegmentFile:             true,
	MesgNumCameraEvent:             true,
	MesgNumTimestampCorrelation:    true,
	MesgNumGyroscopeData:           true,
	MesgNumAccelerometerData:       true,
	MesgNumThreeDSensorCalibration: true,
	MesgNumVideoFrame:              true,
	MesgNumObdiiData:               true,
	MesgNumNmeaSentence:            true,
	MesgNumAviationAttitude:        true,
	MesgNumVideo:                   true,
	MesgNumVideoTitle:              true,
	MesgNumVideoDescription:        true,
	MesgNumVideoClip:               true,
}

var (
	accumuDistance         *uint32Accumulator
	accumuTotalCycles      *uint32Accumulator
	accumuAccumulatedPower *uint32Accumulator
)

// Set length to 256, so that lookup for any
// field 255 (localMesgNumInvalid) will return nil.
var _fields = [...][256]*field{
	MesgNumFileId: {
		0: {0, 0, types.Fit(0)},
		1: {1, 1, types.Fit(4)},
		2: {2, 2, types.Fit(4)},
		3: {3, 3, types.Fit(12)},
		4: {4, 4, types.Fit(38)},
		5: {5, 5, types.Fit(4)},
		8: {6, 8, types.Fit(7)},
	},

	MesgNumFileCreator: {
		0: {0, 0, types.Fit(4)},
		1: {1, 1, types.Fit(2)},
	},

	MesgNumTimestampCorrelation: {},

	MesgNumSoftware: {
		254: {0, 254, types.Fit(4)},
		3:   {1, 3, types.Fit(4)},
		5:   {2, 5, types.Fit(7)},
	},

	MesgNumSlaveDevice: {
		0: {0, 0, types.Fit(4)},
		1: {1, 1, types.Fit(4)},
	},

	MesgNumCapabilities: {
		0:  {0, 0, types.Fit(26)},
		1:  {1, 1, types.Fit(26)},
		21: {2, 21, types.Fit(12)},
		23: {3, 23, types.Fit(12)},
	},

	MesgNumFileCapabilities: {
		254: {0, 254, types.Fit(4)},
		0:   {1, 0, types.Fit(0)},
		1:   {2, 1, types.Fit(10)},
		2:   {3, 2, types.Fit(7)},
		3:   {4, 3, types.Fit(4)},
		4:   {5, 4, types.Fit(6)},
	},

	MesgNumMesgCapabilities: {
		254: {0, 254, types.Fit(4)},
		0:   {1, 0, types.Fit(0)},
		1:   {2, 1, types.Fit(4)},
		2:   {3, 2, types.Fit(0)},
		3:   {4, 3, types.Fit(4)},
	},

	MesgNumFieldCapabilities: {
		254: {0, 254, types.Fit(4)},
		0:   {1, 0, types.Fit(0)},
		1:   {2, 1, types.Fit(4)},
		2:   {3, 2, types.Fit(2)},
		3:   {4, 3, types.Fit(4)},
	},

	MesgNumDeviceSettings: {
		0: {0, 0, types.Fit(2)},
		1: {1, 1, types.Fit(6)},
		5: {2, 5, types.Fit(17)},
	},

	MesgNumUserProfile: {
		254: {0, 254, types.Fit(4)},
		0:   {1, 0, types.Fit(7)},
		1:   {2, 1, types.Fit(0)},
		2:   {3, 2, types.Fit(2)},
		3:   {4, 3, types.Fit(2)},
		4:   {5, 4, types.Fit(4)},
		5:   {6, 5, types.Fit(0)},
		6:   {7, 6, types.Fit(0)},
		7:   {8, 7, types.Fit(0)},
		8:   {9, 8, types.Fit(2)},
		9:   {10, 9, types.Fit(2)},
		10:  {11, 10, types.Fit(2)},
		11:  {12, 11, types.Fit(2)},
		12:  {13, 12, types.Fit(0)},
		13:  {14, 13, types.Fit(0)},
		14:  {15, 14, types.Fit(0)},
		16:  {16, 16, types.Fit(0)},
		17:  {17, 17, types.Fit(0)},
		18:  {18, 18, types.Fit(0)},
		21:  {19, 21, types.Fit(0)},
		22:  {20, 22, types.Fit(4)},
		23:  {21, 23, types.Fit(29)},
		30:  {22, 30, types.Fit(0)},
	},

	MesgNumHrmProfile: {
		254: {0, 254, types.Fit(4)},
		0:   {1, 0, types.Fit(0)},
		1:   {2, 1, types.Fit(11)},
		2:   {3, 2, types.Fit(0)},
		3:   {4, 3, types.Fit(10)},
	},

	MesgNumSdmProfile: {
		254: {0, 254, types.Fit(4)},
		0:   {1, 0, types.Fit(0)},
		1:   {2, 1, types.Fit(11)},
		2:   {3, 2, types.Fit(4)},
		3:   {4, 3, types.Fit(6)},
		4:   {5, 4, types.Fit(0)},
		5:   {6, 5, types.Fit(10)},
		7:   {7, 7, types.Fit(2)},
	},

	MesgNumBikeProfile: {
		254: {0, 254, types.Fit(4)},
		0:   {1, 0, types.Fit(7)},
		1:   {2, 1, types.Fit(0)},
		2:   {3, 2, types.Fit(0)},
		3:   {4, 3, types.Fit(6)},
		4:   {5, 4, types.Fit(11)},
		5:   {6, 5, types.Fit(11)},
		6:   {7, 6, types.Fit(11)},
		7:   {8, 7, types.Fit(11)},
		8:   {9, 8, types.Fit(4)},
		9:   {10, 9, types.Fit(4)},
		10:  {11, 10, types.Fit(4)},
		11:  {12, 11, types.Fit(4)},
		12:  {13, 12, types.Fit(0)},
		13:  {14, 13, types.Fit(0)},
		14:  {15, 14, types.Fit(2)},
		15:  {16, 15, types.Fit(0)},
		16:  {17, 16, types.Fit(0)},
		17:  {18, 17, types.Fit(0)},
		18:  {19, 18, types.Fit(0)},
		19:  {20, 19, types.Fit(2)},
		20:  {21, 20, types.Fit(0)},
		21:  {22, 21, types.Fit(10)},
		22:  {23, 22, types.Fit(10)},
		23:  {24, 23, types.Fit(10)},
		24:  {25, 24, types.Fit(10)},
		37:  {26, 37, types.Fit(2)},
		38:  {27, 38, types.Fit(10)},
		39:  {28, 39, types.Fit(26)},
		40:  {29, 40, types.Fit(10)},
		41:  {30, 41, types.Fit(26)},
		44:  {31, 44, types.Fit(0)},
	},

	MesgNumZonesTarget: {
		1: {0, 1, types.Fit(2)},
		2: {1, 2, types.Fit(2)},
		3: {2, 3, types.Fit(4)},
		5: {3, 5, types.Fit(0)},
		7: {4, 7, types.Fit(0)},
	},

	MesgNumSport: {
		0: {0, 0, types.Fit(0)},
		1: {1, 1, types.Fit(0)},
		3: {2, 3, types.Fit(7)},
	},

	MesgNumHrZone: {
		254: {0, 254, types.Fit(4)},
		1:   {1, 1, types.Fit(2)},
		2:   {2, 2, types.Fit(7)},
	},

	MesgNumSpeedZone: {
		254: {0, 254, types.Fit(4)},
		0:   {1, 0, types.Fit(4)},
		1:   {2, 1, types.Fit(7)},
	},

	MesgNumCadenceZone: {
		254: {0, 254, types.Fit(4)},
		0:   {1, 0, types.Fit(2)},
		1:   {2, 1, types.Fit(7)},
	},

	MesgNumPowerZone: {
		254: {0, 254, types.Fit(4)},
		1:   {1, 1, types.Fit(4)},
		2:   {2, 2, types.Fit(7)},
	},

	MesgNumMetZone: {
		254: {0, 254, types.Fit(4)},
		1:   {1, 1, types.Fit(2)},
		2:   {2, 2, types.Fit(4)},
		3:   {3, 3, types.Fit(2)},
	},

	MesgNumGoal: {
		254: {0, 254, types.Fit(4)},
		0:   {1, 0, types.Fit(0)},
		1:   {2, 1, types.Fit(0)},
		2:   {3, 2, types.Fit(38)},
		3:   {4, 3, types.Fit(38)},
		4:   {5, 4, types.Fit(0)},
		5:   {6, 5, types.Fit(6)},
		6:   {7, 6, types.Fit(0)},
		7:   {8, 7, types.Fit(6)},
		8:   {9, 8, types.Fit(0)},
		9:   {10, 9, types.Fit(4)},
		10:  {11, 10, types.Fit(0)},
	},

	MesgNumActivity: {
		253: {0, 253, types.Fit(38)},
		0:   {1, 0, types.Fit(6)},
		1:   {2, 1, types.Fit(4)},
		2:   {3, 2, types.Fit(0)},
		3:   {4, 3, types.Fit(0)},
		4:   {5, 4, types.Fit(0)},
		5:   {6, 5, types.Fit(70)},
		6:   {7, 6, types.Fit(2)},
	},

	MesgNumSession: {
		254: {0, 254, types.Fit(4)},
		253: {1, 253, types.Fit(38)},
		0:   {2, 0, types.Fit(0)},
		1:   {3, 1, types.Fit(0)},
		2:   {4, 2, types.Fit(38)},
		3:   {5, 3, types.Fit(101)},
		4:   {6, 4, types.Fit(133)},
		5:   {7, 5, types.Fit(0)},
		6:   {8, 6, types.Fit(0)},
		7:   {9, 7, types.Fit(6)},
		8:   {10, 8, types.Fit(6)},
		9:   {11, 9, types.Fit(6)},
		10:  {12, 10, types.Fit(6)},
		11:  {13, 11, types.Fit(4)},
		13:  {14, 13, types.Fit(4)},
		14:  {15, 14, types.Fit(4)},
		15:  {16, 15, types.Fit(4)},
		16:  {17, 16, types.Fit(2)},
		17:  {18, 17, types.Fit(2)},
		18:  {19, 18, types.Fit(2)},
		19:  {20, 19, types.Fit(2)},
		20:  {21, 20, types.Fit(4)},
		21:  {22, 21, types.Fit(4)},
		22:  {23, 22, types.Fit(4)},
		23:  {24, 23, types.Fit(4)},
		24:  {25, 24, types.Fit(2)},
		25:  {26, 25, types.Fit(4)},
		26:  {27, 26, types.Fit(4)},
		27:  {28, 27, types.Fit(2)},
		28:  {29, 28, types.Fit(0)},
		29:  {30, 29, types.Fit(101)},
		30:  {31, 30, types.Fit(133)},
		31:  {32, 31, types.Fit(101)},
		32:  {33, 32, types.Fit(133)},
		34:  {34, 34, types.Fit(4)},
		35:  {35, 35, types.Fit(4)},
		36:  {36, 36, types.Fit(4)},
		37:  {37, 37, types.Fit(4)},
		41:  {38, 41, types.Fit(6)},
		42:  {39, 42, types.Fit(4)},
		43:  {40, 43, types.Fit(0)},
		44:  {41, 44, types.Fit(4)},
		45:  {42, 45, types.Fit(4)},
		46:  {43, 46, types.Fit(0)},
		47:  {44, 47, types.Fit(4)},
		48:  {45, 48, types.Fit(6)},
		49:  {46, 49, types.Fit(4)},
		50:  {47, 50, types.Fit(4)},
		51:  {48, 51, types.Fit(2)},
		52:  {49, 52, types.Fit(3)},
		53:  {50, 53, types.Fit(3)},
		54:  {51, 54, types.Fit(3)},
		55:  {52, 55, types.Fit(3)},
		56:  {53, 56, types.Fit(3)},
		57:  {54, 57, types.Fit(1)},
		58:  {55, 58, types.Fit(1)},
		59:  {56, 59, types.Fit(6)},
		60:  {57, 60, types.Fit(3)},
		61:  {58, 61, types.Fit(3)},
		62:  {59, 62, types.Fit(3)},
		63:  {60, 63, types.Fit(3)},
		64:  {61, 64, types.Fit(2)},
		65:  {62, 65, types.Fit(22)},
		66:  {63, 66, types.Fit(22)},
		67:  {64, 67, types.Fit(22)},
		68:  {65, 68, types.Fit(22)},
		69:  {66, 69, types.Fit(6)},
		70:  {67, 70, types.Fit(4)},
		71:  {68, 71, types.Fit(4)},
		82:  {69, 82, types.Fit(4)},
		83:  {70, 83, types.Fit(4)},
		84:  {71, 84, types.Fit(7)},
		85:  {72, 85, types.Fit(20)},
		86:  {73, 86, types.Fit(20)},
		87:  {74, 87, types.Fit(4)},
		88:  {75, 88, types.Fit(4)},
		89:  {76, 89, types.Fit(4)},
		90:  {77, 90, types.Fit(4)},
		91:  {78, 91, types.Fit(4)},
		92:  {79, 92, types.Fit(2)},
		93:  {80, 93, types.Fit(2)},
		94:  {81, 94, types.Fit(2)},
		111: {82, 111, types.Fit(2)},
		124: {83, 124, types.Fit(6)},
		125: {84, 125, types.Fit(6)},
		126: {85, 126, types.Fit(6)},
		127: {86, 127, types.Fit(6)},
		128: {87, 128, types.Fit(6)},
	},

	MesgNumLap: {
		254: {0, 254, types.Fit(4)},
		253: {1, 253, types.Fit(38)},
		0:   {2, 0, types.Fit(0)},
		1:   {3, 1, types.Fit(0)},
		2:   {4, 2, types.Fit(38)},
		3:   {5, 3, types.Fit(101)},
		4:   {6, 4, types.Fit(133)},
		5:   {7, 5, types.Fit(101)},
		6:   {8, 6, types.Fit(133)},
		7:   {9, 7, types.Fit(6)},
		8:   {10, 8, types.Fit(6)},
		9:   {11, 9, types.Fit(6)},
		10:  {12, 10, types.Fit(6)},
		11:  {13, 11, types.Fit(4)},
		12:  {14, 12, types.Fit(4)},
		13:  {15, 13, types.Fit(4)},
		14:  {16, 14, types.Fit(4)},
		15:  {17, 15, types.Fit(2)},
		16:  {18, 16, types.Fit(2)},
		17:  {19, 17, types.Fit(2)},
		18:  {20, 18, types.Fit(2)},
		19:  {21, 19, types.Fit(4)},
		20:  {22, 20, types.Fit(4)},
		21:  {23, 21, types.Fit(4)},
		22:  {24, 22, types.Fit(4)},
		23:  {25, 23, types.Fit(0)},
		24:  {26, 24, types.Fit(0)},
		25:  {27, 25, types.Fit(0)},
		26:  {28, 26, types.Fit(2)},
		32:  {29, 32, types.Fit(4)},
		33:  {30, 33, types.Fit(4)},
		34:  {31, 34, types.Fit(4)},
		35:  {32, 35, types.Fit(4)},
		37:  {33, 37, types.Fit(4)},
		38:  {34, 38, types.Fit(0)},
		39:  {35, 39, types.Fit(0)},
		40:  {36, 40, types.Fit(4)},
		41:  {37, 41, types.Fit(6)},
		42:  {38, 42, types.Fit(4)},
		43:  {39, 43, types.Fit(4)},
		44:  {40, 44, types.Fit(2)},
		45:  {41, 45, types.Fit(3)},
		46:  {42, 46, types.Fit(3)},
		47:  {43, 47, types.Fit(3)},
		48:  {44, 48, types.Fit(3)},
		49:  {45, 49, types.Fit(3)},
		50:  {46, 50, types.Fit(1)},
		51:  {47, 51, types.Fit(1)},
		52:  {48, 52, types.Fit(6)},
		53:  {49, 53, types.Fit(3)},
		54:  {50, 54, types.Fit(3)},
		55:  {51, 55, types.Fit(3)},
		56:  {52, 56, types.Fit(3)},
		57:  {53, 57, types.Fit(22)},
		58:  {54, 58, types.Fit(22)},
		59:  {55, 59, types.Fit(22)},
		60:  {56, 60, types.Fit(22)},
		61:  {57, 61, types.Fit(4)},
		62:  {58, 62, types.Fit(4)},
		63:  {59, 63, types.Fit(2)},
		71:  {60, 71, types.Fit(4)},
		74:  {61, 74, types.Fit(4)},
		75:  {62, 75, types.Fit(20)},
		76:  {63, 76, types.Fit(20)},
		77:  {64, 77, types.Fit(4)},
		78:  {65, 78, types.Fit(4)},
		79:  {66, 79, types.Fit(4)},
		80:  {67, 80, types.Fit(2)},
		81:  {68, 81, types.Fit(2)},
		82:  {69, 82, types.Fit(2)},
		83:  {70, 83, types.Fit(4)},
		84:  {71, 84, types.Fit(20)},
		85:  {72, 85, types.Fit(20)},
		86:  {73, 86, types.Fit(20)},
		87:  {74, 87, types.Fit(20)},
		88:  {75, 88, types.Fit(20)},
		89:  {76, 89, types.Fit(20)},
		110: {77, 110, types.Fit(6)},
		111: {78, 111, types.Fit(6)},
		112: {79, 112, types.Fit(6)},
		113: {80, 113, types.Fit(6)},
		114: {81, 114, types.Fit(6)},
	},

	MesgNumLength: {
		254: {0, 254, types.Fit(4)},
		253: {1, 253, types.Fit(38)},
		0:   {2, 0, types.Fit(0)},
		1:   {3, 1, types.Fit(0)},
		2:   {4, 2, types.Fit(38)},
		3:   {5, 3, types.Fit(6)},
		4:   {6, 4, types.Fit(6)},
		5:   {7, 5, types.Fit(4)},
		6:   {8, 6, types.Fit(4)},
		7:   {9, 7, types.Fit(0)},
		9:   {10, 9, types.Fit(2)},
		10:  {11, 10, types.Fit(2)},
		11:  {12, 11, types.Fit(4)},
		12:  {13, 12, types.Fit(0)},
		18:  {14, 18, types.Fit(4)},
		19:  {15, 19, types.Fit(4)},
		20:  {16, 20, types.Fit(20)},
		21:  {17, 21, types.Fit(20)},
	},

	MesgNumRecord: {
		253: {0, 253, types.Fit(38)},
		0:   {1, 0, types.Fit(101)},
		1:   {2, 1, types.Fit(133)},
		2:   {3, 2, types.Fit(4)},
		3:   {4, 3, types.Fit(2)},
		4:   {5, 4, types.Fit(2)},
		5:   {6, 5, types.Fit(6)},
		6:   {7, 6, types.Fit(4)},
		7:   {8, 7, types.Fit(4)},
		8:   {9, 8, types.Fit(29)},
		9:   {10, 9, types.Fit(3)},
		10:  {11, 10, types.Fit(2)},
		11:  {12, 11, types.Fit(5)},
		12:  {13, 12, types.Fit(2)},
		13:  {14, 13, types.Fit(1)},
		17:  {15, 17, types.Fit(18)},
		18:  {16, 18, types.Fit(2)},
		19:  {17, 19, types.Fit(6)},
		28:  {18, 28, types.Fit(4)},
		29:  {19, 29, types.Fit(6)},
		30:  {20, 30, types.Fit(2)},
		31:  {21, 31, types.Fit(2)},
		32:  {22, 32, types.Fit(3)},
		33:  {23, 33, types.Fit(4)},
		39:  {24, 39, types.Fit(4)},
		40:  {25, 40, types.Fit(4)},
		41:  {26, 41, types.Fit(4)},
		42:  {27, 42, types.Fit(0)},
		43:  {28, 43, types.Fit(2)},
		44:  {29, 44, types.Fit(2)},
		45:  {30, 45, types.Fit(2)},
		46:  {31, 46, types.Fit(2)},
		47:  {32, 47, types.Fit(2)},
		48:  {33, 48, types.Fit(2)},
		49:  {34, 49, types.Fit(0)},
		50:  {35, 50, types.Fit(2)},
		51:  {36, 51, types.Fit(4)},
		52:  {37, 52, types.Fit(4)},
		53:  {38, 53, types.Fit(2)},
		54:  {39, 54, types.Fit(4)},
		55:  {40, 55, types.Fit(4)},
		56:  {41, 56, types.Fit(4)},
		57:  {42, 57, types.Fit(4)},
		58:  {43, 58, types.Fit(4)},
		59:  {44, 59, types.Fit(4)},
		62:  {45, 62, types.Fit(2)},
		73:  {46, 73, types.Fit(6)},
		78:  {47, 78, types.Fit(6)},
	},

	MesgNumEvent: {
		253: {0, 253, types.Fit(38)},
		0:   {1, 0, types.Fit(0)},
		1:   {2, 1, types.Fit(0)},
		2:   {3, 2, types.Fit(4)},
		3:   {4, 3, types.Fit(6)},
		4:   {5, 4, types.Fit(2)},
		7:   {6, 7, types.Fit(4)},
		8:   {7, 8, types.Fit(4)},
		9:   {8, 9, types.Fit(10)},
		10:  {9, 10, types.Fit(10)},
		11:  {10, 11, types.Fit(10)},
		12:  {11, 12, types.Fit(10)},
	},

	MesgNumDeviceInfo: {
		253: {0, 253, types.Fit(38)},
		0:   {1, 0, types.Fit(2)},
		1:   {2, 1, types.Fit(2)},
		2:   {3, 2, types.Fit(4)},
		3:   {4, 3, types.Fit(12)},
		4:   {5, 4, types.Fit(4)},
		5:   {6, 5, types.Fit(4)},
		6:   {7, 6, types.Fit(2)},
		7:   {8, 7, types.Fit(6)},
		10:  {9, 10, types.Fit(4)},
		11:  {10, 11, types.Fit(2)},
		18:  {11, 18, types.Fit(0)},
		19:  {12, 19, types.Fit(7)},
		20:  {13, 20, types.Fit(10)},
		21:  {14, 21, types.Fit(11)},
		22:  {15, 22, types.Fit(0)},
		25:  {16, 25, types.Fit(0)},
		27:  {17, 27, types.Fit(7)},
	},

	MesgNumTrainingFile: {
		253: {0, 253, types.Fit(38)},
		0:   {1, 0, types.Fit(0)},
		1:   {2, 1, types.Fit(4)},
		2:   {3, 2, types.Fit(4)},
		3:   {4, 3, types.Fit(12)},
		4:   {5, 4, types.Fit(38)},
	},

	MesgNumHrv: {
		0: {0, 0, types.Fit(20)},
	},

	MesgNumCameraEvent: {},

	MesgNumGyroscopeData: {},

	MesgNumAccelerometerData: {},

	MesgNumThreeDSensorCalibration: {},

	MesgNumVideoFrame: {},

	MesgNumObdiiData: {},

	MesgNumNmeaSentence: {
		253: {0, 253, types.Fit(38)},
		0:   {1, 0, types.Fit(4)},
		1:   {2, 1, types.Fit(7)},
	},

	MesgNumAviationAttitude: {
		253: {0, 253, types.Fit(38)},
		0:   {1, 0, types.Fit(4)},
		1:   {2, 1, types.Fit(22)},
		2:   {3, 2, types.Fit(19)},
		3:   {4, 3, types.Fit(19)},
		4:   {5, 4, types.Fit(19)},
		5:   {6, 5, types.Fit(19)},
		6:   {7, 6, types.Fit(19)},
		7:   {8, 7, types.Fit(16)},
		8:   {9, 8, types.Fit(18)},
		9:   {10, 9, types.Fit(20)},
		10:  {11, 10, types.Fit(20)},
	},

	MesgNumVideo: {},

	MesgNumVideoTitle: {
		254: {0, 254, types.Fit(4)},
		0:   {1, 0, types.Fit(4)},
		1:   {2, 1, types.Fit(7)},
	},

	MesgNumVideoDescription: {
		254: {0, 254, types.Fit(4)},
		0:   {1, 0, types.Fit(4)},
		1:   {2, 1, types.Fit(7)},
	},

	MesgNumVideoClip: {},

	MesgNumCourse: {
		4: {0, 4, types.Fit(0)},
		5: {1, 5, types.Fit(7)},
		6: {2, 6, types.Fit(12)},
	},

	MesgNumCoursePoint: {
		254: {0, 254, types.Fit(4)},
		1:   {1, 1, types.Fit(38)},
		2:   {2, 2, types.Fit(101)},
		3:   {3, 3, types.Fit(133)},
		4:   {4, 4, types.Fit(6)},
		5:   {5, 5, types.Fit(0)},
		6:   {6, 6, types.Fit(7)},
		8:   {7, 8, types.Fit(0)},
	},

	MesgNumSegmentId: {
		0: {0, 0, types.Fit(7)},
		1: {1, 1, types.Fit(7)},
		2: {2, 2, types.Fit(0)},
		3: {3, 3, types.Fit(0)},
		4: {4, 4, types.Fit(6)},
		5: {5, 5, types.Fit(6)},
		6: {6, 6, types.Fit(2)},
		7: {7, 7, types.Fit(0)},
		8: {8, 8, types.Fit(0)},
	},

	MesgNumSegmentLeaderboardEntry: {
		254: {0, 254, types.Fit(4)},
		0:   {1, 0, types.Fit(7)},
		1:   {2, 1, types.Fit(0)},
		2:   {3, 2, types.Fit(6)},
		3:   {4, 3, types.Fit(6)},
		4:   {5, 4, types.Fit(6)},
	},

	MesgNumSegmentPoint: {
		254: {0, 254, types.Fit(4)},
		1:   {1, 1, types.Fit(101)},
		2:   {2, 2, types.Fit(133)},
		3:   {3, 3, types.Fit(6)},
		4:   {4, 4, types.Fit(4)},
		5:   {5, 5, types.Fit(22)},
	},

	MesgNumSegmentLap: {
		254: {0, 254, types.Fit(4)},
		253: {1, 253, types.Fit(38)},
		0:   {2, 0, types.Fit(0)},
		1:   {3, 1, types.Fit(0)},
		2:   {4, 2, types.Fit(38)},
		3:   {5, 3, types.Fit(101)},
		4:   {6, 4, types.Fit(133)},
		5:   {7, 5, types.Fit(101)},
		6:   {8, 6, types.Fit(133)},
		7:   {9, 7, types.Fit(6)},
		8:   {10, 8, types.Fit(6)},
		9:   {11, 9, types.Fit(6)},
		10:  {12, 10, types.Fit(6)},
		11:  {13, 11, types.Fit(4)},
		12:  {14, 12, types.Fit(4)},
		13:  {15, 13, types.Fit(4)},
		14:  {16, 14, types.Fit(4)},
		15:  {17, 15, types.Fit(2)},
		16:  {18, 16, types.Fit(2)},
		17:  {19, 17, types.Fit(2)},
		18:  {20, 18, types.Fit(2)},
		19:  {21, 19, types.Fit(4)},
		20:  {22, 20, types.Fit(4)},
		21:  {23, 21, types.Fit(4)},
		22:  {24, 22, types.Fit(4)},
		23:  {25, 23, types.Fit(0)},
		24:  {26, 24, types.Fit(2)},
		25:  {27, 25, types.Fit(101)},
		26:  {28, 26, types.Fit(133)},
		27:  {29, 27, types.Fit(101)},
		28:  {30, 28, types.Fit(133)},
		29:  {31, 29, types.Fit(7)},
		30:  {32, 30, types.Fit(4)},
		31:  {33, 31, types.Fit(4)},
		32:  {34, 32, types.Fit(0)},
		33:  {35, 33, types.Fit(6)},
		34:  {36, 34, types.Fit(4)},
		35:  {37, 35, types.Fit(4)},
		36:  {38, 36, types.Fit(2)},
		37:  {39, 37, types.Fit(3)},
		38:  {40, 38, types.Fit(3)},
		39:  {41, 39, types.Fit(3)},
		40:  {42, 40, types.Fit(3)},
		41:  {43, 41, types.Fit(3)},
		42:  {44, 42, types.Fit(1)},
		43:  {45, 43, types.Fit(1)},
		44:  {46, 44, types.Fit(6)},
		45:  {47, 45, types.Fit(3)},
		46:  {48, 46, types.Fit(3)},
		47:  {49, 47, types.Fit(3)},
		48:  {50, 48, types.Fit(3)},
		49:  {51, 49, types.Fit(22)},
		50:  {52, 50, types.Fit(22)},
		51:  {53, 51, types.Fit(22)},
		52:  {54, 52, types.Fit(22)},
		53:  {55, 53, types.Fit(4)},
		54:  {56, 54, types.Fit(4)},
		55:  {57, 55, types.Fit(2)},
		56:  {58, 56, types.Fit(6)},
		57:  {59, 57, types.Fit(4)},
		58:  {60, 58, types.Fit(0)},
		59:  {61, 59, types.Fit(2)},
		60:  {62, 60, types.Fit(2)},
		61:  {63, 61, types.Fit(2)},
		62:  {64, 62, types.Fit(2)},
		63:  {65, 63, types.Fit(2)},
		64:  {66, 64, types.Fit(0)},
		65:  {67, 65, types.Fit(7)},
		66:  {68, 66, types.Fit(2)},
		67:  {69, 67, types.Fit(2)},
		68:  {70, 68, types.Fit(2)},
		69:  {71, 69, types.Fit(4)},
		70:  {72, 70, types.Fit(4)},
	},

	MesgNumSegmentFile: {
		254: {0, 254, types.Fit(4)},
		1:   {1, 1, types.Fit(7)},
		3:   {2, 3, types.Fit(0)},
		4:   {3, 4, types.Fit(6)},
		7:   {4, 7, types.Fit(16)},
		8:   {5, 8, types.Fit(22)},
		9:   {6, 9, types.Fit(22)},
	},

	MesgNumWorkout: {
		4: {0, 4, types.Fit(0)},
		5: {1, 5, types.Fit(12)},
		6: {2, 6, types.Fit(4)},
		8: {3, 8, types.Fit(7)},
	},

	MesgNumWorkoutStep: {
		254: {0, 254, types.Fit(4)},
		0:   {1, 0, types.Fit(7)},
		1:   {2, 1, types.Fit(0)},
		2:   {3, 2, types.Fit(6)},
		3:   {4, 3, types.Fit(0)},
		4:   {5, 4, types.Fit(6)},
		5:   {6, 5, types.Fit(6)},
		6:   {7, 6, types.Fit(6)},
		7:   {8, 7, types.Fit(0)},
	},

	MesgNumSchedule: {
		0: {0, 0, types.Fit(4)},
		1: {1, 1, types.Fit(4)},
		2: {2, 2, types.Fit(12)},
		3: {3, 3, types.Fit(38)},
		4: {4, 4, types.Fit(0)},
		5: {5, 5, types.Fit(0)},
		6: {6, 6, types.Fit(70)},
	},

	MesgNumTotals: {
		254: {0, 254, types.Fit(4)},
		253: {1, 253, types.Fit(38)},
		0:   {2, 0, types.Fit(6)},
		1:   {3, 1, types.Fit(6)},
		2:   {4, 2, types.Fit(6)},
		3:   {5, 3, types.Fit(0)},
		4:   {6, 4, types.Fit(6)},
		5:   {7, 5, types.Fit(4)},
		6:   {8, 6, types.Fit(6)},
	},

	MesgNumWeightScale: {
		253: {0, 253, types.Fit(38)},
		0:   {1, 0, types.Fit(4)},
		1:   {2, 1, types.Fit(4)},
		2:   {3, 2, types.Fit(4)},
		3:   {4, 3, types.Fit(4)},
		4:   {5, 4, types.Fit(4)},
		5:   {6, 5, types.Fit(4)},
		7:   {7, 7, types.Fit(4)},
		8:   {8, 8, types.Fit(2)},
		9:   {9, 9, types.Fit(4)},
		10:  {10, 10, types.Fit(2)},
		11:  {11, 11, types.Fit(2)},
		12:  {12, 12, types.Fit(4)},
	},

	MesgNumBloodPressure: {
		253: {0, 253, types.Fit(38)},
		0:   {1, 0, types.Fit(4)},
		1:   {2, 1, types.Fit(4)},
		2:   {3, 2, types.Fit(4)},
		3:   {4, 3, types.Fit(4)},
		4:   {5, 4, types.Fit(4)},
		5:   {6, 5, types.Fit(4)},
		6:   {7, 6, types.Fit(2)},
		7:   {8, 7, types.Fit(0)},
		8:   {9, 8, types.Fit(0)},
		9:   {10, 9, types.Fit(4)},
	},

	MesgNumMonitoringInfo: {
		253: {0, 253, types.Fit(38)},
		0:   {1, 0, types.Fit(70)},
	},

	MesgNumMonitoring: {
		253: {0, 253, types.Fit(38)},
		0:   {1, 0, types.Fit(2)},
		1:   {2, 1, types.Fit(4)},
		2:   {3, 2, types.Fit(6)},
		3:   {4, 3, types.Fit(6)},
		4:   {5, 4, types.Fit(6)},
		5:   {6, 5, types.Fit(0)},
		6:   {7, 6, types.Fit(0)},
		8:   {8, 8, types.Fit(4)},
		9:   {9, 9, types.Fit(4)},
		10:  {10, 10, types.Fit(4)},
		11:  {11, 11, types.Fit(70)},
	},

	MesgNumMemoGlob: {},
}

func getField(gmn MesgNum, fdn byte) (*field, bool) {
	if int(gmn) >= len(_fields) {
		return nil, false
	}
	f := _fields[gmn][fdn]
	if f == nil {
		return nil, false
	}
	return f, true
}

var msgsTypes = [...]reflect.Type{
	MesgNumFileId:                  reflect.TypeOf(FileIdMsg{}),
	MesgNumFileCreator:             reflect.TypeOf(FileCreatorMsg{}),
	MesgNumTimestampCorrelation:    reflect.TypeOf(TimestampCorrelationMsg{}),
	MesgNumSoftware:                reflect.TypeOf(SoftwareMsg{}),
	MesgNumSlaveDevice:             reflect.TypeOf(SlaveDeviceMsg{}),
	MesgNumCapabilities:            reflect.TypeOf(CapabilitiesMsg{}),
	MesgNumFileCapabilities:        reflect.TypeOf(FileCapabilitiesMsg{}),
	MesgNumMesgCapabilities:        reflect.TypeOf(MesgCapabilitiesMsg{}),
	MesgNumFieldCapabilities:       reflect.TypeOf(FieldCapabilitiesMsg{}),
	MesgNumDeviceSettings:          reflect.TypeOf(DeviceSettingsMsg{}),
	MesgNumUserProfile:             reflect.TypeOf(UserProfileMsg{}),
	MesgNumHrmProfile:              reflect.TypeOf(HrmProfileMsg{}),
	MesgNumSdmProfile:              reflect.TypeOf(SdmProfileMsg{}),
	MesgNumBikeProfile:             reflect.TypeOf(BikeProfileMsg{}),
	MesgNumZonesTarget:             reflect.TypeOf(ZonesTargetMsg{}),
	MesgNumSport:                   reflect.TypeOf(SportMsg{}),
	MesgNumHrZone:                  reflect.TypeOf(HrZoneMsg{}),
	MesgNumSpeedZone:               reflect.TypeOf(SpeedZoneMsg{}),
	MesgNumCadenceZone:             reflect.TypeOf(CadenceZoneMsg{}),
	MesgNumPowerZone:               reflect.TypeOf(PowerZoneMsg{}),
	MesgNumMetZone:                 reflect.TypeOf(MetZoneMsg{}),
	MesgNumGoal:                    reflect.TypeOf(GoalMsg{}),
	MesgNumActivity:                reflect.TypeOf(ActivityMsg{}),
	MesgNumSession:                 reflect.TypeOf(SessionMsg{}),
	MesgNumLap:                     reflect.TypeOf(LapMsg{}),
	MesgNumLength:                  reflect.TypeOf(LengthMsg{}),
	MesgNumRecord:                  reflect.TypeOf(RecordMsg{}),
	MesgNumEvent:                   reflect.TypeOf(EventMsg{}),
	MesgNumDeviceInfo:              reflect.TypeOf(DeviceInfoMsg{}),
	MesgNumTrainingFile:            reflect.TypeOf(TrainingFileMsg{}),
	MesgNumHrv:                     reflect.TypeOf(HrvMsg{}),
	MesgNumCameraEvent:             reflect.TypeOf(CameraEventMsg{}),
	MesgNumGyroscopeData:           reflect.TypeOf(GyroscopeDataMsg{}),
	MesgNumAccelerometerData:       reflect.TypeOf(AccelerometerDataMsg{}),
	MesgNumThreeDSensorCalibration: reflect.TypeOf(ThreeDSensorCalibrationMsg{}),
	MesgNumVideoFrame:              reflect.TypeOf(VideoFrameMsg{}),
	MesgNumObdiiData:               reflect.TypeOf(ObdiiDataMsg{}),
	MesgNumNmeaSentence:            reflect.TypeOf(NmeaSentenceMsg{}),
	MesgNumAviationAttitude:        reflect.TypeOf(AviationAttitudeMsg{}),
	MesgNumVideo:                   reflect.TypeOf(VideoMsg{}),
	MesgNumVideoTitle:              reflect.TypeOf(VideoTitleMsg{}),
	MesgNumVideoDescription:        reflect.TypeOf(VideoDescriptionMsg{}),
	MesgNumVideoClip:               reflect.TypeOf(VideoClipMsg{}),
	MesgNumCourse:                  reflect.TypeOf(CourseMsg{}),
	MesgNumCoursePoint:             reflect.TypeOf(CoursePointMsg{}),
	MesgNumSegmentId:               reflect.TypeOf(SegmentIdMsg{}),
	MesgNumSegmentLeaderboardEntry: reflect.TypeOf(SegmentLeaderboardEntryMsg{}),
	MesgNumSegmentPoint:            reflect.TypeOf(SegmentPointMsg{}),
	MesgNumSegmentLap:              reflect.TypeOf(SegmentLapMsg{}),
	MesgNumSegmentFile:             reflect.TypeOf(SegmentFileMsg{}),
	MesgNumWorkout:                 reflect.TypeOf(WorkoutMsg{}),
	MesgNumWorkoutStep:             reflect.TypeOf(WorkoutStepMsg{}),
	MesgNumSchedule:                reflect.TypeOf(ScheduleMsg{}),
	MesgNumTotals:                  reflect.TypeOf(TotalsMsg{}),
	MesgNumWeightScale:             reflect.TypeOf(WeightScaleMsg{}),
	MesgNumBloodPressure:           reflect.TypeOf(BloodPressureMsg{}),
	MesgNumMonitoringInfo:          reflect.TypeOf(MonitoringInfoMsg{}),
	MesgNumMonitoring:              reflect.TypeOf(MonitoringMsg{}),
	MesgNumMemoGlob:                reflect.TypeOf(MemoGlobMsg{}),
}

var msgsAllInvalid = [...]reflect.Value{
	MesgNumFileId: reflect.ValueOf(FileIdMsg{
		0xFF,
		0xFFFF,
		0xFFFF,
		0x00000000,
		timeBase,
		0xFFFF,
		"",
	}),
	MesgNumFileCreator: reflect.ValueOf(FileCreatorMsg{
		0xFFFF,
		0xFF,
	}),
	MesgNumTimestampCorrelation: reflect.ValueOf(TimestampCorrelationMsg{}),
	MesgNumSoftware: reflect.ValueOf(SoftwareMsg{
		0xFFFF,
		0xFFFF,
		"",
	}),
	MesgNumSlaveDevice: reflect.ValueOf(SlaveDeviceMsg{
		0xFFFF,
		0xFFFF,
	}),
	MesgNumCapabilities: reflect.ValueOf(CapabilitiesMsg{
		nil,
		nil,
		0x00000000,
		0x00000000,
	}),
	MesgNumFileCapabilities: reflect.ValueOf(FileCapabilitiesMsg{
		0xFFFF,
		0xFF,
		0x00,
		"",
		0xFFFF,
		0xFFFFFFFF,
	}),
	MesgNumMesgCapabilities: reflect.ValueOf(MesgCapabilitiesMsg{
		0xFFFF,
		0xFF,
		0xFFFF,
		0xFF,
		0xFFFF,
	}),
	MesgNumFieldCapabilities: reflect.ValueOf(FieldCapabilitiesMsg{
		0xFFFF,
		0xFF,
		0xFFFF,
		0xFF,
		0xFFFF,
	}),
	MesgNumDeviceSettings: reflect.ValueOf(DeviceSettingsMsg{
		0xFF,
		0xFFFFFFFF,
		nil,
	}),
	MesgNumUserProfile: reflect.ValueOf(UserProfileMsg{
		0xFFFF,
		"",
		0xFF,
		0xFF,
		0xFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFFFF,
		nil,
		0xFF,
	}),
	MesgNumHrmProfile: reflect.ValueOf(HrmProfileMsg{
		0xFFFF,
		0xFF,
		0x0000,
		0xFF,
		0x00,
	}),
	MesgNumSdmProfile: reflect.ValueOf(SdmProfileMsg{
		0xFFFF,
		0xFF,
		0x0000,
		0xFFFF,
		0xFFFFFFFF,
		0xFF,
		0x00,
		0xFF,
	}),
	MesgNumBikeProfile: reflect.ValueOf(BikeProfileMsg{
		0xFFFF,
		"",
		0xFF,
		0xFF,
		0xFFFFFFFF,
		0x0000,
		0x0000,
		0x0000,
		0x0000,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0x00,
		0x00,
		0x00,
		0x00,
		0xFF,
		0x00,
		nil,
		0x00,
		nil,
		0xFF,
	}),
	MesgNumZonesTarget: reflect.ValueOf(ZonesTargetMsg{
		0xFF,
		0xFF,
		0xFFFF,
		0xFF,
		0xFF,
	}),
	MesgNumSport: reflect.ValueOf(SportMsg{
		0xFF,
		0xFF,
		"",
	}),
	MesgNumHrZone: reflect.ValueOf(HrZoneMsg{
		0xFFFF,
		0xFF,
		"",
	}),
	MesgNumSpeedZone: reflect.ValueOf(SpeedZoneMsg{
		0xFFFF,
		0xFFFF,
		"",
	}),
	MesgNumCadenceZone: reflect.ValueOf(CadenceZoneMsg{
		0xFFFF,
		0xFF,
		"",
	}),
	MesgNumPowerZone: reflect.ValueOf(PowerZoneMsg{
		0xFFFF,
		0xFFFF,
		"",
	}),
	MesgNumMetZone: reflect.ValueOf(MetZoneMsg{
		0xFFFF,
		0xFF,
		0xFFFF,
		0xFF,
	}),
	MesgNumGoal: reflect.ValueOf(GoalMsg{
		0xFFFF,
		0xFF,
		0xFF,
		timeBase,
		timeBase,
		0xFF,
		0xFFFFFFFF,
		0xFF,
		0xFFFFFFFF,
		0xFF,
		0xFFFF,
		0xFF,
	}),
	MesgNumActivity: reflect.ValueOf(ActivityMsg{
		timeBase,
		0xFFFFFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFF,
		timeBase,
		0xFF,
	}),
	MesgNumSession: reflect.ValueOf(SessionMsg{
		0xFFFF,
		timeBase,
		0xFF,
		0xFF,
		timeBase,
		NewLatitudeInvalid(),
		NewLongitudeInvalid(),
		0xFF,
		0xFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		NewLatitudeInvalid(),
		NewLongitudeInvalid(),
		NewLatitudeInvalid(),
		NewLongitudeInvalid(),
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFFFFFF,
		0xFFFF,
		0xFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFFFF,
		0xFFFFFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		0x7F,
		0x7F,
		0xFFFFFFFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		0xFF,
		nil,
		nil,
		nil,
		nil,
		0xFFFFFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		"",
		nil,
		nil,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
	}),
	MesgNumLap: reflect.ValueOf(LapMsg{
		0xFFFF,
		timeBase,
		0xFF,
		0xFF,
		timeBase,
		NewLatitudeInvalid(),
		NewLongitudeInvalid(),
		NewLatitudeInvalid(),
		NewLongitudeInvalid(),
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFFFF,
		0xFFFFFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		0x7F,
		0x7F,
		0xFFFFFFFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		nil,
		nil,
		nil,
		nil,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFFFF,
		0xFFFF,
		nil,
		nil,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFF,
		0xFFFF,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
	}),
	MesgNumLength: reflect.ValueOf(LengthMsg{
		0xFFFF,
		timeBase,
		0xFF,
		0xFF,
		timeBase,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFF,
		0xFFFF,
		0xFF,
		0xFFFF,
		0xFFFF,
		nil,
		nil,
	}),
	MesgNumRecord: reflect.ValueOf(RecordMsg{
		timeBase,
		NewLatitudeInvalid(),
		NewLongitudeInvalid(),
		0xFFFF,
		0xFF,
		0xFF,
		0xFFFFFFFF,
		0xFFFF,
		0xFFFF,
		nil,
		0x7FFF,
		0xFF,
		0x7FFFFFFF,
		0xFF,
		0x7F,
		nil,
		0xFF,
		0xFFFFFFFF,
		0xFFFF,
		0xFFFFFFFF,
		0xFF,
		0xFF,
		0x7FFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
	}),
	MesgNumEvent: reflect.ValueOf(EventMsg{
		timeBase,
		0xFF,
		0xFF,
		0xFFFF,
		0xFFFFFFFF,
		0xFF,
		0xFFFF,
		0xFFFF,
		0x00,
		0x00,
		0x00,
		0x00,
	}),
	MesgNumDeviceInfo: reflect.ValueOf(DeviceInfoMsg{
		timeBase,
		0xFF,
		0xFF,
		0xFFFF,
		0x00000000,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFFFFFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		"",
		0x00,
		0x0000,
		0xFF,
		0xFF,
		"",
	}),
	MesgNumTrainingFile: reflect.ValueOf(TrainingFileMsg{
		timeBase,
		0xFF,
		0xFFFF,
		0xFFFF,
		0x00000000,
		timeBase,
	}),
	MesgNumHrv: reflect.ValueOf(HrvMsg{
		nil,
	}),
	MesgNumCameraEvent:             reflect.ValueOf(CameraEventMsg{}),
	MesgNumGyroscopeData:           reflect.ValueOf(GyroscopeDataMsg{}),
	MesgNumAccelerometerData:       reflect.ValueOf(AccelerometerDataMsg{}),
	MesgNumThreeDSensorCalibration: reflect.ValueOf(ThreeDSensorCalibrationMsg{}),
	MesgNumVideoFrame:              reflect.ValueOf(VideoFrameMsg{}),
	MesgNumObdiiData:               reflect.ValueOf(ObdiiDataMsg{}),
	MesgNumNmeaSentence: reflect.ValueOf(NmeaSentenceMsg{
		timeBase,
		0xFFFF,
		"",
	}),
	MesgNumAviationAttitude: reflect.ValueOf(AviationAttitudeMsg{
		timeBase,
		0xFFFF,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	}),
	MesgNumVideo: reflect.ValueOf(VideoMsg{}),
	MesgNumVideoTitle: reflect.ValueOf(VideoTitleMsg{
		0xFFFF,
		0xFFFF,
		"",
	}),
	MesgNumVideoDescription: reflect.ValueOf(VideoDescriptionMsg{
		0xFFFF,
		0xFFFF,
		"",
	}),
	MesgNumVideoClip: reflect.ValueOf(VideoClipMsg{}),
	MesgNumCourse: reflect.ValueOf(CourseMsg{
		0xFF,
		"",
		0x00000000,
	}),
	MesgNumCoursePoint: reflect.ValueOf(CoursePointMsg{
		0xFFFF,
		timeBase,
		NewLatitudeInvalid(),
		NewLongitudeInvalid(),
		0xFFFFFFFF,
		0xFF,
		"",
		0xFF,
	}),
	MesgNumSegmentId: reflect.ValueOf(SegmentIdMsg{
		"",
		"",
		0xFF,
		0xFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFF,
		0xFF,
		0xFF,
	}),
	MesgNumSegmentLeaderboardEntry: reflect.ValueOf(SegmentLeaderboardEntryMsg{
		0xFFFF,
		"",
		0xFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
	}),
	MesgNumSegmentPoint: reflect.ValueOf(SegmentPointMsg{
		0xFFFF,
		NewLatitudeInvalid(),
		NewLongitudeInvalid(),
		0xFFFFFFFF,
		0xFFFF,
		nil,
	}),
	MesgNumSegmentLap: reflect.ValueOf(SegmentLapMsg{
		0xFFFF,
		timeBase,
		0xFF,
		0xFF,
		timeBase,
		NewLatitudeInvalid(),
		NewLongitudeInvalid(),
		NewLatitudeInvalid(),
		NewLongitudeInvalid(),
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		NewLatitudeInvalid(),
		NewLongitudeInvalid(),
		NewLatitudeInvalid(),
		NewLongitudeInvalid(),
		"",
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFFFFFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		0x7F,
		0x7F,
		0xFFFFFFFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		0x7FFF,
		nil,
		nil,
		nil,
		nil,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFFFFFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		0xFF,
		"",
		0xFF,
		0xFF,
		0xFF,
		0xFFFF,
		0xFFFF,
	}),
	MesgNumSegmentFile: reflect.ValueOf(SegmentFileMsg{
		0xFFFF,
		"",
		0xFF,
		0xFFFFFFFF,
		nil,
		nil,
		nil,
	}),
	MesgNumWorkout: reflect.ValueOf(WorkoutMsg{
		0xFF,
		0x00000000,
		0xFFFF,
		"",
	}),
	MesgNumWorkoutStep: reflect.ValueOf(WorkoutStepMsg{
		0xFFFF,
		"",
		0xFF,
		0xFFFFFFFF,
		0xFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFF,
	}),
	MesgNumSchedule: reflect.ValueOf(ScheduleMsg{
		0xFFFF,
		0xFFFF,
		0x00000000,
		timeBase,
		0xFF,
		0xFF,
		timeBase,
	}),
	MesgNumTotals: reflect.ValueOf(TotalsMsg{
		0xFFFF,
		timeBase,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFF,
		0xFFFFFFFF,
		0xFFFF,
		0xFFFFFFFF,
	}),
	MesgNumWeightScale: reflect.ValueOf(WeightScaleMsg{
		timeBase,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFFFF,
	}),
	MesgNumBloodPressure: reflect.ValueOf(BloodPressureMsg{
		timeBase,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		0xFF,
		0xFF,
		0xFF,
		0xFFFF,
	}),
	MesgNumMonitoringInfo: reflect.ValueOf(MonitoringInfoMsg{
		timeBase,
		timeBase,
	}),
	MesgNumMonitoring: reflect.ValueOf(MonitoringMsg{
		timeBase,
		0xFF,
		0xFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFFFFFFFF,
		0xFF,
		0xFF,
		0xFFFF,
		0xFFFF,
		0xFFFF,
		timeBase,
	}),
	MesgNumMemoGlob: reflect.ValueOf(MemoGlobMsg{}),
}

func getMesgAllInvalid(mn MesgNum) reflect.Value {
	val := reflect.New(msgsTypes[mn]).Elem()
	val.Set(msgsAllInvalid[mn])
	return val
}
