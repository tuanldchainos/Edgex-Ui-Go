$(document).ready(function() {
    debugger
})
orgEdgexFoundry.userZigbeeService = (function(){
    "use strict";

    function UserZigbeeService(){
        this.ZigbeeService = 'edgex-device-zigbee'
    }

    UserZigbeeService.prototype = {
        constructor:UserZigbeeService,
        restartService: null,
        putConfig: null,
    }

    var userZigbee = new UserZigbeeService()


    UserZigbeeService.prototype.putConfig = function() {
        var Network_RequestTimeout = document.getElementById("request-timeout").value
        var Network_ResponseTimeout = document.getElementById("response-timeout").value
        var UpdateLastConnected = document.getElementById("update-last-connected").value
        $.ajax({
            url: '/api/v1/user/config/devservice/' + userZigbee.ZigbeeService,
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify({
                "Device": {
                    "UpdateLastConnected": UpdateLastConnected
                },
                "Driver": {
                    "Network_RequestTimeout": Network_RequestTimeout,
                    " Network_ResponseTimeout": Network_ResponseTimeout
                }
            }),
            success: function(data) {
                alert(data)
            },
            error: function() {
                alert("faile to update service config, please try again")
            }
        })
    }

    UserZigbeeService.prototype.restartService = function() {
        $.ajax({
            url: '/api/v1/user/restart/service',
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify({
                "action":"restart",
                "services":[userZigbee.ZigbeeService]
            }),
            success: function(data) {
                let dt = JSON.parse(data)

                if(dt[0].Success){
                    alert("Restart zigbee service successfully")
                }else {
                    alert("fail to restart zigbee service, please try again")
                }
            },
            error: function() {
                alert("fail to restart zigbee service, please try again")
            }
        })
    }

    return userZigbee
})()