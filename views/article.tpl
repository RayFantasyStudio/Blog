[[template "header" .]]
<div class="ui" style="margin: 0 5% 0 5%;">
    <a class="positive ui button" href="/article/create">新建文章</a>
    <table class="ui  basic striped fixed celled  table" style="background: white" >
        <thead>
        <tr>
            <th class="two wide">作者</th>
            <th>标题</th>
            <th>副标题</th>
            <th>分类</th>
            <th>更新时间</th>

        </tr></thead>
        <tbody>
        [[$authors := .Authors]]
        [[range $index, $article := .Articles]]
        <tr>
            <td>
                <h4 class="ui image header">
                    <img src="/static/img/user_avatar/[[index $authors $index]]_avatar.png" class="ui mini rounded image">
                    <a class="content" href="/article?page=1&order=created&desc=true&by_uid=[[$article.Author.Id]]">[[index $authors $index]]
                        <div class="sub header">    </div>
                    </a>
                </h4></td>
            <td>
                <a href="/article/view?id=[[$article.Id]]&quote_rid=0">[[$article.Title]]</a>
            </td>
            <td>
                [[$article.Subtitle]]
            </td>
            <td>
                <a href="/article?page=1&order=[[$.Order_tmp]]&desc=false&by_uid=[[$.User_tmp]]&cate=[[$article.Category]]">[[$article.Category]]</a>
            </td>
            <td>
                [[$article.Updated]]
            </td>
        </tr>
        [[end]]
        </tbody>
    </table>
    [[template "paginator" .Paginator]]
</div>
[[template "footer"]]
