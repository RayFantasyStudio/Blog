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
        <p>RayFantasyStudio官方Blog</p>
        [[if .Message ]]
        <div class="ui message" id="flash">
            <p>[[.Message]]</p>
        </div>
        [[end]]
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
            <a href="/" class="item">Blog</a>
            <a href="/article?page=1&order=updated&desc=true&by_uid=0" class="item">文章</a>
            <div class="right item">
                <div class="ui icon right input" >
                    <input type="text" placeholder="Search..." id="search_content">
                    <i class="inverted circular search link icon" onclick="search()"></i>
                </div>
                <script>
                    function search() {
                        var search_key = $('#search_content').val();
                        location.assign("/search" + "?key=" + search_key);
                    }
                </script>
                <div class="ui [[if .IsLogin]]red[[else]]green[[end]] buttons">
                    [[if .IsLogin]]
                    <a href='[[urlfor "LoginController.Logout"]]' class="ui button">
                        注销
                    </a>
                    <div class="ui floating dropdown icon button">
                        <i class="dropdown icon"></i>
                        <div class="menu">
                            <a class="item" href="/admin"><i class="setting icon"></i> Blog管理</a>
                        </div>
                    </div>
                    [[else]]
                    <a class="ui button" href='[[urlfor "LoginController.Get"]]'>
                        登陆
                    </a>
                    [[end]]

                </div>
            </div>
        </div>
    </div>
</div>
[[end]]