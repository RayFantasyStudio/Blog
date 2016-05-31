[[template "header" .]]

<div class="ui  container">
    <div class="ui relaxed divided items">
        <div class="ui raised segment"><a class="ui red ribbon label">
                <div class="ui left icon input">
                    <input type="text" placeholder="输入分类">
                    <i class="users icon"></i>
                </div>
            </a>
            <form class="ui" style="width: 100% ; padding: 3%" method="post" action="/article">
                <div class="ui form">
                    <div class="field">
                        <label>标题</label>
                        <input type="text" name="title" placeholder="输入你的标题">
                    </div>
                    <div class="field">
                        <label>副标题</label>
                        <input type="text" name="subtitle" placeholder="输入你的副标题">
                    </div>
                    <div class="field">
                        <label>标签</label>
                    </div>
                    <div class="field">
                        <input type="text" name="tags" placeholder="输入标签（空格分割）">
                    </div>

                    <div class="field">
                        <label>正文:</label>
                        <textarea name="content" rows="100"></textarea>
                    </div>
                    <div class="field">
                        <button class="positive ui button" type="submit">发表</button>
                    </div>
                </div>
            </form>
        </div>

    </div>
</div>

[[template "footer"]]
