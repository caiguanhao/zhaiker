# zhaiker

zhaiker.cn API

## Usage

```go
import "github.com/caiguanhao/zhaiker"

ctx := context.Background()
client := zhaiker.NewClient(zhaiker.MustGetKey(ctx, "USERNAME", "PASSWORD"))

var gyms []zhaiker.Gym
client.MustRequest(ctx, zhaiker.API_GET_GYMS, zhaiker.Params("pageSize", 100), &gyms, "DATA.*")

var exams []zhaiker.Exam
client.MustRequest(ctx, zhaiker.API_GET_EXAMS, zhaiker.Params("gymId", gyms[0].Id), &exams, "DATA.*")

var exam zhaiker.ExamWithStandards
client.MustRequest(ctx, zhaiker.API_GET_EXAM, zhaiker.Params("gymId", gyms[0].Id, "id", exams[0].Id), &exam)
```

## ExamWithStandards

```json
{
  "Exam": {
    "AerobicGoal": 398.475,
    "Age": 15,
    "Agility": 482,
    "AnaGoal": 56.925,
    "Balance": 3,
    "BalanceAngle": 73.3039,
    "BloodMaxPressure": 129,
    "BloodMinPressure": 98,
    "BloodOxygen": 95,
    "Bmi": 21.420378,
    "Bmr": 1265,
    "BodyAge": 17,
    "BodyDetect": {
      "BodyDirection": 1,
      "FullBody": true,
      "LeftAnkle": {
        "Score": 0.85002023,
        "X": 264.88184,
        "Y": 532.8994
      },
      "LeftEar": {
        "Score": 0.86575276,
        "X": 255.99512,
        "Y": 164.10059
      },
      "LeftElbow": {
        "Score": 0.8731758,
        "X": 295.98535,
        "Y": 292.958
      },
      "LeftEye": {
        "Score": 0.89518315,
        "X": 247.1084,
        "Y": 155.21387
      },
      "LeftHip": {
        "Score": 0.84058046,
        "X": 264.88184,
        "Y": 364.05176
      },
      "LeftKnee": {
        "Score": 0.8632084,
        "X": 264.88184,
        "Y": 452.91895
      },
      "LeftMouthCorner": {
        "Score": 0.872159,
        "X": 242.66504,
        "Y": 172.9873
      },
      "LeftShoulder": {
        "Score": 0.8567772,
        "X": 282.65527,
        "Y": 221.86426
      },
      "LeftWrist": {
        "Score": 0.8648507,
        "X": 295.98535,
        "Y": 355.16504
      },
      "Neck": {
        "Score": 0.8947183,
        "X": 233.77832,
        "Y": 199.64746
      },
      "Nose": {
        "Score": 0.8996839,
        "X": 238.22168,
        "Y": 164.10059
      },
      "RightAnkle": {
        "Score": 0.9028398,
        "X": 211.56152,
        "Y": 537.3428
      },
      "RightEar": {
        "Score": 0.89784527,
        "X": 216.00488,
        "Y": 164.10059
      },
      "RightElbow": {
        "Score": 0.8529262,
        "X": 176.01465,
        "Y": 292.958
      },
      "RightEye": {
        "Score": 0.8819605,
        "X": 229.33496,
        "Y": 155.21387
      },
      "RightHip": {
        "Score": 0.8557911,
        "X": 211.56152,
        "Y": 364.05176
      },
      "RightKnee": {
        "Score": 0.86116123,
        "X": 207.11816,
        "Y": 457.3623
      },
      "RightMouthCorner": {
        "Score": 0.8761541,
        "X": 229.33496,
        "Y": 177.43066
      },
      "RightShoulder": {
        "Score": 0.8960459,
        "X": 189.34473,
        "Y": 221.86426
      },
      "RightWrist": {
        "Score": 0.8807962,
        "X": 180.45801,
        "Y": 355.16504
      },
      "TopHead": {
        "Score": 0.87812036,
        "X": 238.22168,
        "Y": 132.99707
      }
    },
    "BodyImage": "XXXXXXXXXXXXXXXXXXXXXXXXXX.jpg",
    "BodyShape": 0,
    "BodyShapeRisk": 3,
    "Bone": 3.08,
    "CaloriesInput": 1265,
    "CustomerID": "",
    "DeviceFrom": "XXXXXXXXXXXXXXX",
    "EnduGoal": 113.85,
    "Fat": 25.3,
    "FatControl": -4.59792,
    "FatLeftArm": 1.47,
    "FatLeftLeg": 4.13,
    "FatRightArm": 1.54,
    "FatRightLeg": 4.1,
    "FatTrunk": 12.23,
    "GmtCreate": "2022-02-12T05:15:47Z",
    "GmtModify": "2022-02-12T05:33:25Z",
    "GymID": "XXXXXXXXXXXXXXX",
    "GymName": "XXXXXXXXXXXXX",
    "HeartFun": 1,
    "Height": 161.4,
    "Hips": 84.041,
    "HumpbackRisk": 1,
    "Id": "XXXXXXXXXXXXXXX",
    "InFat": 4,
    "Language": "english",
    "LegRisk": "O3",
    "Muscle": 39.5,
    "MuscleControl": 1.53264,
    "MuscleLeftArm": 2.07,
    "MuscleLeftLeg": 6.56,
    "MuscleRightArm": 2.02,
    "MuscleRightLeg": 6.51,
    "MuscleTrunk": 18.82,
    "Name": "XXXXXXXXXXX",
    "PelvisRisk": 2,
    "PerfectWeight": 51.5722,
    "Phone": "XXXXXXXXXXX",
    "PostureRisk": 1,
    "Protein": 14.88,
    "R100LeftArm": 355.877,
    "R100LeftLeg": 271.359,
    "R100RightArm": 346.548,
    "R100RightLeg": 265.518,
    "R100Trunk": 27.666,
    "R20LeftArm": 392.268,
    "R20LeftLeg": 315.117,
    "R20RightArm": 384.314,
    "R20RightLeg": 309.573,
    "R20Trunk": 32.33,
    "RestingHeartRate": 109,
    "Score": 73,
    "Sex": "M",
    "ShoulderRisk": 1,
    "SideBodyDetect": {
      "BodyDirection": 2,
      "FullBody": true,
      "LeftAnkle": {
        "Score": 0.8364993,
        "X": 259.98633,
        "Y": 541.1621
      },
      "LeftEar": {
        "Score": 0.87245566,
        "X": 273.16992,
        "Y": 167.62695
      },
      "LeftElbow": {
        "Score": 0.84478235,
        "X": 255.5918,
        "Y": 290.67383
      },
      "LeftEye": {
        "Score": 0.8917769,
        "X": 246.80273,
        "Y": 163.23242
      },
      "LeftHip": {
        "Score": 0.7887097,
        "X": 251.19727,
        "Y": 365.38086
      },
      "LeftKnee": {
        "Score": 0.8201479,
        "X": 251.19727,
        "Y": 453.27148
      },
      "LeftMouthCorner": {
        "Score": 0.9146807,
        "X": 246.80273,
        "Y": 180.81055
      },
      "LeftShoulder": {
        "Score": 0.8711852,
        "X": 264.38086,
        "Y": 220.36133
      },
      "LeftWrist": {
        "Score": 0.8892825,
        "X": 229.22461,
        "Y": 352.19727
      },
      "Neck": {
        "Score": 0.87038887,
        "X": 259.98633,
        "Y": 193.99414
      },
      "Nose": {
        "Score": 0.87907565,
        "X": 238.01367,
        "Y": 167.62695
      },
      "RightAnkle": {
        "Score": 0.71923596,
        "X": 255.5918,
        "Y": 523.584
      },
      "RightEar": {
        "Score": 0.67461324,
        "X": 238.01367,
        "Y": 167.62695
      },
      "RightElbow": {
        "Score": 0.721023,
        "X": 255.5918,
        "Y": 290.67383
      },
      "RightEye": {
        "Score": 0.85266376,
        "X": 238.01367,
        "Y": 158.83789
      },
      "RightHip": {
        "Score": 0.75881135,
        "X": 251.19727,
        "Y": 360.98633
      },
      "RightKnee": {
        "Score": 0.76216495,
        "X": 251.19727,
        "Y": 448.87695
      },
      "RightMouthCorner": {
        "Score": 0.8999586,
        "X": 238.01367,
        "Y": 180.81055
      },
      "RightShoulder": {
        "Score": 0.75569737,
        "X": 259.98633,
        "Y": 220.36133
      },
      "RightWrist": {
        "Score": 0.6874447,
        "X": 238.01367,
        "Y": 352.19727
      },
      "TopHead": {
        "Score": 0.8848599,
        "X": 255.5918,
        "Y": 136.86523
      }
    },
    "SideImage": "XXXXXXXXXXXXXXXXXXXXXXXXXX.jpg",
    "SportGoal": 569.25,
    "SportLevel": 4,
    "SportSafeRisk": 3,
    "SubFat": 23.5,
    "UnitType": "metric",
    "UserID": "XXXXXXXXXXXXXXX",
    "VitalCapacity": 0,
    "Waist": 84.7749,
    "Water": 54.3,
    "Wc": 1.00873,
    "Weight": 55.8,
    "WeightControl": -3.06528
  },
  "BasicStandard": {
    "BmaxpStandard": [ 0, 40, 90, 115, 140, 190 ],
    "BmiStandard": [ 15, 15.352855, 17.999847, 19.79742, 21.990438, 27.987003 ],
    "BminpStandard": [ 0, 30, 60, 75, 90, 120 ],
    "BmrStandard": [ 0, 991.5734, 1487.36, 1652.6223, 1817.8846, 1983.1469 ],
    "BoneStandard": [ 0, 0, 2, 4, 6, 8 ],
    "BottomFat": 1.4228104521822784,
    "BottomMuscle": 7.264999713897705,
    "FatStandard": [ 15, 7.58, 12.79, 17.06, 22.67, 28 ],
    "HeartStandard": [ 0, 40, 60, 75, 80, 130 ],
    "HeightStandard": [ 15, 150.1, 163.3, 169.8, 176.3, 189.3 ],
    "MuscleStandard": [ 0, 33.340298, 42.866096, 47.628998, 52.3919, 80.9693 ],
    "ProteinStandard": [ 0, 11, 15, 17, 19, 23 ],
    "ReactionStandard": [ 15, 700, 650, 600, 550, 500 ],
    "SubfatStandard": [ 15, 6.8219995, 11.511, 15.353999, 20.403, 25.199999 ],
    "TopFat": 0.5139603873161316,
    "TopMuscle": 2.541599750518799,
    "TrunkFat": 3.5856129056131367,
    "TrunkMuscle": 20.57959918212891,
    "VisfatStandard": [ 0, -4, 0, 1, 5, 9 ],
    "WaterStandard": [ 0, 45, 55, 55, 65, 75 ],
    "WeightStandard": [ 0, 39.994125, 46.88953, 51.5722, 57.285004, 72.90603 ],
    "WhrStandard": [ 0, 0.65, 0.8, 0.85, 0.95, 1.1 ]
  },
  "WeightGrowthStandard": [
    [ 2, 9.06, 11.24, 12.54, 14.01, 17.54 ],
    [ 3, 10.61, 13.13, 14.65, 16.39, 20.64 ],
    [ 4, 12.01, 14.88, 16.64, 18.67, 23.73 ],
    [ 5, 13.5, 16.87, 18.98, 21.46, 27.85 ],
    [ 6, 14.74, 18.71, 21.26, 24.32, 32.57 ],
    [ 7, 16.01, 20.83, 24.06, 28.05, 39.5 ],
    [ 8, 17.33, 23.23, 27.33, 32.57, 48.57 ],
    [ 9, 18.53, 25.5, 30.46, 36.92, 57.3 ],
    [ 10, 19.81, 27.93, 33.74, 41.31, 65.08 ],
    [ 11, 21.41, 30.95, 37.69, 46.33, 72.39 ],
    [ 12, 23.37, 34.67, 42.49, 52.31, 80.35 ],
    [ 13, 26.21, 39.22, 48.08, 59.04, 89.42 ],
    [ 14, 30.4, 44.08, 53.37, 64.84, 96.8 ],
    [ 15, 34.59, 48, 57.08, 68.35, 100.29 ],
    [ 16, 37.67, 50.62, 59.35, 70.2, 101.25 ],
    [ 17, 39.58, 52.2, 60.68, 71.2, 101.39 ],
    [ 18, 40.65, 53.08, 61.4, 71.73, 101.36 ]
  ],
  "HeightGrowthStandard": [
    [ 2, 78.3, 85.1, 88.5, 92.1, 99.5 ],
    [ 3, 85.6, 93, 96.8, 100.7, 108.7 ],
    [ 4, 92.5, 100.2, 104.1, 108.2, 116.5 ],
    [ 5, 98.7, 107, 111.3, 115.7, 124.7 ],
    [ 6, 104.1, 113.1, 117.7, 122.4, 132.1 ],
    [ 7, 109.2, 119, 124, 129.1, 139.6 ],
    [ 8, 114.1, 124.6, 130, 135.5, 146.8 ],
    [ 9, 118.3, 129.6, 135.4, 141.2, 153.3 ],
    [ 10, 122, 134, 140.2, 146.4, 159.2 ],
    [ 11, 125.7, 138.7, 145.3, 152.1, 165.8 ],
    [ 12, 130, 144.6, 151.9, 159.4, 174.5 ],
    [ 13, 136.3, 151.8, 159.5, 167.3, 183 ],
    [ 14, 144.3, 158.7, 165.9, 173.1, 187.4 ],
    [ 15, 150.1, 163.3, 169.8, 176.3, 189.3 ],
    [ 16, 152.9, 165.4, 171.6, 177.8, 190.1 ],
    [ 17, 154, 166.3, 172.3, 178.4, 190.5 ],
    [ 18, 154.4, 166.6, 172.7, 178.7, 190.6 ]
  ]
}
```
