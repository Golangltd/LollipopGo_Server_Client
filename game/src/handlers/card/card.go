package card

import (
	gamedb "LollipopGo2.8x/data"
	"LollipopGo2.8x/models"
	. "LollipopGo2.8x/msg/drawcard"
	"LollipopGo2.8x/proto/comm_proto"
	"LollipopGo2.8x/temp"
	twLibUser "github.com/Golangltd/Twlib/user"
	"golang.org/x/net/websocket"
)

// 添加卡牌所需结构
type AddCardSt struct {
	Level    int //等级
	ConfigID int // 卡牌配置表ID
	Quality  int // 卡牌品质
}

const initLevel = 1

func newAddCardSt(cardID int, quality int) *AddCardSt {
	m := &AddCardSt{}
	m.Level = initLevel
	m.ConfigID = cardID
	m.Quality = quality
	return m
}

func getAddCardsSt(cardList []*CardData) []*AddCardSt {
	addCardsSt := make([]*AddCardSt, 0)
	for _, cardData := range cardList {
		addCardsSt = append(addCardsSt, newAddCardSt(cardData.CardID, cardData.Quality))
	}
	return addCardsSt
}

// 添加卡牌
func AddCard(game *models.Game, addCards []*AddCardSt) []*twLibUser.CardInfo {
	if len(addCards) < 1 {
		return nil
	}
	cards := make([]*twLibUser.CardInfo, 0)
	for _, c := range addCards {
		// 临时属性 TODO:待修改
		att := &twLibUser.AttributeSt{
			BattlePower:  1000, // 战斗力
			HPPower:      100,  // 血量
			AttackPower:  100,  // 攻击力
			DefensePower: 100,  // 防御力
		}
		cards = append(cards, &twLibUser.CardInfo{
			Level:         c.Level,
			Skills:        make([]*twLibUser.SkillInfo, 0), // TODO:创建卡牌时，技能初始化
			Equips:        make([]*twLibUser.EquipSt, 6),
			AttributeInfo: att,
			Quality:       c.Quality,
			Stars:         0,
			IsShow:        false,
		})
	}
	if game.UserInfo.CardList == nil {
		game.UserInfo.CardList = make([]*twLibUser.CardInfo, 0)
	}
	game.UserInfo.CardList = append(game.UserInfo.CardList, cards...)
	// 添加到缓存
	temp.GMapCardBag[game.AccountId] = game.UserInfo.CardList
	// 添加到数据库
	gamedb.CreateUserCardInfo(game.UserInfo.RoleUid, game.AccountId, cards)
	return cards
}

// 添加卡牌并通知客户端
func AddCardAndNotify(conn *websocket.Conn, openID string, game *models.Game, cardList []*CardData) {
	addCards := getAddCardsSt(cardList)
	// 添加卡牌
	cards := AddCard(game, addCards)
	// 通知客户端
	comm_proto.UpdateCard(conn, openID, nil, cards, nil)
}
