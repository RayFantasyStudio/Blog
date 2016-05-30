<!DOCTYPE html>

<html>
<head>
    <title>[[.Owner]]的Blog</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, minimum-scale=1, maximum-scale=1">
    <link rel="image_src" type="image/jpeg" href="/static/img/logo.png">
    <link href="//cdn.bootcss.com/semantic-ui/2.1.8/semantic.min.css" class="ui" rel="stylesheet">
    <script src="//cdn.bootcss.com/jquery/2.2.4/jquery.js"></script>
    <script src="//cdn.bootcss.com/semantic-ui/2.1.8/semantic.min.js"></script>
    <style type="text/css">

        body {
            background-color: #FFFFFF;
        }

        .main.container {
            margin-top: 2em;
        }

        .main.menu {
            margin-top: 4em;
            border-radius: 0;
            border: none;
            box-shadow: none;
            transition: box-shadow 0.5s ease,
            padding 0.5s ease;
        }

        .main.menu .item img.logo {
            margin-right: 1.5em;
        }

        .overlay {
            float: left;
            margin: 0em 3em 1em 0em;
        }

        .overlay .menu {
            position: relative;
            left: 0;
            transition: left 0.5s ease;
        }

        .main.menu.fixed {
            background-color: #FFFFFF;
            border: 1px solid #DDD;
            box-shadow: 0px 3px 5px rgba(0, 0, 0, 0.2);
        }

        .overlay.fixed .menu {
            left: 800px;
        }

        .text.container .left.floated.image {
            margin: 2em 2em 2em -4em;
        }

        .text.container .right.floated.image {
            margin: 2em -4em 2em 2em;
        }

        .ui.footer.segment {
            margin: 5em 0em 0em;
            padding: 5em 0em;
        }
    </style>
</head>

<body>

<div class="ui borderless main menu">
    <div class="ui text container">
        <div href="#" class="header item">
            user
        </div>
        <a href="#" class="item">Blog</a>
        <a href="#" class="item">Articles</a>
        <a href="#" class="ui right floated dropdown item" tabindex="0">
            Dropdown <i class="dropdown icon"></i>
            <div class="menu transition hidden" tabindex="-1">
                <div class="item">Link Item</div>
                <div class="item">Link Item</div>
                <div class="divider"></div>
                <div class="header">Header Item</div>
                <div class="item">
                    <i class="dropdown icon"></i>
                    Sub Menu
                    <div class="menu transition hidden">
                        <div class="item">Link Item</div>
                        <div class="item">Link Item</div>
                    </div>
                </div>
                <div class="item">Link Item</div>
            </div>
        </a>
    </div>
</div>

<div class="ui  container">
    <div class="ui relaxed divided items">
        <div class="ui raised segment"><a class="ui red ribbon label">
                <div class="ui left icon input">
                    <span style="color: white">我是标签</span>
                </div>
            </a>
            <div class="ui form">
                <div class="field">
                    <h1>标题在这</h1>
                </div>
                <div class="field">
                    <h3><span style="color: gray">副标题在这</span></h3>
                </div>
                <div class="ui divider"></div>
                <div class="field">
                    <label>标签</label>标签在这
                </div>
                <div class="field" style="height: 400px">
                    我是正文
                </div>
            </div>
        </div>
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

<div class="ui inverted vertical footer segment">
    <div class="ui center aligned container">
        <div class="ui stackable inverted divided grid">
            <div class="three wide column">
                <h4 class="ui inverted header">Group 1</h4>
                <div class="ui inverted link list">
                    <a href="#" class="item">Link One</a>
                    <a href="#" class="item">Link Two</a>
                    <a href="#" class="item">Link Three</a>
                    <a href="#" class="item">Link Four</a>
                </div>
            </div>
            <div class="three wide column">
                <h4 class="ui inverted header">Group 2</h4>
                <div class="ui inverted link list">
                    <a href="#" class="item">Link One</a>
                    <a href="#" class="item">Link Two</a>
                    <a href="#" class="item">Link Three</a>
                    <a href="#" class="item">Link Four</a>
                </div>
            </div>
            <div class="three wide column">
                <h4 class="ui inverted header">Group 3</h4>
                <div class="ui inverted link list">
                    <a href="#" class="item">Link One</a>
                    <a href="#" class="item">Link Two</a>
                    <a href="#" class="item">Link Three</a>
                    <a href="#" class="item">Link Four</a>
                </div>
            </div>
            <div class="seven wide column">
                <h4 class="ui inverted header">Footer Header</h4>
                <p>Extra space for a call to action inside the footer that could help re-engage users.</p>
            </div>
        </div>
        <div class="ui inverted section divider"></div>
        <img src="http://semantic-ui.com/examples/assets/images/logo.png" class="ui centered mini image">
        <div class="ui horizontal inverted small divided link list">
            <a class="item" href="#">Site Map</a>
            <a class="item" href="#">Contact Us</a>
            <a class="item" href="#">Terms and Conditions</a>
            <a class="item" href="#">Privacy Policy</a>
        </div>
    </div>
</div>


</body>
</html>