> Sample output:
<pre>
2014/12/08 19:30:42 [INFO] Start building...
2014/12/08 19:30:43 [SUCC] Build was successful
2014/12/08 19:30:43 [INFO] Restarting params ...
2014/12/08 19:30:43 [INFO] ./params is running...
2014/12/08 19:30:43 [WARN] read: interrupted system call
2014/12/08 19:30:43 [app.go:96] [I] http server Running on :8080
2014/12/08 19:30:49 [default.go:18] [D] post id:  19 
2014/12/08 19:30:49 [default.go:27] [D] post id as int64:  19 
2014/12/08 19:30:49 [router.go:845] [D] | GET        | /post/19                                 | 23.281545ms      | match      | /post/:id                                | 
2014/12/08 19:30:53 [default.go:18] [D] post id:  19888 
2014/12/08 19:30:53 [default.go:27] [D] post id as int64:  19888 
2014/12/08 19:30:53 [router.go:845] [D] | GET        | /post/19888                              | 19.135182ms      | match      | /post/:id                                | 
</pre>
