package log

func LogString(code int) (logMsg string) {

	if val, ok := logcodes[code]; ok {
		logMsg = val
	} else {
		logMsg = "Some Issue Occurred"
	}
	return
}

var logcodes map[int]string = map[int]string{
	//1001-1050- SQL ERROR
	1001: "Database Query Failed",
	1002: "Data Source Error",
	1003: "Database Connection Error",
	1004: "Database Config Not Found",

	//1051-1110- REDIS ERROR
	1051: "Cache Connection Failed",
	1052: "Redis - Get Query Failed",
	1053: "Redis - Cache Write Failed",
	1054: "Cache Config Error",
	1055: "Cache Delete Error",

	//1101-1150- AWS ERROR
	1101: "AWS Authorization Failed",
	1102: "TLLMS Secret Not Found In AWS",
	1103: "Unable To Generate Token",
	1104: "Token Validation Failed",
	1105: "Redis Connection Details Not Found In Secret Manager",
	1106: "DB Connection Details Not Found In Secret Manager",

	//1501-1600- General ERROR
	1501: "Params Not Found",
	1502: "Data Not Found",
	1503: "Invalid Parameters Sent",
	1504: "Save API Request Failed",
	1505: "Get API Request Failed",
	1506: "Backend API Save Request Failed",
	1508: "Json Marshal/Unmarshal Error",
	1509: "config not found",
}
