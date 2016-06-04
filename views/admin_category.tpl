[[define "admin_category"]]
<table class="ui striped fixed compact table">
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
            <form action="/admin" method="post">
                <input type="hidden" value="category_delete" name="op">
                <input type="hidden" value="[[.Id]]" name="delete_category_id">
                <button class="ui red button" type="submit">删除</button>
            </form>
            <div class="ui accordion field">
                <div class="title"><i class="icon dropdown"></i> 重命名</div>
                <div class="content field">
                    <form class="form" action="/admin" method="post">
                        <div class="ui input">
                            <input type="text" placeholder="输入新名称" name="new_category">
                            <input type="hidden" value="[[.Name]]" name="former_category">
                            <input type="hidden" value="category_rename" name="op">
                        </div>
                        <button type="submit" class="ui green button">修改</button>
                    </form>
                </div>
            </div>
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
[[end]]