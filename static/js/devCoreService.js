$(document).ready(function() {
    debugger
    orgEdgexFoundry.coreService.loadCoreService()
})

    orgEdgexFoundry.coreService = (function(){
        "use strict";

        function CoreService(){
            this.allCoreServices = ["edgex-core-metadata",'edgex-core-data','edgex-core-command','edgex-support-logging', 'edgex-support-notifications', 'edgex-support-scheduler', 'edgex-sys-mgmt-agent']
        }

        CoreService.prototype = {
            constructor:CoreService,
            restartService: null,
            loadCoreService: null,
            initCoreSelectBox: null,
            getConfig: null,
            putConfig: null,
            renderConfig: null,
        }

        var core = new CoreService()

        CoreService.prototype.initCoreSelectBox = function(allCoreServices) {
            $("#select-bar").empty();
            var str = '';
            $.each(allCoreServices,function(i,s){
                str += '<option value="' + s + '">' + s + '</option>';
            });
            $("#select-bar").append(str);
        }

        CoreService.prototype.loadCoreService = function(){
            core.initCoreSelectBox(core.allCoreServices)
        }

        CoreService.prototype.getConfig = function() {
            var coreService = document.getElementById("select-bar").value;
            $.ajax({
                url:'/' + coreService + '/api/v1/config',
                type:'GET',
                success: function(data){
                    core.renderConfig(data)
                }
            })
        }

        CoreService.prototype.renderConfig = function(data) {
            var dataRender = JSON.stringify(data, null, 4);
            document.getElementById("config-content").value = dataRender;
        }

        CoreService.prototype.putConfig = function() {
            var coreService = document.getElementById("select-bar").value;
            var dataPut = document.getElementById("config-content").value;
            $.ajax({
                url: '/api/v1/dev/config/coreservice/' + coreService,
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

        CoreService.prototype.restartService = function() {
            var coreService = document.getElementById("select-bar").value
            $.ajax({
                url: '/edgex-sys-mgmt-agent/api/v1/operation',
                type: 'POST',
                contentType: 'application/json',
                data: JSON.stringify({
                    "action":"restart",
                    "services":[coreService]
                }),
                success: function(data) {
                    if(data[0].Success){
                        alert("Restart core service successfully")
                    }else {
                        alert("fail to restart core service, please try again")
                    }
                },
            })
        }
        return core
    })()