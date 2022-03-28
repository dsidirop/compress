struct THCurveGenReplyV1 {
	1:    string                        ClientId ,
	2:    THCurveSpecificationV1        Spec     ,
	3:    list<i64>                     Lead1    ,
	4:    list<i64>                     Lead2    ,
	5:    list<i64>                     Lead3    ,
	6:    list<i64>                     Lead4    ,
	7:    list<i64>                     Lead5    ,
	8:    list<i64>                     Lead6    ,
	9:    list<i64>                     Lead7    ,
	10:   list<i64>                     Lead8    ,
	11:   list<i64>                     Lead9    ,
	12:   list<i64>                     Lead10   ,
	13:   list<i64>                     Lead11   ,
	14:   list<i64>                     Lead12   ,
	15:   list<i64>                     Abp      ,
	16:   list<i64>                     Cvp      ,
	17:   list<i64>                     Pap      ,
	18:   list<i64>                     Spo2     ,
	19:   list<i64>                     Wp       ,
	20:   list<THTag>                   Tags     
}

struct THCurveSpecificationV1 {
	1:    string                        SimulatorName  ,
	2:    string                        Tenant         ,
	3:    i64                           StartTime      ,
	4:    i64                           EndTime        ,
	5:    i64                           SampleInterval ,
	6:    i64                           CurveTypes     
}

struct THTag {
	1:    i64                           TagTime,
	2:    i64                           TagType 
}
