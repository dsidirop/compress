struct THVitalsTemplate {
	 1: optional THVitalsTemplateSpec  Spec = 1;
}

struct THVitalsTemplateSpec {
	 1: optional i64                   RespiratoryRate        ,
	 2: optional i64                   SpO2                   ,
	 3: optional THBloodPressure       BloodPressure          ,
	 4: optional i64                   EtCO2                  ,
	 5: optional double                TemperatureBlood       ,
	 6: optional double                TemperaturePeripheral  ,
	 7: optional THEcg                 Ecg                    ,
	 8: optional THEyes                Eyes                   ,
	 9: optional THAirwayObstruction   AirwayObstruction      ,
	10: optional THLungSounds          LungSounds             ,
	11: optional THHeartSounds         HeartSounds            ,
	12: optional THBowelSounds         BowelSounds            ,
	13: optional THPulses              Pulses                 
}

struct THBloodPressure {
	1: optional i64                    Systolic               ,
	2: optional i64                    Diastolic              
}

struct THEcg {
	1: optional i64                    HeartRate              ,
	2: optional string                 BasicRhythm            ,
	3: optional THExtrasystole         Extrasystole           ,
	4: optional i64                    Severity               ,
	5: optional bool                   AllowPacing            ,
	6: optional bool                   EmdPea                 
}

struct THExtrasystole {
	1: optional string                 Type                   ,
	2: optional i64                    Probability            
}

struct THEyes {
	1: optional THEye                  Left                   ,
	2: optional THEye                  Right                  ,
	3: optional string                 BlinkRate              
}

struct THHeartSounds {
	1: optional string                 Aortic                 ,
	2: optional i64                    AorticVolume           ,
	3: optional string                 Pulmonary              ,
	4: optional i64                    PulmonaryVolume        ,
	5: optional string                 Tricuspid              ,
	6: optional i64                    TricuspidVolume        ,
	7: optional string                 Mitral                 ,
	8: optional i64                    MitralVolume           
}


struct THEye {
	1: optional string                 EyelidPosition         ,
	2: optional i64                    PupilSize              ,
	3: optional string                 PupilResponsiveness    
}

struct THAirwayObstruction {
	1: optional string                 TongueEdema            
}

struct THLungSounds {
	1: optional THSingleLungSounds     Left                   ,
	2: optional THSingleLungSounds     Right                  
}

struct THstring {
	1: optional string                 Aortic                 ,
	2: optional i64                    AorticVolume           ,
	3: optional string                 Pulmonary              ,
	4: optional i64                    PulmonaryVolume        ,
	5: optional string                 Tricuspid              ,
	6: optional i64                    TricuspidVolume        ,
	7: optional string                 Mitral                 ,
	8: optional i64                    MitralVolume           
}

struct THSingleLungSounds {
	1: optional string                 AnteriorUpper          ,
	2: optional i64                    AnteriorUpperVolume    ,
	3: optional string                 AnteriorLower          ,
	4: optional i64                    AnteriorLowerVolume    ,
	5: optional string                 AnteriorMiddle         ,
	6: optional i64                    AnteriorMiddleVolume   ,
	7: optional string                 PosteriorUpper         ,
	8: optional i64                    PosteriorUpperVolume   ,
	9: optional string                 PosteriorLower         ,
   10: optional i64                    PosteriorLowerVolume   
}

struct THBowelSounds {
	1: optional string                 Bowel                  ,
	2: optional i64                    BowelVolume            
}

struct THPulses {
     1: optional string                Central               ,
     2: optional string                CentralLimited        ,
     3: optional string                RightLeg              ,
     4: optional string                RightLegLimited       ,
     5: optional string                LeftLeg               ,
     6: optional string                LeftLegLimited        ,
     7: optional string                RightFoot             ,
     8: optional string                RightFootLimited      ,
     9: optional string                LeftFoot              ,
    10: optional string                LeftFootLimited       , 
    11: optional string                RightArm              , 
    12: optional string                RightArmLimited       , 
    13: optional string                LeftArm               , 
    14: optional string                LeftArmLimited        , 
    15: optional string                RightHand             , 
    16: optional string                RightHandLimited      , 
    17: optional string                LeftHand              , 
    18: optional string                LeftHandLimited        
}
