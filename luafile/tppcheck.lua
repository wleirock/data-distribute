local m = require("mymodule")
function sendReport(url,param)
	res = m.sendHTTPPost(url,param,"application/json")
	return res
end
