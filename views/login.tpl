<!DOCTYPE html>
<html lang="en">
<head>
    <title>登录</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link rel="image_src" type="image/jpeg" href="/static/img/logo.png">
    <meta name="viewport" content="width=device-width, initial-scale=1, minimum-scale=1, maximum-scale=1">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0">
    <link href="//cdn.bootcss.com/semantic-ui/2.1.8/semantic.min.css" class="ui" rel="stylesheet">
    <link href="/static/css/login.css" rel="stylesheet">
    <link href="/static/css/base.css" rel="stylesheet">
    <script src="//cdn.bootcss.com/jquery/2.2.4/jquery.js"></script>
    <script src="//cdn.bootcss.com/semantic-ui/2.1.8/semantic.min.js"></script>
</head>
<body>
<div class="ui middle aligned center aligned grid">
    <div class="column">
        <h2 class="ui teal image header">
            <img src="/static/img/RFS_logo_bright.png" class="image">
            <div class="content">
                登录你的账户
            </div>
        </h2>
        <form class="ui large form" method="post" action='[[urlfor "LoginController.Post"]]'>
            <div class="ui stacked segment">
                <div class="field">
                    <div class="ui left icon input">
                        <i class="user icon"></i>
                        <input type="text" name="name" placeholder="用户名">
                    </div>
                </div>
                <div class="field">
                    <div class="ui left icon input">
                        <i class="lock icon"></i>
                        <input type="password" name="pwd" placeholder="密码">
                    </div>
                </div>
                <div class="ui fluid large teal submit button">登录</div>
            </div>

            <div class="ui error message"></div>
            [[if .Message ]]
            <div class="ui message">
                <p>[[.Message]]</p>
            </div>
            [[end]]

        </form>
    </div>
</div>

<script src="/static/js/login.js"></script>

</body>
</html>