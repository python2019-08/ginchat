# 6. 浏览器用什么api 调用 /users？

```go
func Http_template_main() {
	fmt.Println("Http_template_main......start")
	defer fmt.Println("Http_template_main......end")

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		// 模拟数据
		users := []User{
			{ID: 1, Name: "Alice", Email: "alice@example.com", IsActive: true},
			{ID: 2, Name: "Bob", Email: "bob@example.com", IsActive: false},
		}

		// 解析模板
		tmpl := template.Must(template.ParseFiles("test/templates/users.html"))

		// 执行模板，传递数据
		tmpl.Execute(w, users)
	})

	http.ListenAndServe(":8080", nil)
}
```
浏览器用什么api 调用 /users？


### 浏览器调用 `/users` API 的方法

根据您提供的代码，后端使用 Go 的 `html/template` 渲染完整的 HTML 页面，而不是返回 JSON 数据。因此，浏览器直接访问 `/users` 路径即可获取渲染后的页面。以下是几种常用的访问方式：


#### **一、直接在浏览器地址栏输入**
最直接的方式是在浏览器地址栏中输入：
```
http://localhost:8080/users
```
- 浏览器会发送 GET 请求到后端。
- 后端解析 `users.html` 模板，将 `users` 数据填充进去，返回完整的 HTML 页面。
- 浏览器渲染并显示该页面。


#### **二、使用 HTML 链接**
在 HTML 中添加链接：
```html
<a href="/users">查看用户列表</a>
```
- 点击链接时，浏览器会发送 GET 请求到 `/users`。
- 适合在导航栏或其他页面元素中集成。


#### **三、使用 JavaScript（Fetch API）**
如果需要在页面内动态加载 `/users` 内容（例如实现局部刷新），可以使用 Fetch API：
```javascript
fetch('/users')
  .then(response => response.text())  // 因为返回的是 HTML，所以用 text()
  .then(html => {
    // 将返回的 HTML 插入到页面中
    document.getElementById('user-container').innerHTML = html;
  })
  .catch(error => console.error('Error:', error));
```
- 注意：这种方式会获取完整的 HTML 页面，包括 `<html>`、`<body>` 标签，可能导致嵌套问题。
- 更适合后端提供 JSON API 的场景，这里仅作为示例。


#### **四、使用表单提交**
```html
<form action="/users" method="get">
  <button type="submit">查看用户</button>
</form>
```
- 点击提交按钮时，浏览器会发送 GET 请求到 `/users`。
- 适合需要添加额外参数的场景（通过表单字段）。


#### **五、使用浏览器开发者工具**
在 Chrome/Firefox 的开发者工具中：
1. 打开 **Network** 面板。
2. 在地址栏输入 `http://localhost:8080/users` 并回车。
3. 查看请求和响应详情。


#### **六、注意事项**
1. **模板文件路径**：
   - 确保 `test/templates/users.html` 文件存在于正确位置。
   - 相对路径是相对于程序的工作目录，而非源代码文件位置。

2. **模板内容示例**：
   ```html
   <!-- users.html -->
   <!DOCTYPE html>
   <html>
   <head>
       <title>用户列表</title>
   </head>
   <body>
       <h1>用户列表</h1>
       <ul>
           {{range .}}
           <li>{{.Name}} - {{.Email}}</li>
           {{end}}
       </ul>
   </body>
   </html>
   ```

3. **错误处理**：
   - 如果访问时出现 `404 Not Found`，检查：
     - 服务器是否正常运行。
     - 路由是否正确注册。
     - 模板文件路径是否正确。


通过以上方式，您可以在浏览器中成功访问 `/users` 并查看渲染后的用户列表页面。由于这是一个传统的服务端渲染（SSR）应用，浏览器直接接收完整的 HTML，无需额外的前端框架处理。