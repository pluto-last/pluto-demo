-- 支付宝小程序源码获取方案
-- 获取appid: 分享出来的小程序链接里面有appid
-- /data/data/com.eg.android.AlipayGphone/files 目录下进行全局搜索 grep -rl "appid" ./
-- 0.data结尾的就是源码，一般会有多个，一个个的找

-- 最后运行data_extract.py，解压所有0.data

-- 支付宝版本要求
-- 10.6 OK ,10.7 不行