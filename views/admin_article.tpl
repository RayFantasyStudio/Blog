[[define "admin_article"]]
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
<table class="ui striped fixed compact table">
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
        <td>Tag</td>
        <td>创建于:[[.Created]]<br>
            更新于:[[.Updated]]
        </td>
        <td>被查看 [[.ViewCount]] 次<br>
            共 [[.ReplyCount]] 条回复
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
[[end]]