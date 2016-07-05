[[template "header" .]]
<script src="//cdn.bootcss.com/marked/0.3.5/marked.js"></script>
<script src="//cdn.bootcss.com/marked/0.3.5/marked.min.js"></script>
<div class="ui  container">
    [[$author_article := .author]]
    [[$tags := .Tags]]
    [[$isLogin := .IsLogin]]
    [[with .Article]]
    <div class="ui relaxed divided items">
        <div class="ui raised segment"><a class="ui red ribbon label">
            <div class="ui left icon input">
                <span style="color: white">[[.Category]]</span>
            </div>
        </a>
            <div class="ui">
                <div class="field" style="margin-bottom: 3px">
                    <h1>[[.Title]]</h1>
                </div>
                <div class="field">
                    <h3><span style="color: gray">[[.Subtitle]]</span></h3>
                </div>
                <div class="ui divider"></div>
                <div class="field">
                    <h4 class="ui image header">
                        <img src="/static/img/user_avatar/[[$author_article]]_avatar.png" class="ui mini rounded image"
                             onerror="javascript:this.src='/static/img/user_avatar/user.png'">
                        [[$author_article]]
                    </h4>
                </div>
                <div class="field">
                    <label class="ui green label">标签</label>[[range $tags]]<label class="ui label">[[.Name]]</label>[[end]]
                </div>
                <div class="field" id="markdown-content" style="padding: 50px">
                    <script>
                        document.getElementById('markdown-content').innerHTML =
                                '[[.Content]]'
                        ;
                    </script>
                </div>
                [[if $isLogin]]
                <a class="ui  green button" href="/article/modify?id=[[.Id]]">修改</a>
                <a class="ui red button" href="/article?op=del&id=[[.Id]]">删除</a>
                [[end]]
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
                            <span class="date">[[.Time | SinceTime]]</span>
                        </div>
                        [[if lt 0 .QuoteReplyId]]
                        [[$reply_tmp := index $.quote_map .QuoteReplyId]]
                        <div class="ui message">
                            <div class="header">[[$reply_tmp.UserName]]</div>
                            <p>[[$reply_tmp.Content]]</p>
                        </div>
                        [[end]]
                        <div class="text">[[.Content]]</div>
                        <div class="actions">
                            <a class="reply" href="/article/view?id=[[$aid]]&quote_rid=[[.Id]]">回复</a>
                            [[if $IsLogin]]
                            <a class="reply" href="/article/view?op=delrpy&rid=[[.Id]]&aid=[[$aid]]">删除回复</a>
                            [[end]]
                        </div>
                    </div>
                </div>
                [[end]]
                <div class="ui divider"></div>
                [[if lt 0 .quote_rid]]
                [[with .quote_reply]]
                <div class="ui message">
                    <i class="close icon"></i>
                    <div class="header"><span style="color: green">回复</span>[[.UserName]]</div>
                    <p>[[.Content]]</p>
                </div>
                [[end]]
                [[end]]
                <form class="ui reply form" action="/article/view" method="post">
                    <input type="hidden" name="quote_reply_id" value="[[.quote_reply.Id]]">
                    <input type="hidden" name="aid" value="[[.Article.Id]]">
                    <div class="fields">
                        <input type="text" name="reply_name" placeholder="请输入昵称">
                    </div>
                    <div class="field ">
                        <textarea name="reply_content" placeholder="请输入评论内容"></textarea>
                    </div>
                    <button class="ui blue labeled submit icon button" type="submit"><i class="icon edit"></i> 添加评论
                    </button>
                </form>
            </div>
        </div>
    </div>
</div>
[[template "footer"]]
