$(document).ready(function() {
    
    orgEdgexFoundry.devService.loadDevService()
})

    orgEdgexFoundry.devService = (function(){
        "use strict";

        function DevService(){
            this.allDevServices = []
        }

        DevService.prototype = {
            constructor:DevService,
            loadDevService: null,
            initDevSelectBox: null,
            getConfig: null,
            putConfig: null,
            renderConfig: null
        }

        var dev = new DevService()

        DevService.prototype.initDevSelectBox = function(allDevServices) {
            $("#device-service-select-bar").empty();
            var str = '';
            $.each(allDevServices,function(i,s){
                str += '<option value="' + s + '">' + s + '</option>';
            });
            $("#device-service-select-bar").append(str);
        }

        DevService.prototype.loadDevService = function(){
            $.ajax({
                url:'/edgex-core-metadata/api/v1/deviceservice',
                type:'GET',
                success:function(data){
                  $.each(data,function(i,s){
                    dev.allDevServices.push(s.name);
                    dev.initDevSelectBox(dev.allDevServices)
                  });
                }
            });
        }

        DevService.prototype.getConfig = function() {
            var devService = document.getElementById("device-service-select-bar").value;
            $.ajax({
                url:'/edgex-sys-mgmt-agent' + '/api/v1/config/' + devService,
                type:'GET',
                success: function(data){
                    dev.renderConfig(data)
                }
            })
        }

        DevService.prototype.renderConfig = function(data) {
            var dataRender = JSON.stringify(data, null, 4);
            document.getElementById("device-service-config-content").value = dataRender;
        }

        DevService.prototype.putConfig = function() {
            var devService = document.getElementById("device-service-select-bar").value;
            var dataPut = document.getElementById("device-service-config-content").value;
            $.ajax({
                url: '/api/v1/dev/config/devservice/' + devService,
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
        return dev
    })()