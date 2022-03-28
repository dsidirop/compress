typedef string Timestamp # in golang we can grab the utc representation of time as    time.Now().UTC().MarshalText()

struct THSimEventRegisterEventCmd {
	1:            string                 ApiVersion    ,
	2:            string                 SimulatorName ,
	3:            string                 Tenant        ,
	4:            string                 EventID       ,
	5:            Timestamp              TimeStamp     ,
	6: optional   THSimEventArguments    Arguments     
}

union THSimEventArguments {
  1: THPacingArguments        PacingArguments,
  // possibly more later on ...
}

struct THPacingArguments {
  1: Timestamp                TimeStamp,
  2: i32                      MeasuredCurrent,
  3: bool                     Capture,
}
