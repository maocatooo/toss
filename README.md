# toss
文件上传到OSS中, MinIO, awsOSS(亚马逊oss), tencentOSS(腾讯云), aliOSS(阿里云), qiniuOSS(七牛云)...


### 支持
没标记的暂时没有进行单元测试,不保证可用性, 理论上只要是支持s3协议的都可以使用

- [x] MinIO
- [x] tencentOSS
- [ ] aliOSS
- [ ] qiniuOSS
- [ ] awsOSS

### 安装

```shell
go install github.com/maocatooo/toss
```

### 使用

```shell
// 默认当前目录下的config.yaml
toss 

toss --config config_path/config.yaml 
```

### ❗️注意

**toss 会删除oss target目录下的所有文件**

### 配置文件

#### MinIO
```yaml
AK: AKIAIOSFODNN7EXAMPLE // 必填
SK: wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY // 必填
region: cn-hangzhou // MinIO可选
bucket: test123  // bucket名称必填
endpoint: http://127.0.0.1:9000 // 必填
source: ../toss // 本地文件/文件夹路径 必填
target: /toss  //目标oss文件夹路径 必填 
```

#### tencentOSS
bucket 名称已经在endpoint中包含, 无需再次填写

region: 云对象存储地域 必填
```yaml
AK: AKIDKafcXIqBLQLLj3MdlfO4f3DnPfovjcVH
SK: pg4RwLZqxut5GWdXFnImF2f6LVTJHHIG
region: cn-hangzhou
bucket:
endpoint: https://maocat-1309234433.cos.ap-nanjing.myqcloud.com
source: ../oss
target: /
```