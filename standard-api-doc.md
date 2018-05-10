# 标准API文档

- [登陆](#登陆)
- [获取指定帖子(集合)](#获取指定帖子(集合))
- [帖子评论](#帖子评论)

## 登陆

### 应用场景

用户使用树洞就必须登录，否则无法使用包括浏览在内的任何功能。

### 接口链接

POST https://api.shudong.cn/login

### 请求参数

变量名      | 必填 | 类型 | 示例值 | 描述 |
-------  | ------ | ---- | ----| ---- |
username | Y | String(32) | example_username | 用户名 |
password | Y | String(128) | example_password | 密码 |

举例如下：

```text
username=example_username&password=example_passwd
```

### 返回结果

变量名 | 必填 | 类型 | 示例值 | 描述 |
----  | ------ | ---- | ---- | ---- |
msg | Y | String(32) | OK | HTTP Response Code对应的信息 |
data | Y | Object | {"username":"user"} | 包含当前用户信息 |

当登录失败时， `data`为`null`。

### 错误码

名称 | 描述 | 原因 | 解决方案 |
---- | ---- | ---- | ---- |
Unauthorized | 未授权 | 登录失败，用户名或密码错误 | 检查用户名和密码 |
Internal Server Error | 中间服务器出错 | 网络环境出现问题 | 检查网络问题 |

## 获取指定帖子(集合)

### 应用场景

用户进入广场页面之后：需要获取并显示最近的帖子，或是按照一定的条件进行搜索

### 接口链接

URL地址： GET https://api.shudong.cn/posts

### 请求参数

变量名      | 必填 | 类型 | 示例值 | 描述 |
-------  | ------ | ---- | ----| ---- |
postid  | N | Int | 2 | 选择指定postid的帖子。当有效指定此参数时，其他参数均被忽略。(即对/posts/{postId}的简单重复) |
limit  | N | Int | 10 | 限制一次最多返回posts的数量，默认为10 |
offset  | N | Int | 10 | 配合limit使用，设置从第几条查询结果开始返回，缺省为0 |

举例如下：

```text
GET https://api.shudong.cn/posts?limit=15&offset=10
GET https://api.shudong.cn/posts?postid=3
```

### 返回结果

变量名 | 必填 | 类型 | 示例值 | 描述 |
----  | ------ | ---- | ---- | ---- |
msg | Y | String(32) | OK | HTTP Response Code对应的信息 |
data | Y | Array\<Object\> | [{"postId":1,"author":"test", ...}, ...] | 符合查询条件的posts的Object数组

当msg不为2xx(HTTP状态码)对应的信息时，`data`均为`null`

### 错误码

名称 | 描述 | 原因 | 解决方案 |
---- | ---- | ---- | ---- |
Unauthorized | 未授权 | 用户未登陆 | 提示用户登录 |
Internal Server Error | 中间服务器出错 | 网络环境出现问题 | 检查网络问题 |

## 帖子评论

### 应用场景

打开一个帖子之后，在帖子下方进行评论。

### 接口链接

POST https://api.shudong.cn/posts/{postid}/comments

### 请求参数

变量名      | 必填 | 类型 | 示例值 | 描述 |
-------  | ------ | ---- | ----| ---- |
content  | Y | String | "example comment" | 评论内容 |

参数用json转义即可。举例如下：

```json
{ "content": "new example comment" }
```

### 返回结果

变量名 | 必填 | 类型 | 示例值 | 描述 |
----  | ------ | ---- | ---- | ---- |
msg | Y | String(32) | OK | HTTP Response Code对应的信息 |
data | Y | Object | (见下方) | 包含当前用户信息 |

当msg不为2xx(HTTP状态码)对应的信息时，`data`均为`null`。

反之，当状态码为`2xx`时，data将是一个Object，其包含的属性值如下：

变量名 | 必填 | 类型 | 示例值 | 描述 |
----  | ------ | ---- | ---- | ---- |
commentId | Y | Int | 17 | 该新生成的评论的commentId |
author | Y | Object | (见下方) | 包含当前用户信息 |
relatedPostId | Y | Int | 3 | 该条评论与哪一条post相关 |
content | Y | String | "example content" | 新生成的评论的内容 |
like_count | Y | Int | 7 | 当前有多少人给这条评论点赞

### 错误码

名称 | 描述 | 原因 | 解决方案 |
---- | ---- | ---- | ---- |
Unauthorized | 未授权 | 用户未登陆 | 提示用户登录 |
Internal Server Error | 中间服务器出错 | 网络环境出现问题 | 检查网络问题 |

### 系统顺序图

![Comment](https://raw.githubusercontent.com/Chun-Ge/documents/master/docs/model-docs/system-sequence-diagram/15331302-comment.png)
