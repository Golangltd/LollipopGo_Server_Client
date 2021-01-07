shop_proto = {
[1] = { shop_type = 1, money = { {1, 1, 1}, {1, 2, 1}, }, refresh_time = "0 45 18 * * ?", goods_num_max = 8, refresh_byself = 1, refresh_free = 2, refresh_price = { {1, 1, 20}, }, refresh_price_add = { {5, 1, 1, 20}, }, },
[2] = { shop_type = 0, money = { {1, 1, 1}, {1, 10, 1}, }, refresh_time = "0 0 8 1,10,20 * ?", goods_num_max = 4, refresh_byself = 0, refresh_free = 0, },
[3] = { shop_type = 0, money = { {1, 1, 1}, {1, 11, 1}, }, refresh_time = "0 15 10 15 * ？", goods_num_max = 16, refresh_byself = 0, refresh_free = 0, },
[4] = { shop_type = 1, money = { {1, 1, 1}, {1, 8, 1}, }, refresh_time = "0 0 5 * * ? *", goods_num_max = 16, refresh_byself = 0, refresh_free = 0, },
[5] = { shop_type = 1, money = { {1, 1, 1}, {1, 5, 1}, }, refresh_time = "0 0 8 * * ?", goods_num_max = 16, refresh_byself = 0, refresh_free = 0, },
}
return shop_proto
