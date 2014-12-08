# trainining
==========

- Some simple beego samples for learning purposes

- Requirements:
	- Golang 1.3+ ** [https://golang.org/] **

- Beego - a simple & powerful Go web framework (stable v1.4.2):
	- Official site: ** [http://beego.me] **
	- Github repo - ** [https://github.com/astaxie/beego/] **

- Installation:
>go get -u github.com/bakanis/training

- Content:
	- params. 
	This sample test parameters convertion from string to int and int64 types
	To run this sample perfom:
	> bee run github.com/bakanis/training/params

	
- Sample output:
<pre>2014/12/08 19:30:49 [default.go:18] [D] post id:  19 
2014/12/08 19:30:49 [default.go:27] [D] post id as int64:  19 
2014/12/08 19:30:49 [router.go:845] [D] | GET        | /post/19                                 | 23.281545ms      | match      | /post/:id                                | 
2014/12/08 19:30:53 [default.go:18] [D] post id:  19888 
2014/12/08 19:30:53 [default.go:27] [D] post id as int64:  19888 
2014/12/08 19:30:53 [router.go:845] [D] | GET        | /post/19888                              | 19.135182ms      | match      | /post/:id                                | 	
</pre>
	
