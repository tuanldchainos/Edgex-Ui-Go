package helper

import (
	"sync"
	"time"

	"githup.com/tuanldchainos/Edgex-Ui-Go/internal/configs"
	"githup.com/tuanldchainos/Edgex-Ui-Go/internal/registrySupport"
)

var (
	MetaDataUri      string
	CoreDataUri      string
	CommandUri       string
	SchedulerUri     string
	NotificationsUri string
	LoggingUri       string
	SystemAgentUri   string
)

var wg sync.WaitGroup
var mux sync.Mutex

func LoadServiceUri() {
	setServiceUri()
	go updateServiceUri()
}

func setServiceUri() {
	CoreDataUri = HttpProtocol + "://" + configs.StaticProxyConf.CoreDataHost + ":" + configs.StaticProxyConf.CoreDataPort
	MetaDataUri = HttpProtocol + "://" + configs.StaticProxyConf.CoreMetadataHost + ":" + configs.StaticProxyConf.CoreMetadataPort
	CommandUri = HttpProtocol + "://" + configs.StaticProxyConf.CoreCommandHost + ":" + configs.StaticProxyConf.CoreCommandPort
	LoggingUri = HttpProtocol + "://" + configs.StaticProxyConf.SupportLoggingHost + ":" + configs.StaticProxyConf.SupportLoggingPort
	NotificationsUri = HttpProtocol + "://" + configs.StaticProxyConf.SupportNotificationHost + ":" + configs.StaticProxyConf.SupportNotificationPort
	SchedulerUri = HttpProtocol + "://" + configs.StaticProxyConf.SupportSchedulerHost + ":" + configs.StaticProxyConf.SupportSchedulerPort
	SystemAgentUri = HttpProtocol + "://" + configs.StaticProxyConf.SystemAgentHost + ":" + configs.StaticProxyConf.SystemAgentPort
}

func updateServiceUri() {
	for i := 0; ; i++ {
		time.Sleep(30 * time.Second)

		coreDataClient, _ := registrySupport.InitRegistryClientByServiceKey(CoreDataServiceKey, true, ConfigCoreRegistryStem)
		CoreDataUri, _ = registrySupport.GetServiceURLviaRegistry(coreDataClient, CoreDataServiceKey)

		coreMetaDataClient, _ := registrySupport.InitRegistryClientByServiceKey(CoreMetaDataServiceKey, true, ConfigCoreRegistryStem)
		MetaDataUri, _ = registrySupport.GetServiceURLviaRegistry(coreMetaDataClient, CoreMetaDataServiceKey)

		commandClient, _ := registrySupport.InitRegistryClientByServiceKey(CoreCommandServiceKey, true, ConfigCoreRegistryStem)
		CommandUri, _ = registrySupport.GetServiceURLviaRegistry(commandClient, CoreCommandServiceKey)

		loggingClient, _ := registrySupport.InitRegistryClientByServiceKey(SupportLoggingServiceKey, true, ConfigCoreRegistryStem)
		LoggingUri, _ = registrySupport.GetServiceURLviaRegistry(loggingClient, SupportLoggingServiceKey)

		notiClient, _ := registrySupport.InitRegistryClientByServiceKey(SupportNotificationsServiceKey, true, ConfigCoreRegistryStem)
		NotificationsUri, _ = registrySupport.GetServiceURLviaRegistry(notiClient, SupportNotificationsServiceKey)

		schedulerClient, _ := registrySupport.InitRegistryClientByServiceKey(SupportSchedulerServiceKey, true, ConfigCoreRegistryStem)
		SchedulerUri, _ = registrySupport.GetServiceURLviaRegistry(schedulerClient, SupportSchedulerServiceKey)

		agentClient, _ := registrySupport.InitRegistryClientByServiceKey(SystemManagementAgentServiceKey, true, ConfigCoreRegistryStem)
		SystemAgentUri, _ = registrySupport.GetServiceURLviaRegistry(agentClient, SystemManagementAgentServiceKey)
	}
}
