package storefilefixtures

import storefile "github.com/TudorHulban/authentication/infra/stores/store-file"

var ParamsStoreFileTickets storefile.ParamsNewStoreTickets = storefile.ParamsNewStoreTickets{
	PathCacheTickets: "local_cache_ticket.json",
	PathCacheEvent:   "local_cache_event.json",
}

var ParamsStoreFileUsers storefile.ParamsNewStoreUsers = storefile.ParamsNewStoreUsers{
	PathCacheUsers: "local_cache_user.json",
}
