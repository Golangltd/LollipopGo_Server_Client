class GlobalConfig {
    constructor() {
        this.showOpen = true
        this.gameServer = ''      // 运营-游戏服务器
        this.logServer = ''       // 运营-异常捕捉服       
        this.registerServer = ''
        this.gameServerServerID = ''
        this.gameServerServerData = ''
        this.gameServerTocken = ''   // 客户端保存tocken信息
        this.gameServer7_2 = false
        this.gameServerUrl = "http://127.0.0.1:8867"
        this.GameList = {}
        this.showHB = false  //心跳显示
        this.ProxyServerId = "c4ca4238a0b923820dcc509a6f75849b"    
        this.GameServerId  = "c81e728d9d4c2f636f067f89cc14862c"    
        this.BattleServerId = "eccbc87e4b5ce2fe28308fd9f2a7baf3"   
        this.GMServerId    = "a87ff679a2f3e71d9181a67b7542122c"    
        this.DBServerId    = "e4da3b7fbbce2345d7772b0674a318d5"    
        this.CenterServerId  = "1679091c5a880faf6fb5e6087eb1b2dc"  
        this.OpenId = ""  // 玩家唯一Id
        // NOTE: 不要使用ut.isBrowser()，因为有初始化顺序的问题
        if (cc.sys.os != cc.sys.OS_ANDROID && cc.sys.os != cc.sys.OS_IOS) {
            // this.gameServer = 'ws://192.168.2.233:5000'
        }
    }

    get setting() {
        if (!this.__static_val_setting) {
            this.__static_val_setting = {
                soundVolume: 0.5,
                musicVolume: 0.5,
            }

            let str = cc.sys.localStorage.getItem('slots_settings')

            if (str) {
                try {
                    this.__static_val_setting = JSON.parse(str)
                } catch (err) {
                    cc.error(err)
                }
            }
        }

        return this.__static_val_setting
    }

    get appstore() {
        return false
    }

    get hotStorage() {
        return 'hellokitty'
    }

    get memoryCollectImmediate() {
        return false
    }

    saveSetting() {
        cc.sys.localStorage.setItem('slots_settings', JSON.stringify(this.setting))
    }
}

window.globalCfg = new GlobalConfig()
