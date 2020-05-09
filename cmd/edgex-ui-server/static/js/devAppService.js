$(document).ready(function() {
    debugger
    orgEdgexFoundry.appService.loadAppService()
})

    orgEdgexFoundry.appService = (function(){
        "use strict";

        function AppService(){
            this.allAppServices = []
        }

        AppService.prototype = {
            constructor:AppService,
            restartService: null,
            loadAppService: null,
            initAppSelectBox: null,
            getConfig: null,
            putConfig: null,
            renderConfig: null,
        }

        var app = new AppService()

        AppService.prototype.initAppSelectBox = function(allAppServices) {
            $("#app-service-select-bar").empty();
            var str = '';
            $.each(allAppServices,function(i,s){
                str += '<option value="' + s + '">' + s + '</option>';
            });
            $("#app-service-select-bar").append(str);
        }

        AppService.prototype.loadAppService = function(){
            $.ajax({
                url: '/api/v1/dev/appservice/list',
                type: 'GET',
                success:function(services){
                    $.each(services,function (key,value) {
                        app.allAppServices.push(key);
                        app.initAppSelectBox(app.allAppServices)
                    });
                }
            });
        }

        AppService.prototype.getConfig = function() {
            var appService = document.getElementById("app-service-select-bar").value;
            $.ajax({
                // url:'/edgex-sys-mgmt-agent' + '/api/v1/config/' + appService,
                url:'/api/v1/dev/appservice/list',
                type:'GET',
                success: function(data){
                    app.renderConfig(data[appService])
                }
            })
        }

        AppService.prototype.renderConfig = function(data) {
            var dataRender = JSON.stringify(data, null, 4);
            document.getElementById("app-service-config-content").value = dataRender;
        }

        AppService.prototype.putConfig = function() {
            var appService = document.getElementById("app-service-select-bar").value;
            var dataPut = document.getElementById("app-service-config-content").value;
            $.ajax({
                url: '/api/v1/dev/config/appservice/' + appService,
                type: 'POST',
                contentType: 'application/json',
                data: dataPut,
                success: function(data) {
                    alert(data)
                },
                error: function() {
                    alert("faile to update service config, please try again")
                }
            })
        }

        AppService.prototype.restartService = function() {
            var appService = document.getElementById("app-service-select-bar").value
            $.ajax({
                url: '/edgex-sys-mgmt-agent/api/v1/operation',
                type: 'POST',
                contentType: 'application/json',
                data: JSON.stringify({
                    "action":"restart",
                    "services":[appService]
                }),
                success: function(data) {
                    if(data[0].Success){
                        alert("Restart app service successfully")
                    }else {
                        alert("fail to restart app service, please try again")
                    }
                },
            })
        }
        return app
    })()