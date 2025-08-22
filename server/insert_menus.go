package main

import (
	"fmt"
	"log"

	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
)

func main() {
	// 初始化配置
	global.GVA_VP = core.Viper()
	initialize.OtherInit()
	global.GVA_LOG = core.Zap()
	// 初始化数据库
	global.GVA_DB = initialize.Gorm()

	if global.GVA_DB == nil {
		log.Fatal("数据库连接失败")
	}

	fmt.Println("开始插入小说管理菜单...")

	// 插入小说管理主菜单
	mainMenu := system.SysBaseMenu{
		MenuLevel: 0,
		ParentId:  0,
		Path:      "novel",
		Name:      "novel",
		Hidden:    false,
		Component: "view/novel/index.vue",
		Sort:      4,
		Meta: system.Meta{
			Title:     "小说管理",
			Icon:      "reading",
			KeepAlive: true,
		},
	}

	// 检查主菜单是否已存在
	var existingMenu system.SysBaseMenu
	err := global.GVA_DB.Where("name = ? AND path = ?", mainMenu.Name, mainMenu.Path).First(&existingMenu).Error
	if err == nil {
		fmt.Println("小说管理主菜单已存在，跳过创建")
		mainMenu.ID = existingMenu.ID
	} else if err == gorm.ErrRecordNotFound {
		// 创建主菜单
		err = global.GVA_DB.Create(&mainMenu).Error
		if err != nil {
			log.Fatal("创建主菜单失败:", err)
		}
		fmt.Printf("创建主菜单成功，ID: %d\n", mainMenu.ID)
	} else {
		log.Fatal("查询主菜单失败:", err)
	}

	// 定义子菜单
	subMenus := []system.SysBaseMenu{
		{
			MenuLevel: 0,
			ParentId:  mainMenu.ID,
			Path:      "list",
			Name:      "novelList",
			Hidden:    false,
			Component: "view/novel/index.vue",
			Sort:      1,
			Meta: system.Meta{
				Title:     "小说列表",
				Icon:      "list",
				KeepAlive: true,
			},
		},
		{
			MenuLevel: 0,
			ParentId:  mainMenu.ID,
			Path:      "manage",
			Name:      "novelManage",
			Hidden:    false,
			Component: "view/novel/manage.vue",
			Sort:      2,
			Meta: system.Meta{
				Title:     "小说管理",
				Icon:      "management",
				KeepAlive: true,
			},
		},
		{
			MenuLevel: 0,
			ParentId:  mainMenu.ID,
			Path:      "characters",
			Name:      "novelCharacters",
			Hidden:    false,
			Component: "view/novel/characters.vue",
			Sort:      3,
			Meta: system.Meta{
				Title:     "角色管理",
				Icon:      "user",
				KeepAlive: true,
			},
		},
		{
			MenuLevel: 0,
			ParentId:  mainMenu.ID,
			Path:      "chapters",
			Name:      "novelChapters",
			Hidden:    false,
			Component: "view/novel/chapters.vue",
			Sort:      4,
			Meta: system.Meta{
				Title:     "章节管理",
				Icon:      "document",
				KeepAlive: true,
			},
		},
		{
			MenuLevel: 0,
			ParentId:  mainMenu.ID,
			Path:      "foreshadowing",
			Name:      "novelForeshadowing",
			Hidden:    false,
			Component: "view/novel/foreshadowing.vue",
			Sort:      5,
			Meta: system.Meta{
				Title:     "伏笔管理",
				Icon:      "connection",
				KeepAlive: true,
			},
		},
		{
			MenuLevel: 0,
			ParentId:  mainMenu.ID,
			Path:      "outlines",
			Name:      "novelOutlines",
			Hidden:    false,
			Component: "view/novel/outlines.vue",
			Sort:      6,
			Meta: system.Meta{
				Title:     "大纲管理",
				Icon:      "menu",
				KeepAlive: true,
			},
		},
		{
			MenuLevel: 0,
			ParentId:  mainMenu.ID,
			Path:      "ai-assistant",
			Name:      "novelAIAssistant",
			Hidden:    false,
			Component: "view/novel/ai-assistant.vue",
			Sort:      7,
			Meta: system.Meta{
				Title:     "AI助手",
				Icon:      "cpu",
				KeepAlive: true,
			},
		},
		{
			MenuLevel: 0,
			ParentId:  mainMenu.ID,
			Path:      "ai-config",
			Name:      "novelAIConfig",
			Hidden:    false,
			Component: "view/novel/ai-config.vue",
			Sort:      8,
			Meta: system.Meta{
				Title:     "AI配置",
				Icon:      "setting",
				KeepAlive: true,
			},
		},
	}

	// 插入子菜单
	for _, subMenu := range subMenus {
		// 检查子菜单是否已存在
		var existingSubMenu system.SysBaseMenu
		err := global.GVA_DB.Where("name = ? AND parent_id = ?", subMenu.Name, subMenu.ParentId).First(&existingSubMenu).Error
		if err == nil {
			fmt.Printf("子菜单 %s 已存在，跳过创建\n", subMenu.Meta.Title)
			continue
		} else if err == gorm.ErrRecordNotFound {
			// 创建子菜单
			err = global.GVA_DB.Create(&subMenu).Error
			if err != nil {
				log.Printf("创建子菜单 %s 失败: %v\n", subMenu.Meta.Title, err)
				continue
			}
			fmt.Printf("创建子菜单 %s 成功，ID: %d\n", subMenu.Meta.Title, subMenu.ID)
		} else {
			log.Printf("查询子菜单 %s 失败: %v\n", subMenu.Meta.Title, err)
		}
	}

	// 为admin角色分配权限
	fmt.Println("为admin角色分配小说管理权限...")

	// 获取admin用户的角色ID
	var adminUser system.SysUser
	err = global.GVA_DB.Where("username = ?", "admin").First(&adminUser).Error
	if err != nil {
		log.Printf("获取admin用户失败: %v\n", err)
	} else {
		// 获取所有小说管理相关菜单
		var novelMenus []system.SysBaseMenu
		err = global.GVA_DB.Where("id = ? OR parent_id = ?", mainMenu.ID, mainMenu.ID).Find(&novelMenus).Error
		if err != nil {
			log.Printf("获取小说管理菜单失败: %v\n", err)
		} else {
			// 为每个菜单分配权限
			for _, menu := range novelMenus {
				// 检查权限是否已存在
				var existingAuth system.SysAuthorityMenu
				err = global.GVA_DB.Where("sys_authority_authority_id = ? AND sys_base_menu_id = ?", 
					adminUser.AuthorityId, menu.ID).First(&existingAuth).Error
				if err == nil {
					fmt.Printf("菜单 %s 权限已存在\n", menu.Meta.Title)
					continue
				} else if err == gorm.ErrRecordNotFound {
					// 创建权限关联
					auth := system.SysAuthorityMenu{
						MenuId:      fmt.Sprintf("%d", menu.ID),
						AuthorityId: fmt.Sprintf("%d", adminUser.AuthorityId),
					}
					err = global.GVA_DB.Create(&auth).Error
					if err != nil {
						log.Printf("为菜单 %s 分配权限失败: %v\n", menu.Meta.Title, err)
					} else {
						fmt.Printf("为菜单 %s 分配权限成功\n", menu.Meta.Title)
					}
				}
			}
		}
	}

	fmt.Println("\n=== 小说管理菜单插入完成 ===")
	fmt.Println("请刷新浏览器页面查看新菜单")
	fmt.Println("访问地址: http://localhost:8080")
	fmt.Println("登录账号: admin")
	fmt.Println("登录密码: 123456")
}