[[define "header"]]
<!DOCTYPE html>
<html>
<head>
    <title>[[.Title]]</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, minimum-scale=1, maximum-scale=1">
    <link rel="image_src" type="image/jpeg" href="/static/img/logo.png">
    <link href="//cdn.bootcss.com/semantic-ui/2.1.8/semantic.min.css" class="ui" rel="stylesheet">
    <link href="/static/css/base.css" rel="stylesheet">
    <script src="//cdn.bootcss.com/jquery/2.2.4/jquery.js"></script>
    <script src="//cdn.bootcss.com/semantic-ui/2.1.8/semantic.min.js"></script>

</head>
<body>
<div>
    <div class="ui main text container">
        <h1 class="ui header">[[.Owner]]的Blog</h1>
        <p>This example shows how to use lazy loaded images, a sticky menu, and a simple text container</p>
    </div>
    <div class="ui borderless main menu">
        <div class="ui text container">
            <div href="#" class="header item">
                [[if .IsLogin]]
                [[.User.Name]]
                [[else]]
                访客
                [[end]]
            </div>
            <a href="#" class="item">Blog</a>
            <a href="#" class="item">Articles</a>
            [[if .IsLogin]]
            <a href='[[urlfor "LoginController.Logout"]]' class="ui right item">
                注销
            </a>
            [[else]]
            <a href='[[urlfor "LoginController.Get"]]' class="ui right item">
                登陆
            </a>
            [[end]]
            <div class="ui icon input">
                <input type="text" placeholder="Search...">
                <i class="inverted circular search link icon"></i>
            </div>
        </div>
    </div>
</div>
[[end]]