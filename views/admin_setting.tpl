[[define "admin_setting"]]
<div class="ui card" id="filter-card">
    <form class="ui form" enctype="multipart/form-data" action="/admin" method="post">
        <input type="hidden" name="op" value="account_setting">
        <div class="field">
            <label>修改账户</label>
            <div class="three fields">
                <div class="field">
                    <input type="text" name="username" placeholder="用户名" value="[[.UserName]]">
                </div>
                <div class="field">
                    <input type="password" name="pwd" placeholder="密码">
                </div>
            </div>
        </div>
        <div class="field">
            <label>修改头像</label>
            <img src="[[.avatar_path]]" style="width: 100px;height: 100px"
                 onerror="javascript:this.src='/static/img/user_avatar/user.png'">
            <input type="file" name="avatar">
            <input type="hidden" name="user_name" value="[[.UserName]]">
        </div>
        <div class="ui divider"></div>
        <div class=" ui field" style="margin-bottom: 10px">
            <button class="ui green button">保存更改</button>
        </div>
    </form>
</div>
[[end]]