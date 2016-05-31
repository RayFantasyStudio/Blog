[[template "header" .]]
<div class="ui text container">
    <div class="ui relaxed divided items">
        [[range $index, $article := .Articles]]
        <div class="item">
            <div class="content">
                <a class="header">[[$article.Title]]</a>
                <div class="meta">
                    <a>[[$article.Created | SinceTime]]</a>
                    <a>Category</a>
                </div>
                <div class="description">
                    [[$article.Content]]
                </div>
                <div class="extra">
                    <div class="ui right floated primary button">
                        详情
                        <i class="right chevron icon"></i>
                    </div>
                    <div class="ui label">Tag</div>
                </div>
            </div>
        </div>
        [[end]]
    </div>
</div>
[[template "footer"]]
