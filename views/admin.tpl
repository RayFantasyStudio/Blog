<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>管理</title>
    <link href="//cdn.bootcss.com/semantic-ui/2.1.8/semantic.min.css" class="ui" rel="stylesheet">
    <link href="/static/css/admin.css" rel="stylesheet">
    <script src="//cdn.bootcss.com/jquery/2.2.4/jquery.min.js"></script>
    <script src="//cdn.bootcss.com/semantic-ui/2.1.8/semantic.min.js"></script>

</head>
<body>

<div id="main">
    <div class="ui inverted vertical tabs menu" id="nav">
        <a class="item" style="">
            <img src="/static/img/user_avatar/[[.UserName]]_avatar.png"  style="width: 200px;height: 200px" id="nav_user_avatar" onerror="javascript:this.src='/static/img/user_avatar/user.png'">
        </a>
        <a class="item">
            <h4>[[.UserName]]</h4>
        </a>
        <a class="item" data-tab="article">
            <h4>文章管理</h4>
        </a>
        <a class="item" data-tab="category">
            <h4>分类管理</h4>
        </a>
        <a class="item" data-tab="tag">
            <h4>标签管理</h4>
        </a>
        <a class="item" data-tab="reply">
            <h4>回复管理</h4>
        </a>
        <a class="item" data-tab="user">
            <h4>账户信息更改</h4>
        </a>
        <a class="item" id="logout-btn" href='[[urlfor "LoginController.Logout"]]' >
            <h4>登出</h4>
        </a>
    </div>
    <div class="ui container" id="container">
        <div class="ui tab" data-tab="article">
            [[template "admin_article" .]]
        </div>
        <div class="ui tab" data-tab="category">
            [[template "admin_category" .]]
        </div>
        <div class="ui tab" data-tab="tag">
            [[template "admin_tag" .]]
        </div>
        <div class="ui tab" data-tab="reply">
            [[template "admin_reply" .]]

        </div>
        <div class="ui tab" data-tab="user">
            [[template "admin_setting" .]]
        </div>
    </div>
</div>

<!-- tabs支持 -->
<script src="//cdn.bootcss.com/jquery.address/1.6/jquery.address.min.js"></script>
<script src="/static/js/base.js" type="text/javascript"></script>
<script src="/static/js/admin.js"></script>
</body>
</html>