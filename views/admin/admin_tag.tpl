[[define "admin_tag"]]
<table class="ui celled table">
    <thead>
    <tr>
        <th>Id</th>
        <th>名称</th>
        <th>操作</th>
    </tr>
    </thead>
    <tbody>
    [[range .Tags]]
    <tr>
        <td>[[.Id]]</td>
        <td>[[.Name]]</td>
        <td>
            <a href="/admin/tag?op=delete">删除</a><br>
            <a href="/">重命名</a>
        </td>
    </tr>
    [[end]]

    </tbody>
    <tfoot>
    <tr>
        <th colspan="3">
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