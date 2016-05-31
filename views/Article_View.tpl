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
            </div>
        </div>
        [[end]]
        <div class="ui raised segment">
            <div class="ui comments">
                <h3 class="ui dividing header">Comments</h3>

                <div class="comment" style="padding: 0 5px 0 5px">
                    <a class="avatar">
                        <img src="/images/avatar/small/matt.jpg">
                    </a>
                    <div class="content">
                        <a class="author">评论的用户</a>
                        <div class="metadata">
                            <span class="date">评论时间</span>
                        </div>
                        <div class="text">我是评论内容</div>
                        <div class="actions">
                            <a class="reply">回复</a>
                        </div>
                    </div>
                </div>

                <form class="ui reply form">
                    <div class="field">
                        <textarea></textarea>
                    </div>
                    <div class="ui blue labeled submit icon button"><i class="icon edit"></i> 添加评论</div>
                </form>
            </div>
        </div>
    </div>

</div>
[[template "footer"]]
