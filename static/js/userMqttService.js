$(document).ready(function() {
    debugger
})
orgEdgexFoundry.userMqtt = (function(){
    "use strict";

    function UserMqtt(){
        this.Mqtt = 'edgex-core-data'
    }

    UserMqtt.prototype = {
        constructor:UserMqtt,
        restartMqtt: null,
        connectMqtt: null,
    }

    var userMqttService = new UserMqtt()

    // UserMqtt.prototype.connectMqtt = function() {
    //     var host = document.getElementById

    //     var dataMqtt = {

    //     }
    //     $.ajax({
    //         url: '/api/v1/user/config/appservice/' + userMqttService.Mqtt,
    //         type: 'POST',
    //         contentType: 'application/json',
    //         data: ,
    //         success: function(data) {
    //             alert(data)
    //         },
    //         error: function() {
    //             alert("faile to update service config, please try again")
    //         }
    //     })
    // }

    UserMqtt.prototype.restartMqtt = function() {
        $.ajax({
            url: '/api/v1/user/restart/service',
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify({
                "action":"restart",
                "services":[userMqttService.Mqtt]
            }),
            success: function(data) {
                let dt = JSON.parse(data)
                if(dt[0].Success){
                    alert("Restart mqtt service successfully")
                }else {
                    alert("fail to restart mqtt service, please try again")
                }
            },
            error: function() {
                alert("fail to restart mqtt service, please try again")
            }
        })
    }

    return userMqttService
})()
