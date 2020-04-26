$(document).ready(function(){
  $('[data-toggle="log-tooltip"]').tooltip();
  $('[data-toggle="eraser-tooltip"]').tooltip();
  supportLogging.loadAllDeviceServices();
  });
  
  supportLogging = (function(){
    "use strict";
  
    function SupportLogging(){
      this.allMicrosevices = ['edgex-core-metadata','edgex-core-data','edgex-core-command','edgex-support-logging'];
    }
  
    SupportLogging.prototype = {
      constructor:SupportLogging,
      loadLoggingBySearchService: null,
      loadLoggingBySearchKeyword: null,
      renderLoggingBySearch: null,
      searchBtn: null,
      eraseScreenBtn: null,
  
      loadAllDeviceServices: null,
      initLogMiscroseviceSelectPanel: null,
    }
  
    var logging = new SupportLogging();
  
    SupportLogging.prototype.initLogMiscroseviceSelectPanel = function(allMicrosevices){
      $("#edgex-support-logging-tab-main select[name='log_service']").empty();
      var row = '';
      $.each(allMicrosevices,function(i,s){
         row += '<option value="' + s + '">' + s + '</option>';
      });
      $("#edgex-support-logging-tab-main select[name='log_service']").append(row);
    }
  
    SupportLogging.prototype.loadAllDeviceServices = function(){
      $.ajax({
        url:'/edgex-core-metadata/api/v1/deviceservice',
        type:'GET',
        success:function(data){
          $.each(data,function(i,s){
            console.log(data)
            logging.allMicrosevices.push(s.name);
            logging.initLogMiscroseviceSelectPanel(logging.allMicrosevices);
          });
        }
      });
    }
  
    SupportLogging.prototype.eraseScreenBtn = function(){
        $("#log-content div.log_content").empty();
    }
  
    SupportLogging.prototype.searchBtn = function(){
      var service = $("select[name='log_service']").val();
      var keyword = document.getElementById("log_key_word").value;
      var start_str = document.getElementById("log_start_time").value;
      var end_str = document.getElementById("log_end_time").value;
      var limit = $("select[name='log_limit']").val();

      if (keyword === null){
        logging.loadLoggingBySearchService(service,start_str,end_str,limit);
      }else{
        logging.loadLoggingBySearchKeyword(keyword,start_str,end_str,limit);
      }
    }
  
    SupportLogging.prototype.loadLoggingBySearchService = function(service,start_timestamp,end_timestamp,limit){
      $.ajax({
        url:'/edgex-support-logging/api/v1/logs/originServices/'+service+'/'+start_timestamp+'/'+end_timestamp+'/' + limit,
        type:'GET',
        success:function(data){
          console.log(data)
          $("#log-content div.log_content").empty();
          if(data.length == 0) {
              $("#log-content div.log_content").append('<span style="color:white;">No data.</span>');
              return;
          }
          logging.renderLoggingBySearch(data);
        }
      });
    }

    SupportLogging.prototype.loadLoggingBySearchKeyword = function(keyword,start_timestamp,end_timestamp,limit){
        $.ajax({
          url:'/edgex-support-logging/api/v1/logs/keywords/'+service+'/'+start_timestamp+'/'+end_timestamp+'/' + limit,
          type:'GET',
          success:function(data){
            console.log(data)
            $("#log-content div.log_content").empty();
            if(data.length == 0) {
                $("#log-content div.log_content").append('<span style="color:white;">No data.</span>');
                return;
            }
            logging.renderLoggingBySearch(data);
          }
        });
      }
  
    SupportLogging.prototype.renderLoggingBySearch = function(data){
      $.each(data,function(i,v){
        var show_log = '<p>';
        if (v.logLevel == "ERROR") {
          show_log += '<span style="color:red;">'+v.logLevel+'</span>&nbsp;&nbsp;&nbsp;';
          show_log += '<span style="color:red;">'+ dateToString(v.created) + '</span>&nbsp;&nbsp;&nbsp;';
        } else {
          show_log += '<span style="color:green;">'+v.logLevel+'</span>&nbsp;&nbsp;&nbsp;';
          show_log += '<span style="color:green;">'+ dateToString(v.created) + '</span>&nbsp;&nbsp;&nbsp;';
        }
        show_log += '<span style="color:white;">'+ v.message + '</span>';
        show_log += '</p>'
        $("#log-content div.log_content").append(show_log);
      });
    }
    
    return logging;
})();
  