package dns

import (
	"ddns-go/config"
	"time"
)

// DNS interface
type DNS interface {
	Init(conf *config.Config)
	// 添加或更新IPV4/IPV6记录
	AddUpdateDomainRecords() (domains config.Domains)
}

// RunTimer 定时运行
func RunTimer(delay time.Duration) {
	for {
		RunOnce()
		time.Sleep(delay)
	}
}

// RunOnce RunOnce
func RunOnce() {
	conf, err := config.GetConfigCache()
	if err != nil {
		return
	}

	var dnsSelected DNS
	switch conf.DNS.Name {
	case "alidns":
		dnsSelected = &Alidns{}
	case "dnspod":
		dnsSelected = &Dnspod{}
	case "cloudflare":
		dnsSelected = &Cloudflare{}
	case "huaweicloud":
		dnsSelected = &Huaweicloud{}
	default:
		dnsSelected = &Alidns{}
	}
	dnsSelected.Init(&conf)

	domains := dnsSelected.AddUpdateDomainRecords()
	config.ExecWebhook(&domains, &conf)
}
