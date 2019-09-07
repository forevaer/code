结构及说明
===
```text
├── data
│   └── data.json   // 数据
├── main.go         // 程序入口
├── matchers
│   └── rss.go      // rss匹配器
└── search
    ├── default.go  // 默认匹配器
    ├── feed.go     // 文件解析
    ├── match.go    // 匹配过程
    └── search.go   // 程序调度

```
# [main](./main.md)
主程序入口
# [data](./data)
json数据文件
# [matchers](./matchers)
匹配器定义
# [search](./search)
利用匹配器进行搜索匹配

