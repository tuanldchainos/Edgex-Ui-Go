package helper

import (
	"fmt"
	"sync"
	"time"

	"githup.com/tuanldchainos/Edgex-Ui-Go/internal/configs"
	"githup.com/tuanldchainos/Edgex-Ui-Go/internal/core"
	"githup.com/tuanldchainos/Edgex-Ui-Go/internal/handler"
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

func LoadServiceUri() {
	setServiceUri()
	updateServiceUri()
}

func setServiceUri() {
	CoreDataUri = core.HttpProtocol + "://" + configs.StaticProxyConf.CoreDataHost + ":" + configs.StaticProxyConf.CoreDataPort
	MetaDataUri = core.HttpProtocol + "://" + configs.StaticProxyConf.CoreMetadataHost + ":" + configs.StaticProxyConf.CoreMetadataPort
	CommandUri = core.HttpProtocol + "://" + configs.StaticProxyConf.CoreCommandHost + ":" + configs.StaticProxyConf.CoreCommandPort
	LoggingUri = core.HttpProtocol + "://" + configs.StaticProxyConf.SupportLoggingHost + ":" + configs.StaticProxyConf.SupportLoggingPort
	NotificationsUri = core.HttpProtocol + "://" + configs.StaticProxyConf.SupportNotificationHost + ":" + configs.StaticProxyConf.SupportNotificationPort
	SchedulerUri = core.HttpProtocol + "://" + configs.StaticProxyConf.SupportSchedulerHost + ":" + configs.StaticProxyConf.SupportSchedulerPort
	SystemAgentUri = core.HttpProtocol + "://" + configs.StaticProxyConf.SystemAgentHost + ":" + configs.StaticProxyConf.SystemAgentPort
}

func updateServiceUri() {
	for i := 0; ; i++ {
		time.Sleep(10 * time.Second)

		coreDataClient, _ := handler.InitRegistryClientByServiceKey(core.CoreDataServiceKey, true, core.ConfigCoreRegistryStem)
		CoreDataUri, _ = handler.GetServiceURLviaRegistry(coreDataClient, core.CoreDataServiceKey)

		coreMetaDataClient, _ := handler.InitRegistryClientByServiceKey(core.CoreMetaDataServiceKey, true, core.ConfigCoreRegistryStem)
		MetaDataUri, _ = handler.GetServiceURLviaRegistry(coreMetaDataClient, core.CoreMetaDataServiceKey)

		commandClient, _ := handler.InitRegistryClientByServiceKey(core.CoreCommandServiceKey, true, core.ConfigCoreRegistryStem)
		CommandUri, _ = handler.GetServiceURLviaRegistry(commandClient, core.CoreCommandServiceKey)

		loggingClient, _ := handler.InitRegistryClientByServiceKey(core.SupportLoggingServiceKey, true, core.ConfigCoreRegistryStem)
		LoggingUri, _ = handler.GetServiceURLviaRegistry(loggingClient, core.SupportLoggingServiceKey)

		notiClient, _ := handler.InitRegistryClientByServiceKey(core.SupportNotificationsServiceKey, true, core.ConfigCoreRegistryStem)
		NotificationsUri, _ = handler.GetServiceURLviaRegistry(notiClient, core.SupportNotificationsServiceKey)

		schedulerClient, _ := handler.InitRegistryClientByServiceKey(core.SupportSchedulerServiceKey, true, core.ConfigCoreRegistryStem)
		SchedulerUri, _ = handler.GetServiceURLviaRegistry(schedulerClient, core.SupportSchedulerServiceKey)

		agentClient, _ := handler.InitRegistryClientByServiceKey(core.SystemManagementAgentServiceKey, true, core.ConfigCoreRegistryStem)
		SystemAgentUri, _ = handler.GetServiceURLviaRegistry(agentClient, core.SystemManagementAgentServiceKey)

		fmt.Println(CoreDataUri)
	}

}
