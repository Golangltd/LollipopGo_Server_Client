class ProtoGame {
    /*
      1. 游戏主协议，同步到LollipopGo框架
      2. 采用websocket通信
    */
    Proto = {
          GameDataProto : 1,    // 代理服务器主协议
          GameDataDBProto : 2,  // 游戏的DB的主协议
          GameNetProto:3,       // 游戏的Net主协议
          GErrorProto:4,        // 游戏的错误处理
          GGateWayProto:5,      // 网关协议
          GGameHallProto:6,     // 游戏主场景协议
          GGameLoginProto:7,    // 游戏登录服务器主协议
          GGameGMProto:8,       // 游戏GM管理的主协议
          GGamePayProto:9,      // 游戏支付的主协议
          GGameBattleProto:10,  // 游戏战斗主协议
          GGameConfigProto:11,  // 游戏配置主协议
    }
}

window.ProtoGame = new ProtoGame()
