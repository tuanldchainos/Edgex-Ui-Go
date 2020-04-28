
function EdgexFoundry(){
  this.coreData = null;
  this.coreMetadata = null;
  this.coreCommand = null;
  this.coreExport = null;
  this.supportLogging = null;
  this.supportNotification= null;
  this.supportScheduler= null;
  this.supportRuleEngine = null;
  this.deviceService = null;
  this.appService = null;
  this.utils = null;
}

EdgexFoundry.prototype = {
  constructor: EdgexFoundry,
}
var orgEdgexFoundry =  new EdgexFoundry();

bootbox.addLocale('my-locale', {
  OK: 'OK',
  CANCEL : 'No',
  CONFIRM : 'Yes'
});

bootbox.setDefaults({ locale: 'my-locale' });
