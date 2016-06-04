[[define "admin_tag"]]
<table class="ui striped fixed compact table">
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
            <form action="/admin" method="post">
                <input type="hidden" value="tag_delete" name="op">
                <input type="hidden" value="[[.Id]]" name="delete_tag_id">
                <button class="ui red button" type="submit">删除</button>
            </form>
            <div class="ui accordion field">
                <div class="title"><i class="icon dropdown"></i> 重命名</div>
                <div class="content field">
                    <form class="form" action="/admin" method="post">
                        <div class="ui input">
                            <input type="hidden" value="tag_rename" name="op">
                            <input type="text" placeholder="输入新名称" name="rename_tag_name">
                            <input type="hidden" value="[[.Id]]" name="rename_tag_id">

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