package log

//日志
const (
	LOG_DEFAULT_LEVEL  = "trace"
	LOG_DEFAULT_FORMAT = "%Date/%Time [%LEV] %Msg%n"
	LOG_DEFAULT_PATH   = "./tmp/dhaiy.log"
	LOG_DEFAULT_ROLL   = "2006-01-02"
)

//获取日志配置
func getLogConfig(level string, format string, path string, roll string, consoleOn bool) string {
	config := `
	<seelog>
	    <outputs formatid="main">`
	if consoleOn {
		config += `
			<filter levels="` + level + `">
	        	<console />
	    	</filter>
		`
	}
	config += `
			<filter levels="` + level + `">
				<rollingfile type="date" filename="` + path + `" datepattern="` + roll + `" maxrolls="7" />
	        </filter>
	    </outputs>
	    <formats>
	        <format id="main" format="` + format + `"/>
	    </formats>
	</seelog>
	`
	return config
}
