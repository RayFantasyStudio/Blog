[[template "header" .]]

<div class="ui  container">
    [[with .Article]]
    <div class="ui relaxed divided items">
        <form method="post" action="/article/modify">
            <div class="ui raised segment"><a class="ui red ribbon label">
                    <div class="ui left icon input">
                        <input type="text" name="category" value="[[.Category]]">
                        <i class="users icon"></i>
                    </div>
                </a>
                <div class="ui" style="width: 100% ; padding: 3%">
                    <input type="hidden" value="[[.Id]]" name="id">
                    <input type="hidden" value="[[.Category]]" name="former_category">
                    <div class="ui form">
                        <div class="field">
                            <label>标题</label>
                            <input type="text" name="title" value="[[.Title]]">
                        </div>
                        <div class="field">
                            <label>副标题</label>
                            <input type="text" name="subtitle" value="[[.Subtitle]]">
                        </div>
                        <div class="field">
                            <label>标签</label>
                        </div>
                        <div class="field">
                            <input type="text" name="tags" placeholder="输入标签（空格分割）">
                        </div>

                        <div class="field">
                            <label>正文:</label>
                            <textarea name="content" rows="100">[[.Content]]</textarea>
                        </div>
                        <div class="field">
                            <button class="positive ui button" type="submit">修改</button>
                        </div>
                    </div>
                </div>

            </div>
        </form>
    </div>
    [[end]]
</div>

[[template "footer"]]
