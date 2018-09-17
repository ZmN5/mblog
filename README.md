# mblog

---

这是一个用`Golang`写的用`Markdown`生成静态网页的博客，其中`Markdown`转换`Html`用的工具是[blackfriday](https://github.com/russross/blackfriday)
```

compile:
make build-linux  (Linux)
make build-darwin  (Mac)

release:
make release

run:
docker run -v /path/to/markdown:/data/blog -e AUTH=yoursecret -p 8000:8000 -d fucangyu/mblog

```


Todo:
---
* 添加测试
* 支持图片
* 可配置HTTPS
* 删除文章
* 更新文章
