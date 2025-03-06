# ai-proxy

根据配置文件，转发 openai api 请求到不同的服务提供商

```yaml
addr: localhost:8080
providers:
  - name: siliconflow
    baseURL: https://api.siliconflow.cn/v1
    apiKey: sk-xxx
    timeoutMs: 10000
```
