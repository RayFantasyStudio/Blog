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
                <div class="description" id="description">
                    <script>
                        document.getElementById('description').innerHTML =
                                '[[$article.Content]]'
                        ;
                    </script>
                </div>
                <div class="extra">
                    <a class="ui right floated primary button" href="/article/view?id=[[$article.Id]]">
                        详情
                        <i class="right chevron icon"></i>
                    </a>
                    <div class="ui label">Tag</div>
                </div>
            </div>
        </div>
        [[end]]
        [[template "paginator" .Paginator]]
    </div>
</div>
[[template "footer"]]
