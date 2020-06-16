$(document).ready(function() {
    debugger
})
orgEdgexFoundry.userZigbeeService = (function(){
    "use strict";

    function UserZigbeeService(){
        this.ZigbeeService = 'device-zigbee'
    }

    UserZigbeeService.prototype = {
        constructor:UserZigbeeService,
        restartService: null,
        getConfig: null,
        putConfig: null,
        renderConfig: null
    }

    var userZigbee = new UserZigbeeService()


    UserZigbeeService.prototype.putConfig = function() {
        var NetworkType = document.getElementById("network-type").value
        var RequestTimeout = document.getElementById("request-timeout").value
        var CommandResponseTimeout = document.getElementById("command-response-timeout").value
        var ObjectResponseTimeout = document.getElementById("object-response-timeout").value
        $.ajax({
            url: '/api/v1/user/config/devservice/' + userZigbee.ZigbeeService,
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify({
                "Driver": {
                    "NetworkType": NetworkType,
                    "NetWorkRequestTimeout": RequestTimeout, 
                    "NetworkCommandResponseTimeout": CommandResponseTimeout,
                    "NetworkAddObjectResponseTimeout": ObjectResponseTimeout
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

    UserZigbeeService.prototype.getConfig = function() {
        $.ajax({
            url: '/api/v1/user/config/devservice/' + userZigbee.ZigbeeService,
            type: 'GET',
            dataType: 'json',
            success: function(data) {
                let dt = data.Driver
                userZigbee.renderConfig(dt)
                alert("Get current device zigbee config successfully")
            }
        })
    }

    UserZigbeeService.prototype.renderConfig = function(dt) {
        document.getElementById("network-type").value = dt.NetworkType
        document.getElementById("request-timeout").value = dt.NetWorkRequestTimeout
        document.getElementById("command-response-timeout").value = dt.NetworkCommandResponseTimeout
        document.getElementById("object-response-timeout").value = dt.NetworkAddObjectResponseTimeout
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