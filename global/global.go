package global

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"mini_shop/config"
	"mini_shop/model"
)

var DBClient *gorm.DB

func InitMysql() {
	mysqlConfig := config.AppConfig.Database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Name)
	client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("连接数据库失败", err)
	}
	DBClient = client
	log.Println("连接数据库成功")

	// 自动创建表（根据model下定义模型）
	if err := client.AutoMigrate(&model.User{}); err != nil {
		log.Fatalln("创建 users 表失败: ", err)
	}
	if err := client.AutoMigrate(&model.Product{}); err != nil {
		log.Fatalln("创建 products 表失败: ", err)
	}
	if err := client.AutoMigrate(&model.CartItem{}); err != nil {
		log.Fatalln("创建 cartItems 表失败: ", err)
	}
	if err := client.AutoMigrate(&model.Order{}); err != nil {
		log.Fatalln("创建 order 表失败: ", err)
	}
	if err := client.AutoMigrate(&model.OrderItem{}); err != nil {
		log.Fatalln("创建 orderItems 表失败: ", err)
	}

}

func InitProducts() {
	db := GetDB()
	var count int64
	db.Model(&model.Product{}).Count(&count)
	if count == 0 {
		products := []model.Product{
			// 手机
			{Name: "苹果 iPhone 15 Pro", Description: "A17 Pro 芯片，钛金属边框，影像系统升级", Price: 10999.00, Stock: 80, ImageURL: "/images/iphone15pro.jpg", Status: "on_sale"},
			{Name: "三星 Galaxy S24 Ultra", Description: "200MP 主摄，支持夜景拍摄", Price: 11999.00, Stock: 60, ImageURL: "/images/galaxy_s24ultra.jpg", Status: "on_sale"},
			{Name: "小米 14 Pro", Description: "徕卡影像系统，2K 屏幕，性能旗舰", Price: 6999.00, Stock: 100, ImageURL: "/images/xiaomi14pro.jpg", Status: "on_sale"},
			{Name: "Google Pixel 8", Description: "原生安卓系统，AI 拍照增强", Price: 7999.00, Stock: 70, ImageURL: "/images/pixel8.jpg", Status: "on_sale"},

			// 笔记本电脑
			{Name: "苹果 MacBook Air M3", Description: "轻薄便携，M3 芯片，续航 18 小时", Price: 12999.00, Stock: 50, ImageURL: "/images/macbook_air_m3.jpg", Status: "on_sale"},
			{Name: "联想 ThinkPad X1 Carbon", Description: "碳纤维机身，商务旗舰笔电", Price: 15999.00, Stock: 40, ImageURL: "/images/thinkpad_x1.jpg", Status: "on_sale"},
			{Name: "华硕 ROG 幻 14", Description: "RTX4060 显卡，电竞性能本", Price: 18999.00, Stock: 30, ImageURL: "/images/rog_g14.jpg", Status: "on_sale"},
			{Name: "惠普 Spectre x360", Description: "可翻转触控屏笔电，OLED 屏幕", Price: 14999.00, Stock: 35, ImageURL: "/images/hp_spectre.jpg", Status: "on_sale"},

			// 耳机音响
			{Name: "索尼 WH-1000XM5", Description: "头戴式蓝牙降噪耳机", Price: 3999.00, Stock: 90, ImageURL: "/images/sony_wh1000xm5.jpg", Status: "on_sale"},
			{Name: "苹果 AirPods Pro 2", Description: "主动降噪无线耳机，音质升级", Price: 2499.00, Stock: 200, ImageURL: "/images/airpods_pro2.jpg", Status: "on_sale"},
			{Name: "Bose QuietComfort Ultra", Description: "旗舰级降噪耳机，沉浸式音效", Price: 3799.00, Stock: 75, ImageURL: "/images/bose_qc_ultra.jpg", Status: "on_sale"},
			{Name: "Marshall Emberton II", Description: "便携蓝牙音箱，复古设计", Price: 1799.00, Stock: 110, ImageURL: "/images/marshall_emberton2.jpg", Status: "on_sale"},

			// 家电
			{Name: "戴森 V15 Detect", Description: "无线吸尘器，激光除尘技术", Price: 6999.00, Stock: 65, ImageURL: "/images/dyson_v15.jpg", Status: "on_sale"},
			{Name: "米家空气净化器 4", Description: "智能空气净化器，APP控制", Price: 1999.00, Stock: 120, ImageURL: "/images/mijia_air4.jpg", Status: "on_sale"},
			{Name: "飞利浦 HD9860 空气炸锅", Description: "健康无油炸，智能控温", Price: 3299.00, Stock: 90, ImageURL: "/images/philips_airfryer.jpg", Status: "on_sale"},
			{Name: "松下 EH-NA9 吹风机", Description: "纳米水离子护发技术", Price: 1999.00, Stock: 150, ImageURL: "/images/panasonic_hairdryer.jpg", Status: "on_sale"},

			// 游戏设备
			{Name: "任天堂 Switch OLED", Description: "掌上游戏机，7 吋 OLED 屏幕", Price: 3499.00, Stock: 90, ImageURL: "/images/switch_oled.jpg", Status: "on_sale"},
			{Name: "索尼 PlayStation 5", Description: "次世代主机，4K HDR 支持", Price: 4999.00, Stock: 50, ImageURL: "/images/ps5.jpg", Status: "on_sale"},
			{Name: "微软 Xbox Series X", Description: "1TB SSD，高速加载", Price: 4999.00, Stock: 45, ImageURL: "/images/xbox_seriesx.jpg", Status: "on_sale"},
			{Name: "Steam Deck 掌机", Description: "掌上 PC 游戏机，支持 SteamOS", Price: 5499.00, Stock: 60, ImageURL: "/images/steamdeck.jpg", Status: "on_sale"},

			// 外设配件
			{Name: "罗技 MX Master 3S 鼠标", Description: "人体工学设计，静音快充", Price: 999.00, Stock: 180, ImageURL: "/images/mx_master3s.jpg", Status: "on_sale"},
			{Name: "Keychron K8 Pro 键盘", Description: "机械键盘，RGB 背光", Price: 1299.00, Stock: 140, ImageURL: "/images/keychron_k8pro.jpg", Status: "on_sale"},
			{Name: "Anker 100W 氮化镓充电器", Description: "多设备快充，便携高效", Price: 899.00, Stock: 160, ImageURL: "/images/anker_gan100w.jpg", Status: "on_sale"},
			{Name: "三星 T7 便携固态硬盘 1TB", Description: "防震防水，高速传输", Price: 1399.00, Stock: 130, ImageURL: "/images/samsung_t7.jpg", Status: "on_sale"},

			// 生活用品
			{Name: "Kindle Paperwhite", Description: "电子书阅读器，防水设计", Price: 1399.00, Stock: 130, ImageURL: "/images/kindle_paperwhite.jpg", Status: "on_sale"},
			{Name: "Instant Pot 多功能电压力锅", Description: "智能烹饪，7 合 1 功能", Price: 1199.00, Stock: 100, ImageURL: "/images/instantpot_duo.jpg", Status: "on_sale"},
			{Name: "GoPro Hero 12", Description: "运动相机，5.3K 视频录制", Price: 3999.00, Stock: 70, ImageURL: "/images/gopro_hero12.jpg", Status: "on_sale"},
			{Name: "DJI Mini 4 Pro 无人机", Description: "轻型无人机，支持4K HDR航拍", Price: 9999.00, Stock: 30, ImageURL: "/images/dji_mini4pro.jpg", Status: "on_sale"},
			{Name: "Fitbit Charge 6", Description: "智能手环，健康监测", Price: 1499.00, Stock: 150, ImageURL: "/images/fitbit_charge6.jpg", Status: "on_sale"},
			{Name: "苹果手表 Series 9", Description: "健康检测与运动追踪", Price: 3999.00, Stock: 90, ImageURL: "/images/apple_watch9.jpg", Status: "on_sale"},
		}

		db.Create(&products)
	}
}

func GetDB() *gorm.DB {
	return DBClient
}
func CloseDB() {
	if DBClient != nil {
		sqlDB, err := DBClient.DB()
		if err == nil {
			sqlDB.Close()
		}
	}
}
