## 移远接口调用说明

## Usage

```
func main(){
    client := NewClient("123", "AAAAAAAAAA")
    resp, err := client.GetCardInfo("89860478102050071500")
    if err != nil{
        log.Println(err)
    }
}
```

## API 文档

[API 文档](https://iot.quectel.com/connect_api_doc.html#singleCardRestore) 是移远物联卡平台接口文档

## 申请API key / secret
### 申请说明
- 用户首先需要在 [物联卡管理平台](https://iot.quectel.com)上申请自己的APP_KEY和APP_SECRET.

## License

```
Copyright 2021 MingGuang

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```