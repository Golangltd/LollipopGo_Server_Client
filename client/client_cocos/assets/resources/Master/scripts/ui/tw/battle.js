
cc.Class({
    extends: cc.Component,

    properties: {
    },

    // 框架设计
    // -------------------------------------------------------------------------------------
    onLoad() {
        this.node.onenter = this.onenter.bind(this)
        this.node.onleave = this.onleave.bind(this)
    },

    onenter: function (data) {
        eventMgr.on('Battle', this, this.onMessage)
    },

    onleave: function () {
        eventMgr.off('Battle', this)
    },
    // -------------------------------------------------------------------------------------
    onMessage(data) {
        cc.log('战斗服 消息处理！', data)
        connMgr.reset(false);

        // 数据处理 -- 协议字段更新 
        if (data.Protocol == 10 && data.Protocol2 == 96) {
            cc.log("接收到进入战斗返回的协议 -- ",data);
            // 解析数据 操作
            cc.log("房间Id  -- ",data.RoomID);
            cc.log("父回合  -- ",data.ParentRound);
            cc.log("子回合  -- ",data.ChildRound);
            cc.log("出战列表  -- ",data.PlayList);    // 出战列表
            cc.log("攻击卡牌  -- ",data.AttackCard);  // 当前攻击卡牌
    
            if (globalCfg.OpenId== data.AttackCard.OpenId){
                // 对比相等就是自己的数据
                // 激活箭头设置
            }

            return

            // 发送转发协议给，game server 服务器
            var json = {
                Protocol: 6,    
                Protocol2: 1,
                Token: globalCfg.gameServerTocken,
                OpenId: data.OpenID,
            };
            client.sendProxyServer(globalCfg.GameServerId,json)
            
            return
        }
    },
    // -------------------------------------------------------------------------------------
    // 点击播放按钮 
    btnEvent(evt, data) {
        cc.log(data)
        audioMgr.playSound('clickbutton')

        if (data == 'play') {
            // 发送消息到战斗服务器
            var json = {
                Protocol: 10,    
                Protocol2: 95,
                IsPvP: false,
                UserA: null,
                UserB: null,
            };
            // 补充 用户A 和用户B 数据
            // 用户数据结构

            /*
                //用户信息
                type UserInfo struct {
                    OpenId string            // 用户唯一标识符
                    Level  int               // 用户等级
                    Cards  map[int]*CardInfo // 每个座位对应的卡牌信息
                }

                //角色信息
                type CardInfo struct {
                    CardID   uint64 // 卡牌唯一ID
                    Level    int    // 卡牌等级
                    RoleID   int    // 角色ID
                    Position int    // 站位
                    Skills   []int  // 使用技能列表
                }
            */
             
            //--------------------------------------------------------------------------
            var UserStA = {
                OpenId:globalCfg.OpenId,
                Level:1,
                Cards:[],
            }

            UserStA.Cards.push({
                CardID:   11,
				Level:    2,
				RoleID:   1021,
				Position: 0,
				Skills:   [10210101, 10210201, 10210301],
            })

            UserStA.Cards[1] ={
                CardID:   12,
				Level:    3,
				RoleID:   1022,
				Position: 1,
				Skills:   [10220101, 10220201, 10220301],
            }
            
            UserStA.Cards[2] ={
				CardID:   13,
				Level:    3,
				RoleID:   1005,
				Position: 2,
				Skills:   [10050101, 10050201, 10050301],
            }

            UserStA.Cards[3] ={
				CardID:   14,
				Level:    3,
				RoleID:   1006,
				Position: 3,
				Skills:   [10060101, 10060201, 10060301],
            }

            UserStA.Cards[5] ={
				CardID:   15,
				Level:    3,
				RoleID:   1033,
				Position: 4,
				Skills:   [10330101, 10330201, 10330301],
            }

            json.UserA = UserStA
            //--------------------------------------------------------------------------
            // 关卡怪物
            var UserStB = {
                OpenId:globalCfg.OpenId,
                Level:1,
                Cards:[],
            }

            UserStB.Cards[0] ={
                CardID:   11,
				Level:    2,
				RoleID:   1021,
				Position: 0,
				Skills:   [10210101, 10210201, 10210301],
            }

            UserStB.Cards[1] ={
                CardID:   12,
				Level:    3,
				RoleID:   1022,
				Position: 1,
				Skills:   [10220101, 10220201, 10220301],
            }

            UserStB.Cards[2] ={
				CardID:   13,
				Level:    3,
				RoleID:   1005,
				Position: 2,
				Skills:   [10050101, 10050201, 10050301],
            }

            UserStB.Cards[3] ={
				CardID:   14,
				Level:    3,
				RoleID:   1006,
				Position: 3,
				Skills:   [10060101, 10060201, 10060301],
            }

            UserStB.Cards[5] ={
				CardID:   15,
				Level:    3,
				RoleID:   1033,
				Position: 4,
				Skills:   [10330101, 10330201, 10330301],
            }

            json.UserB = UserStB
            //--------------------------------------------------------------------------
            // 发送给 战斗服务器
            cc.log(json)
            client.sendProxyServer(globalCfg.BattleServerId,json)
        } 
    },
});
