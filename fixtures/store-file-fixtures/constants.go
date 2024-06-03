package storefilefixtures

import storefile "github.com/TudorHulban/authentication/infra/stores/store-file"

var ParamsStoreFile storefile.ParamsNewStoreTask = storefile.ParamsNewStoreTask{
	PathCacheTask:  "local_cache_task.json",
	PathCacheEvent: "local_cache_event.json",
}
