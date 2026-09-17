# go-weather-cli

命令行查天气。给一个城市名，去 Open-Meteo 取当前气温和天气，打印一行出来。

```bash
$ go run . 上海
上海 2026-09-17 当前气温 29.4℃，天气：多云
```

城市名中英文都能填，`go run . Tokyo` 一样跑得通。

这是我练手写的，主要想摸清 Go 里怎么发 HTTP 请求、怎么解 JSON，还有 error 怎么一层层往回传。

## 跑起来

需要 Go 1.26 以上。没有第三方依赖，拉下来就能跑。

```bash
git clone https://github.com/Abeautiful-crypto/go-weather-cli.git
cd go-weather-cli
go run . 上海
```

想编译成二进制

```bash
go build -o weather .
./weather 上海
```

## 文件分工

```
main.go      入口。读命令行参数，串起下面两次调用，最后把结果格式化输出
geo.go       城市名换经纬度
weather.go   按经纬度查实时天气
```

两个查询函数都返回 `(值, error)`，出错就往上传，统一在 `main` 里打印。

## 用了哪两个接口

Open-Meteo 的公开接口，免费，不用注册也不用申请 key。

城市名换坐标

```
https://geocoding-api.open-meteo.com/v1/search?name=上海&count=1&language=zh
```

坐标查天气

```
https://api.open-meteo.com/v1/forecast?latitude=31.2222&longitude=121.4581&current=temperature_2m,weather_code
```

`language=zh` 这个参数省不掉。第一版漏了它，中文城市名一个都查不到，返回的 JSON 里连 `results` 字段都没有，程序老老实实走进「找不到城市」分支。上海、北京、广州、成都都试过，不加就查不到，加上立刻正常。

## weather_code 得自己查表

接口返回的 `weather_code` 是整数，得自己转成文字。目前映射了这些

```
0   晴        1   基本晴     2   局部多云    3   多云
45  雾       61   小雨       63  中雨       65  大雨
80  阵雨
```

没命中的会打印「未知」。广州有天返回 51，就落到这个分支了。

## 已知问题

- 51 / 53 / 55 这组毛毛雨的代码还没补进映射表
- 只取搜索结果的第一条，重名城市会认错
- `http.Get` 没设超时，网络不通的时候会一直等

## 后面想加的

- 加个 `-json` 开关，输出结构化数据，方便被别的程序调用
- 支持一次查多个城市
- 补全 WMO 天气代码
