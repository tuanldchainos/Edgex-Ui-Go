$(document).ready(function() {
    debugger
})
orgEdgexFoundry.userMqtt = (function(){
    "use strict";

    function UserMqtt(){
        this.Mqtt = 'mqtt-app'
    }

    UserMqtt.prototype = {
        constructor:UserMqtt,
        restartMqtt: null,
        connectMqtt: null,
    }

    var userMqttService = new UserMqtt()

    UserMqtt.prototype.connectMqtt = function() {
        var username = document.getElementById("client-name").value
        var password = document.getElementById("client-pass").value
        var clientId = document.getElementById("client-pass").value
        var qos = document.getElementById("connect-qos").value
        var autoReconnect = document.getElementById("connect-arc").value
        var retained = document.getElementById("connect-retained").value
        var keepAlive = document.getElementById("connect-kal").value
        var connectTimeout = document.getElementById("connect-ct").value
        var skipCert = document.getElementById("sercure-scv").value
        var key = document.getElementById("sercure-key").value
        var cert = document.getElementById("sercure-cert").value
        var pubHost = document.getElementById("publish-host").value
        var pubPort = document.getElementById("publish-port").value
        var pubPro = document.getElementById("publish-protocol").value
        var subHost = document.getElementById("subscribe-host").value
        var subPort = document.getElementById("subscribe-port").value
        var subPro = document.getElementById("subscribe-protocol").value
        var request = document.getElementById("topic-request").value
        var response = document.getElementById("topic-response").value

        var dataMqtt = {
            "Binding": {
                "SubscribeTopic": request,
                "PublishTopic": response
            },
            "MessageBus": {
                "SubscribeHost": {
                    "Host": pubHost,
                    "Port": pubPort,
                    "Protocol": pubPro
                },
                "PublishHost": {
                    "Host": subHost,
                    "Port": subPort,
                    "Protocol": subPro
                },
                "Optional": {
                    "Username": username,
                    "Password": password,
                    "ClientId": clientId,
                    "Qos": qos,
                    "KeepAlive": keepAlive,
                    "Retained": retained,
                    "AutoReconnect": autoReconnect,
                    "ConnectTimeout": connectTimeout,
                    "SkipCertVerify": skipCert,
                    "KeyPEMBlock": key,
                    "CertPEMBlock": cert
                }
            }
        }
        console.log(dataMqtt)
        $.ajax({
            url: '/api/v1/user/config/appservice/' + userMqttService.Mqtt,
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(dataMqtt),
            success: function(data) {
                alert(data)
            },
            error: function() {
                alert("faile to update service config, please try again")
            }
        })
    }

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
