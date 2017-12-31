练习要求：
-----

依据文档图6-1，用中文描述 Reactive 动机
使用 go HTTPClient 实现图 6-2 的 Naive Approach
为每个 HTTP 请求设计一个 goroutine ，利用 Channel 搭建基于消息的异步机制，实现图 6-3
对比两种实现，用数据说明 go 异步 REST 服务协作的优势
思考： 是否存在一般性的解决方案？

![这里写图片描述](http://img.blog.csdn.net/20171231223811482?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSE9NRVJVTklU/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)


![这里写图片描述](http://img.blog.csdn.net/20171231223828097?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSE9NRVJVTklU/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

![这里写图片描述](http://img.blog.csdn.net/20171231223845908?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSE9NRVJVTklU/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

运行
--
![这里写图片描述](http://img.blog.csdn.net/20171231223914682?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvSE9NRVJVTklU/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
