package api

/************************************************
  generated by IDE. for [PublicPriceAPI]
************************************************/

import (
	"github.com/yamamoto-febc/libsacloud/sacloud"
)

/************************************************
   To support influent interface for Find()
************************************************/

func (api *PublicPriceAPI) Reset() *PublicPriceAPI {
	api.reset()
	return api
}

func (api *PublicPriceAPI) Offset(offset int) *PublicPriceAPI {
	api.offset(offset)
	return api
}

func (api *PublicPriceAPI) Limit(limit int) *PublicPriceAPI {
	api.limit(limit)
	return api
}

func (api *PublicPriceAPI) Include(key string) *PublicPriceAPI {
	api.include(key)
	return api
}

func (api *PublicPriceAPI) Exclude(key string) *PublicPriceAPI {
	api.exclude(key)
	return api
}

func (api *PublicPriceAPI) FilterBy(key string, value interface{}) *PublicPriceAPI {
	api.filterBy(key, value, false)
	return api
}

// func (api *PublicPriceAPI) FilterMultiBy(key string, value interface{}) *PublicPriceAPI {
// 	api.filterBy(key, value, true)
// 	return api
// }

func (api *PublicPriceAPI) WithNameLike(name string) *PublicPriceAPI {
	return api.FilterBy("Name", name)
}

func (api *PublicPriceAPI) WithTag(tag string) *PublicPriceAPI {
	return api.FilterBy("Tags.Name", tag)
}
func (api *PublicPriceAPI) WithTags(tags []string) *PublicPriceAPI {
	return api.FilterBy("Tags.Name", tags)
}

// func (api *PublicPriceAPI) WithSizeGib(size int) *PublicPriceAPI {
// 	api.FilterBy("SizeMB", size*1024)
// 	return api
// }

// func (api *PublicPriceAPI) WithSharedScope() *PublicPriceAPI {
// 	api.FilterBy("Scope", "shared")
// 	return api
// }

// func (api *PublicPriceAPI) WithUserScope() *PublicPriceAPI {
// 	api.FilterBy("Scope", "user")
// 	return api
// }

func (api *PublicPriceAPI) SortBy(key string, reverse bool) *PublicPriceAPI {
	api.sortBy(key, reverse)
	return api
}

func (api *PublicPriceAPI) SortByName(reverse bool) *PublicPriceAPI {
	api.sortByName(reverse)
	return api
}

// func (api *PublicPriceAPI) SortBySize(reverse bool) *PublicPriceAPI {
// 	api.sortBy("SizeMB", reverse)
// 	return api
// }

/************************************************
  To support CRUD(Create/Read/Update/Delete)
************************************************/

// func (api *PublicPriceAPI) New() *sacloud.PublicPrice {
// 	return &sacloud.PublicPrice{}
// }

// func (api *PublicPriceAPI) Create(value *sacloud.PublicPrice) (*sacloud.PublicPrice, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.create(api.createRequest(value), res)
// 	})
// }

// func (api *PublicPriceAPI) Read(id string) (*sacloud.PublicPrice, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.read(id, nil, res)
// 	})
// }

// func (api *PublicPriceAPI) Update(id string, value *sacloud.PublicPrice) (*sacloud.PublicPrice, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.update(id, api.createRequest(value), res)
// 	})
// }

// func (api *PublicPriceAPI) Delete(id string) (*sacloud.PublicPrice, error) {
// 	return api.request(func(res *sacloud.Response) error {
// 		return api.delete(id, nil, res)
// 	})
// }

/************************************************
  Inner functions
************************************************/

func (api *PublicPriceAPI) setStateValue(setFunc func(*sacloud.Request)) *PublicPriceAPI {
	api.baseAPI.setStateValue(setFunc)
	return api
}

//func (api *PublicPriceAPI) request(f func(*sacloud.Response) error) (*sacloud.PublicPrice, error) {
//	res := &sacloud.Response{}
//	err := f(res)
//	if err != nil {
//		return nil, err
//	}
//	return res.ServiceClass, nil
//}
//
//func (api *PublicPriceAPI) createRequest(value *sacloud.PublicPrice) *sacloud.Request {
//	req := &sacloud.Request{}
//	req.ServiceClass = value
//	return req
//}
