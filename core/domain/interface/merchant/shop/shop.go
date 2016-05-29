/**
 * Copyright 2015 @ z3q.net.
 * name : shop
 * author : jarryliu
 * date : 2016-05-29 11:12
 * description :
 * history :
 */
package shop

const (
	// 线上商店
	TypeOnlineShop = 1
	// 线下实体店
	TypeOfflineShop = 2
)

var (
	DefaultOnlineShop = OnlineShop{
		// 通讯地址
		Address: "",
		// 联系电话
		Tel: "",
		//别名,用于在商店域名
		Alias: "",
		//域名
		Host: "",
		//前台Logo
		Logo: "res/shop_logo.png",
		//首页标题
		IndexTitle: "",
		//子页面标题
		SubTitle: "",
		// Notice
		Notice: "",
	}

	DefaultOfflineShop = OfflineShop{
		// 联系电话
		Tel: "",
		// 通讯地址
		Address: "",
		// 经度
		Lng: 0,
		// 纬度
		Lat: 0,
		// 配送最大半径(公里)
		DeliverRadius: 0,
	}
)

type (
	IShop interface {
		GetDomainId() int

		// 商店类型
		Type() int

		// 获取值
		GetValue() Shop

		// 设置值
		SetValue(*Shop) error

		// 保存
		Save() (int, error)

		// 数据
		Data() *ShopDto
	}

	// 线上商城
	IOnlineShop interface {
		// 设置值
		SetShopValue(*OnlineShop) error

		// 获取值
		GetShopValue() OnlineShop
	}

	// 线下商店
	IOfflineShop interface {
		// 设置值
		SetShopValue(*OfflineShop) error

		// 获取值
		GetShopValue() OfflineShop

		// 获取经维度
		GetLngLat() (float64, float64)

		// 是否可以配送
		// 返回是否可以配送，以及距离(米)
		CanDeliver(lng, lat float64) (bool, int)

		// 是否可以配送
		// 返回是否可以配送，以及距离(米)
		CanDeliverTo(address string) (bool, int)
	}

	// 商店
	Shop struct {
		Id         int    `db:"id" pk:"yes" auto:"yes"`
		MerchantId int    `db:"mch_id"`
		ShopType   int    `db:"shop_type"`
		Name       string `db:"name"`
		State      int    `db:"state"`
		SortNumber int    `db:"sort_number"`
		CreateTime int64  `db:"create_time"`
	}

	// 商店数据传输对象
	ShopDto struct {
		Id         int
		MerchantId int
		ShopType   int
		Name       string
		// 线上/线下商店的数据
		Data interface{}
	}

	// 商城
	OnlineShop struct {
		// 商店编号
		ShopId int `db:"shop_id" pk:"yes" auto:"no"`
		// 通讯地址
		Address string `db:"addr"`
		// 联系电话
		Tel string `db:"tel"`

		//别名,用于在商店域名
		Alias string `db:"alias"`

		//域名
		Host string `db:"host"`

		//前台Logo
		Logo string `db:"logo"`

		//首页标题
		IndexTitle string `db:"index_tit"`

		//子页面标题
		SubTitle string `db:"sub_tit"`

		// Notice
		Notice string `db:"notice_html"`
	}

	// 门店
	OfflineShop struct {
		// 商店编号
		ShopId int `db:"shop_id" pk:"yes" auto:"no"`
		// 联系电话
		Tel string `db:"tel"`
		// 通讯地址
		Address string `db:"addr"`

		// 经度
		Lng float32 `db:"lng"`

		// 纬度
		Lat float32 `db:"lat"`

		// 配送最大半径(公里)
		DeliverRadius int `db:"deliver_radius"`
	}
)