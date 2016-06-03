[[define "admin_category"]]
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
[[end]]