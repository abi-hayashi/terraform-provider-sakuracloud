package api

/************************************************
  generated by IDE. for [DNSAPI]
************************************************/

import (
	"github.com/yamamoto-febc/libsacloud/sacloud"
)

/************************************************
   To support influent interface for Find()
************************************************/

func (api *DNSAPI) Reset() *DNSAPI {
	api.reset()
	return api
}

func (api *DNSAPI) Offset(offset int) *DNSAPI {
	api.offset(offset)
	return api
}

func (api *DNSAPI) Limit(limit int) *DNSAPI {
	api.limit(limit)
	return api
}

func (api *DNSAPI) Include(key string) *DNSAPI {
	api.include(key)
	return api
}

func (api *DNSAPI) Exclude(key string) *DNSAPI {
	api.exclude(key)
	return api
}

func (api *DNSAPI) FilterBy(key string, value interface{}) *DNSAPI {
	api.filterBy(key, value, false)
	return api
}

// func (api *DNSAPI) FilterMultiBy(key string, value interface{}) *DNSAPI {
// 	api.filterBy(key, value, true)
// 	return api
// }

func (api *DNSAPI) WithNameLike(name string) *DNSAPI {
	return api.FilterBy("Name", name)
}

func (api *DNSAPI) WithTag(tag string) *DNSAPI {
	return api.FilterBy("Tags.Name", tag)
}
func (api *DNSAPI) WithTags(tags []string) *DNSAPI {
	return api.FilterBy("Tags.Name", tags)
}

// func (api *DNSAPI) WithSizeGib(size int) *DNSAPI {
// 	api.FilterBy("SizeMB", size*1024)
// 	return api
// }

// func (api *DNSAPI) WithSharedScope() *DNSAPI {
// 	api.FilterBy("Scope", "shared")
// 	return api
// }

// func (api *DNSAPI) WithUserScope() *DNSAPI {
// 	api.FilterBy("Scope", "user")
// 	return api
// }

func (api *DNSAPI) SortBy(key string, reverse bool) *DNSAPI {
	api.sortBy(key, reverse)
	return api
}

func (api *DNSAPI) SortByName(reverse bool) *DNSAPI {
	api.sortByName(reverse)
	return api
}

// func (api *DNSAPI) SortBySize(reverse bool) *DNSAPI {
// 	api.sortBy("SizeMB", reverse)
// 	return api
// }

/************************************************
  To support CRUD(Create/Read/Update/Delete)
************************************************/

// func (api *DNSAPI) New() *sacloud.DNS {
// 	return &sacloud.DNS{}
// }

// func (api *DNSAPI) Create(value *sacloud.DNS) (*sacloud.DNS, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.create(api.createRequest(value), res)
// 	})
// }

// func (api *DNSAPI) Read(id string) (*sacloud.DNS, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.read(id, nil, res)
// 	})
// }

// func (api *DNSAPI) Update(id string, value *sacloud.DNS) (*sacloud.DNS, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.update(id, api.createRequest(value), res)
// 	})
// }

// func (api *DNSAPI) Delete(id string) (*sacloud.DNS, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.delete(id, nil, res)
// 	})
// }

/************************************************
  Inner functions
************************************************/

func (api *DNSAPI) setStateValue(setFunc func(*sacloud.Request)) *DNSAPI {
	api.baseAPI.setStateValue(setFunc)
	return api
}

//func (api *DNSAPI) request(f func(*sacloud.Response) error) (*sacloud.DNS, error) {
//	res := &sacloud.Response{}
//	err := f(res)
//	if err != nil {
//		return nil, err
//	}
//	return res.DNS, nil
//}
//
//func (api *DNSAPI) createRequest(value *sacloud.DNS) *dnsRequest {
//	req := &dnsRequest{}
//	req.CommonServiceDNSItem = value
//	return req
//}
