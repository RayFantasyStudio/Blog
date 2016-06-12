[[template "header" .]]
<div class="ui" style="margin: 0 5% 0 5%;">
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
                    <img src="/images/avatar2/small/lena.png" class="ui mini rounded image">
                    <a class="content" href="/article?page=1&order=created&desc=true&by_uid=[[$article.Author.Id]]">[[index $authors $index]]
                        <div class="sub header">人力资源 </div>
                    </a>
                </h4></td>
            <td>
                <a href="/article/view?id=[[$article.Id]]">[[$article.Title]]</a>
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
