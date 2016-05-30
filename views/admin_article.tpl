<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>管理文章</title>
    <link href="//cdn.bootcss.com/semantic-ui/2.1.8/semantic.min.css" rel="stylesheet">
    <link href="../static/css/admin.css" rel="stylesheet">
    <script src="//cdn.bootcss.com/jquery/3.0.0-rc1/jquery.js"></script>
    <script src="//cdn.bootcss.com/semantic-ui/2.1.8/semantic.min.js"></script>
</head>
<body>
<div id="main">
    <div class="ui  inverted vertical menu" id="nav">
        <a class="item" style="">
            我是头像
        </a>
        <a class="item">
            <h4>userName</h4>
        </a>
        <a class="active item">
            <h4>文章管理</h4>
        </a>
        <a class="item">
            <h4>分类管理</h4>
        </a>
        <a class="item">
            <h4>标签管理</h4>
        </a>
        <a class="item">
            <h4>回复管理</h4>
        </a>
        <a class="item" style="background-color: red">
            <h4>登出</h4>
        </a>
    </div>
    <div id="container">
        <table class="ui celled table">
            <thead>
            <tr>
                <th>标题</th>
                <th>副标题</th>
                <th>标签</th>
                <th>时间</th>
                <th>回复数/查看数</th>
                <th>操作</th>
            </tr>
            </thead>
            <tbody>
            [[range .articlelist]]
            <tr>
                <td>[[.Title]]</td>
                <td>[[.Subtitle]]</td>
                <td>[[.tag]]</td>
                <td>创建于:[[.Created]]<br>
                    更新于:[[.Updated]]
                </td>
                <td>被查看 [[.Created]] 次<br>
                    共 [[.Updated]] 条回复
                </td>
                <td>
                    <a href="/admin/article?op=delete">删除</a><br>
                    <a href="/">修改</a>
                </td>
            </tr>
            [[end]]
            </tbody>
            <tfoot>
            <tr>
                <th colspan="6">
                    <div class="ui right floated pagination menu">
                        <a class="icon item">
                            <i class="left chevron icon"></i>
                        </a>
                        <a class="item">1</a>
                        <a class="item">2</a>
                        <a class="item">3</a>
                        <a class="item">4</a>
                        <a class="icon item">
                            <i class="right chevron icon"></i>
                        </a>
                    </div>
                </th>
            </tr>
            </tfoot>
        </table>
    </div>
</div>

</body>
</html>