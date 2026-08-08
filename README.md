# go-pingdom

想临时起个服务、查个 IP、探个端口，还要装一堆东西？没必要。

对多个 URL 做 HTTP 健康检查，报告每个的可用状态和响应耗时。

## 安装

```bash
go build -o go-pingdom.exe
```

## 用法

```bash
go-pingdom https://example.com https://gitee.com
go-pingdom -t 3s https://example.com     # 超时 3 秒
```

输出示例：

```
OK    https://example.com  200 84ms
FAIL  https://不存在的域名   dial tcp: ...
```

## 说明

零依赖纯 Go，用标准库 `net/http` 发 GET。
