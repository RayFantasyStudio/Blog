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
            我是头像
        </a>
        <a class="item">
            <h4>userName</h4>
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
        <a class="item" id="logout-btn">
            <h4>登出</h4>
        </a>
    </div>
    <div class="ui container" id="container">
        <div class="ui card" id="filter-card">
            <div class="ui accordion field">
                <div class="title active"><i class="icon dropdown"></i> 过滤器</div>
                <div class="content field" id="filter-container">
                    <div class="ui divider"></div>
                    <div class="ui floating dropdown labeled icon button">
                        <i class="filter icon"></i>
                        <span class="text">分类过滤器</span>
                        <div class="menu">
                            <div class="ui icon search input">
                                <i class="search icon"></i>
                                <input type="text" placeholder="搜索分类">
                            </div>
                            <div class="divider"></div>
                            <div class="header">
                                <i class="tags icon"></i>
                                分类
                            </div>
                            <div class="scrolling menu">
                                <div class="item">
                                    <div class="ui red empty circular label"></div>
                                    全部
                                </div>
                                [[range .Categories]]
                                <div class="item">
                                    <div class="ui red empty circular label"></div>
                                    [[.Name]]
                                </div>
                                [[end]]
                            </div>
                        </div>
                    </div>
                    <br>
                    <label>标签过滤</label>
                    <select class="ui fluid search dropdown" multiple="">
                        <option value="">全部</option>
                        <option value="AL">Alabama</option>
                    </select>


                </div>
            </div>
        </div>
        <div class="ui tab" data-tab="article">
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
                [[range .Articles]]
                <tr>
                    <td>[[.Title]]</td>
                    <td>[[.Subtitle]]</td>
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
        <div class="ui tab" data-tab="category">
            <table class="ui celled table">
                <thead>
                <tr>
                    <th>Id</th>
                    <th>名称</th>
                    <th>创建日期</th>
                    <th>文章数</th>
                    <th>操作</th>
                </tr>
                </thead>
                <tbody>
                [[range .Categories]]
                <tr>
                    <td>[[.Id]]</td>
                    <td>[[.Name]]</td>
                    <td>创建于:[[.Created]]</td>
                    <td>共有 [[.ArticleCount]] 篇文章</td>
                    <td>
                        <a href="/admin/category?op=delete">删除</a><br>
                        <a href="/">重命名</a>
                    </td>
                </tr>
                [[end]]
                </tbody>
                <tfoot>
                <tr>
                    <th colspan="5">
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
        <div class="ui tab" data-tab="tag">
            <table class="ui celled table">
                <thead>
                <tr>
                    <th>Id</th>
                    <th>名称</th>
                    <th>文章数</th>
                    <th>操作</th>
                </tr>
                </thead>
                <tbody>
                {*TODO 标签页面待完善*}
                <tr>
                    <td>[[.Id]]</td>
                    <td>[[.Name]]</td>
                    <td>共有 [[.ArticleCount]] 篇文章</td>
                    <td>
                        <a href="/admin/tag?op=delete">删除</a><br>
                        <a href="/">重命名</a>
                    </td>
                </tr>

                </tbody>
                <tfoot>
                <tr>
                    <th colspan="5">
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
        <div class="ui tab" data-tab="reply">
            <table class="ui celled table">
                <thead>
                <tr>
                    <th>Id</th>
                    <th>文章Id</th>
                    <th>用户名</th>
                    <th>内容</th>
                    <th>时间</th>
                    <th>点赞数</th>
                    <th>引用于</th>
                    <th>操作</th>
                </tr>
                </thead>
                <tbody>
                [[range .Replies]]
                <tr>
                    <td>[[.Id]]</td>
                    <td>[[.ArticleId]]</td>
                    <td>[[.UserName]]</td>
                    <td>[[.Content]]</td>
                    <td>[[.Time]]</td>
                    <td>[[.LikeCount]]</td>
                    <td>[[.QuoteReplyId]]</td>
                    <td>
                        <a href="/admin/tag?op=delete">删除</a><br>
                        <a href="/">重命名</a>
                    </td>
                </tr>
                [[end]]

                </tbody>
                <tfoot>
                <tr>
                    <th colspan="8">
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
        <div class="ui tab" data-tab="user">
        </div>
    </div>
</div>

<!-- tabs支持 -->
<script src="//cdn.bootcss.com/jquery.address/1.6/jquery.address.min.js"></script>
<script src="/static/js/admin.js" type="text/javascript"></script>
</body>
</html>