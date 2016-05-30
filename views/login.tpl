<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>登录</title>
    <link href="//cdn.bootcss.com/semantic-ui/2.1.8/semantic.min.css" rel="stylesheet">
    <link href="/static/css/login.css" rel="stylesheet">
    <link href="../static/css/admin.css" rel="stylesheet">
    <script src="//cdn.bootcss.com/jquery/3.0.0-rc1/jquery.js"></script>
    <script src="//cdn.bootcss.com/semantic-ui/2.1.8/semantic.min.js"></script>
</head>
<body >
<div id="container">
    <div style=""class="ui card" id="head-title">
        <h1>登录你的账户</h1>
    </div>
    <div  class="ui card" id="login-form">
        <form method="post" action="/login">
            <div class="ui form">
                <div class="field">
                    <div class="ui left icon input">
                        <input type="text" placeholder="输入你的用户名" name="username">
                        <i class="user icon"></i>
                    </div>
                </div>
                <div class="field">
                    <div class="ui left icon input">
                        <input type="password" placeholder="输入你的密码" name="pwd">
                        <i class="lock icon"></i>
                    </div>
                </div>
                <div class="field">
                    <div class="ui checkbox">
                        <input type="checkbox" tabindex="0" class="hidden">
                        <label>记住我</label>
                    </div>
                </div>
                <div class="field">
                    <div class="ui buttons">
                        <a class="ui button" href="/">返回</a>
                        <div class="or"></div>
                        <button class="ui positive button" type="submit">登录</button>
                    </div>
                </div>
            </div>
        </form>
    </div>
</div>

</body>
</html>