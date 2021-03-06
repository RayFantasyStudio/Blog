[[define "admin_reply"]]
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
            <form action="/admin" method="post">
                <input type="hidden" value="reply_delete" name="op">
                <input type="hidden" value="[[.Id]]" name="delete_reply_id">
                <button class="ui red button" type="submit">删除</button>
            </form>
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
[[end]]