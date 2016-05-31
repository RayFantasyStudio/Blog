[[template "header" .]]
<div class="ui  container">

    [[with .Article]]
    <div class="ui relaxed divided items">
        <div class="ui raised segment"><a class="ui red ribbon label">
                <div class="ui left icon input">
                    <span style="color: white">[[.Category]]</span>
                </div>
            </a>
            <div class="ui form">
                <div class="field">
                    <h1>[[.Title]]</h1>
                </div>
                <div class="field">
                    <h3><span style="color: gray">[[.Subtitle]]</span></h3>
                </div>
                <div class="ui divider"></div>
                <div class="field">
                    <label>标签</label>标签在这
                </div>
                <div class="field" style="height: 400px">
                    [[.Content]]
                </div>
                <a class="ui  green button" href="/article/modify?id=[[.Id]]">修改</a>
                <a class="ui red button" href="/article?op=del&id=[[.Id]]">删除</a>
            </div>
        </div>
        [[end]]
        <div class="ui raised segment">
            <div class="ui comments">
                <h3 class="ui dividing header">Comments</h3>
                [[$aid := .Article.Id]]
                [[range .Replies]]
                <div class="comment" style="padding: 0 5px 0 5px">
                    <a class="avatar">
                        <img src="/images/avatar/small/matt.jpg">
                    </a>
                    <div class="content">
                        <a class="author">[[.UserName]]</a>
                        <div class="metadata">
                            <span class="date">[[.Time]]</span>
                        </div>
                        <div class="text">[[.Content]]</div>
                        <div class="actions">
                            <a class="reply">回复</a>
                            <a class="reply" href="/article/view?op=delrpy&rid=[[.Id]]&aid=[[$aid]]">删除回复</a>
                        </div>
                    </div>
                </div>
                [[end]]

                <form class="ui reply form" action="/article/view" method="post">
                    <input type="hidden" name="aid" value="[[.Article.Id]]">
                    <div class="fields">
                        <input type="text" name="reply_name" placeholder="请输入昵称">
                    </div>
                    <div class="field ">
                        <textarea name="reply_content" placeholder="请输入评论内容"></textarea>
                    </div>
                    <button class="ui blue labeled submit icon button" type="submit"><i class="icon edit"></i> 添加评论</button>
                </form>
            </div>
        </div>
    </div>

</div>
[[template "footer"]]
