{{ define "header"}}
<!doctype html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport"
		  content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
	<meta http-equiv="X-UA-Compatible" content="ie=edge">
	<title>{{ .Title}} - marunのbeego博客</title>
	<link rel="shortcu icon" href="/static/img/favicon.jpg">
	<link rel="stylesheet" href="/static/css/bootstrap.min.css">
</head>
<body>

<header class="navbar navbar-static-top bs-docs-nav" id="top">
	<div class="container">
		<div class="navbar-header">
			<button class="navbar-toggle collapsed" type="button" data-toggle="collapse" data-target="#bs-navbar" aria-controls="bs-navbar" aria-expanded="false">
				<span class="sr-only">Toggle navigation</span>
				<span class="icon-bar"></span>
				<span class="icon-bar"></span>
				<span class="icon-bar"></span>
			</button>
			<a href="/" class="navbar-brand">我的博客</a>
		</div>
		<nav id="bs-navbar" class="collapse navbar-collapse">
			<ul class="nav navbar-nav">
				<li class="active"><a href="/" >首页</a></li>
				<li><a href="/category">分类</a></li>
				<li><a href="/topic">文章</a></li>
			</ul>

			<ul class="nav navbar-nav navbar-right">
				<li><a href="/login" target="_blank">登陆</a></li>
			</ul>
		</nav>
	</div>
</header>
{{ end }}