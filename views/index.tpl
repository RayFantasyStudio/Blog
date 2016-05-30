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
    <script>
        $(document)
                .ready(function () {

                    // fix main menu to page on passing
                    $('.main.menu').visibility({
                        type: 'fixed'
                    });
                    $('.overlay').visibility({
                        type: 'fixed',
                        offset: 80
                    });

                    // lazy load images
                    $('.image').visibility({
                        type: 'image',
                        transition: 'vertical flip in',
                        duration: 500
                    });

                    // show dropdown on hover
                    $('.main.menu  .ui.dropdown').dropdown({
                        on: 'hover'
                    });
                })
        ;
    </script>
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

<div class="ui main text container">
    <h1 class="ui header">[[.Owner]]的Blog</h1>
    <p>This example shows how to use lazy loaded images, a sticky menu, and a simple text container</p>
</div>

<div class="ui borderless main menu">
    <div class="ui text container">
        <div href="#" class="header item">
            [[.User]]
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
