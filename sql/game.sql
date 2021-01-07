/*
Navicat MySQL Data Transfer

Source Server         : 80
Source Server Version : 50727
Source Host           : 120.25.80.44:3306
Source Database       : game

Target Server Type    : MYSQL
Target Server Version : 50727
File Encoding         : 65001

Date: 2020-11-12 09:06:29
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `gl_card`
-- ----------------------------
DROP TABLE IF EXISTS `gl_card`;
CREATE TABLE `gl_card` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '普通id',
  `cardid` varchar(255) DEFAULT NULL COMMENT '卡牌配置id',
  `carduid` int(11) DEFAULT NULL COMMENT '卡牌uid',
  `lev` varchar(255) DEFAULT NULL COMMENT '等级',
  `accountid` int(11) DEFAULT NULL COMMENT '账号id',
  `roleuid` int(11) DEFAULT NULL COMMENT '角色uid',
  `quality` varchar(255) DEFAULT NULL COMMENT '  ',
  `skilllist` varchar(2000) DEFAULT NULL COMMENT '鎶€鑳藉垪琛?',
  `equiplist` varchar(2000) DEFAULT NULL COMMENT '瑁呭鍒楄〃',
  `attribute` varchar(2000) DEFAULT NULL COMMENT '闁诲繒鍋熼崑鐐哄焵椤戭剙鍊风换鍡涙煙?',
  `star` int(11) DEFAULT NULL COMMENT '星级',
  `isshow` int(11) DEFAULT NULL COMMENT '是否打开过图鉴系统',
  `createtime` varchar(255) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3095 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of gl_card
-- ----------------------------
INSERT INTO `gl_card` VALUES ('2851', '304001', '2851', '1', '119', '263', '4', '{}', '[{\"UID\":3,\"ConfID\":11001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":630},null,null,null,null,null]', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603113125492516212');
INSERT INTO `gl_card` VALUES ('2852', '107002', '2852', '1', '119', '263', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603113307378856216');
INSERT INTO `gl_card` VALUES ('2853', '207003', '2853', '1', '119', '263', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603113307492821691');
INSERT INTO `gl_card` VALUES ('2854', '204001', '2854', '1', '119', '263', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603113307618131581');
INSERT INTO `gl_card` VALUES ('2855', '107002', '2855', '1', '119', '263', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603113307751497917');
INSERT INTO `gl_card` VALUES ('2856', '207003', '2856', '1', '119', '263', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603113307884929120');
INSERT INTO `gl_card` VALUES ('2857', '207001', '2857', '1', '119', '263', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603113307992888801');
INSERT INTO `gl_card` VALUES ('2858', '407002', '2858', '1', '119', '263', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603113308109436641');
INSERT INTO `gl_card` VALUES ('2859', '110002', '2859', '1', '119', '263', '10', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603113308240098184');
INSERT INTO `gl_card` VALUES ('2860', '204001', '2860', '1', '119', '263', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603113308350550454');
INSERT INTO `gl_card` VALUES ('2861', '407002', '2861', '1', '119', '263', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603113308460303389');
INSERT INTO `gl_card` VALUES ('2862', '307002', '2862', '1', '113', '264', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116848968365560');
INSERT INTO `gl_card` VALUES ('2863', '404001', '2863', '1', '113', '264', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116849098203672');
INSERT INTO `gl_card` VALUES ('2864', '304001', '2864', '1', '113', '264', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116849207072137');
INSERT INTO `gl_card` VALUES ('2865', '407002', '2865', '1', '113', '264', '7', '{}', '[{\"UID\":10,\"ConfID\":21001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":770},{\"UID\":14,\"ConfID\":22001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":280},{\"UID\":12,\"ConfID\":23001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":420},{\"UID\":16,\"ConfID\":24001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":210},null,null]', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116849335963185');
INSERT INTO `gl_card` VALUES ('2866', '307001', '2866', '1', '113', '264', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116849445598364');
INSERT INTO `gl_card` VALUES ('2867', '104001', '2867', '1', '113', '264', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116849569248381');
INSERT INTO `gl_card` VALUES ('2868', '107001', '2868', '1', '113', '264', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116849684924093');
INSERT INTO `gl_card` VALUES ('2869', '204001', '2869', '1', '113', '264', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116849807781199');
INSERT INTO `gl_card` VALUES ('2870', '204001', '2870', '1', '113', '264', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116849936140551');
INSERT INTO `gl_card` VALUES ('2871', '407001', '2871', '1', '113', '264', '7', '{}', '[{\"UID\":10,\"ConfID\":21001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":770},null,null,null,null,null]', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116850051388793');
INSERT INTO `gl_card` VALUES ('2872', '113002', '2872', '1', '113', '264', '13', '{}', '[{\"UID\":9,\"ConfID\":11001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":630},{\"UID\":13,\"ConfID\":12001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":210},{\"UID\":11,\"ConfID\":13001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":490},{\"UID\":15,\"ConfID\":14001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":210},null,null]', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116854983191697');
INSERT INTO `gl_card` VALUES ('2873', '404001', '2873', '1', '113', '264', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116855093625985');
INSERT INTO `gl_card` VALUES ('2874', '407001', '2874', '1', '113', '264', '7', '{}', '[{\"UID\":10,\"ConfID\":21001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":770},{\"UID\":14,\"ConfID\":22001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":280},{\"UID\":12,\"ConfID\":23001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":420},{\"UID\":16,\"ConfID\":24001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":210},null,null]', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116855213986019');
INSERT INTO `gl_card` VALUES ('2875', '404001', '2875', '1', '113', '264', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116855332056581');
INSERT INTO `gl_card` VALUES ('2876', '204001', '2876', '1', '113', '264', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116855454378140');
INSERT INTO `gl_card` VALUES ('2877', '404001', '2877', '1', '113', '264', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116855574474176');
INSERT INTO `gl_card` VALUES ('2878', '204001', '2878', '1', '113', '264', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116855702741196');
INSERT INTO `gl_card` VALUES ('2879', '204001', '2879', '1', '113', '264', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116855814164721');
INSERT INTO `gl_card` VALUES ('2880', '204001', '2880', '1', '113', '264', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116855932698037');
INSERT INTO `gl_card` VALUES ('2881', '404001', '2881', '1', '113', '264', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116856061643381');
INSERT INTO `gl_card` VALUES ('2882', '304001', '2882', '1', '120', '265', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116856575584830');
INSERT INTO `gl_card` VALUES ('2883', '204001', '2883', '1', '120', '265', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116856683846261');
INSERT INTO `gl_card` VALUES ('2884', '204001', '2884', '1', '120', '265', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116856811485338');
INSERT INTO `gl_card` VALUES ('2885', '210001', '2885', '1', '120', '265', '10', '{}', '[{\"UID\":17,\"ConfID\":11001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":630},{\"UID\":21,\"ConfID\":12001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":210},{\"UID\":19,\"ConfID\":13001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":490},null,null,null]', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116856930537813');
INSERT INTO `gl_card` VALUES ('2886', '304001', '2886', '1', '120', '265', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116857061763805');
INSERT INTO `gl_card` VALUES ('2887', '204001', '2887', '1', '120', '265', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116857181739544');
INSERT INTO `gl_card` VALUES ('2888', '404001', '2888', '1', '120', '265', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116857315178642');
INSERT INTO `gl_card` VALUES ('2889', '207003', '2889', '1', '120', '265', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116857430008941');
INSERT INTO `gl_card` VALUES ('2890', '407001', '2890', '1', '120', '265', '7', '{}', '[{\"UID\":18,\"ConfID\":21001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":770},{\"UID\":22,\"ConfID\":22001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":280},{\"UID\":20,\"ConfID\":23001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":420},null,null,null]', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116857530987248');
INSERT INTO `gl_card` VALUES ('2891', '110001', '2891', '1', '120', '265', '10', '{}', '[{\"UID\":17,\"ConfID\":11001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":630},null,{\"UID\":19,\"ConfID\":13001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":490},null,null,null]', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116857643122914');
INSERT INTO `gl_card` VALUES ('2892', '304001', '2892', '1', '113', '264', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116858874367820');
INSERT INTO `gl_card` VALUES ('2893', '110001', '2893', '1', '113', '264', '10', '{}', '[{\"UID\":9,\"ConfID\":11001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":630},{\"UID\":13,\"ConfID\":12001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":210},{\"UID\":11,\"ConfID\":13001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":490},{\"UID\":15,\"ConfID\":14001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":210},null,null]', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116858983082575');
INSERT INTO `gl_card` VALUES ('2894', '404001', '2894', '1', '113', '264', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116859087186330');
INSERT INTO `gl_card` VALUES ('2895', '104001', '2895', '1', '113', '264', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116859202228950');
INSERT INTO `gl_card` VALUES ('2896', '404001', '2896', '1', '113', '264', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116859317066625');
INSERT INTO `gl_card` VALUES ('2897', '304001', '2897', '1', '113', '264', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116859431294965');
INSERT INTO `gl_card` VALUES ('2898', '304001', '2898', '1', '113', '264', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116859547054374');
INSERT INTO `gl_card` VALUES ('2899', '104001', '2899', '1', '113', '264', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116859658616435');
INSERT INTO `gl_card` VALUES ('2900', '304001', '2900', '1', '113', '264', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116859777308413');
INSERT INTO `gl_card` VALUES ('2901', '304001', '2901', '1', '113', '264', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603116859889106907');
INSERT INTO `gl_card` VALUES ('2902', '204001', '2902', '1', '120', '265', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117159211198996');
INSERT INTO `gl_card` VALUES ('2903', '107001', '2903', '1', '120', '265', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117159336088964');
INSERT INTO `gl_card` VALUES ('2904', '407001', '2904', '1', '120', '265', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117159472109031');
INSERT INTO `gl_card` VALUES ('2905', '404001', '2905', '1', '120', '265', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117159596310080');
INSERT INTO `gl_card` VALUES ('2906', '404001', '2906', '1', '120', '265', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117159729251787');
INSERT INTO `gl_card` VALUES ('2907', '307002', '2907', '1', '120', '265', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117159857306889');
INSERT INTO `gl_card` VALUES ('2908', '104001', '2908', '1', '120', '265', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117159973858657');
INSERT INTO `gl_card` VALUES ('2909', '404001', '2909', '1', '120', '265', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117160085212693');
INSERT INTO `gl_card` VALUES ('2910', '307001', '2910', '1', '120', '265', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117160196085756');
INSERT INTO `gl_card` VALUES ('2911', '304001', '2911', '1', '120', '265', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117160323687235');
INSERT INTO `gl_card` VALUES ('2912', '207003', '2912', '1', '120', '265', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117493621087918');
INSERT INTO `gl_card` VALUES ('2913', '107003', '2913', '1', '120', '265', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117493740086121');
INSERT INTO `gl_card` VALUES ('2914', '107001', '2914', '1', '120', '265', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117493853863790');
INSERT INTO `gl_card` VALUES ('2915', '204001', '2915', '1', '120', '265', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117493979850956');
INSERT INTO `gl_card` VALUES ('2916', '404001', '2916', '1', '120', '265', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117494094410649');
INSERT INTO `gl_card` VALUES ('2917', '210002', '2917', '1', '120', '265', '10', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117494223052718');
INSERT INTO `gl_card` VALUES ('2918', '610001', '2918', '1', '120', '265', '10', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117494354213018');
INSERT INTO `gl_card` VALUES ('2919', '404001', '2919', '1', '120', '265', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117494478988205');
INSERT INTO `gl_card` VALUES ('2920', '407001', '2920', '1', '120', '265', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117494580291458');
INSERT INTO `gl_card` VALUES ('2921', '304001', '2921', '1', '120', '265', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117494691228299');
INSERT INTO `gl_card` VALUES ('2922', '404001', '2922', '1', '120', '265', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117548391409975');
INSERT INTO `gl_card` VALUES ('2923', '304001', '2923', '1', '120', '265', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117548504239372');
INSERT INTO `gl_card` VALUES ('2924', '110005', '2924', '1', '120', '265', '10', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117548625569673');
INSERT INTO `gl_card` VALUES ('2925', '204001', '2925', '1', '120', '265', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117548748168538');
INSERT INTO `gl_card` VALUES ('2926', '407001', '2926', '1', '120', '265', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117548875678172');
INSERT INTO `gl_card` VALUES ('2927', '213001', '2927', '1', '120', '265', '13', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117548989481802');
INSERT INTO `gl_card` VALUES ('2928', '404001', '2928', '1', '120', '265', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117549120869517');
INSERT INTO `gl_card` VALUES ('2929', '304001', '2929', '1', '120', '265', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117549238214557');
INSERT INTO `gl_card` VALUES ('2930', '410004', '2930', '1', '120', '265', '10', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117549338213172');
INSERT INTO `gl_card` VALUES ('2931', '407001', '2931', '1', '120', '265', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117549462151502');
INSERT INTO `gl_card` VALUES ('2932', '407002', '2932', '1', '121', '266', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117782834154303');
INSERT INTO `gl_card` VALUES ('2933', '307001', '2933', '1', '121', '266', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117799061930230');
INSERT INTO `gl_card` VALUES ('2934', '104001', '2934', '1', '121', '266', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117799186330306');
INSERT INTO `gl_card` VALUES ('2935', '204001', '2935', '1', '121', '266', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117799309643251');
INSERT INTO `gl_card` VALUES ('2936', '107003', '2936', '1', '121', '266', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117799429336450');
INSERT INTO `gl_card` VALUES ('2937', '407002', '2937', '1', '121', '266', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117799542260498');
INSERT INTO `gl_card` VALUES ('2938', '407001', '2938', '1', '121', '266', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117799670485766');
INSERT INTO `gl_card` VALUES ('2939', '204001', '2939', '1', '121', '266', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117799790618502');
INSERT INTO `gl_card` VALUES ('2940', '307001', '2940', '1', '121', '266', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117799916967356');
INSERT INTO `gl_card` VALUES ('2941', '207003', '2941', '1', '121', '266', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117800029995212');
INSERT INTO `gl_card` VALUES ('2942', '204001', '2942', '1', '121', '266', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117800148075516');
INSERT INTO `gl_card` VALUES ('2943', '407002', '2943', '1', '122', '267', '7', '{}', '[{\"UID\":2,\"ConfID\":21001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":770},{\"UID\":7,\"ConfID\":22001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":280},{\"UID\":3,\"ConfID\":23001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":420},{\"UID\":9,\"ConfID\":24001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":210},null,null]', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117802066277633');
INSERT INTO `gl_card` VALUES ('2944', '107002', '2944', '1', '122', '267', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117802180666459');
INSERT INTO `gl_card` VALUES ('2945', '104001', '2945', '1', '122', '267', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117802288119843');
INSERT INTO `gl_card` VALUES ('2946', '404001', '2946', '1', '122', '267', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117802424534120');
INSERT INTO `gl_card` VALUES ('2947', '610003', '2947', '1', '122', '267', '10', '{}', '[{\"UID\":1,\"ConfID\":11001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":630},{\"UID\":6,\"ConfID\":12001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":210},{\"UID\":4,\"ConfID\":13001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":490},{\"UID\":8,\"ConfID\":14001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":210},null,null]', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117802538370313');
INSERT INTO `gl_card` VALUES ('2948', '204001', '2948', '1', '122', '267', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117802655049282');
INSERT INTO `gl_card` VALUES ('2949', '410003', '2949', '1', '122', '267', '10', '{}', '[{\"UID\":1,\"ConfID\":11001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":630},{\"UID\":6,\"ConfID\":12001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":210},{\"UID\":5,\"ConfID\":13001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":490},null,null,null]', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117802766625781');
INSERT INTO `gl_card` VALUES ('2950', '307002', '2950', '1', '122', '267', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117802889418040');
INSERT INTO `gl_card` VALUES ('2951', '104001', '2951', '1', '122', '267', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117803007270671');
INSERT INTO `gl_card` VALUES ('2952', '107001', '2952', '1', '122', '267', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117803124963933');
INSERT INTO `gl_card` VALUES ('2953', '104001', '2953', '1', '121', '266', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117805555758624');
INSERT INTO `gl_card` VALUES ('2954', '510002', '2954', '1', '121', '266', '10', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117805680096434');
INSERT INTO `gl_card` VALUES ('2955', '104001', '2955', '1', '121', '266', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117805806017340');
INSERT INTO `gl_card` VALUES ('2956', '404001', '2956', '1', '121', '266', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117805942550515');
INSERT INTO `gl_card` VALUES ('2957', '407001', '2957', '1', '121', '266', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117806071285877');
INSERT INTO `gl_card` VALUES ('2958', '104001', '2958', '1', '121', '266', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117806204461448');
INSERT INTO `gl_card` VALUES ('2959', '207002', '2959', '1', '121', '266', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117806331393174');
INSERT INTO `gl_card` VALUES ('2960', '207003', '2960', '1', '121', '266', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117806443267687');
INSERT INTO `gl_card` VALUES ('2961', '407001', '2961', '1', '121', '266', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117806562333821');
INSERT INTO `gl_card` VALUES ('2962', '304001', '2962', '1', '121', '266', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117806677164211');
INSERT INTO `gl_card` VALUES ('2963', '310004', '2963', '1', '121', '266', '10', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117810838615011');
INSERT INTO `gl_card` VALUES ('2964', '404001', '2964', '1', '121', '266', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117810952085159');
INSERT INTO `gl_card` VALUES ('2965', '107001', '2965', '1', '121', '266', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117811083243504');
INSERT INTO `gl_card` VALUES ('2966', '207003', '2966', '1', '121', '266', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117811197603866');
INSERT INTO `gl_card` VALUES ('2967', '104001', '2967', '1', '121', '266', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117811334355687');
INSERT INTO `gl_card` VALUES ('2968', '104001', '2968', '1', '121', '266', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117811452807799');
INSERT INTO `gl_card` VALUES ('2969', '110002', '2969', '1', '121', '266', '10', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117811572528262');
INSERT INTO `gl_card` VALUES ('2970', '104001', '2970', '1', '121', '266', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117811698115889');
INSERT INTO `gl_card` VALUES ('2971', '107003', '2971', '1', '121', '266', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117811817132393');
INSERT INTO `gl_card` VALUES ('2972', '407001', '2972', '1', '121', '266', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117811933700982');
INSERT INTO `gl_card` VALUES ('2973', '204001', '2973', '1', '121', '266', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117815649705901');
INSERT INTO `gl_card` VALUES ('2974', '104001', '2974', '1', '121', '266', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117815773010417');
INSERT INTO `gl_card` VALUES ('2975', '104001', '2975', '1', '121', '266', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117815900053456');
INSERT INTO `gl_card` VALUES ('2976', '204001', '2976', '1', '121', '266', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117816028857915');
INSERT INTO `gl_card` VALUES ('2977', '104001', '2977', '1', '121', '266', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117816148205115');
INSERT INTO `gl_card` VALUES ('2978', '207003', '2978', '1', '121', '266', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117816280971579');
INSERT INTO `gl_card` VALUES ('2979', '304001', '2979', '1', '121', '266', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117816399478529');
INSERT INTO `gl_card` VALUES ('2980', '104001', '2980', '1', '121', '266', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117816516969290');
INSERT INTO `gl_card` VALUES ('2981', '310002', '2981', '1', '121', '266', '10', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117816630806639');
INSERT INTO `gl_card` VALUES ('2982', '410004', '2982', '1', '121', '266', '10', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603117816743321845');
INSERT INTO `gl_card` VALUES ('2983', '404001', '2983', '1', '123', '268', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603248232092350376');
INSERT INTO `gl_card` VALUES ('2984', '207002', '2984', '1', '123', '268', '7', '{}', '[{\"UID\":11,\"ConfID\":11001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":630},null,null,null,null,null]', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603248235301763565');
INSERT INTO `gl_card` VALUES ('2985', '407001', '2985', '1', '123', '268', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603248235421150534');
INSERT INTO `gl_card` VALUES ('2986', '407001', '2986', '1', '123', '268', '7', '{}', '[null,null,null,null,null,null]', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603248235542115824');
INSERT INTO `gl_card` VALUES ('2987', '204001', '2987', '1', '123', '268', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603248235672128516');
INSERT INTO `gl_card` VALUES ('2988', '404001', '2988', '1', '123', '268', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603248235786430658');
INSERT INTO `gl_card` VALUES ('2989', '304001', '2989', '1', '123', '268', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603248235897222041');
INSERT INTO `gl_card` VALUES ('2990', '207002', '2990', '1', '123', '268', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603248236017176493');
INSERT INTO `gl_card` VALUES ('2991', '307002', '2991', '1', '123', '268', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603248236143398132');
INSERT INTO `gl_card` VALUES ('2992', '104001', '2992', '1', '123', '268', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603248236256875716');
INSERT INTO `gl_card` VALUES ('2993', '407002', '2993', '1', '123', '268', '7', '{}', '[{\"UID\":14,\"ConfID\":21001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":770},null,null,null,null,null]', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603248236385407456');
INSERT INTO `gl_card` VALUES ('2994', '410003', '2994', '1', '123', '268', '10', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603248957847631518');
INSERT INTO `gl_card` VALUES ('2995', '310003', '2995', '1', '124', '270', '10', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603255933299443254');
INSERT INTO `gl_card` VALUES ('2996', '104001', '2996', '1', '124', '270', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603255980975141658');
INSERT INTO `gl_card` VALUES ('2997', '304001', '2997', '1', '124', '270', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603255981101672445');
INSERT INTO `gl_card` VALUES ('2998', '307002', '2998', '1', '124', '270', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603255981217813408');
INSERT INTO `gl_card` VALUES ('2999', '307002', '2999', '1', '124', '270', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603255981350123286');
INSERT INTO `gl_card` VALUES ('3000', '304001', '3000', '1', '124', '270', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603255981478113671');
INSERT INTO `gl_card` VALUES ('3001', '107002', '3001', '1', '124', '270', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603255981607018566');
INSERT INTO `gl_card` VALUES ('3002', '207001', '3002', '1', '124', '270', '7', '{}', '[null,null,null,null,null,null]', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603255981716259580');
INSERT INTO `gl_card` VALUES ('3003', '207002', '3003', '1', '124', '270', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603255981838203387');
INSERT INTO `gl_card` VALUES ('3004', '404001', '3004', '1', '124', '270', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603255981961184087');
INSERT INTO `gl_card` VALUES ('3005', '110004', '3005', '1', '124', '270', '10', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603255982065608969');
INSERT INTO `gl_card` VALUES ('3006', '104001', '3006', '1', '126', '272', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603290002764239389');
INSERT INTO `gl_card` VALUES ('3007', '407002', '3007', '1', '126', '272', '7', '{}', '[{\"UID\":22,\"ConfID\":21001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":770},null,{\"UID\":24,\"ConfID\":23001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":420},null,null,null]', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603290004624921579');
INSERT INTO `gl_card` VALUES ('3008', '204001', '3008', '1', '126', '272', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603290004741933644');
INSERT INTO `gl_card` VALUES ('3009', '104001', '3009', '1', '126', '272', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603290004869387230');
INSERT INTO `gl_card` VALUES ('3010', '304001', '3010', '1', '126', '272', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603290004992432110');
INSERT INTO `gl_card` VALUES ('3011', '404001', '3011', '1', '126', '272', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603290005118772536');
INSERT INTO `gl_card` VALUES ('3012', '207002', '3012', '1', '126', '272', '7', '{}', '[null,null,null,null,null,null]', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603290005239170292');
INSERT INTO `gl_card` VALUES ('3013', '304001', '3013', '1', '126', '272', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603290005336798396');
INSERT INTO `gl_card` VALUES ('3014', '307002', '3014', '1', '126', '272', '7', '{}', '[{\"UID\":21,\"ConfID\":11001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":630},null,{\"UID\":23,\"ConfID\":13001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":490},null,null,null]', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603290005442265075');
INSERT INTO `gl_card` VALUES ('3015', '207002', '3015', '1', '126', '272', '7', '{}', '[{\"UID\":21,\"ConfID\":11001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":630},null,null,null,null,null]', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603290005550109021');
INSERT INTO `gl_card` VALUES ('3016', '207001', '3016', '1', '126', '272', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603290005662249028');
INSERT INTO `gl_card` VALUES ('3017', '204001', '3017', '1', '127', '274', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603355273405245236');
INSERT INTO `gl_card` VALUES ('3018', '107001', '3018', '1', '127', '274', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603355273523563932');
INSERT INTO `gl_card` VALUES ('3019', '404001', '3019', '1', '127', '274', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603355273648657320');
INSERT INTO `gl_card` VALUES ('3020', '304001', '3020', '1', '127', '274', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603355273769369595');
INSERT INTO `gl_card` VALUES ('3021', '207001', '3021', '1', '127', '274', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603355273876946053');
INSERT INTO `gl_card` VALUES ('3022', '204001', '3022', '1', '127', '274', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603355273998262268');
INSERT INTO `gl_card` VALUES ('3023', '407001', '3023', '1', '127', '274', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603355274123882398');
INSERT INTO `gl_card` VALUES ('3024', '510002', '3024', '1', '127', '274', '10', '{}', '[{\"UID\":25,\"ConfID\":11001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":630},null,null,null,null,null]', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603355274247209295');
INSERT INTO `gl_card` VALUES ('3025', '304001', '3025', '1', '127', '274', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603355274378997125');
INSERT INTO `gl_card` VALUES ('3026', '204001', '3026', '1', '127', '274', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603355274498583060');
INSERT INTO `gl_card` VALUES ('3027', '407002', '3027', '1', '128', '275', '7', '{}', '[{\"UID\":31,\"ConfID\":21001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":770},null,null,null,null,null]', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603548089796291747');
INSERT INTO `gl_card` VALUES ('3028', '104001', '3028', '1', '128', '275', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603548089910661211');
INSERT INTO `gl_card` VALUES ('3029', '207003', '3029', '1', '128', '275', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603548090021306166');
INSERT INTO `gl_card` VALUES ('3030', '107002', '3030', '1', '128', '275', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603548090147102768');
INSERT INTO `gl_card` VALUES ('3031', '304001', '3031', '1', '128', '275', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603548090256946314');
INSERT INTO `gl_card` VALUES ('3032', '207003', '3032', '1', '128', '275', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603548090365510591');
INSERT INTO `gl_card` VALUES ('3033', '107001', '3033', '1', '128', '275', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603548090483480602');
INSERT INTO `gl_card` VALUES ('3034', '104001', '3034', '1', '128', '275', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603548090601214837');
INSERT INTO `gl_card` VALUES ('3035', '407002', '3035', '1', '128', '275', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603548090728487690');
INSERT INTO `gl_card` VALUES ('3036', '510003', '3036', '1', '128', '275', '10', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603548090842156454');
INSERT INTO `gl_card` VALUES ('3037', '204001', '3037', '1', '129', '276', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603636861458297138');
INSERT INTO `gl_card` VALUES ('3038', '304001', '3038', '1', '129', '276', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603636861591354073');
INSERT INTO `gl_card` VALUES ('3039', '513003', '3039', '1', '129', '276', '13', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603636861722314014');
INSERT INTO `gl_card` VALUES ('3040', '104001', '3040', '1', '129', '276', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603636861839860879');
INSERT INTO `gl_card` VALUES ('3041', '110004', '3041', '1', '129', '276', '10', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603636861959840467');
INSERT INTO `gl_card` VALUES ('3042', '207003', '3042', '1', '129', '276', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603636862085718865');
INSERT INTO `gl_card` VALUES ('3043', '104001', '3043', '1', '129', '276', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603636862214187347');
INSERT INTO `gl_card` VALUES ('3044', '107002', '3044', '1', '129', '276', '7', '{}', '[{\"UID\":36,\"ConfID\":21001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":770},null,null,null,null,null]', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603636862355212942');
INSERT INTO `gl_card` VALUES ('3045', '107001', '3045', '1', '129', '276', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603636862478361411');
INSERT INTO `gl_card` VALUES ('3046', '207001', '3046', '1', '129', '276', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603636862592055005');
INSERT INTO `gl_card` VALUES ('3047', '207002', '3047', '1', '130', '277', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603710485166445368');
INSERT INTO `gl_card` VALUES ('3048', '307001', '3048', '1', '130', '277', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603710488202269144');
INSERT INTO `gl_card` VALUES ('3049', '110005', '3049', '1', '130', '277', '10', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603710488315718859');
INSERT INTO `gl_card` VALUES ('3050', '304001', '3050', '1', '130', '277', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603710488426208421');
INSERT INTO `gl_card` VALUES ('3051', '304001', '3051', '1', '130', '277', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603710488542165330');
INSERT INTO `gl_card` VALUES ('3052', '207001', '3052', '1', '130', '277', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603710488653058524');
INSERT INTO `gl_card` VALUES ('3053', '104001', '3053', '1', '130', '277', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603710488767155871');
INSERT INTO `gl_card` VALUES ('3054', '404001', '3054', '1', '130', '277', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603710488877265454');
INSERT INTO `gl_card` VALUES ('3055', '207002', '3055', '1', '130', '277', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603710488977603360');
INSERT INTO `gl_card` VALUES ('3056', '207002', '3056', '1', '130', '277', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603710489090570249');
INSERT INTO `gl_card` VALUES ('3057', '110001', '3057', '1', '130', '277', '10', '{}', '[{\"UID\":37,\"ConfID\":11001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":630},null,{\"UID\":39,\"ConfID\":13001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":490},null,null,null]', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603710489203989214');
INSERT INTO `gl_card` VALUES ('3058', '404001', '3058', '1', '131', '278', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603854221413516009');
INSERT INTO `gl_card` VALUES ('3059', '304001', '3059', '1', '131', '278', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603854224191348000');
INSERT INTO `gl_card` VALUES ('3060', '207001', '3060', '1', '131', '278', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603854224320811883');
INSERT INTO `gl_card` VALUES ('3061', '413001', '3061', '1', '131', '278', '13', '{}', '[]', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603854224445137294');
INSERT INTO `gl_card` VALUES ('3062', '407002', '3062', '1', '131', '278', '7', '{}', '[{\"UID\":42,\"ConfID\":21001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":770},null,{\"UID\":45,\"ConfID\":23001,\"Star\":0,\"Camp\":0,\"CampRate\":0,\"Num\":1,\"Exp\":0,\"Power\":420},null,null,null]', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603854224572015824');
INSERT INTO `gl_card` VALUES ('3063', '204001', '3063', '1', '131', '278', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603854224692851744');
INSERT INTO `gl_card` VALUES ('3064', '107003', '3064', '1', '131', '278', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603854224811008111');
INSERT INTO `gl_card` VALUES ('3065', '204001', '3065', '1', '131', '278', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603854224945272624');
INSERT INTO `gl_card` VALUES ('3066', '204001', '3066', '1', '131', '278', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603854225054199058');
INSERT INTO `gl_card` VALUES ('3067', '210002', '3067', '1', '131', '278', '10', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603854225181355469');
INSERT INTO `gl_card` VALUES ('3068', '404001', '3068', '1', '131', '278', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603854225300978061');
INSERT INTO `gl_card` VALUES ('3069', '204001', '3069', '1', '131', '278', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1603854339962783575');
INSERT INTO `gl_card` VALUES ('3070', '404001', '3070', '1', '131', '278', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1604741943053218575');
INSERT INTO `gl_card` VALUES ('3071', '407002', '3071', '1', '131', '278', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1604745411551235484');
INSERT INTO `gl_card` VALUES ('3072', '307002', '3072', '1', '131', '278', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1604745411676378074');
INSERT INTO `gl_card` VALUES ('3073', '104001', '3073', '1', '131', '278', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1604745411790335143');
INSERT INTO `gl_card` VALUES ('3074', '204001', '3074', '1', '131', '278', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1604745411906280335');
INSERT INTO `gl_card` VALUES ('3075', '407001', '3075', '1', '131', '278', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1604745412036155531');
INSERT INTO `gl_card` VALUES ('3076', '204001', '3076', '1', '131', '278', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1604745412162314113');
INSERT INTO `gl_card` VALUES ('3077', '510001', '3077', '1', '131', '278', '10', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1604745412287200683');
INSERT INTO `gl_card` VALUES ('3078', '404001', '3078', '1', '131', '278', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1604745412419949717');
INSERT INTO `gl_card` VALUES ('3079', '207002', '3079', '1', '131', '278', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1604745412544611292');
INSERT INTO `gl_card` VALUES ('3080', '104001', '3080', '1', '131', '278', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1604745412663157669');
INSERT INTO `gl_card` VALUES ('3081', '407001', '3081', '1', '132', '279', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1604746311066871275');
INSERT INTO `gl_card` VALUES ('3082', '210001', '3082', '1', '132', '279', '10', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1604746326478563323');
INSERT INTO `gl_card` VALUES ('3083', '304001', '3083', '1', '132', '279', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1604746995436487621');
INSERT INTO `gl_card` VALUES ('3084', '207001', '3084', '1', '126', '272', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1605103842041579774');
INSERT INTO `gl_card` VALUES ('3085', '207001', '3085', '1', '126', '272', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1605103844032431560');
INSERT INTO `gl_card` VALUES ('3086', '404001', '3086', '1', '126', '272', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1605103844153297121');
INSERT INTO `gl_card` VALUES ('3087', '107002', '3087', '1', '126', '272', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1605103844281220328');
INSERT INTO `gl_card` VALUES ('3088', '104001', '3088', '1', '126', '272', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1605103844383240126');
INSERT INTO `gl_card` VALUES ('3089', '204001', '3089', '1', '126', '272', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1605103844483556824');
INSERT INTO `gl_card` VALUES ('3090', '410001', '3090', '1', '126', '272', '10', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1605103844597720861');
INSERT INTO `gl_card` VALUES ('3091', '304001', '3091', '1', '126', '272', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1605103844719988936');
INSERT INTO `gl_card` VALUES ('3092', '307001', '3092', '1', '126', '272', '7', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1605103844854539036');
INSERT INTO `gl_card` VALUES ('3093', '210001', '3093', '1', '126', '272', '10', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1605103844988156706');
INSERT INTO `gl_card` VALUES ('3094', '304001', '3094', '1', '126', '272', '4', '{}', '{}', '{\"BattlePower\":1000,\"HPPower\":100,\"AttackPower\":100,\"DefensePower\":100,\"MagicPower\":0}', '0', '0', '1605103845106328818');

-- ----------------------------
-- Table structure for `gl_chapter`
-- ----------------------------
DROP TABLE IF EXISTS `gl_chapter`;
CREATE TABLE `gl_chapter` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `roleuid` int(11) DEFAULT NULL COMMENT '玩家ID',
  `chapterid` int(11) DEFAULT NULL COMMENT '章节ID',
  `chapterid2` int(11) DEFAULT NULL COMMENT 'zi章节ID',
  `roundid` int(11) DEFAULT NULL COMMENT '管卡ID',
  `createtime` varchar(20) DEFAULT NULL COMMENT '鍒涘缓鏃堕棿',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=49 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of gl_chapter
-- ----------------------------
INSERT INTO `gl_chapter` VALUES ('35', '263', '1', '2', '13', '1603113120089849893');
INSERT INTO `gl_chapter` VALUES ('36', '264', '1', '3', '22', '1603116837077999756');
INSERT INTO `gl_chapter` VALUES ('37', '265', '1', '2', '11', '1603117169908311069');
INSERT INTO `gl_chapter` VALUES ('38', '267', '1', '3', '22', '1603117830163988242');
INSERT INTO `gl_chapter` VALUES ('39', '268', '1', '1', '7', '1603248240791530802');
INSERT INTO `gl_chapter` VALUES ('40', '269', '1', '1', '3', '1603250010366567144');
INSERT INTO `gl_chapter` VALUES ('41', '270', '1', '1', '2', '1603255928356356371');
INSERT INTO `gl_chapter` VALUES ('42', '272', '1', '2', '11', '1603290019909121225');
INSERT INTO `gl_chapter` VALUES ('43', '274', '1', '1', '6', '1603355289504289649');
INSERT INTO `gl_chapter` VALUES ('44', '275', '1', '1', '6', '1603548109510591474');
INSERT INTO `gl_chapter` VALUES ('45', '276', '1', '1', '4', '1603636872370674187');
INSERT INTO `gl_chapter` VALUES ('46', '277', '1', '1', '6', '1603710493695389670');
INSERT INTO `gl_chapter` VALUES ('47', '278', '1', '1', '9', '1603854230047828070');
INSERT INTO `gl_chapter` VALUES ('48', '279', '1', '1', '3', '1604748419431268211');

-- ----------------------------
-- Table structure for `gl_equip`
-- ----------------------------
DROP TABLE IF EXISTS `gl_equip`;
CREATE TABLE `gl_equip` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `roleuid` int(11) DEFAULT NULL COMMENT '鐟欐帟澹奿d',
  `accountid` int(11) DEFAULT NULL COMMENT '璐﹀彿id',
  `equipuid` bigint(20) DEFAULT NULL COMMENT '鍞竴ID锛屽叏鏈嶅敮涓€',
  `equipid` int(11) DEFAULT NULL COMMENT '装备唯一ID',
  `num` int(11) DEFAULT NULL COMMENT '瑁呭鏁伴噺',
  `createtime` varchar(20) DEFAULT NULL COMMENT '鍒涘缓鏃堕棿',
  `star` int(11) DEFAULT NULL,
  `camp` int(11) DEFAULT NULL,
  `camprate` int(11) DEFAULT NULL,
  `exp` int(11) DEFAULT NULL,
  `power` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=71 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of gl_equip
-- ----------------------------
INSERT INTO `gl_equip` VALUES ('11', '263', '119', '11', '11001', '3', '1603113135677314567', '0', '0', '0', '0', '630');
INSERT INTO `gl_equip` VALUES ('12', '263', '119', '12', '21001', '2', '1603113135878228913', '0', '0', '0', '0', '770');
INSERT INTO `gl_equip` VALUES ('13', '263', '119', '13', '13001', '2', '1603113422015345450', '0', '0', '0', '0', '490');
INSERT INTO `gl_equip` VALUES ('14', '263', '119', '14', '23001', '2', '1603113422216175783', '0', '0', '0', '0', '420');
INSERT INTO `gl_equip` VALUES ('15', '263', '119', '15', '12001', '1', '1603113435049854305', '0', '0', '0', '0', '210');
INSERT INTO `gl_equip` VALUES ('16', '263', '119', '16', '22001', '1', '1603113435234607031', '0', '0', '0', '0', '280');
INSERT INTO `gl_equip` VALUES ('17', '264', '113', '17', '11001', '3', '1603116840570504786', '0', '0', '0', '0', '630');
INSERT INTO `gl_equip` VALUES ('18', '264', '113', '18', '21001', '3', '1603116840756626407', '0', '0', '0', '0', '770');
INSERT INTO `gl_equip` VALUES ('19', '264', '113', '19', '13001', '2', '1603116891688334688', '0', '0', '0', '0', '490');
INSERT INTO `gl_equip` VALUES ('20', '264', '113', '20', '23001', '2', '1603116891893581369', '0', '0', '0', '0', '420');
INSERT INTO `gl_equip` VALUES ('21', '264', '113', '21', '12001', '2', '1603116936785970701', '0', '0', '0', '0', '210');
INSERT INTO `gl_equip` VALUES ('22', '264', '113', '22', '22001', '2', '1603116936977717266', '0', '0', '0', '0', '280');
INSERT INTO `gl_equip` VALUES ('23', '264', '113', '23', '14001', '2', '1603116944020587898', '0', '0', '0', '0', '210');
INSERT INTO `gl_equip` VALUES ('24', '264', '113', '24', '24001', '2', '1603116944215016965', '0', '0', '0', '0', '210');
INSERT INTO `gl_equip` VALUES ('25', '265', '120', '25', '11001', '2', '1603117173299535149', '0', '0', '0', '0', '630');
INSERT INTO `gl_equip` VALUES ('26', '265', '120', '26', '21001', '2', '1603117173463466070', '0', '0', '0', '0', '770');
INSERT INTO `gl_equip` VALUES ('27', '265', '120', '27', '13001', '2', '1603117180160530448', '0', '0', '0', '0', '490');
INSERT INTO `gl_equip` VALUES ('28', '265', '120', '28', '23001', '2', '1603117180385827469', '0', '0', '0', '0', '420');
INSERT INTO `gl_equip` VALUES ('29', '265', '120', '29', '12001', '1', '1603117190642640941', '0', '0', '0', '0', '210');
INSERT INTO `gl_equip` VALUES ('30', '265', '120', '30', '22001', '1', '1603117190817935037', '0', '0', '0', '0', '280');
INSERT INTO `gl_equip` VALUES ('31', '267', '122', '31', '11001', '3', '1603117831757055311', '0', '0', '0', '0', '630');
INSERT INTO `gl_equip` VALUES ('32', '267', '122', '32', '21001', '3', '1603117831942028503', '0', '0', '0', '0', '770');
INSERT INTO `gl_equip` VALUES ('33', '267', '122', '33', '23001', '2', '1603117835832876281', '0', '0', '0', '0', '420');
INSERT INTO `gl_equip` VALUES ('34', '267', '122', '34', '13001', '2', '1603117836157731137', '0', '0', '0', '0', '490');
INSERT INTO `gl_equip` VALUES ('35', '267', '122', '35', '12001', '2', '1603117855598222465', '0', '0', '0', '0', '210');
INSERT INTO `gl_equip` VALUES ('36', '267', '122', '36', '22001', '2', '1603117855786760849', '0', '0', '0', '0', '280');
INSERT INTO `gl_equip` VALUES ('37', '267', '122', '37', '14001', '2', '1603117884731019925', '0', '0', '0', '0', '210');
INSERT INTO `gl_equip` VALUES ('38', '267', '122', '38', '24001', '2', '1603117884925600017', '0', '0', '0', '0', '210');
INSERT INTO `gl_equip` VALUES ('39', '268', '123', '39', '11001', '2', '1603248244344647299', '0', '0', '0', '0', '630');
INSERT INTO `gl_equip` VALUES ('40', '268', '123', '40', '21001', '4', '1603248244531320277', '0', '0', '0', '0', '770');
INSERT INTO `gl_equip` VALUES ('41', '268', '123', '41', '13001', '1', '1603248981338754460', '0', '0', '0', '0', '490');
INSERT INTO `gl_equip` VALUES ('42', '268', '123', '42', '23001', '1', '1603248981509433794', '0', '0', '0', '0', '420');
INSERT INTO `gl_equip` VALUES ('43', '269', '108', '43', '11001', '1', '1603250019036142037', '0', '0', '0', '0', '630');
INSERT INTO `gl_equip` VALUES ('44', '269', '108', '44', '21001', '1', '1603250019222539693', '0', '0', '0', '0', '770');
INSERT INTO `gl_equip` VALUES ('45', '272', '126', '45', '11001', '2', '1603290021605635343', '0', '0', '0', '0', '630');
INSERT INTO `gl_equip` VALUES ('46', '272', '126', '46', '21001', '3', '1603290021791656542', '0', '0', '0', '0', '770');
INSERT INTO `gl_equip` VALUES ('47', '272', '126', '47', '13001', '2', '1603290031234297878', '0', '0', '0', '0', '490');
INSERT INTO `gl_equip` VALUES ('48', '272', '126', '48', '23001', '2', '1603290031504377837', '0', '0', '0', '0', '420');
INSERT INTO `gl_equip` VALUES ('49', '274', '127', '49', '11001', '3', '1603355294649118665', '0', '0', '0', '0', '630');
INSERT INTO `gl_equip` VALUES ('50', '274', '127', '50', '21001', '2', '1603355294849327768', '0', '0', '0', '0', '770');
INSERT INTO `gl_equip` VALUES ('51', '274', '127', '51', '13001', '2', '1603355306421785453', '0', '0', '0', '0', '490');
INSERT INTO `gl_equip` VALUES ('52', '274', '127', '52', '23001', '1', '1603355306604130781', '0', '0', '0', '0', '420');
INSERT INTO `gl_equip` VALUES ('53', '275', '128', '53', '11001', '2', '1603548112465638989', '0', '0', '0', '0', '630');
INSERT INTO `gl_equip` VALUES ('54', '275', '128', '54', '21001', '2', '1603548112639024743', '0', '0', '0', '0', '770');
INSERT INTO `gl_equip` VALUES ('55', '275', '128', '55', '13001', '1', '1603548123908669617', '0', '0', '0', '0', '490');
INSERT INTO `gl_equip` VALUES ('56', '275', '128', '56', '23001', '2', '1603548124082152724', '0', '0', '0', '0', '420');
INSERT INTO `gl_equip` VALUES ('57', '276', '129', '57', '11001', '2', '1603636874843951660', '0', '0', '0', '0', '630');
INSERT INTO `gl_equip` VALUES ('58', '276', '129', '58', '21001', '3', '1603636875030083385', '0', '0', '0', '0', '770');
INSERT INTO `gl_equip` VALUES ('59', '277', '130', '59', '11001', '2', '1603710498184164049', '0', '0', '0', '0', '630');
INSERT INTO `gl_equip` VALUES ('60', '277', '130', '60', '21001', '2', '1603710498360412530', '0', '0', '0', '0', '770');
INSERT INTO `gl_equip` VALUES ('61', '277', '130', '61', '13001', '1', '1603710507561624575', '0', '0', '0', '0', '490');
INSERT INTO `gl_equip` VALUES ('62', '277', '130', '62', '23001', '1', '1603710507722600119', '0', '0', '0', '0', '420');
INSERT INTO `gl_equip` VALUES ('63', '278', '131', '63', '11001', '2', '1603854234611693502', '0', '0', '0', '0', '630');
INSERT INTO `gl_equip` VALUES ('64', '278', '131', '64', '21001', '3', '1603854234798330321', '0', '0', '0', '0', '770');
INSERT INTO `gl_equip` VALUES ('65', '278', '131', '65', '13001', '2', '1603854245256826394', '0', '0', '0', '0', '490');
INSERT INTO `gl_equip` VALUES ('66', '278', '131', '66', '23001', '3', '1603854245492511270', '0', '0', '0', '0', '420');
INSERT INTO `gl_equip` VALUES ('67', '279', '132', '67', '11001', '1', '1604748426801788453', '0', '0', '0', '0', '630');
INSERT INTO `gl_equip` VALUES ('68', '279', '132', '68', '21001', '1', '1604748427002367949', '0', '0', '0', '0', '770');
INSERT INTO `gl_equip` VALUES ('69', '272', '126', '69', '12001', '1', '1605103860211553973', '0', '0', '0', '0', '210');
INSERT INTO `gl_equip` VALUES ('70', '272', '126', '70', '22001', '1', '1605103860421989974', '0', '0', '0', '0', '280');

-- ----------------------------
-- Table structure for `gl_item`
-- ----------------------------
DROP TABLE IF EXISTS `gl_item`;
CREATE TABLE `gl_item` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `uid` bigint(20) DEFAULT NULL,
  `accountid` int(11) DEFAULT NULL COMMENT '账号ID',
  `roleuid` int(11) DEFAULT NULL COMMENT '角色UID',
  `itemid` int(11) DEFAULT NULL COMMENT '道具唯一ID',
  `itemtype` int(11) DEFAULT NULL COMMENT '鍔熻兘id',
  `itemnum` int(11) DEFAULT NULL COMMENT '道具数量',
  `creattime` varchar(20) DEFAULT NULL COMMENT '鍒涘缓鏃堕棿',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=327 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of gl_item
-- ----------------------------
INSERT INTO `gl_item` VALUES ('271', '271', '121', '266', '1', '1', '100000', '1603117773247962091');
INSERT INTO `gl_item` VALUES ('272', '272', '121', '266', '2', '1', '20000000', '1603117773594504415');
INSERT INTO `gl_item` VALUES ('273', '273', '121', '266', '4', '1', '150000000', '1603117773841656031');
INSERT INTO `gl_item` VALUES ('274', '274', '121', '266', '101', '1', '1500000', '1603117774080315009');
INSERT INTO `gl_item` VALUES ('275', '275', '122', '267', '1', '1', '100000', '1603117790805372404');
INSERT INTO `gl_item` VALUES ('276', '276', '122', '267', '2', '1', '20000000', '1603117791110114230');
INSERT INTO `gl_item` VALUES ('277', '277', '122', '267', '4', '1', '150000000', '1603117791351527253');
INSERT INTO `gl_item` VALUES ('278', '278', '122', '267', '101', '1', '1500000', '1603117791605808285');
INSERT INTO `gl_item` VALUES ('279', '279', '123', '268', '1', '1', '100000', '1603248222371691577');
INSERT INTO `gl_item` VALUES ('280', '280', '123', '268', '2', '1', '20000000', '1603248222706854668');
INSERT INTO `gl_item` VALUES ('281', '281', '123', '268', '4', '1', '150000000', '1603248222947087750');
INSERT INTO `gl_item` VALUES ('282', '282', '123', '268', '101', '1', '1500000', '1603248223212828482');
INSERT INTO `gl_item` VALUES ('283', '283', '108', '269', '1', '1', '100000', '1603250004552155994');
INSERT INTO `gl_item` VALUES ('284', '284', '108', '269', '2', '1', '20000000', '1603250004852487020');
INSERT INTO `gl_item` VALUES ('285', '285', '108', '269', '4', '1', '150000000', '1603250005108377027');
INSERT INTO `gl_item` VALUES ('286', '286', '108', '269', '101', '1', '1500000', '1603250005381969693');
INSERT INTO `gl_item` VALUES ('287', '287', '124', '270', '101', '1', '1500000', '1603255920890470385');
INSERT INTO `gl_item` VALUES ('288', '288', '124', '270', '1', '1', '100000', '1603255921243592321');
INSERT INTO `gl_item` VALUES ('289', '289', '124', '270', '2', '1', '20000000', '1603255921555123654');
INSERT INTO `gl_item` VALUES ('290', '290', '124', '270', '4', '1', '150000000', '1603255921805342501');
INSERT INTO `gl_item` VALUES ('291', '291', '125', '271', '1', '1', '100000', '1603289236307327809');
INSERT INTO `gl_item` VALUES ('292', '292', '125', '271', '2', '1', '20000000', '1603289236607574567');
INSERT INTO `gl_item` VALUES ('293', '293', '125', '271', '4', '1', '150000000', '1603289236838822872');
INSERT INTO `gl_item` VALUES ('294', '294', '125', '271', '101', '1', '1500000', '1603289237078422588');
INSERT INTO `gl_item` VALUES ('295', '295', '126', '272', '1', '1', '100000', '1603289396683208296');
INSERT INTO `gl_item` VALUES ('296', '296', '126', '272', '2', '1', '20000000', '1603289396963989023');
INSERT INTO `gl_item` VALUES ('297', '297', '126', '272', '4', '1', '150000000', '1603289397216560731');
INSERT INTO `gl_item` VALUES ('298', '298', '126', '272', '101', '1', '1500000', '1603289397454432534');
INSERT INTO `gl_item` VALUES ('299', '299', '117', '273', '1', '1', '100000', '1603346651100545473');
INSERT INTO `gl_item` VALUES ('300', '300', '117', '273', '2', '1', '20000000', '1603346651453673807');
INSERT INTO `gl_item` VALUES ('301', '301', '117', '273', '4', '1', '150000000', '1603346651700613033');
INSERT INTO `gl_item` VALUES ('302', '302', '117', '273', '101', '1', '1500000', '1603346651949525297');
INSERT INTO `gl_item` VALUES ('303', '303', '127', '274', '1', '1', '100000', '1603355257445624014');
INSERT INTO `gl_item` VALUES ('304', '304', '127', '274', '2', '1', '20000000', '1603355257805442220');
INSERT INTO `gl_item` VALUES ('305', '305', '127', '274', '4', '1', '150000000', '1603355258046721384');
INSERT INTO `gl_item` VALUES ('306', '306', '127', '274', '101', '1', '1500000', '1603355258287676970');
INSERT INTO `gl_item` VALUES ('307', '307', '128', '275', '1', '1', '100000', '1603548082005229274');
INSERT INTO `gl_item` VALUES ('308', '308', '128', '275', '2', '1', '20000000', '1603548082343623285');
INSERT INTO `gl_item` VALUES ('309', '309', '128', '275', '4', '1', '150000000', '1603548082565912370');
INSERT INTO `gl_item` VALUES ('310', '310', '128', '275', '101', '1', '1500000', '1603548082779273656');
INSERT INTO `gl_item` VALUES ('311', '311', '129', '276', '1', '1', '100000', '1603636319453226140');
INSERT INTO `gl_item` VALUES ('312', '312', '129', '276', '2', '1', '20000000', '1603636319801761312');
INSERT INTO `gl_item` VALUES ('313', '313', '129', '276', '4', '1', '150000000', '1603636320072328431');
INSERT INTO `gl_item` VALUES ('314', '314', '129', '276', '101', '1', '1500000', '1603636320328006843');
INSERT INTO `gl_item` VALUES ('315', '315', '130', '277', '1', '1', '100000', '1603710477844118413');
INSERT INTO `gl_item` VALUES ('316', '316', '130', '277', '2', '1', '20000000', '1603710478148700240');
INSERT INTO `gl_item` VALUES ('317', '317', '130', '277', '4', '1', '150000000', '1603710478404759959');
INSERT INTO `gl_item` VALUES ('318', '318', '130', '277', '101', '1', '1500000', '1603710478650364994');
INSERT INTO `gl_item` VALUES ('319', '319', '131', '278', '1', '1', '100000', '1603854183149710676');
INSERT INTO `gl_item` VALUES ('320', '320', '131', '278', '2', '1', '20000000', '1603854183511648931');
INSERT INTO `gl_item` VALUES ('321', '321', '131', '278', '4', '1', '150000000', '1603854183766549024');
INSERT INTO `gl_item` VALUES ('322', '322', '131', '278', '101', '1', '1500000', '1603854184004964114');
INSERT INTO `gl_item` VALUES ('323', '323', '132', '279', '1', '1', '100000', '1604731649944576960');
INSERT INTO `gl_item` VALUES ('324', '324', '132', '279', '2', '1', '20000000', '1604731650317128575');
INSERT INTO `gl_item` VALUES ('325', '325', '132', '279', '4', '1', '150000000', '1604731650577680171');
INSERT INTO `gl_item` VALUES ('326', '326', '132', '279', '101', '1', '1500000', '1604731650816675619');

-- ----------------------------
-- Table structure for `gl_mail`
-- ----------------------------
DROP TABLE IF EXISTS `gl_mail`;
CREATE TABLE `gl_mail` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `userid` int(11) NOT NULL,
  `maildata` tinytext,
  `createtime` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of gl_mail
-- ----------------------------

-- ----------------------------
-- Table structure for `gl_role`
-- ----------------------------
DROP TABLE IF EXISTS `gl_role`;
CREATE TABLE `gl_role` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `roleuid` int(11) NOT NULL COMMENT '唯一UID',
  `avatar` int(11) DEFAULT NULL COMMENT '头像',
  `name` varchar(255) DEFAULT NULL COMMENT '角色名字',
  `lev` int(11) DEFAULT NULL COMMENT '角色等级',
  `channelid` int(11) NOT NULL COMMENT '渠道id',
  `accountid` int(11) DEFAULT NULL COMMENT '账号id',
  `sex` int(11) DEFAULT NULL COMMENT '0、1区分',
  `coin` varchar(255) DEFAULT NULL COMMENT '金币',
  `diamond` varchar(255) DEFAULT NULL COMMENT '砖石',
  `creattime` varchar(255) DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=280 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of gl_role
-- ----------------------------
INSERT INTO `gl_role` VALUES ('262', '262', '1', '惜灵の尔', '1', '1', '118', '0', '0', '0', '1603112550');
INSERT INTO `gl_role` VALUES ('263', '263', '1', '伍莱冷语', '1', '1', '119', '0', '40508', '100', '1603113116');
INSERT INTO `gl_role` VALUES ('264', '264', '1', '静芙の门口', '1', '1', '113', '1', '217588', '200', '1603116825');
INSERT INTO `gl_role` VALUES ('265', '265', '1', '忻得分王', '1', '1', '120', '0', '62877', '100', '1603116851');
INSERT INTO `gl_role` VALUES ('266', '266', '1', '楮の寻梦', '1', '1', '121', '1', '100000', '19989200', '1603117772');
INSERT INTO `gl_role` VALUES ('267', '267', '1', '若雁の寂寞', '1', '1', '122', '1', '177548', '19997500', '1603117790');
INSERT INTO `gl_role` VALUES ('268', '268', '1', '衫の上帝', '1', '1', '123', '1', '104896', '19997050', '1603248222');
INSERT INTO `gl_role` VALUES ('269', '269', '1', '卫菖蒲', '1', '1', '108', '1', '100692', '20000000', '1603250004');
INSERT INTO `gl_role` VALUES ('270', '270', '1', '莫尔顿の王者', '1', '1', '124', '0', '1190908', '19997300', '1603255920');
INSERT INTO `gl_role` VALUES ('271', '271', '1', '罗伊の小傻瓜', '1', '1', '125', '0', '100000', '20000000', '1603289236');
INSERT INTO `gl_role` VALUES ('272', '272', '1', '乌简慕', '1', '1', '126', '0', '112944', '19994400', '1603289396');
INSERT INTO `gl_role` VALUES ('273', '273', '1', '拉金の道尔', '1', '1', '117', '1', '100000', '20000000', '1603346650');
INSERT INTO `gl_role` VALUES ('274', '274', '1', '吕の波比', '1', '1', '127', '0', '103484', '19997350', '1603355257');
INSERT INTO `gl_role` VALUES ('275', '275', '1', '问梅の亡灵', '1', '1', '128', '0', '103484', '19997350', '1603548081');
INSERT INTO `gl_role` VALUES ('276', '276', '1', 'golang', '1', '1', '129', '0', '101384', '19997300', '1603636319');
INSERT INTO `gl_role` VALUES ('277', '277', '1', '玛吉の单身', '1', '1', '130', '1', '103484', '19997350', '1603710477');
INSERT INTO `gl_role` VALUES ('278', '278', '1', '萨罗扬卡', '1', '1', '131', '0', '108426', '19994050', '1603854182');
INSERT INTO `gl_role` VALUES ('279', '279', '1', 'asdfasdf', '1', '1', '132', '1', '2964575', '19999400', '1604731649');

-- ----------------------------
-- Table structure for `gl_runtime_event`
-- ----------------------------
DROP TABLE IF EXISTS `gl_runtime_event`;
CREATE TABLE `gl_runtime_event` (
  `id` int(10) unsigned NOT NULL,
  `configid` int(11) DEFAULT NULL COMMENT '配置表ID',
  `state` int(11) DEFAULT NULL COMMENT '活动状态',
  `loopnum` int(10) unsigned DEFAULT '0' COMMENT '循环类活动批次号，非循环类活动默认0',
  `begintime` datetime DEFAULT NULL COMMENT '开始时间',
  `endtime` datetime DEFAULT NULL COMMENT '结束时间',
  `closetime` datetime DEFAULT NULL COMMENT '关闭时间',
  `createtime` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='当前服务器运行的全局活动数据表';

-- ----------------------------
-- Records of gl_runtime_event
-- ----------------------------
