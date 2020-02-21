# git robot to github or gogs ...
根据git的一些事件去触发处理一些事

处理issue中的一些诸如`/pay 100`的指令，可以自定义指令和注册指令的处理器。

使用场景：

> 自动回复issue

比如issue一天没回复，自动回复以让提issue的人莫着急

> 自动提醒

如kubernetes中/sig 指令会把消息推到对应的兴趣小组

> 自动化测试

`/test e2e test/test.sh`

> 触发drone promote事件

`/promote 1 test` 在issue里回复这个，就会触发CI/CD里面下面定义的一个pipeline

```
      git      robot     drone
 issue |         |         |
------>| event   |         |
       |-------->|         |
       |         | promote |
       |         |-------->|
       |         |         | do what ever you want
       |         |         |
       V         V         V
```

> 自动merge代码

`/merge` 指令可以自动merge代码，还可以在merge之前之后做一些事，比如记录下PR的作者，发邮件，等等

> 其它

打标签，关闭超时issue等