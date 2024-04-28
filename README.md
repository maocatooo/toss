# toss
文件上传到OSS中, MinIO, awsOSS(亚马逊oss), tencentOSS(腾讯云), aliOSS(阿里云), qiniuOSS(七牛云)...


### 支持
没标记的暂时没有进行单元测试,不保证可用性, 理论上只要是支持s3协议的都可以使用

- [x] MinIO
- [x] tencentOSS
- [x] aliOSS
- [ ] qiniuOSS
- [x] awsOSS

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
**bucket 前面要加 /**
```yaml
AK: AKIAIOSFODNN7EXAMPLE // 必填
SK: wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY // 必填
region: cn-hangzhou // MinIO可选
bucket: /test123  // bucket名称必填
endpoint: http://127.0.0.1:9000 // 必填
source: ../toss // 本地文件/文件夹路径 必填
target: /toss  //目标oss文件夹路径 必填 
```

#### tencentOSS
**bucket 名称已经在endpoint中包含, 固定为 /**

**region: 云对象存储地域endpoint中包含 必填**

```yaml
AK: AKIAIOSFODNN7EXAMPLE
SK: wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
region: cn-hangzhou
bucket: /
endpoint: https://maocat-1309234433.cos.ap-nanjing.myqcloud.com
source: ../oss
target: /
```
#### aliOSS

```yaml
AK: AKIAIOSFODNN7EXAMPLE
SK: wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
region: oss-cn-beijing
bucket: maocat
endpoint: https://oss-cn-beijing.aliyuncs.com
source: ../oss
target: /
```

#### aws
**bucket 必须
**endpoint 非必须
**region: 必须
```yaml
AK: AKIAIOSFODNN7EXAMPLE
SK: wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
region: oss-cn-beijing
bucket: maocat
endpoint: 
source: ../oss
target: /
```