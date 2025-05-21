# xcheckhub

## 项目定义
1. 一堆web应用程序
2. 每个都能正常运行：docker-compose
3. 统一的页面/API接口
  - /                    页面，展示所有api
  - /bug_key/get         漏洞key+method
4. 统一的测试脚本
  - 语言补充
  - 框架补充
  - 自动打分

高级特性：
5. 不同框架间的重复代码：service、rink点等可以复用
6. 自动生成代码，支持修改函数名、变量名，插入无效代码，混淆等，防止友商hardcode


## 目录结构
语言-框架_version-可运行的docker-compose目录
```
.
├── php
│   ├── symfony_5.x
│   │   └── docker-compose.yaml
│   └── symfony_4.x
│       └── docker-compose.yaml
└── py
    ├── django_4.x
    │   └── docker-compose.yaml
    └── flask_3.x
        └── docker-compose.yaml
```
## 开发规范
沿用xchecker