
[[template "header" .]]
    <div class="ui text container">
        <div class="ui relaxed divided items" style="background-color: #EEEEEE">
            [[$authors := .Authors]]
            [[range $index, $article := .Articles]]
            <div class="item">
                <div class="content">
                    <a class="header" href="/article/view?id=[[$article.Id]]&quote_rid=0">[[$article.Title]]</a>
                    <div class="meta ui right floated">
                        <span>[[index $authors $index]]</span>
                        <span>[[$article.Created | SinceTime]]</span>
                    </div>
                    <div class="description" id="description">
                        <script>
                            document.getElementById('description').innerHTML =
                                    '[[$article.Content]]'
                            ;
                        </script>
                    </div>
                    <div class="extra">
                        <a class="ui right floated primary button" href="/article/view?id=[[$article.Id]]&quote_rid=0">
                            详情
                            <i class="right chevron icon"></i>
                        </a>
                    <label class="ui label left floated">[[$article.Category]]</label>
                    </div>
                </div>
            </div>
            [[end]]
            [[template "paginator" .Paginator]]
        </div>
    </div>

[[template "footer"]]
