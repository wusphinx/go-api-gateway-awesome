# go-api-gateway-awesome
api gateway implemented by go in real world

| API Gateway              | 
| -------------------------------       | 
| [lura](/lura)         | 
| [caddy](/caddy)     | 
| [traefik](/traefik)   | 
| [easegress](/easegress)   | 

# 对比
这几个网关均是笔者过去一年多在实际工作中调研过的网关，其中lura、traefik均有在生产环境中使用过，生产环境日均请求总量不超过10w，因为主要是公司内部使用，QPS也很小。就实际使用体验来讲，这几个网关各有优势

- lura 最易上手，但配置冗余
- traefik 配置可读性高，文档丰富，配置支持热更新，社区活跃
- caddy 配置最为简洁，官方文档示例不是很全(这也是当初没有采用的原因)
- easegress是国产网关，目前还是快速迭代中，使用体验也挺不错，与traefik风格类似，当然因为项目比较新，示例较少，喜欢尝鲜的同学可以一试

如果流量较小的项目，笔者首推traefik，定制插件的开发也相对容易一些。


refer:
- [如何使用 docker buildx 构建跨平台 Go 镜像](https://waynerv.com/posts/building-multi-architecture-images-with-docker-buildx/)