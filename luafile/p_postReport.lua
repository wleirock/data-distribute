local m = require("mymodule")
function sendData(url,param)
	res = m.sendHTTPPost("http://"..url.."/index",param,"application/json")
	return res
end