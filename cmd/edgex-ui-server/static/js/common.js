
var dateToString = function (num){
	var date = new Date();
	date.setTime(num);
	var y = date.getFullYear();
	var M = date.getMonth() + 1;
	M = (M < 10) ? ('0' + M) : M ;
	var d = date.getDate();
	d = (d < 10) ? ('0' + d) : d ;
	var hh = date.getHours();
	hh = (hh < 10 )? ('0' + hh) : hh ;
	var mm = date.getMinutes();
	mm = (mm < 10 )? ('0' + mm) : mm ;
	var ss = date.getSeconds();
	ss = (ss < 10) ?('0' + ss) : ss ;

	var str = y + '-' + M + '-' + d + ' ' + hh + ':' + mm + ':' + ss
	return str;
};

Date.prototype.Format = function (fmt) {
    var o = {
        "M+": this.getMonth() + 1, // Month
        "d+": this.getDate(), // Day
        "h+": this.getHours(), // Hour
        "m+": this.getMinutes(), // Minute
        "s+": this.getSeconds(), // Second
        "q+": Math.floor((this.getMonth() + 3) / 3), // Quarter
        "S": this.getMilliseconds() // Millisecond
    };
    if (/(y+)/.test(fmt))
        fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
    for (var k in o)
        if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
    return fmt;
};

//ISO 8601 format (YYYYMMDD'T'HHmmss)
var ISO8601FormatToDate = function (dateString) {
    var year = dateString.substring(0,4);
    var month = dateString.substring(4,6);
    var day = dateString.substring(6,8);
    var hour = dateString.substring(9,11);
    var minute = dateString.substring(11,13);
    var second = dateString.substring(13);
    return new Date(year,month,day,hour,minute,second);
};
