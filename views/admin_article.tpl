[[define "admin_article"]]
<div class="ui card" id="filter-card">
    <div class="ui accordion field">
        <div class="title active"><i class="icon dropdown"></i> 过滤器</div>
        <form class="content field" id="filter-container">
            <div class="ui divider"></div>
            <div class="ui grid">
                <div class="four wide column">
                    <label class="label">文章分类：</label>
                    <select class="ui search selection dropdown" id="search-select">
                        <option value="All">All</option>
                        [[range .Categories]]
                        <option value="[[.Name]]">[[.Name]]</option>
                        [[end]]
                    </select>
                </div>
                <div class="twelve wide column">
                    <div class="label">文章标签：</div>
                    <select class="ui search selection dropdown" multiple="">
                        <option value="">全部</option>
                        [[range .Tags]]
                        <option value="[[.Name]]">[[.Name]]</option>
                        [[end]]
                    </select>
                </div>
            </div>
            <div class="ui divider"></div>
            <div class="ui grid">
                <div class="four wide column">
                    <div class="ui radio checkbox">
                        <input type="radio" name="order_create" checked="" tabindex="0" class="hidden">
                        <label>按照创建日期升序</label>
                    </div>
                    <div class="ui radio checkbox">
                        <input type="radio" name="order_create" checked="" tabindex="0" class="hidden">
                        <label>按照创建日期降序</label>
                    </div>
                </div>
                <div class="four wide column">
                    <div class="ui radio checkbox">
                        <input type="radio" name="order_update" checked="" tabindex="0" class="hidden">
                        <label>按照更新日期升序</label>
                    </div>
                    <div class="ui radio checkbox">
                        <input type="radio" name="order_update" checked="" tabindex="0" class="hidden">
                        <label>按照更新日期降序</label>
                    </div>
                </div>
                <div class="four wide column">
                    <div class="ui radio checkbox">
                        <input type="radio" name="order_last_reply_time" checked="" tabindex="0" class="hidden">
                        <label>按照最后回复升序</label>
                    </div>
                    <div class="ui radio checkbox">
                        <input type="radio" name="order_last_reply_time" checked="" tabindex="0" class="hidden">
                        <label>按照最后回复降序</label>
                    </div>
                </div>
                <div class="four wide column">
                    <div class="ui radio checkbox">
                        <input type="radio" name="order_view_count" checked="" tabindex="0" class="hidden">
                        <label>按照查看数升序</label>
                    </div>
                    <div class="ui radio checkbox">
                        <input type="radio" name="order_view_count" checked="" tabindex="0" class="hidden">
                        <label>按照查看数降序</label>
                    </div>
                </div>
            </div>
            <div class="ui divider"></div>
            <div style="margin-bottom: 70px">
                <button class="ui red right floated button">重置</button>
                <button class="ui green right floated button">确认</button>
            </div>
        </form>
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